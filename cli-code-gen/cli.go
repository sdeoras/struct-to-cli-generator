package main

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
	app.Commands = []cli.Command{
		 {
			 Name: "config",
			 Usage: "usage",
			 Description: "description",
			 Hidden: false ,
			 Action: setConfigValues,
			 Flags: []cli.Flag{
				 cli.StringFlag{
					 Name: "description",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "mode",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "version",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "created",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "cluster_id",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "logging_url",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "alerting_url",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "scheduler",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.BoolFlag{
					 Name: "multicontainer",
					 Usage: "(Bool)\tusage to be added",
					 Hidden: false,
				 },
				 cli.BoolFlag{
					 Name: "nolh",
					 Usage: "(Bool)\tusage to be added",
					 Hidden: false,
				 },
				 cli.BoolFlag{
					 Name: "callhome",
					 Usage: "(Bool)\tusage to be added",
					 Hidden: false,
				 },
				 cli.BoolFlag{
					 Name: "bootstrap",
					 Usage: "(Bool)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "tunnel_end_point",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringSliceFlag{
					 Name: "tunnel_certs",
					 Usage: "(Str...)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "driver",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "debug_level",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
				 cli.StringFlag{
					 Name: "domain",
					 Usage: "(Str)\tusage to be added",
					 Hidden: false,
				 },
			 },
			 Subcommands: []cli.Command{
				 {
					 Name: "show",
					 Usage: "Show values",
					 Description: "Show values",
					 Action: showConfigValues,
					 Flags: []cli.Flag{
						 cli.BoolFlag{
							 Name: "all, a",
							 Usage: "(Bool)\tShow all data",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "description",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "mode",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "version",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "created",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "cluster_id",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "logging_url",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "alerting_url",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "scheduler",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "multicontainer",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "nolh",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "callhome",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "bootstrap",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "tunnel_end_point",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "tunnel_certs",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "driver",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "debug_level",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
						 cli.BoolFlag{
							 Name: "domain",
							 Usage: "(Bool)\tusage to be added",
							 Hidden: false,
						 },
					 },
				 },
				 {
					 Name: "nodeConfig",
					 Usage: "node usage",
					 Description: "node description",
					 Hidden: false ,
					 Action: setNodeConfigValues,
					 Flags: []cli.Flag{
						 cli.StringFlag{
							 Name: "node_id",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
					 },
					 Subcommands: []cli.Command{
						 {
							 Name: "show",
							 Usage: "Show values",
							 Description: "Show values",
							 Action: showNodeConfigValues,
							 Flags: []cli.Flag{
								 cli.BoolFlag{
									 Name: "all, a",
									 Usage: "(Bool)\tShow all data",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "node_id",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
							 },
						 },
						 {
							 Name: "network",
							 Usage: "usage to be added",
							 Description: "description to be added",
							 Hidden: false ,
							 Action: setNetworkValues,
							 Flags: []cli.Flag{
								 cli.StringFlag{
									 Name: "mgt_iface",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "data_iface",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
							 },
							 Subcommands: []cli.Command{
								 {
									 Name: "show",
									 Usage: "Show values",
									 Description: "Show values",
									 Action: showNetworkValues,
									 Flags: []cli.Flag{
										 cli.BoolFlag{
											 Name: "all, a",
											 Usage: "(Bool)\tShow all data",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "mgt_iface",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "data_iface",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
									 },
								 },
							 },
						 },
						 {
							 Name: "storage",
							 Usage: "usage to be added",
							 Description: "description to be added",
							 Hidden: false ,
							 Action: setStorageValues,
							 Flags: []cli.Flag{
								 cli.StringSliceFlag{
									 Name: "devices_md",
									 Usage: "(Str...)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringSliceFlag{
									 Name: "devices",
									 Usage: "(Str...)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "raid_level",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "raid_level_md",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "async_io",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
							 },
							 Subcommands: []cli.Command{
								 {
									 Name: "show",
									 Usage: "Show values",
									 Description: "Show values",
									 Action: showStorageValues,
									 Flags: []cli.Flag{
										 cli.BoolFlag{
											 Name: "all, a",
											 Usage: "(Bool)\tShow all data",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "devices_md",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "devices",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "raid_level",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "raid_level_md",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "async_io",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
									 },
								 },
							 },
						 },
					 },
				 },
				 {
					 Name: "secrets",
					 Usage: "usage to be added",
					 Description: "description to be added",
					 Hidden: false ,
					 Action: setSecretsValues,
					 Flags: []cli.Flag{
						 cli.StringFlag{
							 Name: "secret_type",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
						 cli.StringFlag{
							 Name: "cluster_secret_key",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
					 },
					 Subcommands: []cli.Command{
						 {
							 Name: "show",
							 Usage: "Show values",
							 Description: "Show values",
							 Action: showSecretsValues,
							 Flags: []cli.Flag{
								 cli.BoolFlag{
									 Name: "all, a",
									 Usage: "(Bool)\tShow all data",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "secret_type",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "cluster_secret_key",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
							 },
						 },
						 {
							 Name: "vault",
							 Usage: "usage to be added",
							 Description: "none yet",
							 Hidden: false ,
							 Action: setVaultValues,
							 Flags: []cli.Flag{
								 cli.StringFlag{
									 Name: "vault_token",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "vault_addr",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "vault_cacert",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "vault_capath",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "vault_client_cert",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "vault_client_key",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "vault_skip_verify",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "vault_tls_server_name",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "vault_base_path",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
							 },
							 Subcommands: []cli.Command{
								 {
									 Name: "show",
									 Usage: "Show values",
									 Description: "Show values",
									 Action: showVaultValues,
									 Flags: []cli.Flag{
										 cli.BoolFlag{
											 Name: "all, a",
											 Usage: "(Bool)\tShow all data",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_token",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_addr",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_cacert",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_capath",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_client_cert",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_client_key",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_skip_verify",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_tls_server_name",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "vault_base_path",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
									 },
								 },
							 },
						 },
						 {
							 Name: "aws",
							 Usage: "usage to be added",
							 Description: "none yet",
							 Hidden: false ,
							 Action: setAwsValues,
							 Flags: []cli.Flag{
								 cli.StringFlag{
									 Name: "aws_access_key_id",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "aws_secret_access_key",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "aws_secret_token_key",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "aws_cmk",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
								 cli.StringFlag{
									 Name: "aws_region",
									 Usage: "(Str)\tusage to be added",
									 Hidden: false,
								 },
							 },
							 Subcommands: []cli.Command{
								 {
									 Name: "show",
									 Usage: "Show values",
									 Description: "Show values",
									 Action: showAwsValues,
									 Flags: []cli.Flag{
										 cli.BoolFlag{
											 Name: "all, a",
											 Usage: "(Bool)\tShow all data",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "aws_access_key_id",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "aws_secret_access_key",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "aws_secret_token_key",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "aws_cmk",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
										 cli.BoolFlag{
											 Name: "aws_region",
											 Usage: "(Bool)\tusage to be added",
											 Hidden: false,
										 },
									 },
								 },
							 },
						 },
					 },
				 },
				 {
					 Name: "kvdb",
					 Usage: "usage to be added",
					 Description: "description to be added",
					 Hidden: false ,
					 Action: setKvdbValues,
					 Flags: []cli.Flag{
						 cli.StringFlag{
							 Name: "username",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
						 cli.StringFlag{
							 Name: "password",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
						 cli.StringFlag{
							 Name: "ca_file",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
						 cli.StringFlag{
							 Name: "cert_file",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
						 cli.StringFlag{
							 Name: "trusted_ca_file",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
						 cli.StringFlag{
							 Name: "client_cert_auth",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
						 cli.StringFlag{
							 Name: "acl_token",
							 Usage: "(Str)\tusage to be added",
							 Hidden: false,
						 },
						 cli.StringSliceFlag{
							 Name: "kvdb_addr",
							 Usage: "(Str...)\tusage to be added",
							 Hidden: false,
						 },
					 },
					 Subcommands: []cli.Command{
						 {
							 Name: "show",
							 Usage: "Show values",
							 Description: "Show values",
							 Action: showKvdbValues,
							 Flags: []cli.Flag{
								 cli.BoolFlag{
									 Name: "all, a",
									 Usage: "(Bool)\tShow all data",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "username",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "password",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "ca_file",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "cert_file",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "trusted_ca_file",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "client_cert_auth",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "acl_token",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
								 cli.BoolFlag{
									 Name: "kvdb_addr",
									 Usage: "(Bool)\tusage to be added",
									 Hidden: false,
								 },
							 },
						 },
					 },
				 },
			 },
		 },
	 }
	 app.Run(os.Args)
}
func setConfigValues(c *cli.Context) {
	 if c.IsSet("description") {
		 description := c.String("description")
		 config.Description  =  description
	 }
	 if c.IsSet("mode") {
		 mode := c.String("mode")
		 config.Mode  =  mode
	 }
	 if c.IsSet("version") {
		 version := c.String("version")
		 config.Version  =  version
	 }
	 if c.IsSet("created") {
		 created := c.String("created")
		 config.Created  =  created
	 }
	 if c.IsSet("cluster_id") {
		 clusterId := c.String("cluster_id")
		 config.ClusterId  =  clusterId
	 }
	 if c.IsSet("logging_url") {
		 loggingUrl := c.String("logging_url")
		 config.LoggingUrl  =  loggingUrl
	 }
	 if c.IsSet("alerting_url") {
		 alertingUrl := c.String("alerting_url")
		 config.AlertingUrl  =  alertingUrl
	 }
	 if c.IsSet("scheduler") {
		 scheduler := c.String("scheduler")
		 config.Scheduler  =  scheduler
	 }
	 if c.IsSet("multicontainer") {
		 multicontainer := c.Bool("multicontainer")
		 config.Multicontainer  =  multicontainer
	 }
	 if c.IsSet("nolh") {
		 nolh := c.Bool("nolh")
		 config.Nolh  =  nolh
	 }
	 if c.IsSet("callhome") {
		 callhome := c.Bool("callhome")
		 config.Callhome  =  callhome
	 }
	 if c.IsSet("bootstrap") {
		 bootstrap := c.Bool("bootstrap")
		 config.Bootstrap  =  bootstrap
	 }
	 if c.IsSet("tunnel_end_point") {
		 tunnelEndPoint := c.String("tunnel_end_point")
		 config.TunnelEndPoint  =  tunnelEndPoint
	 }
	 if c.IsSet("tunnel_certs") {
		 tunnelCerts := c.StringSlice("tunnel_certs")
		 config.TunnelCerts  =  tunnelCerts
	 }
	 if c.IsSet("driver") {
		 driver := c.String("driver")
		 config.Driver  =  driver
	 }
	 if c.IsSet("debug_level") {
		 debugLevel := c.String("debug_level")
		 config.DebugLevel  =  debugLevel
	 }
	 if c.IsSet("domain") {
		 domain := c.String("domain")
		 config.Domain  =  domain
	 }
	 jb, _ := json.Marshal(config); ioutil.WriteFile("/tmp/config.json", jb, 0666)
	 jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666)
}

func showConfigValues(c *cli.Context) {
	 if c.IsSet("all") || c.IsSet("description") {
		 fmt.Println( "Description:", config.Description )
	 }
	 if c.IsSet("all") || c.IsSet("mode") {
		 fmt.Println( "Mode:", config.Mode )
	 }
	 if c.IsSet("all") || c.IsSet("version") {
		 fmt.Println( "Version:", config.Version )
	 }
	 if c.IsSet("all") || c.IsSet("created") {
		 fmt.Println( "Created:", config.Created )
	 }
	 if c.IsSet("all") || c.IsSet("cluster_id") {
		 fmt.Println( "ClusterId:", config.ClusterId )
	 }
	 if c.IsSet("all") || c.IsSet("logging_url") {
		 fmt.Println( "LoggingUrl:", config.LoggingUrl )
	 }
	 if c.IsSet("all") || c.IsSet("alerting_url") {
		 fmt.Println( "AlertingUrl:", config.AlertingUrl )
	 }
	 if c.IsSet("all") || c.IsSet("scheduler") {
		 fmt.Println( "Scheduler:", config.Scheduler )
	 }
	 if c.IsSet("all") || c.IsSet("multicontainer") {
		 fmt.Println( "Multicontainer:", config.Multicontainer )
	 }
	 if c.IsSet("all") || c.IsSet("nolh") {
		 fmt.Println( "Nolh:", config.Nolh )
	 }
	 if c.IsSet("all") || c.IsSet("callhome") {
		 fmt.Println( "Callhome:", config.Callhome )
	 }
	 if c.IsSet("all") || c.IsSet("bootstrap") {
		 fmt.Println( "Bootstrap:", config.Bootstrap )
	 }
	 if c.IsSet("all") || c.IsSet("tunnel_end_point") {
		 fmt.Println( "TunnelEndPoint:", config.TunnelEndPoint )
	 }
	 if c.IsSet("all") || c.IsSet("tunnel_certs") {
		 fmt.Println( "TunnelCerts:", config.TunnelCerts )
	 }
	 if c.IsSet("all") || c.IsSet("driver") {
		 fmt.Println( "Driver:", config.Driver )
	 }
	 if c.IsSet("all") || c.IsSet("debug_level") {
		 fmt.Println( "DebugLevel:", config.DebugLevel )
	 }
	 if c.IsSet("all") || c.IsSet("domain") {
		 fmt.Println( "Domain:", config.Domain )
	 }
}

func setNodeConfigValues(c *cli.Context) {
	 if c.IsSet("node_id") {
		 nodeId := c.String("node_id")
		 nodeConfig.NodeId  =  nodeId
	 }
	 jb, _ := json.Marshal(config); ioutil.WriteFile("/tmp/config.json", jb, 0666)
	 jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666)
}

func showNodeConfigValues(c *cli.Context) {
	 if c.IsSet("all") || c.IsSet("node_id") {
		 fmt.Println( "NodeId:", nodeConfig.NodeId )
	 }
}

func setNetworkValues(c *cli.Context) {
	 if c.IsSet("mgt_iface") {
		 mgtIface := c.String("mgt_iface")
		 nodeConfig.Network.MgtIface  =  mgtIface
	 }
	 if c.IsSet("data_iface") {
		 dataIface := c.String("data_iface")
		 nodeConfig.Network.DataIface  =  dataIface
	 }
	 jb, _ := json.Marshal(config); ioutil.WriteFile("/tmp/config.json", jb, 0666)
	 jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666)
}

func showNetworkValues(c *cli.Context) {
	 if c.IsSet("all") || c.IsSet("mgt_iface") {
		 fmt.Println( "MgtIface:", nodeConfig.Network.MgtIface )
	 }
	 if c.IsSet("all") || c.IsSet("data_iface") {
		 fmt.Println( "DataIface:", nodeConfig.Network.DataIface )
	 }
}

func setStorageValues(c *cli.Context) {
	 if c.IsSet("devices_md") {
		 devicesMd := c.StringSlice("devices_md")
		 nodeConfig.Storage.DevicesMd  =  devicesMd
	 }
	 if c.IsSet("devices") {
		 devices := c.StringSlice("devices")
		 nodeConfig.Storage.Devices  =  devices
	 }
	 if c.IsSet("raid_level") {
		 raidLevel := c.String("raid_level")
		 nodeConfig.Storage.RaidLevel  =  raidLevel
	 }
	 if c.IsSet("raid_level_md") {
		 raidLevelMd := c.String("raid_level_md")
		 nodeConfig.Storage.RaidLevelMd  =  raidLevelMd
	 }
	 if c.IsSet("async_io") {
		 asyncIo := c.Bool("async_io")
		 nodeConfig.Storage.AsyncIo  =  asyncIo
	 }
	 jb, _ := json.Marshal(config); ioutil.WriteFile("/tmp/config.json", jb, 0666)
	 jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666)
}

func showStorageValues(c *cli.Context) {
	 if c.IsSet("all") || c.IsSet("devices_md") {
		 fmt.Println( "DevicesMd:", nodeConfig.Storage.DevicesMd )
	 }
	 if c.IsSet("all") || c.IsSet("devices") {
		 fmt.Println( "Devices:", nodeConfig.Storage.Devices )
	 }
	 if c.IsSet("all") || c.IsSet("raid_level") {
		 fmt.Println( "RaidLevel:", nodeConfig.Storage.RaidLevel )
	 }
	 if c.IsSet("all") || c.IsSet("raid_level_md") {
		 fmt.Println( "RaidLevelMd:", nodeConfig.Storage.RaidLevelMd )
	 }
	 if c.IsSet("all") || c.IsSet("async_io") {
		 fmt.Println( "AsyncIo:", nodeConfig.Storage.AsyncIo )
	 }
}

func setSecretsValues(c *cli.Context) {
	 if c.IsSet("secret_type") {
		 secretType := c.String("secret_type")
		 config.Secrets.SecretType  =  secretType
	 }
	 if c.IsSet("cluster_secret_key") {
		 clusterSecretKey := c.String("cluster_secret_key")
		 config.Secrets.ClusterSecretKey  =  clusterSecretKey
	 }
	 jb, _ := json.Marshal(config); ioutil.WriteFile("/tmp/config.json", jb, 0666)
	 jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666)
}

func showSecretsValues(c *cli.Context) {
	 if c.IsSet("all") || c.IsSet("secret_type") {
		 fmt.Println( "SecretType:", config.Secrets.SecretType )
	 }
	 if c.IsSet("all") || c.IsSet("cluster_secret_key") {
		 fmt.Println( "ClusterSecretKey:", config.Secrets.ClusterSecretKey )
	 }
}

func setVaultValues(c *cli.Context) {
	 if c.IsSet("vault_token") {
		 vaultToken := c.String("vault_token")
		 config.Secrets.Vault.VaultToken  =  vaultToken
	 }
	 if c.IsSet("vault_addr") {
		 vaultAddr := c.String("vault_addr")
		 config.Secrets.Vault.VaultAddr  =  vaultAddr
	 }
	 if c.IsSet("vault_cacert") {
		 vaultCacert := c.String("vault_cacert")
		 config.Secrets.Vault.VaultCacert  =  vaultCacert
	 }
	 if c.IsSet("vault_capath") {
		 vaultCapath := c.String("vault_capath")
		 config.Secrets.Vault.VaultCapath  =  vaultCapath
	 }
	 if c.IsSet("vault_client_cert") {
		 vaultClientCert := c.String("vault_client_cert")
		 config.Secrets.Vault.VaultClientCert  =  vaultClientCert
	 }
	 if c.IsSet("vault_client_key") {
		 vaultClientKey := c.String("vault_client_key")
		 config.Secrets.Vault.VaultClientKey  =  vaultClientKey
	 }
	 if c.IsSet("vault_skip_verify") {
		 vaultSkipVerify := c.String("vault_skip_verify")
		 config.Secrets.Vault.VaultSkipVerify  =  vaultSkipVerify
	 }
	 if c.IsSet("vault_tls_server_name") {
		 vaultTlsServerName := c.String("vault_tls_server_name")
		 config.Secrets.Vault.VaultTlsServerName  =  vaultTlsServerName
	 }
	 if c.IsSet("vault_base_path") {
		 vaultBasePath := c.String("vault_base_path")
		 config.Secrets.Vault.VaultBasePath  =  vaultBasePath
	 }
	 jb, _ := json.Marshal(config); ioutil.WriteFile("/tmp/config.json", jb, 0666)
	 jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666)
}

func showVaultValues(c *cli.Context) {
	 if c.IsSet("all") || c.IsSet("vault_token") {
		 fmt.Println( "VaultToken:", config.Secrets.Vault.VaultToken )
	 }
	 if c.IsSet("all") || c.IsSet("vault_addr") {
		 fmt.Println( "VaultAddr:", config.Secrets.Vault.VaultAddr )
	 }
	 if c.IsSet("all") || c.IsSet("vault_cacert") {
		 fmt.Println( "VaultCacert:", config.Secrets.Vault.VaultCacert )
	 }
	 if c.IsSet("all") || c.IsSet("vault_capath") {
		 fmt.Println( "VaultCapath:", config.Secrets.Vault.VaultCapath )
	 }
	 if c.IsSet("all") || c.IsSet("vault_client_cert") {
		 fmt.Println( "VaultClientCert:", config.Secrets.Vault.VaultClientCert )
	 }
	 if c.IsSet("all") || c.IsSet("vault_client_key") {
		 fmt.Println( "VaultClientKey:", config.Secrets.Vault.VaultClientKey )
	 }
	 if c.IsSet("all") || c.IsSet("vault_skip_verify") {
		 fmt.Println( "VaultSkipVerify:", config.Secrets.Vault.VaultSkipVerify )
	 }
	 if c.IsSet("all") || c.IsSet("vault_tls_server_name") {
		 fmt.Println( "VaultTlsServerName:", config.Secrets.Vault.VaultTlsServerName )
	 }
	 if c.IsSet("all") || c.IsSet("vault_base_path") {
		 fmt.Println( "VaultBasePath:", config.Secrets.Vault.VaultBasePath )
	 }
}

func setAwsValues(c *cli.Context) {
	 if c.IsSet("aws_access_key_id") {
		 awsAccessKeyId := c.String("aws_access_key_id")
		 config.Secrets.Aws.AwsAccessKeyId  =  awsAccessKeyId
	 }
	 if c.IsSet("aws_secret_access_key") {
		 awsSecretAccessKey := c.String("aws_secret_access_key")
		 config.Secrets.Aws.AwsSecretAccessKey  =  awsSecretAccessKey
	 }
	 if c.IsSet("aws_secret_token_key") {
		 awsSecretTokenKey := c.String("aws_secret_token_key")
		 config.Secrets.Aws.AwsSecretTokenKey  =  awsSecretTokenKey
	 }
	 if c.IsSet("aws_cmk") {
		 awsCmk := c.String("aws_cmk")
		 config.Secrets.Aws.AwsCmk  =  awsCmk
	 }
	 if c.IsSet("aws_region") {
		 awsRegion := c.String("aws_region")
		 config.Secrets.Aws.AwsRegion  =  awsRegion
	 }
	 jb, _ := json.Marshal(config); ioutil.WriteFile("/tmp/config.json", jb, 0666)
	 jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666)
}

func showAwsValues(c *cli.Context) {
	 if c.IsSet("all") || c.IsSet("aws_access_key_id") {
		 fmt.Println( "AwsAccessKeyId:", config.Secrets.Aws.AwsAccessKeyId )
	 }
	 if c.IsSet("all") || c.IsSet("aws_secret_access_key") {
		 fmt.Println( "AwsSecretAccessKey:", config.Secrets.Aws.AwsSecretAccessKey )
	 }
	 if c.IsSet("all") || c.IsSet("aws_secret_token_key") {
		 fmt.Println( "AwsSecretTokenKey:", config.Secrets.Aws.AwsSecretTokenKey )
	 }
	 if c.IsSet("all") || c.IsSet("aws_cmk") {
		 fmt.Println( "AwsCmk:", config.Secrets.Aws.AwsCmk )
	 }
	 if c.IsSet("all") || c.IsSet("aws_region") {
		 fmt.Println( "AwsRegion:", config.Secrets.Aws.AwsRegion )
	 }
}

func setKvdbValues(c *cli.Context) {
	 if c.IsSet("username") {
		 username := c.String("username")
		 config.Kvdb.Username  =  username
	 }
	 if c.IsSet("password") {
		 password := c.String("password")
		 config.Kvdb.Password  =  password
	 }
	 if c.IsSet("ca_file") {
		 caFile := c.String("ca_file")
		 config.Kvdb.CaFile  =  caFile
	 }
	 if c.IsSet("cert_file") {
		 certFile := c.String("cert_file")
		 config.Kvdb.CertFile  =  certFile
	 }
	 if c.IsSet("trusted_ca_file") {
		 trustedCaFile := c.String("trusted_ca_file")
		 config.Kvdb.TrustedCaFile  =  trustedCaFile
	 }
	 if c.IsSet("client_cert_auth") {
		 clientCertAuth := c.String("client_cert_auth")
		 config.Kvdb.ClientCertAuth  =  clientCertAuth
	 }
	 if c.IsSet("acl_token") {
		 aclToken := c.String("acl_token")
		 config.Kvdb.AclToken  =  aclToken
	 }
	 if c.IsSet("kvdb_addr") {
		 kvdbAddr := c.StringSlice("kvdb_addr")
		 config.Kvdb.KvdbAddr  =  kvdbAddr
	 }
	 jb, _ := json.Marshal(config); ioutil.WriteFile("/tmp/config.json", jb, 0666)
	 jb, _ = json.Marshal(nodeConfig); ioutil.WriteFile("/tmp/nodeConfig.json", jb, 0666)
}

func showKvdbValues(c *cli.Context) {
	 if c.IsSet("all") || c.IsSet("username") {
		 fmt.Println( "Username:", config.Kvdb.Username )
	 }
	 if c.IsSet("all") || c.IsSet("password") {
		 fmt.Println( "Password:", config.Kvdb.Password )
	 }
	 if c.IsSet("all") || c.IsSet("ca_file") {
		 fmt.Println( "CaFile:", config.Kvdb.CaFile )
	 }
	 if c.IsSet("all") || c.IsSet("cert_file") {
		 fmt.Println( "CertFile:", config.Kvdb.CertFile )
	 }
	 if c.IsSet("all") || c.IsSet("trusted_ca_file") {
		 fmt.Println( "TrustedCaFile:", config.Kvdb.TrustedCaFile )
	 }
	 if c.IsSet("all") || c.IsSet("client_cert_auth") {
		 fmt.Println( "ClientCertAuth:", config.Kvdb.ClientCertAuth )
	 }
	 if c.IsSet("all") || c.IsSet("acl_token") {
		 fmt.Println( "AclToken:", config.Kvdb.AclToken )
	 }
	 if c.IsSet("all") || c.IsSet("kvdb_addr") {
		 fmt.Println( "KvdbAddr:", config.Kvdb.KvdbAddr )
	 }
}

