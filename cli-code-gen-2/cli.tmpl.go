package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/sdeoras/configio"
	"github.com/sdeoras/configio/configfile"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"path/filepath"
)

var manager configio.ConfigManager
var configManager configio.ConfigManager

type configData struct {
	File string `json:"file"`
}

func (c *configData) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *configData) Unmarshal(b []byte) error {
	return json.Unmarshal(b, c)
}

func main() {
    logrus.SetLevel(logrus.ErrorLevel)

	appName := "cli"
	d := new(configData)

	var err error
	configManager, err = configfile.NewManager(context.Background(), "file",
		filepath.Join(os.Getenv("HOME"), ".config", appName, "config.json"))
	if err != nil {
		logrus.Fatal(err)
	}

	if err := configManager.Unmarshal(d); err == nil {
		manager, err = configfile.NewManager(context.Background(), "file", d.File)
		if err != nil {
			logrus.Fatal(err)
		}
	} else {
		manager, err = configfile.NewManager(context.Background())
		if err != nil {
			logrus.Fatal(err)
		}
	}

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "json, j",
			Usage: "Print JSON output",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "init",
			Usage: "initialize config backend",
			Description: "initialize config backend",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "file",
					Usage: "(Str)\tpath of config file",
				},
			},
			Action: initConfigValues,
		},
		{{.Commands}}
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func initConfigValues(c *cli.Context) error {
	config := new({{.Dtype}}).Init()
	if c.IsSet("file") {
		d := new(configData)
		d.File = c.String("file")
		if err := configManager.Marshal(d); err != nil {
			return err
		}

		if m, err := configfile.NewManager(context.Background(), "file", c.String("file")); err != nil {
			return err
		} else {
			manager = m
		}
	} else {
		if m, err := configfile.NewManager(context.Background()); err != nil {
			return err
		} else {
			manager = m
		}
	}

	if manager == nil {
		return fmt.Errorf("config manager could not be instantiated")
	}

	return manager.Marshal(config)
}

{{.Functions}}