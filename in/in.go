package in

import (
	"os"
	"strings"
)

func Lines(path string) ([]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	d := string(b)
	lines := strings.Split(d, "\n")
	return lines, nil
}
