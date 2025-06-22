package cipher

import (
	"path/filepath"
	"strings"
)

func getNameWithoutExt(filename string) string {
	base := filepath.Base(filename)
	for {
		ext := filepath.Ext(base)
		if ext == "" {
			break
		}
		base = strings.TrimSuffix(base, ext)
	}
	return base
}
