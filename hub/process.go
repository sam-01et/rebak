package hub

import (
	"fmt"
	"log"
	"rebak/helpers"
	"strings"
	"github.com/go-git/go-git/v5"
)

func StartCloning(repos []string, dir string, gitAccountUsername string) {
	for _, repo := range repos {
		newDir := helpers.CreateDir(dir, repo, true)
		baseGitUrl := strings.Join([]string{"https://github.com/", gitAccountUsername, "/", repo, ".git"}, "")
		fmt.Printf("Cloning %s \n", repo)
		_, _ = git.PlainClone(newDir, false, &git.CloneOptions{
			URL:               baseGitUrl,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Progress:          log.Writer(),
		})
	}
}
