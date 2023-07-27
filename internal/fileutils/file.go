package fileutils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/tim-pipi/cloudwego-api-gateway/internal/dir"
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
func CopyAllTemplateFiles(to string, relativePath string) error {
	// Get the directory listing from the embedded FS
	files, err := templates.ReadDir(filepath.Join("templates", relativePath))
	if err != nil {
		klog.Errorf("Could not read directory from embedded FS: %v", err)
		return err
	}

	// Loop through each file and copy it to the destination directory
	for _, file := range files {
		// Recursively copy all files in the nested directory
		if file.IsDir() {
			nestedDir := filepath.Join(to, relativePath, file.Name())
			// Check if the directory exists
			ok, _ := dir.Exists(nestedDir)

			newRelativePath := filepath.Join(relativePath, file.Name())
			if !ok {
				os.MkdirAll(newRelativePath, 0777)
			}

			if err := CopyAllTemplateFiles(to, newRelativePath); err != nil {
				klog.Errorf("Could not copy files in nested directory %s: %v", nestedDir, err)
			}

			continue
		}

		data, err := templates.ReadFile(filepath.Join("templates", relativePath, file.Name()))
		if err != nil {
			klog.Errorf("Could not read template file %s: %v", file.Name(), err)
			return err
		}

		// Construct the destination path for each file
		destination := filepath.Join(to, relativePath, file.Name())

		err = os.WriteFile(destination, data, 0o600)
		if err != nil {
			klog.Errorf("Could not write template file %s: %v", destination, err)
			return err
		}
	}

	return nil
}
