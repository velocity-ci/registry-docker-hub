package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const registryAddress = "https://index.docker.io"

func main() {
	out := &output{
		Error:         "",
		State:         "success",
		ServerAddress: registryAddress,
	}
	defer finish(out)

	username := strings.TrimSpace(os.Getenv("username"))
	password := strings.TrimSpace(os.Getenv("password"))

	if len(username) < 1 {
		out.State = "error"
		out.Error = "missing username"
		return
	}

	if len(password) < 1 {
		out.State = "error"
		out.Error = "missing password"
		return
	}

	out.Username = username
	out.Password = password
}

func finish(out *output) {
	o, _ := json.Marshal(out)

	fmt.Printf("%s\n", o)
}

type output struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	ServerAddress string `json:"serverAddress"`
	Error         string `json:"error"`
	State         string `json:"state"`
}
