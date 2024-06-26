package update

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThis(t *testing.T) {
	inc := "0001-01-01"
	tm, err := time.Parse("2006-01-02", inc)
	assert.NoError(t, err)
	fmt.Println(tm.IsZero())
}
