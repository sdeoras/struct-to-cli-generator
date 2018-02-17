package main

import (
	"github.com/codegangsta/cli"
	"os"
	"reflect"
	"fmt"
	"github.com/libopenstorage/openstorage/osdconfig"
)

func main() {
	C := make([]cli.Command, 0, 0)
	config := new(osdconfig.ClusterConfig)
	config.Secrets = new(osdconfig.SecretsConfig)
	config.Kvdb = new(osdconfig.KvdbConfig)
	printFields(reflect.Indirect(reflect.ValueOf(config)), "cluster", &C)
	app := cli.NewApp()
	app.Commands = C
	app.Run(os.Args)
}

func printFields(v reflect.Value, prefix string, C *[]cli.Command) {
	var c cli.Command
	c.Name = prefix
	c.Usage = c.Name + ": Usage to be added"
	c.Description = c.Name + ": Description to be added"
	c.Flags = make([]cli.Flag, 0, 0)
	c.Subcommands = make([]cli.Command, 0, 0)
	for i :=0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Ptr:
			if !field.IsNil() {
				subCommands := make([]cli.Command, 0, 0)
				printFields(field.Elem(), v.Type().Field(i).Name, &subCommands)
				c.Subcommands = append(c.Subcommands, subCommands...)
			}
		case reflect.Slice:
			switch field.Type(){
			case reflect.TypeOf([]string{}):
				c.Flags = append(c.Flags,
					cli.StringSliceFlag{
						Name:v.Type().Field(i).Name,
						Usage: fmt.Sprintf("(Str...)\t%s: Description to be added", v.Type().Field(i).Name),
						})
			case reflect.TypeOf([]int{}):
				c.Flags = append(c.Flags,
					cli.IntSliceFlag{
						Name:v.Type().Field(i).Name,
						Usage: fmt.Sprintf("(Int...)\t%s: Description to be added", v.Type().Field(i).Name),
						})
			default:
				fmt.Println("ignoring", prefix+v.Type().Field(i).Name, "of type", v.Field(i).Type().String())
			}
		case reflect.String:
			c.Flags = append(c.Flags,
				cli.StringFlag{
					Name:v.Type().Field(i).Name,
					Usage: fmt.Sprintf("(Str)\t%s: Description to be added", v.Type().Field(i).Name),
					})
		case reflect.Int:
			c.Flags = append(c.Flags,
				cli.IntFlag{
					Name:v.Type().Field(i).Name,
					Usage: fmt.Sprintf("(Int)\t%s: Description to be added", v.Type().Field(i).Name),
					})
		default:
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}
	*C = append(*C, c)
}