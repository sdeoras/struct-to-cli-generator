package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/sdeoras/configio"
	"github.com/sdeoras/configio/configfile"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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
	logrus.SetLevel(logrus.FatalLevel)
	log := logrus.WithField("func", "main").WithField("manager", "cli")

	appName := "cli"
	d := new(configData)

	var err error
	configManager, err = configfile.NewManager(context.Background(), "file",
		filepath.Join(os.Getenv("HOME"), ".config", appName, "config.json"))
	if err != nil {
		log.Fatal(err)
	}

	if err := configManager.Unmarshal(d); err == nil {
		if len(d.File) > 0 {
			manager, err = configfile.NewManager(context.Background(), "file", d.File)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			manager, err = configfile.NewManager(context.Background())
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		if err := configManager.Marshal(d); err != nil {
			log.Fatal(err)
		}
		manager, err = configfile.NewManager(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

	logrus.SetLevel(logrus.ErrorLevel)
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "json, j",
			Usage: "Print JSON output",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:        "init",
			Usage:       "initialize config backend",
			Description: "initialize config backend",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "(Str)\tpath of config file",
				},
			},
			Action: initConfigValues,
		},
		{
			Name:        "watch",
			Usage:       "watch for config changes",
			Description: "watch for config changes",
			Action:      watchConfigValues,
		},
		// header
		{
			Name:        "config",
			Usage:       "Config params",
			Description: "Manage config parameters",
			Hidden:      false,
			Action:      setConfigValues,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "my-name",
					Usage:  "(Str)\tmy name",
					Hidden: false,
				},
				cli.IntFlag{
					Name:   "my-value",
					Usage:  "(Int)\tmy value",
					Hidden: false,
				},
			},
			Subcommands: []cli.Command{
				{
					Name:        "show",
					Usage:       "Show values",
					Description: "Show values",
					Action:      showConfigValues,
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:   "all, a",
							Usage:  "(Bool)\tShow all data",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "my-name",
							Usage:  "(Bool)\tmy name",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "my-value",
							Usage:  "(Bool)\tmy value",
							Hidden: false,
						},
					},
				},
				{
					Name:        "a",
					Usage:       "command a",
					Description: "command a",
					Hidden:      false,
					Action:      setAValues,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "name",
							Usage:  "(Str)\tname of data",
							Hidden: false,
						},
						cli.IntFlag{
							Name:   "value",
							Usage:  "(Int)\tvalue of data",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "present",
							Usage:  "(Bool)\tpresent in data",
							Hidden: false,
						},
					},
					Subcommands: []cli.Command{
						{
							Name:        "show",
							Usage:       "Show values",
							Description: "Show values",
							Action:      showAValues,
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name:   "all, a",
									Usage:  "(Bool)\tShow all data",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "name",
									Usage:  "(Bool)\tname of data",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "value",
									Usage:  "(Bool)\tvalue of data",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "present",
									Usage:  "(Bool)\tpresent in data",
									Hidden: false,
								},
							},
						},
					},
				},
				{
					Name:        "b",
					Usage:       "command b",
					Description: "command b",
					Hidden:      false,
					Action:      setBValues,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "name",
							Usage:  "(Str)\tname of data",
							Hidden: false,
						},
						cli.IntFlag{
							Name:   "value",
							Usage:  "(Int)\tvalue of data",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "present",
							Usage:  "(Bool)\tpresent in data",
							Hidden: false,
						},
					},
					Subcommands: []cli.Command{
						{
							Name:        "show",
							Usage:       "Show values",
							Description: "Show values",
							Action:      showBValues,
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name:   "all, a",
									Usage:  "(Bool)\tShow all data",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "name",
									Usage:  "(Bool)\tname of data",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "value",
									Usage:  "(Bool)\tvalue of data",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "present",
									Usage:  "(Bool)\tpresent in data",
									Hidden: false,
								},
							},
						},
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func watchConfigValues(c *cli.Context) error {
	config := new(Data).Init()
	f := func(ctx context.Context, data interface{}, err error) <-chan error {
		done := make(chan error)
		if err := manager.Unmarshal(config); err != nil {
			done <- err
			return done
		}

		if jb, err := config.Marshal(); err != nil {
			done <- err
			return done
		} else {
			fmt.Println(string(jb))
		}

		done <- nil
		return done
	}

	trigger := manager.Watch("cliWatch", nil, f)
	cc := make(chan os.Signal, 2)
	signal.Notify(cc, os.Interrupt, syscall.SIGTERM)

	logrus.SetLevel(logrus.InfoLevel)
	for {
		select {
		case <-trigger:
		case <-cc:
			logrus.SetLevel(logrus.ErrorLevel)
			return nil
		}
	}
}

