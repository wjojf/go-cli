package key

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type PrivateKey string

type PrivateKeyLoader struct {
	Filepath string
}

func GetPrivateKeyLoaders() []PrivateKeyLoader {

	var dir = getSSHDir()
	var loaders []PrivateKeyLoader

	filepath.WalkDir(
		dir,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			if d.Name()[:3] != "id_" {
				return nil
			}

			if filepath.Ext(path) == ".pub" {
				return nil
			}

			loaders = append(loaders, getLoader(path))

			return nil
		},
	)

	return loaders
}

func (l PrivateKeyLoader) Load() (PrivateKey, error) {
	fp := filepath.Join(l.Filepath)
	key, err := os.ReadFile(fp)
	if err != nil {
		return "", err
	}

	return PrivateKey(key), nil
}

func getLoader(path string) PrivateKeyLoader {
	return PrivateKeyLoader{
		Filepath: path,
	}
}

func getSSHDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error getting home dir", err)
		return ""
	}

	return filepath.Join(home, ".ssh")
}
