package fileutils

import (
	"fmt"
	"os"
	"path/filepath"

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
	kitexDir := "kitex-template"

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

// Copy all template files from the embedded FS to the specified path
func CopyAllTemplateFiles(to string) error {
	// Get the directory listing from the embedded FS
	files, err := templates.ReadDir("templates")
	if err != nil {
		klog.Errorf("Could not read directory from embedded FS: %v", err)
		return err
	}

	// Loop through each file and copy it to the destination directory
	for _, file := range files {
		data, err := templates.ReadFile(fmt.Sprintf("templates/%s", file.Name()))
		if err != nil {
			klog.Errorf("Could not read template file %s: %v", file.Name(), err)
			return err
		}

		// Construct the destination path for each file
		destination := filepath.Join(to, file.Name())

		err = os.WriteFile(destination, data, 0o600)
		if err != nil {
			klog.Errorf("Could not write template file %s: %v", destination, err)
			return err
		}
	}

	return nil
}
