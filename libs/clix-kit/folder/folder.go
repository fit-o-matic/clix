package folder

import (
	"os"
	"path/filepath"
	"strings"
)

type Folder struct {
	path string
}

func (f *Folder) Exists() bool {
	if _, err := os.Stat(f.path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (f *Folder) Contains(s string) bool {
	if _, err := os.Stat(filepath.Join(f.path, s)); os.IsNotExist(err) {
		return false
	}
	return true
}

func UserHome() (*Folder, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return &Folder{path: homeDir}, nil
}

func New(path string) *Folder {
	return &Folder{path: path}
}

func (f *Folder) GetPath() string {
	return f.path
}

func (f *Folder) GetName() string {
	parts := strings.Split(f.path, string(os.PathSeparator))
	return parts[len(parts)-1]
}

func (f *Folder) GetParent() *Folder {
	parentPath := filepath.Dir(f.path)
	return &Folder{path: parentPath}
}

func (f *Folder) GetSubfolders() ([]*Folder, error) {
	entries, err := os.ReadDir(f.path)
	if err != nil {
		return nil, err
	}

	var subfolders []*Folder
	for _, entry := range entries {
		if entry.IsDir() {
			subfolders = append(subfolders, &Folder{path: filepath.Join(f.path, entry.Name())})
		}
	}
	return subfolders, nil
}

func (f *Folder) GetFiles() ([]string, error) {
	entries, err := os.ReadDir(f.path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func (f *Folder) GetSubfolder(name string) *Folder {
	subfolderPath := filepath.Join(f.path, name)
	return &Folder{path: subfolderPath}
}

func (f *Folder) EnsureExists() error {
	return os.MkdirAll(f.path, 0755)
}

func (f *Folder) Delete() error {
	return os.RemoveAll(f.path)
}
