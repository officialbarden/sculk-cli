/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"sculk/cmd"
	"sculk/src/commands/config"
)

func main() {
	config.ConfigExists()
	cmd.Execute()
}
