package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
	"time"
)

var prefix2VarName map[string]string
var prefixOrigin map[string]string
var prefixDepth map[string]int
var cw *bufio.Writer
var fw *bufio.Writer
var dtype = "Data"

func main() {
	prefix2VarName = make(map[string]string)
	prefixOrigin = make(map[string]string)
	prefixDepth = make(map[string]int)
	config := new(Data).Init()

	var cb bytes.Buffer
	cw = bufio.NewWriter(&cb)

	var fb bytes.Buffer
	fw = bufio.NewWriter(&fb)

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	fmt.Fprintln(cw, "// header")
	prefix2VarName["config"] = "config"
	prefixOrigin["config"] = "config"
	prefixDepth["config"] = 0
	printFields(reflect.Indirect(reflect.ValueOf(config)), false, "config", "Config params",
		"Manage config parameters", "\t\t")

	fmt.Fprintln(fw, "func printJson(obj interface{}) error {")
	fmt.Fprintln(fw, "\t", "if b, err := json.MarshalIndent(obj, \"\", \"  \"); err != nil {")
	fmt.Fprintln(fw, "\t\t", "return err")
	fmt.Fprintln(fw, "\t", "} else {")
	fmt.Fprintln(fw, "\t\t", "fmt.Println(string(b))")
	fmt.Fprintln(fw, "\t\t", "return nil")
	fmt.Fprintln(fw, "\t", "}")
	fmt.Fprintln(fw, "}")
	fmt.Fprintln(fw)

	cw.Flush()
	fw.Flush()

	tmpl, err := template.ParseFiles("cli.tmpl.go")
	if err != nil {
		log.Fatal(err)
	}

	type Data struct {
		Commands  string
		Functions string
		Dtype     string
	}

	d := new(Data)
	d.Commands = string(cb.Bytes())
	d.Functions = string(fb.Bytes())
	d.Dtype = dtype

	if err := tmpl.Execute(w, d); err != nil {
		log.Fatal(err)
	}

	w.Flush()
	if err := ioutil.WriteFile("cli.go", b.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote cli.go")
}

// define a function for getter field behavior
func getFunc(cw, fw *bufio.Writer, v reflect.Value, i int, tab, prefix string) {
	if !isEnabled(v.Type(), i) {
		return
	}
	fmt.Fprintln(cw, tabs(tab, 4), "cli.BoolFlag{")
	fmt.Fprintln(cw, tabs(tab, 5), "Name:", "\""+getTag(v.Type(), i)+"\",")

	fmt.Fprintln(fw, "\t", "if c.IsSet(\"all\") || c.IsSet(\""+getTag(v.Type(), i)+"\") {")
	fmt.Fprintln(fw, "\t\t",
		"fmt.Println(", "\""+getJsonName(v.Type(), i)+":\",", prefix2VarName[prefix]+"."+v.Type().Field(i).Name, ")")
	fmt.Fprintln(fw, "\t", "}")

	fmt.Fprintln(cw, tabs(tab, 5), "Usage: \"(Bool)\\t"+getUsage(v.Type(), i)+"\",")
	if isHidden(v.Type(), i) {
		fmt.Fprintln(cw, tabs(tab, 5), "Hidden: true,")
	} else {
		fmt.Fprintln(cw, tabs(tab, 5), "Hidden: false,")
	}
	fmt.Fprintln(cw, tabs(tab, 4), "},")
}

// define a function  for setter field behavior
func setFunc(cw, fw *bufio.Writer, v reflect.Value, i int, tab, prefix, castType, flagType, flagTag string) {
	if !isEnabled(v.Type(), i) {
		return
	}
	fmt.Fprintln(cw, tabs(tab, 2), "cli."+flagType+"Flag{")
	fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+getTag(v.Type(), i)+"\",")

	fmt.Fprintln(fw, "\t", "if c.IsSet(\""+getTag(v.Type(), i)+"\") {")
	if castType == "" {
		fmt.Fprintln(fw, "\t\t",
			prefix2VarName[prefix]+"."+v.Type().Field(i).Name, "=", "c."+flagType+"(\""+getTag(v.Type(), i)+"\")")
	} else {
		fmt.Fprintln(fw, "\t\t",
			prefix2VarName[prefix]+"."+v.Type().Field(i).Name, "=", castType+"(c."+flagType+"(\""+getTag(v.Type(), i)+"\"))")
	}
	fmt.Fprintln(fw, "\t", "}")

	fmt.Fprintln(cw, tabs(tab, 3), "Usage: \"("+flagTag+")\\t"+getUsage(v.Type(), i)+"\",")
	if isHidden(v.Type(), i) {
		fmt.Fprintln(cw, tabs(tab, 3), "Hidden: true,")
	} else {
		fmt.Fprintln(cw, tabs(tab, 3), "Hidden: false,")
	}
	fmt.Fprintln(cw, tabs(tab, 2), "},")
}

