package file

import (
	"github.com/goppuchino/oag/pkg/generator"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func SaveAsYAML(spec *generator.Spec, filename string) error {
	data, err := yaml.Marshal(spec)
	if err != nil {
		return err
	}

	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
