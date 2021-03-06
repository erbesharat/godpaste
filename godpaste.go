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
		os.Exit(10)
	}
}

func main() {
	var file string
	var expire string
	var syntax string
	v := url.Values{}

	app := cli.NewApp()

	app.Name = "godpaste"
	app.Usage = "A command line tool for creating items on dpaste"
	app.UsageText = "godpaste FILE_NAME"
	app.Authors = []cli.Author{
		{
			Name:  "Erfan Besharat \"@erbesharat\"",
			Email: "erbesharat@gmail.com",
		},
	}
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Value:       "none",
			Usage:       "Your text's file",
			Destination: &file,
		},
		cli.StringFlag{
			Name:        "expire, e",
			Value:       "none",
			Usage:       "Item's expiry date",
			Destination: &expire,
		},
		cli.StringFlag{
			Name:        "syntax, s",
			Value:       "none",
			Usage:       "Set syntax for your item",
			Destination: &syntax,
		},
	}

	app.Action = func(c *cli.Context) error {
		if file != "none" {
			text, err := ioutil.ReadFile(file)
			checkerr(err, "Couldn't find your file")
			v.Add("content", string(text))
			if syntax != "none" {
				v.Add("syntax", syntax)
			}
			if expire != "none" {
				v.Add("expiry_days", expire)
			}
			resp, err := http.PostForm("http://dpaste.com/api/v2/", v)
			checkerr(err, "Couldn't create your dpaste")
			output, _ := resp.Location()
			if expire != "none" {
				fmt.Printf("link ==> %q ---- Expires in %q days", output, expire)
			} else {
				fmt.Printf("link ==> %q", output)
			}
		} else {
			fmt.Println("Please give a file. For more information check --help")
		}
		return nil
	}

	app.Run(os.Args)
}
