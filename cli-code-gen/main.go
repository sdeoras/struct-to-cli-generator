package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/libopenstorage/openstorage/osdconfig"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

var prefix2VarName map[string]string
var cw *bufio.Writer
var fw *bufio.Writer
var header = `package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/libopenstorage/openstorage/osdconfig"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

var config *osdconfig.ClusterConfig
var nodeConfig *osdconfig.NodeConfig

func main() {
	config = new(osdconfig.ClusterConfig)
	nodeConfig = new(osdconfig.NodeConfig)

	if b, err := ioutil.ReadFile("/tmp/config.json"); err != nil {
		logrus.Warn(err)
	} else {
		if err := json.Unmarshal(b, config); err != nil {
			logrus.Warn(err)
		}
	}

	if b, err := ioutil.ReadFile("/tmp/nodeConfig.json"); err != nil {
		logrus.Warn(err)
	} else {
		if err := json.Unmarshal(b, nodeConfig); err != nil {
			logrus.Warn(err)
		}
	}

	app := cli.NewApp()
	app.Commands = []cli.Command{`

func main() {
	prefix2VarName = make(map[string]string)
	config := new(osdconfig.ClusterConfig)
	config.Secrets = new(osdconfig.SecretsConfig)
	config.Kvdb = new(osdconfig.KvdbConfig)
	config.Secrets.Vault = new(osdconfig.VaultConfig)
	config.Secrets.Aws = new(osdconfig.AWSConfig)

	var cb bytes.Buffer
	cw = bufio.NewWriter(&cb)

	var fb bytes.Buffer
	fw = bufio.NewWriter(&fb)

	fmt.Fprintln(cw, header)
	prefix2VarName["config"] = "config"
	printFields(reflect.Indirect(reflect.ValueOf(config)), false, "config", "usage", "description", "\t\t")
	fmt.Fprintln(cw, tabs("\t", 0), "}")
	fmt.Fprintln(cw, tabs("\t", 0), "app.Run(os.Args)")
	fmt.Fprintln(cw, "}")

	cw.Flush()
	fw.Flush()
	if err := ioutil.WriteFile("cli.go", append(cb.Bytes(), fb.Bytes()...), 0666); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote cli.go")
}

