package cmd

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.New("isosec", "help", "version")
	x.Usage = `[help|version]`
	x.Summary = `Current UTC time in YYYYMMDDhhmmss format`
	x.Version = `v0.0.1`

	x.Description = `
		The *isosec* command returnd the current time in the following format:
		YYYYMMDDhhmmss. All time is displayed as UTC time.`

	x.Method = func(args []string) (err error) {

		if len(args) > 0 {
			switch args[0] {
			case "help":
				fmt.Println(x.Description)
				return nil
			case "version":
				fmt.Println(x.Version)
				return nil
			default:
				return x.UsageError()
			}
			return nil
		}

		fmt.Println(time.Now().UTC().Format("20060102150405"))
		return nil

	}
}
