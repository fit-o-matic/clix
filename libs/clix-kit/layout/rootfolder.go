package layout

import (
	"path/filepath"

	"github.com/finkt/clix-kit/folder"
)

type RootFolder struct {
	folder *folder.Folder
}

func NewRoot(cliName string) (*RootFolder, error) {
	userHome, err := folder.UserHome()
	if err != nil {
		return nil, err
	}
	cliHome := filepath.Join(userHome.GetPath(), ".", cliName)
	return &RootFolder{
		folder: folder.New(cliHome),
	}, nil
}
