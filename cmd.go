package cmd

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdbox"
	_ "github.com/rwxrob/cmdbox-version"
)

func init() {
	x := cmdbox.Add("isosec", "help", "version")
	x.Usage = `[h|help|version]`
	x.Summary = `Current UTC time in YYYYMMDDhhmmss format`
	x.Version = `v0.0.7`

	x.Description = `
		The *isosec* command returnd the current time in the following format:
		YYYYMMDDhhmmss. All time is displayed as UTC time.`

	x.Method = func(args []string) (err error) {

		if len(args) > 0 {
			switch args[0] {
			case "h", "help":
				fmt.Println(x.Description)
				return nil
			case "version":
				fmt.Println(x.VersionLine())
				return nil
			default:
				return x.UsageError()
			}
		}

		fmt.Println(time.Now().UTC().Format("20060102150405"))
		return nil

	}
}
