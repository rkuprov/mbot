package ui

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type PrintCfg struct {
	Title  string
	Header table.Row
	Body   []table.Row
}

func Tabular(cfg PrintCfg) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(cfg.Title)
	if len(cfg.Header) > 0 {
		t.AppendHeader(cfg.Header)
	}
	if len(cfg.Body) == 0 {
		t.Render()
	}

	rows := make([]table.Row, 0, 4)
	for _, row := range cfg.Body {
		rows = append(rows, row)
		if len(rows) == 4 {
			t.AppendRows(rows)
			rows = rows[:0]
			t.AppendSeparator()
		}
	}
	t.AppendRows(rows)

	t.Render()
}
