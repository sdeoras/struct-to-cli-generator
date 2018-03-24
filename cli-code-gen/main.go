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
	"time"

	"github.com/libopenstorage/openstorage/osdconfig"
)

var prefix2VarName map[string]string
var prefixOrigin map[string]string
var prefixDepth map[string]int
var cw *bufio.Writer
var fw *bufio.Writer
var header = `package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/libopenstorage/openstorage/osdconfig"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"strconv"
)

type manager struct{}

func (m *manager) GetClusterConf() (*osdconfig.ClusterConfig, error) {
	config := new(osdconfig.ClusterConfig)
	if b, err := ioutil.ReadFile("/tmp/config.json"); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(b, config); err != nil {
			return nil, err
		}
	}
	return config, nil
}
func (m *manager) GetNodeConf(id string) (*osdconfig.NodeConfig, error) {
	configMap := make(map[string]*osdconfig.NodeConfig)
	if b, err := ioutil.ReadFile("/tmp/nodeConfig.json"); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(b, &configMap); err != nil {
			return nil, err
		}
	}
	return configMap[id], nil
}
func (m *manager) SetClusterConf(config *osdconfig.ClusterConfig) error {
	if b, err := json.Marshal(config); err != nil {
		return err
	} else {
		if err := ioutil.WriteFile("/tmp/config.json", b, 0666); err != nil {
			return err
		}
	}
	return nil
}
func (m *manager) SetNodeConf(config *osdconfig.NodeConfig) error {
	configMap := make(map[string]*osdconfig.NodeConfig)
	if b, err := ioutil.ReadFile("/tmp/nodeConfig.json"); err != nil {
		logrus.Warn(err)
	} else {
		if err := json.Unmarshal(b, &configMap); err != nil {
			logrus.Warn(err)
		}
	}
	configMap[config.NodeId] = config
	if b, err := json.Marshal(configMap); err != nil {
		return err
	} else {
		if err := ioutil.WriteFile("/tmp/nodeConfig.json", b, 0666); err != nil {
			return err
		}
	}
	return nil
}

func (m *manager) EnumerateNodeConf() (*osdconfig.NodesConfig, error) {
	configMap := make(map[string]*osdconfig.NodeConfig)
	if b, err := ioutil.ReadFile("/tmp/nodeConfig.json"); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(b, &configMap); err != nil {
			return nil, err
		}
	}
	nodesConfig := new(osdconfig.NodesConfig)
	for _, val := range configMap {
		*nodesConfig = append(*nodesConfig, val)
	}
	return nodesConfig, nil
}

func (m *manager) DeleteNodeConf(nodeId string) error {
	return nil
}

var osdconfigCaller *manager

func main() {
	osdconfigCaller = new(manager)

	fileInfo, err := os.Stat("/tmp/config.json")
	if err == nil && fileInfo.IsDir() {
		logrus.Fatal("log file is a dir")
	}

	if (err == nil && fileInfo.Size() == 0) || err != nil {
		clusterConfig := new(osdconfig.ClusterConfig).Init()
		clusterConfig.ClusterId = "myCluster"
		clusterConfig.Description = "myDescription"
		if jb, err := json.MarshalIndent(clusterConfig, "", "  "); err != nil {
			logrus.Fatal(err)
		} else {
			if err := ioutil.WriteFile("/tmp/config.json", jb, 0666); err != nil {
				logrus.Fatal(err)
			}
		}
	}

	fileInfo, err = os.Stat("/tmp/nodeConfig.json")
	if err == nil && fileInfo.IsDir() {
		logrus.Fatal("log file is a dir")
	}
	if (err == nil && fileInfo.Size() == 0) || err != nil {
		configMap := make(map[string]*osdconfig.NodeConfig)
		for i :=0; i < 3; i++ {
			config := new(osdconfig.NodeConfig).Init()
			config.NodeId = "nodeid_"+strconv.FormatInt(int64(i), 10)
			config.Storage.Devices = []string{"/dev/sda", "/dev/sdb"}
			configMap[config.NodeId] = config
		}
		if jb, err := json.Marshal(configMap); err != nil {
			logrus.Fatal(err)
		} else {
			if err := ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666); err != nil {
				logrus.Fatal(err)
			}
		}
	}

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "json, j",
			Usage: "Print JSON output",
		},
	}
	app.Commands = []cli.Command{`

