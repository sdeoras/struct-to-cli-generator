package main

import (
	"reflect"
	"fmt"
	"github.com/libopenstorage/openstorage/osdconfig"
	"os"
)

var callbacks map[string][]string

func main() {
	callbacks = make(map[string][]string)
	config := new(osdconfig.ClusterConfig)
	config.Secrets = new(osdconfig.SecretsConfig)
	config.Kvdb = new(osdconfig.KvdbConfig)
	config.Secrets.Vault = new(osdconfig.VaultConfig)
	config.Secrets.Aws = new(osdconfig.AWSConfig)
	s := `
package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
`
	fmt.Println(s)
	printFields(reflect.Indirect(reflect.ValueOf(config)), "config", "\t\t")
	fmt.Println("}")
	fmt.Println("app.Run(os.Args)")
	fmt.Println("}")
	for _, fs := range callbacks {
		for _, f := range fs {
			fmt.Println("func", f, "(c *cli.Context) {}")
		}
	}
}

func printFields(v reflect.Value, prefix, tab string) {
	x := "\t"
	fmt.Println(tab, "{")
	fmt.Println(tab, x, "Name:", "\""+prefix+"\",")
	fmt.Println(tab, x, "Usage:", "\"To be added\",")
	fmt.Println(tab, x, "Description:", "\"To be added\",")

	fmt.Println(tab, x, "Subcommands: []cli.Command{")
	for i :=0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Ptr:
			if !field.IsNil() {
				printFields(field.Elem(), v.Type().Field(i).Name, tab+x+x)
			}
		default:
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}

	if prefix == "config" {
		config := new(osdconfig.NodeConfig)
		config.Network = new(osdconfig.NetworkConfig)
		config.Storage = new(osdconfig.StorageConfig)
		printFields(reflect.Indirect(reflect.ValueOf(config)), "node", tab+x+x)
	}

	fmt.Println(tab, x, x, x, "{")
	fmt.Println(tab, x, x, x, x, "Name:", "\""+"get"+"\",")
	fmt.Println(tab, x, x, x, x, "Usage:", "\"Get values\",")
	fmt.Println(tab, x, x, x, x, "Description:", "\"Get values\",")
	fmt.Println(tab, x, x, x, x, "Action:", "get_"+prefix+"_values,")
	callbacks[prefix] = append(callbacks[prefix], "get_"+prefix+"_values")
	fmt.Println(tab, x, x, x, x, "Flags: []cli.Flag{")
	for i :=0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Slice:
			switch field.Type(){
			case reflect.TypeOf([]string{}):
				fmt.Println(tab, x, x, x, x, x, "cli.BoolFlag{")
				fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
				fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Bool)\\tTo be added\",")
				fmt.Println(tab, x, x, x, x, x, "},")
			case reflect.TypeOf([]int{}):
				fmt.Println(tab, x, x, x, x, x, "cli.BoolFlag{")
				fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
				fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Bool)\\tTo be added\",")
				fmt.Println(tab, x, x, x, x, x, "},")
			default:
				fmt.Fprintln(os.Stderr, "ignoring", prefix+v.Type().Field(i).Name, "of type", v.Field(i).Type().String())
			}
		case reflect.String:
			fmt.Println(tab, x, x, x, x, x, "cli.BoolFlag{")
			fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
			fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Bool)\\tTo be added\",")
			fmt.Println(tab, x, x, x, x, x, "},")
		case reflect.Int:
			fmt.Println(tab, x, x, x, x, x, "cli.BoolFlag{")
			fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
			fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Bool)\\tTo be added\",")
			fmt.Println(tab, x, x, x, x, x, "},")
		case reflect.Bool:
			fmt.Println(tab, x, x, x, x, x, "cli.BoolFlag{")
			fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
			fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Bool)\\tTo be added\",")
			fmt.Println(tab, x, x, x, x, x, "},")
		default:
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}
	fmt.Println(tab, x, x, x, x, "},")
	fmt.Println(tab, x, x, "},")

	fmt.Println(tab, x, x, x, "{")
	fmt.Println(tab, x, x, x, x, "Name:", "\""+"set"+"\",")
	fmt.Println(tab, x, x, x, x, "Usage:", "\"Set values\",")
	fmt.Println(tab, x, x, x, x, "Description:", "\"Set values\",")
	fmt.Println(tab, x, x, x, x, "Action:", "set_"+prefix+"_values,")
	callbacks[prefix] = append(callbacks[prefix],"set_"+prefix+"_values")
	fmt.Println(tab, x, x, x, x, "Flags: []cli.Flag{")
	for i :=0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Slice:
			switch field.Type(){
			case reflect.TypeOf([]string{}):
				fmt.Println(tab, x, x, x, x, x, "cli.StringSliceFlag{")
				fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
				fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Str...)\\tTo be added\",")
				fmt.Println(tab, x, x, x, x, x, "},")
			case reflect.TypeOf([]int{}):
				fmt.Println(tab, x, x, x, x, x, "cli.IntSliceFlag{")
				fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
				fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Int...)\\tTo be added\",")
				fmt.Println(tab, x, x, x, x, x, "},")
			default:
				fmt.Fprintln(os.Stderr, "ignoring", prefix+v.Type().Field(i).Name, "of type", v.Field(i).Type().String())
			}
		case reflect.String:
			fmt.Println(tab, x, x, x, x, x, "cli.StringFlag{")
			fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
			fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Str)\\tTo be added\",")
			fmt.Println(tab, x, x, x, x, x, "},")
		case reflect.Int:
			fmt.Println(tab, x, x, x, x, x, "cli.IntFlag{")
			fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
			fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Int)\\tTo be added\",")
			fmt.Println(tab, x, x, x, x, x, "},")
		case reflect.Bool:
			fmt.Println(tab, x, x, x, x, x, "cli.BoolFlag{")
			fmt.Println(tab, x, x, x, x, x, x, "Name:", "\""+v.Type().Field(i).Name+"\",")
			fmt.Println(tab, x, x, x, x, x, x, "Usage: \"(Bool)\\tTo be added\",")
			fmt.Println(tab, x, x, x, x, x, "},")
		default:
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}
	fmt.Println(tab, x, x, x, x, "},")
	fmt.Println(tab, x, x, "},")

	fmt.Println(tab, x, "},")

	if prefix == "node" {
		fmt.Println(tab, x, "Flags: []cli.Flag{")
		fmt.Println(tab, x, x, "cli.StringSliceFlag{")
		fmt.Println(tab, x, x, x, "Name:", "\""+"id"+"\",")
		fmt.Println(tab, x, x, x, "Usage: \"(Str...)\\tNode id\",")
		fmt.Println(tab, x, x, "},")
		fmt.Println(tab, x, "},")
	}

	fmt.Println(tab, "},")
}
