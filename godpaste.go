package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func checkerr(e error, message string) {
	if e != nil {
		fmt.Println(message)
	}
}

func main() {

	app := cli.NewApp()

	app.Name = "godpaste"
	app.Usage = "A command line tool for creating items on dpaste"
	app.UsageText = "godpaste FILE_NAME"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Erfan Besharat \"@erbesharat\"",
			Email: "erbesharat@gmail.com",
		},
	}
	app.Version = "0.1"

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) != 0 {
			input := c.Args().First()
			text, err := ioutil.ReadFile(input)
			checkerr(err, "Couldn't find your file")

			v := url.Values{}
			v.Add("content", string(text))
			resp, err := http.PostForm("http://dpaste.com/api/v2/", v)
			checkerr(err, "Couldn't create your dpaste")
			output, _ := resp.Location()
			fmt.Printf("Check ==> %q", output)
		} else {
			fmt.Println("Please give a file. For more information check --help")
		}
		return nil
	}

	app.Run(os.Args)
}
