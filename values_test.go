package xhttp

import (
	"fmt"
	"testing"
)

func TestValues(t *testing.T) {
	m := make(Values)
	m.Set("main", "1")
	fmt.Println(m)
}
