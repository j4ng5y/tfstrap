package tfstrap

import (
	"os"
	"path"
)

// DirectoryStructure holds all releveant issues to building out the directory
type DirectoryStructure struct {
	RootDirPath          string
	DirectoryPermissions os.FileMode
	ConfigFile           string
	VariablesFile        string
	ModulesDirectoryName string
}

func (D *DirectoryStructure) Write() error {
	if err := os.MkdirAll(D.RootDirPath, D.DirectoryPermissions); err != nil {
		return err
	}
	if err := os.Mkdir(path.Join(D.RootDirPath, D.ModulesDirectoryName), D.DirectoryPermissions); err != nil {
		return err
	}
	if err := writeTFV12File(D.RootDirPath); err != nil {
		return err
	}
	if err := writeTFV12File(path.Join(D.RootDirPath, D.ModulesDirectoryName)); err != nil {
		return err
	}

	return nil
}

func writeTFV12File(dirPath string) error {
	versiontfFileContents := `
terraform {
  required_version = ">= 0.12"
}

`
	f, err := os.Create(path.Join(dirPath, "versions.tf"))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(versiontfFileContents))
	if err != nil {
		return err
	}

	return nil
}