func printFields(v reflect.Value, hidden bool, prefix, usage, description, tab string) {
	fmt.Fprintln(cw, tabs(tab, 0), "{")
	fmt.Fprintln(cw, tabs(tab, 1), "Name:", "\""+prefix+"\",")
	fmt.Fprintln(cw, tabs(tab, 1), "Usage:", "\""+usage+"\",")
	fmt.Fprintln(cw, tabs(tab, 1), "Description:", "\""+description+"\",")
	fmt.Fprintln(cw, tabs(tab, 1), "Hidden:", hidden, ",")
	fmt.Fprintln(cw, tabs(tab, 1), "Action:", getCamelCase("set_"+prefix+"_values,"))

	fmt.Fprintln(fw, "func", getCamelCase("set_"+prefix+"_values(c *cli.Context) error {"))

	getConfig(fw, dtype)

	fmt.Fprintln(fw, nullChecker(prefix2VarName[prefix]+".field"))
	fmt.Fprintln(cw, tabs(tab, 1), "Flags: []cli.Flag{")
	if prefix == "node" {
		fmt.Fprintln(cw, tabs(tab, 2), "cli.BoolFlag{")
		fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+"all, a"+"\",")
		fmt.Fprintln(cw, tabs(tab, 3), "Usage: \"(Bool)\\tFor all nodes on cluster\",")
		fmt.Fprintln(cw, tabs(tab, 2), "},")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Slice:
			switch field.Type() {
			case reflect.TypeOf([]string{}):
				setFunc(cw, fw, v, i, tab, prefix, "", "StringSlice", "Str...")

			case reflect.TypeOf([]int{}):
				setFunc(cw, fw, v, i, tab, prefix, "", "IntSlice", "Int...")
			default:
				fmt.Fprintln(os.Stderr, "ignoring",
					prefix+v.Type().Field(i).Name, "of type", v.Field(i).Type().String())
			}
		case reflect.String:
			setFunc(cw, fw, v, i, tab, prefix, "", "String", "Str")
		case reflect.Int:
			setFunc(cw, fw, v, i, tab, prefix, "", "Int", "Int")
		case reflect.Uint32:
			setFunc(cw, fw, v, i, tab, prefix, "uint32", "Uint", "Uint")
		case reflect.Bool:
			setFunc(cw, fw, v, i, tab, prefix, "", "Bool", "Bool")
		case reflect.Ptr:
		default:
			fmt.Println("skipping in set: ", prefix+"."+v.Type().Field(i).Name, " of type: ", v.Field(i).Kind())
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}

	fmt.Fprintln(fw, "\t", "if err := manager.Marshal(config); err != nil {")
	fmt.Fprintln(fw, "\t\t", "logrus.Error(\"Set config for cluster\")")
	fmt.Fprintln(fw, "\t\t", "return err")
	fmt.Fprintln(fw, "\t", "}")
	fmt.Fprintln(fw, "\t", "logrus.Info(\"Set config for cluster\")")
	fmt.Fprintln(fw, "\t", "return nil")

	fmt.Fprintln(fw, "}")
	fmt.Fprintln(fw)
	fmt.Fprintln(cw, tabs(tab, 1), "},")

	fmt.Fprintln(cw, tabs(tab, 1), "Subcommands: []cli.Command{")
	fmt.Fprintln(cw, tabs(tab, 2), "{")
	fmt.Fprintln(cw, tabs(tab, 3), "Name:", "\""+"show"+"\",")
	fmt.Fprintln(cw, tabs(tab, 3), "Usage:", "\"Show values\",")
	fmt.Fprintln(cw, tabs(tab, 3), "Description:", "\"Show values\",")
	fmt.Fprintln(cw, tabs(tab, 3), "Action:", getCamelCase("show_"+prefix+"_values,"))

	fmt.Fprintln(fw, "func", getCamelCase("show_"+prefix+"_values(c *cli.Context) error {"))

	getConfig(fw, dtype)

	fmt.Fprintln(fw, nullChecker(prefix2VarName[prefix]+".field"))
	fmt.Fprintln(fw, "\t", "if c.GlobalBool(\"json\") {")
	fmt.Fprintln(fw, "\t\t", "return printJson("+prefix2VarName[prefix]+")")
	fmt.Fprintln(fw, "\t", "}")

	fmt.Fprintln(cw, tabs(tab, 3), "Flags: []cli.Flag{")
	fmt.Fprintln(cw, tabs(tab, 4), "cli.BoolFlag{")
	fmt.Fprintln(cw, tabs(tab, 5), "Name:", "\"all, a\",")
	fmt.Fprintln(cw, tabs(tab, 5), "Usage:", "\"(Bool)\\tShow all data\",")
	fmt.Fprintln(cw, tabs(tab, 5), "Hidden:", "false,")
	fmt.Fprintln(cw, tabs(tab, 4), "},")
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Struct:
			switch field.Type() {
			case reflect.TypeOf(time.Now()):
				getFunc(cw, fw, v, i, tab, prefix)
			default:
				fmt.Println("skipping in get: ", prefix+"."+v.Type().Field(i).Name, " of type: ", v.Field(i).Kind())
			}
		case reflect.Slice:
			switch field.Type() {
			case reflect.TypeOf([]string{}):
				getFunc(cw, fw, v, i, tab, prefix)
			case reflect.TypeOf([]int{}):
				getFunc(cw, fw, v, i, tab, prefix)
			default:
				fmt.Fprintln(os.Stderr, "ignoring",
					prefix+v.Type().Field(i).Name, "of type", v.Field(i).Type().String())
			}
		case reflect.String:
			getFunc(cw, fw, v, i, tab, prefix)
		case reflect.Int:
			getFunc(cw, fw, v, i, tab, prefix)
		case reflect.Uint32:
			getFunc(cw, fw, v, i, tab, prefix)
		case reflect.Bool:
			getFunc(cw, fw, v, i, tab, prefix)
		case reflect.Ptr:
		default:
			fmt.Println("skipping in get: ", prefix+"."+v.Type().Field(i).Name, " of type: ", v.Field(i).Kind())
			//M[v.Field(i).Kind()] = append(M[v.Field(i).Kind()], prefix+v.Type().Field(i).Name)
		}
	}
	fmt.Fprintln(fw, "\t", "return nil")
	fmt.Fprintln(fw, "}")
	fmt.Fprintln(fw)
	fmt.Fprintln(cw, tabs(tab, 3), "},")
	fmt.Fprintln(cw, tabs(tab, 2), "},")

	// subcommand based on nested struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch k := field.Kind(); k {
		case reflect.Ptr:
			if !field.IsNil() && isEnabled(v.Type(), i) {
				prefix2VarName[getTag(v.Type(), i)] = prefix2VarName[prefix] + "." + v.Type().Field(i).Name
				prefixOrigin[getTag(v.Type(), i)] = prefixOrigin[prefix]
				prefixDepth[getTag(v.Type(), i)] = prefixDepth[prefix] + 1
				printFields(field.Elem(),
					isHidden(v.Type(), i), getTag(v.Type(), i), getUsage(v.Type(), i),
					getDescription(v.Type(), i), tabs(tab, 2))
			}
		default:
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

