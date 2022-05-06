package hub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"rebak/helpers"
	"strings"

	"github.com/go-git/go-git/v5"
)

type GitJsonResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

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

//Fetches repositories for the given github account
func FetchRepositories(accountUsername string) []string {
	url := helpers.CreateUrl(accountUsername)
	response, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gitResponses []GitJsonResponse
	json.Unmarshal(responseData, &gitResponses)

	var repos []string
	for _, value := range gitResponses {
		repos = append(repos, value.Name)
	}
	return repos
}