func printFields(v reflect.Value, hidden bool, prefix, usage, description, tab string) {
	fmt.Fprintln(cw, tabs(tab, 0), "{")
	fmt.Fprintln(cw, tabs(tab, 1), "Name:", "\""+prefix+"\",")
	fmt.Fprintln(cw, tabs(tab, 1), "Usage:", "\""+usage+"\",")
	fmt.Fprintln(cw, tabs(tab, 1), "Description:", "\""+description+"\",")
	fmt.Fprintln(cw, tabs(tab, 1), "Hidden:", hidden, ",")
	fmt.Fprintln(cw, tabs(tab, 1), "Action:", getCamelCase("set_"+prefix+"_values,"))

	fmt.Fprintln(fw, "func", getCamelCase("set_"+prefix+"_values(c *cli.Context) {"))
	fmt.Fprintln(cw, tabs(tab, 1), "Flags: []cli.Flag{")
	if prefix == "node" {
		fmt.Fprintln(cw, tabs(tab, 2), "cli.StringSliceFlag{")
		fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+"id"+"\",")
		fmt.Fprintln(cw, tabs(tab, 3), "Usage: \"(Str...)\\tNode id\",")
		fmt.Fprintln(cw, tabs(tab, 2), "},")
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Slice:
			switch field.Type() {
			case reflect.TypeOf([]string{}):
				fmt.Fprintln(cw, tabs(tab, 2), "cli.StringSliceFlag{")
				fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+getTag(v.Type(), i)+"\",")

				fmt.Fprintln(fw, "\t", "if c.IsSet(\""+getTag(v.Type(), i)+"\") {")
				fmt.Fprintln(fw, "\t\t", getCamelCase(getTag(v.Type(), i)), ":=", "c.StringSlice(\""+getTag(v.Type(), i)+"\")")
				fmt.Fprintln(fw, "\t\t", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, " = ", getCamelCase(getTag(v.Type(), i)))
				fmt.Fprintln(fw, "\t", "}")

				fmt.Fprintln(cw, tabs(tab, 3), "Usage: \"(Str...)\\t"+getUsage(v.Type(), i)+"\",")
				if isHidden(v.Type(), i) {
					fmt.Fprintln(cw, tabs(tab, 3), "Hidden: true,")
				} else {
					fmt.Fprintln(cw, tabs(tab, 3), "Hidden: false,")
				}
				fmt.Fprintln(cw, tabs(tab, 2), "},")
			case reflect.TypeOf([]int{}):
				fmt.Fprintln(cw, tabs(tab, 2), "cli.IntSliceFlag{")
				fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+getTag(v.Type(), i)+"\",")

				fmt.Fprintln(fw, "\t", "if c.IsSet(\""+getTag(v.Type(), i)+"\") {")
				fmt.Fprintln(fw, "\t\t", getCamelCase(getTag(v.Type(), i)), ":=", "c.IntSlice(\""+getTag(v.Type(), i)+"\")")
				fmt.Fprintln(fw, "\t\t", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, " = ", getCamelCase(getTag(v.Type(), i)))
				fmt.Fprintln(fw, "\t", "}")

				fmt.Fprintln(cw, tabs(tab, 3), "Usage: \"(Int...)\\t"+getUsage(v.Type(), i)+"\",")
				if isHidden(v.Type(), i) {
					fmt.Fprintln(cw, tabs(tab, 3), "Hidden: true,")
				} else {
					fmt.Fprintln(cw, tabs(tab, 3), "Hidden: false,")
				}
				fmt.Fprintln(cw, tabs(tab, 2), "},")
			default:
				fmt.Fprintln(os.Stderr, "ignoring", prefix+v.Type().Field(i).Name, "of type", v.Field(i).Type().String())
			}
		case reflect.String:
			fmt.Fprintln(cw, tabs(tab, 2), "cli.StringFlag{")
			fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+getTag(v.Type(), i)+"\",")

			fmt.Fprintln(fw, "\t", "if c.IsSet(\""+getTag(v.Type(), i)+"\") {")
			fmt.Fprintln(fw, "\t\t", getCamelCase(getTag(v.Type(), i)), ":=", "c.String(\""+getTag(v.Type(), i)+"\")")
			fmt.Fprintln(fw, "\t\t", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, " = ", getCamelCase(getTag(v.Type(), i)))
			fmt.Fprintln(fw, "\t", "}")

			fmt.Fprintln(cw, tabs(tab, 3), "Usage: \"(Str)\\t"+getUsage(v.Type(), i)+"\",")
			if isHidden(v.Type(), i) {
				fmt.Fprintln(cw, tabs(tab, 3), "Hidden: true,")
			} else {
				fmt.Fprintln(cw, tabs(tab, 3), "Hidden: false,")
			}
			fmt.Fprintln(cw, tabs(tab, 2), "},")
		case reflect.Int:
			fmt.Fprintln(cw, tabs(tab, 2), "cli.IntFlag{")
			fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+getTag(v.Type(), i)+"\",")

			fmt.Fprintln(fw, "\t", "if c.IsSet(\""+getTag(v.Type(), i)+"\") {")
			fmt.Fprintln(fw, "\t\t", getCamelCase(getTag(v.Type(), i)), ":=", "c.Int(\""+getTag(v.Type(), i)+"\")")
			fmt.Fprintln(fw, "\t\t", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, " = ", getCamelCase(getTag(v.Type(), i)))
			fmt.Fprintln(fw, "\t", "}")

			fmt.Fprintln(cw, tabs(tab, 3), "Usage: \"(Int)\\t"+getUsage(v.Type(), i)+"\",")
			if isHidden(v.Type(), i) {
				fmt.Fprintln(cw, tabs(tab, 3), "Hidden: true,")
			} else {
				fmt.Fprintln(cw, tabs(tab, 3), "Hidden: false,")
			}
			fmt.Fprintln(cw, tabs(tab, 2), "},")
		case reflect.Bool:
			fmt.Fprintln(cw, tabs(tab, 2), "cli.BoolFlag{")
			fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+getTag(v.Type(), i)+"\",")

			fmt.Fprintln(fw, "\t", "if c.IsSet(\""+getTag(v.Type(), i)+"\") {")
			fmt.Fprintln(fw, "\t\t", getCamelCase(getTag(v.Type(), i)), ":=", "c.Bool(\""+getTag(v.Type(), i)+"\")")
			fmt.Fprintln(fw, "\t\t", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, " = ", getCamelCase(getTag(v.Type(), i)))
			fmt.Fprintln(fw, "\t", "}")

			fmt.Fprintln(cw, tabs(tab, 3), "Usage: \"(Bool)\\t"+getUsage(v.Type(), i)+"\",")
			if isHidden(v.Type(), i) {
				fmt.Fprintln(cw, tabs(tab, 3), "Hidden: true,")
			} else {
				fmt.Fprintln(cw, tabs(tab, 3), "Hidden: false,")
			}
			fmt.Fprintln(cw, tabs(tab, 2), "},")
		default:
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}

	fmt.Fprintln(fw, "\t", "jb, _ := json.Marshal(config); ioutil.WriteFile(\"/tmp/config.json\", jb, 0666)")
	fmt.Fprintln(fw, "\t", "jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile(\"/tmp/nodeConfig.json\", jb, 0666)")
	fmt.Fprintln(fw, "}")
	fmt.Fprintln(fw)
	fmt.Fprintln(cw, tabs(tab, 1), "},")

	fmt.Fprintln(cw, tabs(tab, 1), "Subcommands: []cli.Command{")
	fmt.Fprintln(cw, tabs(tab, 2), "{")
	fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+"show"+"\",")
	fmt.Fprintln(cw, tabs(tab, 3), "Usage:", "\"Show values\",")
	fmt.Fprintln(cw, tabs(tab, 3), "Description:", "\"Show values\",")
	fmt.Fprintln(cw, tabs(tab, 3), "Action:", getCamelCase("show_"+prefix+"_values,"))

	fmt.Fprintln(fw, "func", getCamelCase("show_"+prefix+"_values(c *cli.Context) {"))
	fmt.Fprintln(cw, tabs(tab, 3), "Flags: []cli.Flag{")
	fmt.Fprintln(cw, tabs(tab, 4), "cli.BoolFlag{")
	fmt.Fprintln(cw, tabs(tab, 5), "Name:", "\"all, a\",")
	fmt.Fprintln(cw, tabs(tab, 5), "Usage:", "\"(Bool)\\tShow all data\",")
	fmt.Fprintln(cw, tabs(tab, 5), "Hidden:", "false,")
	fmt.Fprintln(cw, tabs(tab, 4), "},")
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Slice:
			switch field.Type() {
			case reflect.TypeOf([]string{}):
				fmt.Fprintln(cw, tabs(tab, 4), "cli.BoolFlag{")
				fmt.Fprintln(cw, tabs(tab, 5), "Name:", "\""+getTag(v.Type(), i)+"\",")

				fmt.Fprintln(fw, "\t", "if c.IsSet(\"all\") || c.IsSet(\""+getTag(v.Type(), i)+"\") {")
				fmt.Fprintln(fw, "\t\t", "fmt.Println(", "\""+v.Type().Field(i).Name+":\",", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, ")")
				fmt.Fprintln(fw, "\t", "}")

				fmt.Fprintln(cw, tabs(tab, 5), "Usage: \"(Bool)\\t"+getUsage(v.Type(), i)+"\",")
				if isHidden(v.Type(), i) {
					fmt.Fprintln(cw, tabs(tab, 5), "Hidden: true,")
				} else {
					fmt.Fprintln(cw, tabs(tab, 5), "Hidden: false,")
				}
				fmt.Fprintln(cw, tabs(tab, 4), "},")
			case reflect.TypeOf([]int{}):
				fmt.Fprintln(cw, tabs(tab, 4), "cli.BoolFlag{")
				fmt.Fprintln(cw, tabs(tab, 5), "Name:", "\""+getTag(v.Type(), i)+"\",")

				fmt.Fprintln(fw, "\t", "if c.IsSet(\"all\") || c.IsSet(\""+getTag(v.Type(), i)+"\") {")
				fmt.Fprintln(fw, "\t\t", "fmt.Println(", "\""+v.Type().Field(i).Name+":\",", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, ")")
				fmt.Fprintln(fw, "\t", "}")

				fmt.Fprintln(cw, tabs(tab, 5), "Usage: \"(Bool)\\t"+getUsage(v.Type(), i)+"\",")
				if isHidden(v.Type(), i) {
					fmt.Fprintln(cw, tabs(tab, 5), "Hidden: true,")
				} else {
					fmt.Fprintln(cw, tabs(tab, 5), "Hidden: false,")
				}
				fmt.Fprintln(cw, tabs(tab, 4), "},")
			default:
				fmt.Fprintln(os.Stderr, "ignoring", prefix+v.Type().Field(i).Name, "of type", v.Field(i).Type().String())
			}
		case reflect.String:
			fmt.Fprintln(cw, tabs(tab, 4), "cli.BoolFlag{")
			fmt.Fprintln(cw, tabs(tab, 5), "Name:", "\""+getTag(v.Type(), i)+"\",")

			fmt.Fprintln(fw, "\t", "if c.IsSet(\"all\") || c.IsSet(\""+getTag(v.Type(), i)+"\") {")
			fmt.Fprintln(fw, "\t\t", "fmt.Println(", "\""+v.Type().Field(i).Name+":\",", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, ")")
			fmt.Fprintln(fw, "\t", "}")

			fmt.Fprintln(cw, tabs(tab, 5), "Usage: \"(Bool)\\t"+getUsage(v.Type(), i)+"\",")
			if isHidden(v.Type(), i) {
				fmt.Fprintln(cw, tabs(tab, 5), "Hidden: true,")
			} else {
				fmt.Fprintln(cw, tabs(tab, 5), "Hidden: false,")
			}
			fmt.Fprintln(cw, tabs(tab, 4), "},")
		case reflect.Int:
			fmt.Fprintln(cw, tabs(tab, 4), "cli.BoolFlag{")
			fmt.Fprintln(cw, tabs(tab, 5), "Name:", "\""+getTag(v.Type(), i)+"\",")

			fmt.Fprintln(fw, "\t", "if c.IsSet(\"all\") || c.IsSet(\""+getTag(v.Type(), i)+"\") {")
			fmt.Fprintln(fw, "\t\t", "fmt.Println(", "\""+v.Type().Field(i).Name+":\",", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, ")")
			fmt.Fprintln(fw, "\t", "}")

			fmt.Fprintln(cw, tabs(tab, 5), "Usage: \"(Bool)\\t"+getUsage(v.Type(), i)+"\",")
			if isHidden(v.Type(), i) {
				fmt.Fprintln(cw, tabs(tab, 5), "Hidden: true,")
			} else {
				fmt.Fprintln(cw, tabs(tab, 5), "Hidden: false,")
			}
			fmt.Fprintln(cw, tabs(tab, 4), "},")
		case reflect.Bool:
			fmt.Fprintln(cw, tabs(tab, 4), "cli.BoolFlag{")
			fmt.Fprintln(cw, tabs(tab, 5), "Name:", "\""+getTag(v.Type(), i)+"\",")

			fmt.Fprintln(fw, "\t", "if c.IsSet(\"all\") || c.IsSet(\""+getTag(v.Type(), i)+"\") {")
			fmt.Fprintln(fw, "\t\t", "fmt.Println(", "\""+v.Type().Field(i).Name+":\",", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, ")")
			fmt.Fprintln(fw, "\t", "}")

			fmt.Fprintln(cw, tabs(tab, 5), "Usage: \"(Bool)\\t"+getUsage(v.Type(), i)+"\",")
			if isHidden(v.Type(), i) {
				fmt.Fprintln(cw, tabs(tab, 5), "Hidden: true,")
			} else {
				fmt.Fprintln(cw, tabs(tab, 5), "Hidden: false,")
			}
			fmt.Fprintln(cw, tabs(tab, 4), "},")
		default:
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}
	fmt.Fprintln(fw, "}")
	fmt.Fprintln(fw)
	fmt.Fprintln(cw, tabs(tab, 3), "},")
	fmt.Fprintln(cw, tabs(tab, 2), "},")

	// subcommand for node
	if prefix == "config" {
		config := new(osdconfig.NodeConfig)
		config.Network = new(osdconfig.NetworkConfig)
		config.Storage = new(osdconfig.StorageConfig)
		prefix2VarName["nodeConfig"] = "nodeConfig"
		printFields(reflect.Indirect(reflect.ValueOf(config)), false, "nodeConfig", "node usage", "node description", tabs(tab, 2))
	}

	// subcommand based on nested struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Ptr:
			if !field.IsNil() {
				prefix2VarName[getTag(v.Type(), i)] = prefix2VarName[prefix] + "." + v.Type().Field(i).Name
				printFields(field.Elem(), isHidden(v.Type(), i), getTag(v.Type(), i), getUsage(v.Type(), i), getDescription(v.Type(), i), tabs(tab, 2))
			}
		default:
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}

	fmt.Fprintln(cw, tabs(tab, 1), "},")
	fmt.Fprintln(cw, tabs(tab, 0), "},")
}

