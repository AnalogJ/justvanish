package main

import (
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/analogj/justvanish/pkg/actions/list"
	"github.com/analogj/justvanish/pkg/actions/request"
	"github.com/analogj/justvanish/pkg/config"
	"github.com/analogj/justvanish/pkg/version"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

var goos string
var goarch string

func main() {

	configuration, err := config.Create()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}
	//we're going to load the config file manually, since we need to validate it.
	_ = configuration.ReadConfig("config.yaml") // Find and read the config file if it exists

	app := &cli.App{
		Name:     "vanish",
		Usage:    "Tell databrokers to F#@% Off. Your data is your data, they shouldn't be monetizing your personal information without your knowledge.",
		Version:  version.VERSION,
		Compiled: time.Now(),
		Authors: []cli.Author{
			cli.Author{
				Name:  "Jason Kulatunga",
				Email: "jason@thesparktree.com",
			},
		},
		Before: func(c *cli.Context) error {

			capsuleUrl := "AnalogJ/justvanish"

			versionInfo := fmt.Sprintf("%s.%s-%s\n", goos, goarch, version.VERSION)

			subtitle := capsuleUrl + utils.LeftPad2Len(versionInfo, " ", 53-len(capsuleUrl))
			fmt.Fprintf(c.App.Writer, subtitle)
			//fmt.Fprintf(c.App.Writer, fmt.Sprintf(utils.StripIndent(
			//	`
			// _   _    __   ____  ___  _   _  ____  ____
			//( )_( )  /__\ (_  _)/ __)( )_( )( ___)(_  _)
			// ) _ (  /(__)\  )( ( (__  ) _ (  )__)   )(
			//(_) (_)(__)(__)(__) \___)(_) (_)(____) (__)
			//%s
			//`), subtitle))
			return nil
		},

		Commands: []cli.Command{
			{
				Name:  "list",
				Usage: "List all known organizations that store your personal information",
				Action: func(c *cli.Context) error {

					actionLogger := logrus.WithFields(logrus.Fields{
						"type": "list",
					})

					if c.IsSet("regulation-type") {
						configuration.Set("action.regulation-type", c.String("regulation-type"))
					}
					if c.IsSet("org-type") {
						configuration.Set("action.org-type", c.String("org-type"))
					}
					if c.IsSet("org-id") {
						configuration.Set("action.org-id", c.String("org-id"))
					}

					if c.IsSet("dry-run") {
						configuration.Set("action.dry-run", c.Bool("dry-run"))
					}
					//
					if configuration.GetBool("debug") {
						logrus.SetLevel(logrus.DebugLevel)
					} else {
						logrus.SetLevel(logrus.InfoLevel)
					}

					action, err := list.New(actionLogger, configuration)
					if err != nil {
						return err
					}
					return action.Start()
				},

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "regulation-type",
						Usage: "Filter by regulation type",
					},
					&cli.StringFlag{
						Name:  "org-type",
						Usage: "Filter by organization type",
					},
					&cli.StringFlag{
						Name:  "org-id",
						Usage: "Filter by organization id",
					},
					&cli.BoolFlag{
						Name:  "dry-run",
						Usage: "Dry run mode",
					},

					&cli.BoolFlag{
						Name:  "debug",
						Usage: "Enable debug logging",
					},
				},
			},

			{
				Name:  "request",
				Usage: "Request a copy of your of your personal information",
				Action: func(c *cli.Context) error {

					actionLogger := logrus.WithFields(logrus.Fields{
						"type": "request",
					})

					configuration.Set("action.regulation-type", c.String("regulation-type"))
					if c.IsSet("org-type") {
						configuration.Set("action.org-type", c.String("org-type"))
					}
					if c.IsSet("org-id") {
						configuration.Set("action.org-id", c.String("org-id"))
					}

					if c.IsSet("dry-run") {
						configuration.Set("action.dry-run", c.Bool("dry-run"))
					}
					//
					if configuration.GetBool("debug") {
						logrus.SetLevel(logrus.DebugLevel)
					} else {
						logrus.SetLevel(logrus.InfoLevel)
					}

					action, err := request.New(actionLogger, configuration)
					if err != nil {
						return err
					}
					return action.Start()
				},

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "regulation-type",
						Usage: "Filter by regulation type",
						Value: "ccpa",
					},
					&cli.StringFlag{
						Name:  "org-type",
						Usage: "Filter by organization type",
					},
					&cli.StringFlag{
						Name:  "org-id",
						Usage: "Filter by organization id",
					},
					&cli.BoolFlag{
						Name:  "dry-run",
						Usage: "Dry run mode",
					},

					&cli.BoolFlag{
						Name:  "debug",
						Usage: "Enable debug logging",
					},
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(color.HiRedString("ERROR: %v", err))
	}
}
