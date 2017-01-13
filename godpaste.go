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
		cli.Author{
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
		if file != "none" && expire == "none" {
			text, err := ioutil.ReadFile(file)
			checkerr(err, "Couldn't find your file")
			v.Add("content", string(text))
			v.Add("syntax", syntax)
			resp, err := http.PostForm("http://dpaste.com/api/v2/", v)
			checkerr(err, "Couldn't create your dpaste")
			output, _ := resp.Location()
			fmt.Printf("Check ==> %q", output)
		} else if file != "none" && expire != "none" {
			text, err := ioutil.ReadFile(file)
			checkerr(err, "Couldn't find your file")
			v.Add("content", string(text))
			v.Add("expiry_days", expire)
			v.Add("syntax", syntax)
			resp, err := http.PostForm("http://dpaste.com/api/v2/", v)
			checkerr(err, "Couldn't create your dpaste")
			link, _ := resp.Location()
			fmt.Printf("link ==> %q ---- Expires in %q days", link, expire)
		} else {
			fmt.Println("Please give a file. For more information check --help")
		}
		return nil
	}

	app.Run(os.Args)
}
