package helpers

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const newFileDefaultPermission = 0775
const defaultRebakDirectory = "rebak_dir"

func CreateUrl(accountUsername string) string {
	s := []string{"https://api.github.com/users/", accountUsername, "/repos"}
	return strings.Join(s, "")
}

func CreateDir(dir string, repo string, createRepositoryDir bool) string {
	var newDir string
	var fileSeparator = string(filepath.Separator)

	if len(dir) > 0 {
		if createRepositoryDir && len(repo) > 0 {
			newDir = strings.Join([]string{dir, fileSeparator, repo}, "")
		} else if !createRepositoryDir && len(repo) == 0 {
			newDir = strings.Join([]string{dir, fileSeparator, defaultRebakDirectory}, "")
		}
	} else {
		log.Fatalln("")
	}

	err := os.Mkdir(newDir, newFileDefaultPermission)
	if err != nil {
		log.Print(err)
	}
	return newDir
}
