package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/urfave/cli"
)

func BuildFlags(x interface{}) []cli.Flag {
	//w := os.Stdout
	r := reflect.Indirect(reflect.ValueOf(x))
	F := make([]cli.Flag, 0, 0)
	for i := 0; i < r.NumField(); i++ {
		fmt.Print(r.Type().Field(i).Name, ", ")
		fmt.Print(r.Field(i).Kind().String(), ", ")
		fmt.Println(r.Field(i).Type().String())

		Tags := GetCliTags(r.Type().Field(i).Tag.Get("cli"))

		switch r.Field(i).Kind() {
		case reflect.String:
			f := cli.StringFlag{
				Name:   Tags["name"],
				Usage:  fmt.Sprintf("(Str)\t%s", Tags["usage"]),
				Hidden: Tags["hidden"] == "true",
				EnvVar: fmt.Sprintf("CLI_%s", strings.ToUpper(Tags["name"])),
			}
			F = append(F, f)
		case reflect.Int:
			f := cli.IntFlag{
				Name:   Tags["name"],
				Usage:  fmt.Sprintf("(Int)\t%s", Tags["usage"]),
				Hidden: Tags["hidden"] == "true",
				EnvVar: fmt.Sprintf("CLI_%s", strings.ToUpper(Tags["name"])),
			}
			F = append(F, f)
		case reflect.Bool:
			f := cli.BoolFlag{
				Name:   Tags["name"],
				Usage:  fmt.Sprintf("(Bool)\t%s", Tags["usage"]),
				Hidden: Tags["hidden"] == "true",
				EnvVar: fmt.Sprintf("CLI_%s", strings.ToUpper(Tags["name"])),
			}
			F = append(F, f)
		case reflect.Slice:
			switch r.Field(i).Type() {
			case reflect.TypeOf([]string{}):
				f := cli.StringSliceFlag{
					Name:   Tags["name"],
					Usage:  fmt.Sprintf("(Str...)\t%s", Tags["usage"]),
					Hidden: Tags["hidden"] == "true",
					EnvVar: fmt.Sprintf("CLI_%s", strings.ToUpper(Tags["name"])),
				}
				F = append(F, f)
			case reflect.TypeOf([]int{}):
				f := cli.IntSliceFlag{
					Name:   Tags["name"],
					Usage:  fmt.Sprintf("(Int...)\t%s", Tags["usage"]),
					Hidden: Tags["hidden"] == "true",
					EnvVar: fmt.Sprintf("CLI_%s", strings.ToUpper(Tags["name"])),
				}
				F = append(F, f)
			}
		}
	}

	return F
}

func BuildCommand(x interface{}) []cli.Command {
	r := reflect.Indirect(reflect.ValueOf(x))
	C := make([]cli.Command, 0, 0)
	for i := 0; i < r.NumField(); i++ {
		Tags := GetCliTags(r.Type().Field(i).Tag.Get("cli"))

		switch r.Field(i).Kind() {
		case reflect.Struct, reflect.Ptr:
			c := cli.Command{
				Name:        Tags["name"],
				Usage:       fmt.Sprintf("(Str)\t%s", Tags["usage"]),
				Hidden:      Tags["hidden"] == "true",
				Description: Tags["description"],
				Flags:       BuildFlags(r.Field(i).Elem()),
			}
			C = append(C, c)
		}
	}
	return C
}

func main() {
	type Data struct {
		Name   string   `cli:"name=name,usage=name usage,hidden=false"`
		Value  int      `cli:"name=value,usage=name usage,hidden=false"`
		Names  []string `cli:"name=names,usage=name usage,hidden=false"`
		Values []int    `cli:"name=values,usage=name usage,hidden=false"`
		Flag   bool     `cli:"name=flag,usage=name usage,hidden=false"`
	}

	type Data2 struct {
		A Data
		B Data
	}

	app := cli.NewApp()
	app.Flags = BuildFlags(Data{})
	app.Commands = BuildCommand(Data2{})
	app.Run(os.Args)
}

func GetCliTags(input string) map[string]string {
	M := make(map[string]string)
	inputs := strings.Split(input, ",")
	for i := range inputs {
		mapValues := strings.Split(inputs[i], "=")
		if len(mapValues) == 2 {
			M[strings.TrimSpace(mapValues[0])] = strings.TrimSpace(mapValues[1])
		}
	}
	return M
}