func name2tag(name string) string {
	name2 := make([]byte, 0, 0)
	for i := 0; i < len(name); i++ {
		if i > 0 && name[i] >= 'A' && name[i] <= 'Z' && name[i-1] >= 'a' && name[i-1] <= 'z' {
			name2 = append(name2, '-')
		}
		name2 = append(name2, name[i])
	}

	return strings.ToLower(string(name2))
}

func tabs(base string, n int) string {
	t := base
	for i := 0; i < n; i++ {
		t += "\t"
	}
	return t
}

func parseJsonTags(input string) string {
	s := strings.Split(input, ",")
	if len(s) > 0 {
		return s[0]
	}
	return ""
}

func isHidden(t reflect.Type, n int) bool {
	if n >= t.NumField() {
		panic("index out of bounds")
	}

	tag := t.Field(n).Tag.Get("hidden")
	if tag == "true" {
		return true
	}
	if tag == "false" {
		return false
	}
	return false
}

func getUsage(t reflect.Type, n int) string {
	if n > t.NumField() {
		panic("index out of bounds")
	}

	tag := t.Field(n).Tag.Get("usage")
	if len(tag) == 0 {
		tag = "none yet"
	}
	return tag
}

func getDescription(t reflect.Type, n int) string {
	if n > t.NumField() {
		panic("index out of bounds")
	}

	tag := t.Field(n).Tag.Get("description")
	if len(tag) == 0 {
		tag = "none yet"
	}
	return tag
}

func getTag(t reflect.Type, n int) string {
	if n >= t.NumField() {
		panic("index out of bounds")
	}

	tag := parseJsonTags(t.Field(n).Tag.Get("json"))
	if len(tag) == 0 {
		tag = name2tag(t.Field(n).Name)
	}

	return tag
}

func getCamelCase(input string) string {
	upper := []byte(strings.ToUpper(input))
	out := make([]byte, 0, 0)
	for i :=0; i < len(input); i++ {
		if i > 0 && input[i] == '_' {
			continue
		}
		if i > 0 && input[i-1] == '_' {
			out = append(out, upper[i])
		} else {
			out = append(out, input[i])
		}
	}
	return string(out)
}
