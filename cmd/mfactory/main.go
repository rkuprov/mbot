package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/rkuprov/mbot/pkg/cfg"
)

//go:embed panel_calculator.xlsm
var fs embed.FS

func main() {
	now := time.Now()
	ctx := context.Background()
	errors := make(chan error, 1)

	err := checkSubscription(ctx)
	if err != nil {
		panic(err)
	}

	err = checkUpdates(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(now))
	}()
	name, err := createFile(ctx)
	if err != nil {
		panic(err)
	}
	defer os.Remove(name)

	go func() {
		var cmd *exec.Cmd
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("Opening file with default mac application")
			cmd = exec.Command("open", "-a", "Numbers", name)
		case "linux":
			fmt.Println("Opening file with default linux application")
		case "windows":
			fmt.Println("Opening file with default windows application")
			cmd = exec.Command("cmd", "start", name)
		default:
			fmt.Println(os)
		}
		if err = cmd.Run(); err != nil {
			fmt.Println(err)
			panic(err)
		}

		pid := cmd.Process.Pid
		fmt.Printf("Process ID: %d\n", pid)
		time.Sleep(5 * time.Second)
		ctx.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case err = <-errors:
			if err != nil {
				panic(err)
			}
		}
	}
}

func createFile(_ context.Context) (string, error) {
	bytes, err := fs.ReadFile("panel_calculator.xlsm")
	if err != nil {
		return "", err
	}
	tmpFile, err := os.CreateTemp("", "panel_calculator_temp.xlsm")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	_, err = tmpFile.Write(bytes)
	if err != nil {
		return "", err
	}

	fmt.Println("File created:", tmpFile.Name())
	return tmpFile.Name(), nil
}

func checkSubscription(ctx context.Context) error {
	client := &http.Client{}
	body := bytes.NewBufferString("id")
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/status", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d, error: %s", resp.StatusCode, out)
	}

	return nil
}

func checkUpdates(ctx context.Context) error {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/version", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d, error: %s", resp.StatusCode, out)
	}

	version, err := os.ReadFile(filepath.Join(cfg.InstallPath(), "VERSION"))
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}

	soft, firm := toUpdate(version, out)
	if soft {
		err = updateSoft(ctx)
	}
	if firm {
		err = updateFirm(ctx)
	}

	return nil
}

func toUpdate(current, received []byte) (bool, bool) {
	splitCurrent := bytes.Split(current, []byte("\n"))
	splitReceived := bytes.Split(received, []byte("\n"))

	return !bytes.Equal(splitCurrent[0], splitReceived[0]), !bytes.Equal(splitCurrent[1], splitReceived[1])
}

func updateSoft(ctx context.Context) error {
	return update(ctx, "soft")
}
func updateFirm(ctx context.Context) error {
	return update(ctx, "firm")
}

func update(ctx context.Context, target string) error {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(`%s/update%s`, cfg.UpdateURL(), target), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", cfg.SubscriptionToken())
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d, error: %s", resp.StatusCode, err)
	}

	checkSum := resp.Header.Get("mbot-checksum")
	// update version
	err = updateFileContents(ctx, filepath.Join(cfg.InstallPath(), "VERSION"), bytes.NewReader([]byte(target)))
	if err != nil {
		return err
	}

	// update binary
	var updatePath string
	switch target {
	case "soft":
		updatePath = filepath.Join(cfg.InstallPath(), "data.mbot")
	case "firm":
		updatePath = filepath.Join(cfg.BinaryPath(), cfg.BinaryName())

	}
	err = updateFileContents(ctx, updatePath, bytes.NewReader(out))
	if err != nil {
		return err
	}
	err = validateFile(ctx, updatePath, checkSum)
	if err != nil {
		return err
	}

	fmt.Println("Update successful")
	return nil
}

func updateFileContents(ctx context.Context, name string, payload io.Reader) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, payload)

	return err
}

func validateFile(ctx context.Context, name, checksum string) error {
	f, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	if checksum != fmt.Sprintf("%x", sha256.Sum256(f)) {
		return fmt.Errorf("checksums do not match")
	}

	return nil
}
