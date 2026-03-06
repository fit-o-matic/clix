package cache

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/finkt/clix-kit/folder"
)

type Cache struct {
	folder *folder.Folder
}

func (c *Cache) Contains(s string) bool {
	return c.folder.Contains(s)
}

func New(parentFolder *folder.Folder) *Cache {
	return &Cache{folder: parentFolder.GetSubfolder(".cache")}
}

func (c *Cache) Clear() error {
	return c.folder.Delete()
}

func (c *Cache) ReadJson(filename string, v any) error {
	filePath := filepath.Join(c.folder.GetPath(), filename)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (c *Cache) WriteJson(filename string, v any) error {
	if err := c.folder.EnsureExists(); err != nil {
		return err
	}
	filePath := filepath.Join(c.folder.GetPath(), filename)
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

// ReadOrCreateJson reads JSON from a file into v. If the file does not exist,
// it calls create to get the initial value, writes it to the file, and unmarshals it into v.
func (c *Cache) ReadOrCreateJson(filename string, v any, create func() (any, error)) error {
	err := c.ReadJson(filename, v)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}

	initial, err := create()
	if err != nil {
		return err
	}
	if err := c.WriteJson(filename, initial); err != nil {
		return err
	}
	return c.ReadJson(filename, v)
}