func initConfigValues(c *cli.Context) error {
	config := new(Data).Init()
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

func setConfigValues(c *cli.Context) error {
	log := logrus.WithField("func", "setConfigValues").WithField("manager", "cli")
	config := new(Data).Init()
	if err := manager.Unmarshal(config); err != nil {
		log.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}

	if c.IsSet("my-name") {
		config.MyName = c.String("my-name")
	}
	if c.IsSet("my-value") {
		config.MyValue = c.Int("my-value")
	}
	if err := manager.Marshal(config); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func showConfigValues(c *cli.Context) error {
	log := logrus.WithField("func", "showConfigValues").WithField("manager", "cli")
	config := new(Data).Init()
	if err := manager.Unmarshal(config); err != nil {
		log.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config)
	}
	if c.IsSet("all") || c.IsSet("my-name") {
		fmt.Println("my-name:", config.MyName)
	}
	if c.IsSet("all") || c.IsSet("my-value") {
		fmt.Println("my-value:", config.MyValue)
	}
	return nil
}

func setAValues(c *cli.Context) error {
	log := logrus.WithField("func", "setAValues").WithField("manager", "cli")
	config := new(Data).Init()
	if err := manager.Unmarshal(config); err != nil {
		log.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}
	if config.A == nil {
		err := errors.New("config.A" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}

	if c.IsSet("name") {
		config.A.Name = c.String("name")
	}
	if c.IsSet("value") {
		config.A.Value = c.Int("value")
	}
	if c.IsSet("present") {
		config.A.Present = c.Bool("present")
	}
	if err := manager.Marshal(config); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func showAValues(c *cli.Context) error {
	log := logrus.WithField("func", "showAValues").WithField("manager", "cli")
	config := new(Data).Init()
	if err := manager.Unmarshal(config); err != nil {
		log.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}
	if config.A == nil {
		err := errors.New("config.A" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config.A)
	}
	if c.IsSet("all") || c.IsSet("name") {
		fmt.Println("name:", config.A.Name)
	}
	if c.IsSet("all") || c.IsSet("value") {
		fmt.Println("value:", config.A.Value)
	}
	if c.IsSet("all") || c.IsSet("present") {
		fmt.Println("present:", config.A.Present)
	}
	return nil
}

func setBValues(c *cli.Context) error {
	log := logrus.WithField("func", "setBValues").WithField("manager", "cli")
	config := new(Data).Init()
	if err := manager.Unmarshal(config); err != nil {
		log.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}
	if config.B == nil {
		err := errors.New("config.B" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}

	if c.IsSet("name") {
		config.B.Name = c.String("name")
	}
	if c.IsSet("value") {
		config.B.Value = c.Int("value")
	}
	if c.IsSet("present") {
		config.B.Present = c.Bool("present")
	}
	if err := manager.Marshal(config); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func showBValues(c *cli.Context) error {
	log := logrus.WithField("func", "showBValues").WithField("manager", "cli")
	config := new(Data).Init()
	if err := manager.Unmarshal(config); err != nil {
		log.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}
	if config.B == nil {
		err := errors.New("config.B" + ": no data found, received nil pointer")
		log.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config.B)
	}
	if c.IsSet("all") || c.IsSet("name") {
		fmt.Println("name:", config.B.Name)
	}
	if c.IsSet("all") || c.IsSet("value") {
		fmt.Println("value:", config.B.Value)
	}
	if c.IsSet("all") || c.IsSet("present") {
		fmt.Println("present:", config.B.Present)
	}
	return nil
}

func printJson(obj interface{}) error {
	if b, err := json.MarshalIndent(obj, "", "  "); err != nil {
		return err
	} else {
		fmt.Println(string(b))
		return nil
	}
}
