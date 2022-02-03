package main

import (
	"testing"
)

func Test_createUrl(t *testing.T) {
	type args struct {
		accountUsername string
		want string
	}
	tests := [] args {
		{accountUsername: "samson", want: "https://api.github.com/users/samson/repos"},
		{accountUsername: "john", want: "https://api.github.com/users/john/repos"},
		{accountUsername: "gang", want: "https://api.github.com/users/gang/repos"},
	}

	for _, tt := range tests {
		t.Run(tt.accountUsername, func(t *testing.T) {
			if got := createUrl(tt.accountUsername); got != tt.want {
				t.Errorf("createUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}