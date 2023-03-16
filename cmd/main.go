package main

import (
	"divo/cmd/boot"
	"divo/cmd/cmd"
)

func main() {
	if err := boot.OnBoot(); err != nil {
		panic(err)
	}
	cmd.Execute()
}
