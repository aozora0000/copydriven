package main

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
)

var version = "unknown"
var revision = "unknown"

func main() {
	app := cli.NewApp()
	app.Usage = "Copy Driven: automate manipulate when copy"
	app.Version = version + "@" + revision
	app.EnableBashCompletion = true
	app.Action = func(context *cli.Context) error {
		if context.NArg() == 0 {
			return errors.New("required execute command")
		}
		command := context.Args().Get(0)
		str := ""
		for {
			input, _ := clipboard.ReadAll()
			if str != input {
				cmd := exec.Command(os.Getenv("SHELL"), []string{"-c", fmt.Sprintf(command, input)}...)
				cmd.Stdout = os.Stdout
				fmt.Println(cmd.String())
				err := cmd.Run()
				if err != nil {
					return err
				}
				str = input
			}
		}
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