func main() {
	prefix2VarName = make(map[string]string)
	prefixOrigin = make(map[string]string)
	prefixDepth = make(map[string]int)
	config := new(osdconfig.ClusterConfig).Init()

	var cb bytes.Buffer
	cw = bufio.NewWriter(&cb)

	var fb bytes.Buffer
	fw = bufio.NewWriter(&fb)

	fmt.Fprintln(cw, header)
	prefix2VarName["config"] = "config"
	prefixOrigin["config"] = "config"
	prefixDepth["config"] = 0
	printFields(reflect.Indirect(reflect.ValueOf(config)), false, "config", "Configure cluster",
		"Configure cluster and nodes. Node ID is required for node configuration. "+
			"Get node id using pxctl status", "\t\t")
	fmt.Fprintln(cw, tabs("\t", 0), "}")
	fmt.Fprintln(cw, tabs("\t", 0), "app.Run(os.Args)")
	fmt.Fprintln(cw, "}")

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
	if err := ioutil.WriteFile("cli.go", append(cb.Bytes(), fb.Bytes()...), 0666); err != nil {
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
	if prefixOrigin[prefix] == "node" {
		s := "c"
		for i := 0; i < prefixDepth[prefix]; i++ {
			s += ".Parent()"
		}
		fmt.Fprintln(fw, "\t", "if !"+s+".IsSet(\"node_id\") && !"+s+".IsSet(\"all\") {")
		fmt.Fprintln(fw, "\t\t", "err := errors.New(\"--node_id must be provided or --all must be set\")")
		fmt.Fprintln(fw, "\t\t", "logrus.Error(err)")
		fmt.Fprintln(fw, "\t\t", "return err")
		fmt.Fprintln(fw, "\t", "}")
		fmt.Fprintln(fw, "\t", "configs := new(osdconfig.NodesConfig)")
		fmt.Fprintln(fw, "\t", "var err error")
		fmt.Fprintln(fw, "\t", "if "+s+".IsSet(\"all\") {")
		fmt.Fprintln(fw, "\t\t", "configs, err = osdconfigCaller.EnumerateNodeConf()")
		fmt.Fprintln(fw, "\t\t", "if err != nil {")
		fmt.Fprintln(fw, "\t\t\t", "logrus.Error(err)")
		fmt.Fprintln(fw, "\t\t\t", "return err")
		fmt.Fprintln(fw, "\t\t", "}")
		fmt.Fprintln(fw, "\t", "} else {")
		fmt.Fprintln(fw, "\t\t", "config, err := osdconfigCaller.GetNodeConf("+s+".String(\"node_id\")"+")")
		fmt.Fprintln(fw, "\t\t", "if err != nil {")
		fmt.Fprintln(fw, "\t\t\t", "logrus.Error(err)")
		fmt.Fprintln(fw, "\t\t\t", "return err")
		fmt.Fprintln(fw, "\t\t", "}")
		fmt.Fprintln(fw, "\t\t", "*configs = append(*configs, config)")
		fmt.Fprintln(fw, "\t", "}")
		fmt.Fprintln(fw, "\t\t", "for _, config := range *configs {")
		fmt.Fprintln(fw, "\t\t\t", "config := config")
	} else {
		fmt.Fprintln(fw, "\t", "config, err := osdconfigCaller.GetClusterConf()")
		fmt.Fprintln(fw, "\t", "if err != nil {")
		fmt.Fprintln(fw, "\t\t", "logrus.Error(err)")
		fmt.Fprintln(fw, "\t\t", "return err")
		fmt.Fprintln(fw, "\t", "}")
	}
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

	if prefixOrigin[prefix] == "node" {
		fmt.Fprintln(fw, "\t\t\t", "if err := osdconfigCaller.SetNodeConf(config); err != nil {")
		fmt.Fprintln(fw, "\t\t\t\t", "logrus.Error(\"Set config for node: \", config.NodeId)")
		fmt.Fprintln(fw, "\t\t\t\t", "return err")
		fmt.Fprintln(fw, "\t\t\t", "}")
		fmt.Fprintln(fw, "\t\t\t", "logrus.Info(\"Set config for node: \", config.NodeId)")
		fmt.Fprintln(fw, "\t\t", "}")
		fmt.Fprintln(fw, "\t", "return nil")
	} else {
		fmt.Fprintln(fw, "\t", "if err := osdconfigCaller.SetClusterConf(config); err != nil {")
		fmt.Fprintln(fw, "\t\t", "logrus.Error(\"Set config for cluster\")")
		fmt.Fprintln(fw, "\t\t", "return err")
		fmt.Fprintln(fw, "\t", "}")
		fmt.Fprintln(fw, "\t", "logrus.Info(\"Set config for cluster\")")
		fmt.Fprintln(fw, "\t", "return nil")
	}
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
	if prefixOrigin[prefix] == "node" {
		s := "c.Parent()"
		for i := 0; i < prefixDepth[prefix]; i++ {
			s += ".Parent()"
		}
		fmt.Fprintln(fw, "\t", "if !"+s+".IsSet(\"node_id\") && !"+s+".IsSet(\"all\") {")
		fmt.Fprintln(fw, "\t\t", "err := errors.New(\"--node_id must be provided or --all must be set\")")
		fmt.Fprintln(fw, "\t\t", "logrus.Error(err)")
		fmt.Fprintln(fw, "\t\t", "return err")
		fmt.Fprintln(fw, "\t", "}")
		fmt.Fprintln(fw, "\t", "configs := new(osdconfig.NodesConfig)")
		fmt.Fprintln(fw, "\t", "var err error")
		fmt.Fprintln(fw, "\t", "if "+s+".IsSet(\"all\") {")
		fmt.Fprintln(fw, "\t\t", "configs, err = osdconfigCaller.EnumerateNodeConf()")
		fmt.Fprintln(fw, "\t\t", "if err != nil {")
		fmt.Fprintln(fw, "\t\t\t", "logrus.Error(err)")
		fmt.Fprintln(fw, "\t\t\t", "return err")
		fmt.Fprintln(fw, "\t\t", "}")
		fmt.Fprintln(fw, "\t", "} else {")
		fmt.Fprintln(fw, "\t\t", "config, err := osdconfigCaller.GetNodeConf("+s+".String(\"node_id\")"+")")
		fmt.Fprintln(fw, "\t\t", "if err != nil {")
		fmt.Fprintln(fw, "\t\t\t", "logrus.Error(err)")
		fmt.Fprintln(fw, "\t\t\t", "return err")
		fmt.Fprintln(fw, "\t\t", "}")
		fmt.Fprintln(fw, "\t\t", "*configs = append(*configs, config)")
		fmt.Fprintln(fw, "\t", "}")
		fmt.Fprintln(fw, "\t", "for _, config := range *configs {")
	} else {
		fmt.Fprintln(fw, "\t", "config, err := osdconfigCaller.GetClusterConf()")
		fmt.Fprintln(fw, "\t", "if err != nil {")
		fmt.Fprintln(fw, "\t\t", "logrus.Error(err)")
		fmt.Fprintln(fw, "\t\t", "return err")
		fmt.Fprintln(fw, "\t", "}")
	}
	fmt.Fprintln(fw, nullChecker(prefix2VarName[prefix]+".field"))
	fmt.Fprintln(fw, "\t", "if c.GlobalBool(\"json\") {")
	if prefixOrigin[prefix] == "node" {
		fmt.Fprintln(fw, "\t\t", "if err := printJson(struct{NodeId string `json:\"node_id\"`; Config interface{} `json:\"config\"`}{config.NodeId, "+prefix2VarName[prefix]+"}); err != nil {")
		fmt.Fprintln(fw, "\t\t\t", "return err")
		fmt.Fprintln(fw, "\t\t", "}")
		fmt.Fprintln(fw, "\t", "} else {")
		fmt.Fprintln(fw, "\t\t", "fmt.Println(\"node_id:\", config.NodeId)")
	} else {
		fmt.Fprintln(fw, "\t\t", "return printJson("+prefix2VarName[prefix]+")")
	}
	if prefixOrigin[prefix] != "node" {
		fmt.Fprintln(fw, "\t", "}")
	}
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
	if prefixOrigin[prefix] == "node" {
		fmt.Fprintln(fw, "\t\t\t", "fmt.Println()")
		fmt.Fprintln(fw, "\t\t", "}")
		fmt.Fprintln(fw, "\t", "}")
	}
	fmt.Fprintln(fw, "\t", "return nil")
	fmt.Fprintln(fw, "}")
	fmt.Fprintln(fw)
	fmt.Fprintln(cw, tabs(tab, 3), "},")
	fmt.Fprintln(cw, tabs(tab, 2), "},")

	// subcommand for node
	if prefix == "config" {
		config := new(osdconfig.NodeConfig).Init()
		prefix2VarName["node"] = "config"
		prefixOrigin["node"] = "node"
		prefixDepth["node"] = 0
		printFields(reflect.Indirect(reflect.ValueOf(config)),
			false, "node", "node usage", "node description", tabs(tab, 2))
	}

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
