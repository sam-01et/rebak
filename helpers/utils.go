package helpers

import "strings"

func CreateUrl(accountUsername string) string {
	s := []string{"https://api.github.com/users/", accountUsername, "/repos"}
	return strings.Join(s, "")
}
