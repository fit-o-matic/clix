package command

import (
	"github.com/finkt/clix-kit/cache"
	"github.com/finkt/clix-kit/folder"
)

type Registry struct {
	folder *folder.Folder
	cache  *cache.Cache
}

func NewRegistry(parentFolder *folder.Folder) *Registry {
	var registryFolder = parentFolder.GetSubfolder("registry")

	return &Registry{
		folder: registryFolder,
		cache:  cache.New(registryFolder),
	}
}

type Manifest struct {
	Summaries []Summary `json:"commands"`
}

type Commands []*Command

func (r *Registry) GetManifest() (*Manifest, error) {

	var manifest Manifest

	err := r.cache.ReadOrCreateJson("manifest.json", &manifest, r.metadataSupplier())
	if err != nil {
		return nil, err
	}
	return &manifest, nil
}

func (r *Registry) ClearCache() error {
	return r.cache.Clear()
}

func (r *Registry) GetCommand(name string) (*Command, error) {
	cmdFolder := r.folder.GetSubfolder(name)
	if !cmdFolder.Exists() {
		return nil, nil
	}
	cmd, err := Load(cmdFolder)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func (c Commands) createManifest() *Manifest {
	var summaries []Summary
	for _, command := range c {
		summaries = append(summaries, Summary{
			Name:        command.GetName(),
			Description: command.GetDescription(),
		})
	}
	return &Manifest{Summaries: summaries}

}

func (r *Registry) loadCommands() (Commands, error) {
	subfolders, err := r.folder.GetSubfolders()
	if err != nil {
		return nil, err
	}

	var commands Commands
	for _, subfolder := range subfolders {
		command, err := Load(subfolder)
		if err != nil {
			return nil, err
		}
		commands = append(commands, command)
	}
	return commands, nil
}

func (r *Registry) metadataSupplier() func() (any, error) {
	return func() (any, error) {
		commands, err := r.loadCommands()
		if err != nil {
			return nil, err
		}
		return commands.createManifest(), nil
	}
}
