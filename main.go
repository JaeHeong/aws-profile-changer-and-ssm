package main

import (
	"github.com/JaeHeong/aws-profile-changer-and-ssm/cmd"
)

var gossmVersion string

func main() {
	cmd.Execute(gossmVersion)
}
