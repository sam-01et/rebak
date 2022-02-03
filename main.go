package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-git/go-git/v5"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type GitJsonResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Private bool `json:"private"`
}

//default permission for a newly created file
const newFileDefaultPermission = 0775
const defaultRebakDirectory = "rebak_dir"

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal("There seems to be an issue with the current directory")
	}
	var gitAccountUsername = flag.String("gituser", "", "Github user account")
	var cloneDir = flag.String("dir", cwd, "Directory to clone to")

	flag.Parse()

	dir := *cloneDir

	if !argsAreValid(*gitAccountUsername) {
		fmt.Println("1")
		os.Exit(1)
	}

	if len(dir) <= 0 {
		log.Fatalln("Empty directory.")
	}

	//If specified directory to backup to does not exist, create it
	createDirIfNotExists(dir)

	if *cloneDir == cwd {
		//to prevent overcrowding current dir, create a new dir inside the current dir- by default name it rebak_dir
		dir = createDir(dir, "", false)
		log.Printf("No directory specified, using dafault location %s \n", dir)
	}

	fmt.Println("Fetching repositories")

	repos := fetchRepositories(*gitAccountUsername)

	if len(repos) == 0 {
		log.Println("No repositories were found.")
		os.Exit(1)
	}

	fmt.Printf("Found repos: %s \n", repos)
	fmt.Println("**********************************************")

	for _, repo := range repos {
		newDir := createDir(dir, repo, true)
		baseGitUrl := strings.Join([]string{"https://github.com/", *gitAccountUsername, "/", repo, ".git"}, "")
		fmt.Printf("Cloning %s \n", repo)
		_, _ = git.PlainClone(newDir, false, &git.CloneOptions{
			URL: baseGitUrl,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			Progress: log.Writer(),
		})
	}

	fmt.Println("Cloning completed")
}

func createDirIfNotExists(dir string)  {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, newFileDefaultPermission)
		if err != nil {
			log.Fatal("Failed to create a directory that does not exist.")
		}
	}
}

func createDir(dir string, repo string, createRepositoryDir bool) string {
	var newDir string
	var fileSeparator = string(filepath.Separator)

	if len(dir) > 0 {
		if createRepositoryDir && len(repo) > 0 {
			newDir = strings.Join([]string{dir, fileSeparator , repo}, "")
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

//Validates the commandline arguments
func argsAreValid(gitAccountUsername string) bool {
	if len(gitAccountUsername) == 0 {
		log.Println("Please specify git username")
		return false
	}
	return true
}

func createUrl(accountUsername string) string{
	s := []string{"https://api.github.com/users/", accountUsername, "/repos"}
	return strings.Join(s, "")
}

//Fetches repositories for the given github account
func fetchRepositories (accountUsername string) []string{
	url := createUrl(accountUsername)
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