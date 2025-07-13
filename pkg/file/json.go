package file

import (
	"encoding/json"
	"github.com/goppuchino/oag/pkg/generator"
	"os"
	"path/filepath"
)

func SaveAsJSON(spec *generator.Spec, filename string) error {
	data, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(filename)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
