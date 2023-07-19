package fileutils

import (
	"fmt"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
)

var templateNames []string = []string{
	"handler_tpl.yaml",
	"main_tpl.yaml",
	"middleware_tpl.yaml",
	"readme_tpl.yaml",
}

// Copies all template files from kitex to the specified directory
func CopyTemplateKitexDir(dir string) error {
	kitexDir := "kitex"

	for _, name := range templateNames {
		filename := fmt.Sprintf("%s/%s", kitexDir, name)
		if err := CopyTemplateFile(filename, fmt.Sprintf("%s/%s", dir, name)); err != nil {
			return err
		}
	}

	return nil
}

// Copy a template file to the specified path
func CopyTemplateFile(name, to string) error {
	data, err := templates.ReadFile(fmt.Sprintf("templates/%s", name))

	if err != nil {
		klog.Errorf("Could not find template file %s: %v", name, err)
		return err
	}

	err = os.WriteFile(to, data, 0o600)
	if err != nil {
		klog.Errorf("Could not find template file %s: %v", to, err)
		return err
	}

	return nil
}