func isEnabled(t reflect.Type, n int) bool {
	if n >= t.NumField() {
		panic("index out of bounds")
	}

	tag := t.Field(n).Tag.Get("enable")
	if tag == "true" {
		return true
	}
	if tag == "false" {
		return false
	}
	return true
}

func getJsonName(t reflect.Type, n int) string {
	if n > t.NumField() {
		panic("index out of bounds")
	}

	tag := t.Field(n).Tag.Get("json")
	if len(tag) == 0 {
		tag = "none yet"
	}
	return strings.Split(tag, ",")[0]
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
	for i := 0; i < len(input); i++ {
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

func nullChecker(input string) string {
	fields := strings.Split(input, ".")
	s := ""
	field := fields[0]
	for i := 0; i < len(fields)-1; i++ {
		s += fmt.Sprintln("\t",
			"if",
			field, "== nil {\n\t\terr := errors.New(\""+field+"\"+\": no data found, received nil pointer\")\n"+
				"\t\tlogrus.Error(err)\n\t\treturn err\n\t}")
		field += "."
		field += fields[i+1]
	}

	return s
}

func getConfig(fw *bufio.Writer, dtype string) {
	fmt.Fprintln(fw, "\t", "config := new("+dtype+").Init()")
	fmt.Fprintln(fw, "\t", "if err := manager.Unmarshal(config); err != nil {")
	fmt.Fprintln(fw, "\t\t", "logrus.Error(err)")
	fmt.Fprintln(fw, "\t\t", "return err")
	fmt.Fprintln(fw, "\t", "}")
}
