package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/libopenstorage/openstorage/osdconfig"
	"github.com/urfave/cli"
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
		if err := json.Unmarshal(b, configMap); err != nil {
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

var clusterManager *manager

func main() {
	clusterManager = new(manager)

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "json, j",
			Usage: "Print JSON output",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:        "config",
			Usage:       "Configure cluster",
			Description: "Configure cluster and nodes. Node ID is required for node configuration. Get node id using pxctl status",
			Hidden:      false,
			Action:      setConfigValues,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "description",
					Usage:  "(Str)\tCluster description",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "mode",
					Usage:  "(Str)\tMode for cluster",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "version",
					Usage:  "(Str)\tVersion info for cluster",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "created",
					Usage:  "(Str)\tCreation info for cluster",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "cluster_id",
					Usage:  "(Str)\tCluster ID info",
					Hidden: false,
				},
				cli.StringSliceFlag{
					Name:   "node_id",
					Usage:  "(Str...)\tNode ID info",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "logging_url",
					Usage:  "(Str)\tURL for logging",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "alerting_url",
					Usage:  "(Str)\tURL for altering",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "scheduler",
					Usage:  "(Str)\tCluster scheduler",
					Hidden: false,
				},
				cli.BoolFlag{
					Name:   "multicontainer",
					Usage:  "(Bool)\tusage to be added",
					Hidden: false,
				},
				cli.BoolFlag{
					Name:   "nolh",
					Usage:  "(Bool)\tusage to be added",
					Hidden: false,
				},
				cli.BoolFlag{
					Name:   "callhome",
					Usage:  "(Bool)\tusage to be added",
					Hidden: false,
				},
				cli.BoolFlag{
					Name:   "bootstrap",
					Usage:  "(Bool)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "tunnel_end_point",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringSliceFlag{
					Name:   "tunnel_certs",
					Usage:  "(Str...)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "driver",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "debug_level",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "domain",
					Usage:  "(Str)\tusage to be added",
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
							Name:   "description",
							Usage:  "(Bool)\tCluster description",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "mode",
							Usage:  "(Bool)\tMode for cluster",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "version",
							Usage:  "(Bool)\tVersion info for cluster",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "created",
							Usage:  "(Bool)\tCreation info for cluster",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "cluster_id",
							Usage:  "(Bool)\tCluster ID info",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "node_id",
							Usage:  "(Bool)\tNode ID info",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "logging_url",
							Usage:  "(Bool)\tURL for logging",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "alerting_url",
							Usage:  "(Bool)\tURL for altering",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "scheduler",
							Usage:  "(Bool)\tCluster scheduler",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "multicontainer",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "nolh",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "callhome",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "bootstrap",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "tunnel_end_point",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "tunnel_certs",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "driver",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "debug_level",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "domain",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
					},
				},
				{
					Name:        "node",
					Usage:       "node usage",
					Description: "node description",
					Hidden:      false,
					Action:      setNodeValues,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "node_id",
							Usage:  "(Str)\tID for the node",
							Hidden: false,
						},
					},
					Subcommands: []cli.Command{
						{
							Name:        "show",
							Usage:       "Show values",
							Description: "Show values",
							Action:      showNodeValues,
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name:   "all, a",
									Usage:  "(Bool)\tShow all data",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "node_id",
									Usage:  "(Bool)\tID for the node",
									Hidden: false,
								},
							},
						},
						{
							Name:        "network",
							Usage:       "Network configuration",
							Description: "Configure network values for a node",
							Hidden:      false,
							Action:      setNetworkValues,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:   "mgt_iface",
									Usage:  "(Str)\tManagement interface",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "data_iface",
									Usage:  "(Str)\tData interface",
									Hidden: false,
								},
							},
							Subcommands: []cli.Command{
								{
									Name:        "show",
									Usage:       "Show values",
									Description: "Show values",
									Action:      showNetworkValues,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name:   "all, a",
											Usage:  "(Bool)\tShow all data",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "mgt_iface",
											Usage:  "(Bool)\tManagement interface",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "data_iface",
											Usage:  "(Bool)\tData interface",
											Hidden: false,
										},
									},
								},
							},
						},
						{
							Name:        "storage",
							Usage:       "Storage configuration",
							Description: "Configure storage values for a node",
							Hidden:      false,
							Action:      setStorageValues,
							Flags: []cli.Flag{
								cli.StringSliceFlag{
									Name:   "devices_md",
									Usage:  "(Str...)\tDevices MD",
									Hidden: false,
								},
								cli.StringSliceFlag{
									Name:   "devices",
									Usage:  "(Str...)\tDevices list",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "raid_level",
									Usage:  "(Str)\tRAID level info",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "raid_level_md",
									Usage:  "(Str)\tRAID level MD",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "async_io",
									Usage:  "(Bool)\tAsync I/O",
									Hidden: false,
								},
							},
							Subcommands: []cli.Command{
								{
									Name:        "show",
									Usage:       "Show values",
									Description: "Show values",
									Action:      showStorageValues,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name:   "all, a",
											Usage:  "(Bool)\tShow all data",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "devices_md",
											Usage:  "(Bool)\tDevices MD",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "devices",
											Usage:  "(Bool)\tDevices list",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "raid_level",
											Usage:  "(Bool)\tRAID level info",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "raid_level_md",
											Usage:  "(Bool)\tRAID level MD",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "async_io",
											Usage:  "(Bool)\tAsync I/O",
											Hidden: false,
										},
									},
								},
							},
						},
					},
				},
				{
					Name:        "secrets",
					Usage:       "usage to be added",
					Description: "description to be added",
					Hidden:      false,
					Action:      setSecretsValues,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "secret_type",
							Usage:  "(Str)\tSecret type",
							Hidden: false,
						},
						cli.StringFlag{
							Name:   "cluster_secret_key",
							Usage:  "(Str)\tSecret key",
							Hidden: false,
						},
					},
					Subcommands: []cli.Command{
						{
							Name:        "show",
							Usage:       "Show values",
							Description: "Show values",
							Action:      showSecretsValues,
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name:   "all, a",
									Usage:  "(Bool)\tShow all data",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "secret_type",
									Usage:  "(Bool)\tSecret type",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "cluster_secret_key",
									Usage:  "(Bool)\tSecret key",
									Hidden: false,
								},
							},
						},
						{
							Name:        "vault",
							Usage:       "Vault configuration",
							Description: "none yet",
							Hidden:      false,
							Action:      setVaultValues,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:   "vault_token",
									Usage:  "(Str)\tVault token",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_addr",
									Usage:  "(Str)\tVault address",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_cacert",
									Usage:  "(Str)\tVault CA certificate",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_capath",
									Usage:  "(Str)\tVault CA path",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_client_cert",
									Usage:  "(Str)\tVault client certificate",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_client_key",
									Usage:  "(Str)\tVault client key",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_skip_verify",
									Usage:  "(Str)\tVault skip verification",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_tls_server_name",
									Usage:  "(Str)\tVault TLS server name",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_base_path",
									Usage:  "(Str)\tVault base path",
									Hidden: false,
								},
							},
							Subcommands: []cli.Command{
								{
									Name:        "show",
									Usage:       "Show values",
									Description: "Show values",
									Action:      showVaultValues,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name:   "all, a",
											Usage:  "(Bool)\tShow all data",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_token",
											Usage:  "(Bool)\tVault token",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_addr",
											Usage:  "(Bool)\tVault address",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_cacert",
											Usage:  "(Bool)\tVault CA certificate",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_capath",
											Usage:  "(Bool)\tVault CA path",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_client_cert",
											Usage:  "(Bool)\tVault client certificate",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_client_key",
											Usage:  "(Bool)\tVault client key",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_skip_verify",
											Usage:  "(Bool)\tVault skip verification",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_tls_server_name",
											Usage:  "(Bool)\tVault TLS server name",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_base_path",
											Usage:  "(Bool)\tVault base path",
											Hidden: false,
										},
									},
								},
							},
						},
						{
							Name:        "aws",
							Usage:       "AWS configuration",
							Description: "none yet",
							Hidden:      false,
							Action:      setAwsValues,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:   "aws_access_key_id",
									Usage:  "(Str)\tAWS access key ID",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "aws_secret_access_key",
									Usage:  "(Str)\tAWS secret access key",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "aws_secret_token_key",
									Usage:  "(Str)\tAWS secret token key",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "aws_cmk",
									Usage:  "(Str)\tAWS CMK",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "aws_region",
									Usage:  "(Str)\tAWS region",
									Hidden: false,
								},
							},
							Subcommands: []cli.Command{
								{
									Name:        "show",
									Usage:       "Show values",
									Description: "Show values",
									Action:      showAwsValues,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name:   "all, a",
											Usage:  "(Bool)\tShow all data",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_access_key_id",
											Usage:  "(Bool)\tAWS access key ID",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_secret_access_key",
											Usage:  "(Bool)\tAWS secret access key",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_secret_token_key",
											Usage:  "(Bool)\tAWS secret token key",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_cmk",
											Usage:  "(Bool)\tAWS CMK",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_region",
											Usage:  "(Bool)\tAWS region",
											Hidden: false,
										},
									},
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
func setConfigValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.IsSet("description") {
		config.Description = c.String("description")
	}
	if c.IsSet("mode") {
		config.Mode = c.String("mode")
	}
	if c.IsSet("version") {
		config.Version = c.String("version")
	}
	if c.IsSet("created") {
		config.Created = c.String("created")
	}
	if c.IsSet("cluster_id") {
		config.ClusterId = c.String("cluster_id")
	}
	if c.IsSet("node_id") {
		config.NodeId = c.StringSlice("node_id")
	}
	if c.IsSet("logging_url") {
		config.LoggingUrl = c.String("logging_url")
	}
	if c.IsSet("alerting_url") {
		config.AlertingUrl = c.String("alerting_url")
	}
	if c.IsSet("scheduler") {
		config.Scheduler = c.String("scheduler")
	}
	if c.IsSet("multicontainer") {
		config.Multicontainer = c.Bool("multicontainer")
	}
	if c.IsSet("nolh") {
		config.Nolh = c.Bool("nolh")
	}
	if c.IsSet("callhome") {
		config.Callhome = c.Bool("callhome")
	}
	if c.IsSet("bootstrap") {
		config.Bootstrap = c.Bool("bootstrap")
	}
	if c.IsSet("tunnel_end_point") {
		config.TunnelEndPoint = c.String("tunnel_end_point")
	}
	if c.IsSet("tunnel_certs") {
		config.TunnelCerts = c.StringSlice("tunnel_certs")
	}
	if c.IsSet("driver") {
		config.Driver = c.String("driver")
	}
	if c.IsSet("debug_level") {
		config.DebugLevel = c.String("debug_level")
	}
	if c.IsSet("domain") {
		config.Domain = c.String("domain")
	}
	return clusterManager.SetClusterConf(config)
}

func showConfigValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config)
	}
	if c.IsSet("all") || c.IsSet("description") {
		fmt.Println("description:", config.Description)
	}
	if c.IsSet("all") || c.IsSet("mode") {
		fmt.Println("mode:", config.Mode)
	}
	if c.IsSet("all") || c.IsSet("version") {
		fmt.Println("version:", config.Version)
	}
	if c.IsSet("all") || c.IsSet("created") {
		fmt.Println("created:", config.Created)
	}
	if c.IsSet("all") || c.IsSet("cluster_id") {
		fmt.Println("cluster_id:", config.ClusterId)
	}
	if c.IsSet("all") || c.IsSet("node_id") {
		fmt.Println("node_id:", config.NodeId)
	}
	if c.IsSet("all") || c.IsSet("logging_url") {
		fmt.Println("logging_url:", config.LoggingUrl)
	}
	if c.IsSet("all") || c.IsSet("alerting_url") {
		fmt.Println("alerting_url:", config.AlertingUrl)
	}
	if c.IsSet("all") || c.IsSet("scheduler") {
		fmt.Println("scheduler:", config.Scheduler)
	}
	if c.IsSet("all") || c.IsSet("multicontainer") {
		fmt.Println("multicontainer:", config.Multicontainer)
	}
	if c.IsSet("all") || c.IsSet("nolh") {
		fmt.Println("nolh:", config.Nolh)
	}
	if c.IsSet("all") || c.IsSet("callhome") {
		fmt.Println("callhome:", config.Callhome)
	}
	if c.IsSet("all") || c.IsSet("bootstrap") {
		fmt.Println("bootstrap:", config.Bootstrap)
	}
	if c.IsSet("all") || c.IsSet("tunnel_end_point") {
		fmt.Println("tunnel_end_point:", config.TunnelEndPoint)
	}
	if c.IsSet("all") || c.IsSet("tunnel_certs") {
		fmt.Println("tunnel_certs:", config.TunnelCerts)
	}
	if c.IsSet("all") || c.IsSet("driver") {
		fmt.Println("driver:", config.Driver)
	}
	if c.IsSet("all") || c.IsSet("debug_level") {
		fmt.Println("debug_level:", config.DebugLevel)
	}
	if c.IsSet("all") || c.IsSet("domain") {
		fmt.Println("domain:", config.Domain)
	}
	return nil
}

func setNodeValues(c *cli.Context) error {
	if !c.IsSet("node_id") {
		err := errors.New("--node_id must be set")
		logrus.Error(err)
		return err
	}
	config, err := clusterManager.GetNodeConf(c.String("node_id"))
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.IsSet("node_id") {
		config.NodeId = c.String("node_id")
	}
	return clusterManager.SetNodeConf(config)
}

func showNodeValues(c *cli.Context) error {
	if !c.Parent().IsSet("node_id") {
		err := errors.New("--node_id must be set")
		logrus.Error(err)
		return err
	}
	config, err := clusterManager.GetNodeConf(c.Parent().String("node_id"))
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config)
	}
	if c.IsSet("all") || c.IsSet("node_id") {
		fmt.Println("node_id:", config.NodeId)
	}
	return nil
}

func setNetworkValues(c *cli.Context) error {
	if !c.Parent().IsSet("node_id") {
		err := errors.New("--node_id must be set")
		logrus.Error(err)
		return err
	}
	config, err := clusterManager.GetNodeConf(c.Parent().String("node_id"))
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Network == nil {
		err := errors.New("config.Network" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.IsSet("mgt_iface") {
		config.Network.MgtIface = c.String("mgt_iface")
	}
	if c.IsSet("data_iface") {
		config.Network.DataIface = c.String("data_iface")
	}
	return clusterManager.SetNodeConf(config)
}

func showNetworkValues(c *cli.Context) error {
	if !c.Parent().Parent().IsSet("node_id") {
		err := errors.New("--node_id must be set")
		logrus.Error(err)
		return err
	}
	config, err := clusterManager.GetNodeConf(c.Parent().Parent().String("node_id"))
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Network == nil {
		err := errors.New("config.Network" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config.Network)
	}
	if c.IsSet("all") || c.IsSet("mgt_iface") {
		fmt.Println("mgt_iface:", config.Network.MgtIface)
	}
	if c.IsSet("all") || c.IsSet("data_iface") {
		fmt.Println("data_iface:", config.Network.DataIface)
	}
	return nil
}

func setStorageValues(c *cli.Context) error {
	if !c.Parent().IsSet("node_id") {
		err := errors.New("--node_id must be set")
		logrus.Error(err)
		return err
	}
	config, err := clusterManager.GetNodeConf(c.Parent().String("node_id"))
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Storage == nil {
		err := errors.New("config.Storage" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.IsSet("devices_md") {
		config.Storage.DevicesMd = c.StringSlice("devices_md")
	}
	if c.IsSet("devices") {
		config.Storage.Devices = c.StringSlice("devices")
	}
	if c.IsSet("raid_level") {
		config.Storage.RaidLevel = c.String("raid_level")
	}
	if c.IsSet("raid_level_md") {
		config.Storage.RaidLevelMd = c.String("raid_level_md")
	}
	if c.IsSet("async_io") {
		config.Storage.AsyncIo = c.Bool("async_io")
	}
	return clusterManager.SetNodeConf(config)
}

func showStorageValues(c *cli.Context) error {
	if !c.Parent().Parent().IsSet("node_id") {
		err := errors.New("--node_id must be set")
		logrus.Error(err)
		return err
	}
	config, err := clusterManager.GetNodeConf(c.Parent().Parent().String("node_id"))
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Storage == nil {
		err := errors.New("config.Storage" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config.Storage)
	}
	if c.IsSet("all") || c.IsSet("devices_md") {
		fmt.Println("devices_md:", config.Storage.DevicesMd)
	}
	if c.IsSet("all") || c.IsSet("devices") {
		fmt.Println("devices:", config.Storage.Devices)
	}
	if c.IsSet("all") || c.IsSet("raid_level") {
		fmt.Println("raid_level:", config.Storage.RaidLevel)
	}
	if c.IsSet("all") || c.IsSet("raid_level_md") {
		fmt.Println("raid_level_md:", config.Storage.RaidLevelMd)
	}
	if c.IsSet("all") || c.IsSet("async_io") {
		fmt.Println("async_io:", config.Storage.AsyncIo)
	}
	return nil
}

func setSecretsValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets == nil {
		err := errors.New("config.Secrets" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.IsSet("secret_type") {
		config.Secrets.SecretType = c.String("secret_type")
	}
	if c.IsSet("cluster_secret_key") {
		config.Secrets.ClusterSecretKey = c.String("cluster_secret_key")
	}
	return clusterManager.SetClusterConf(config)
}

func showSecretsValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets == nil {
		err := errors.New("config.Secrets" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config.Secrets)
	}
	if c.IsSet("all") || c.IsSet("secret_type") {
		fmt.Println("secret_type:", config.Secrets.SecretType)
	}
	if c.IsSet("all") || c.IsSet("cluster_secret_key") {
		fmt.Println("cluster_secret_key:", config.Secrets.ClusterSecretKey)
	}
	return nil
}

func setVaultValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets == nil {
		err := errors.New("config.Secrets" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets.Vault == nil {
		err := errors.New("config.Secrets.Vault" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.IsSet("vault_token") {
		config.Secrets.Vault.VaultToken = c.String("vault_token")
	}
	if c.IsSet("vault_addr") {
		config.Secrets.Vault.VaultAddr = c.String("vault_addr")
	}
	if c.IsSet("vault_cacert") {
		config.Secrets.Vault.VaultCacert = c.String("vault_cacert")
	}
	if c.IsSet("vault_capath") {
		config.Secrets.Vault.VaultCapath = c.String("vault_capath")
	}
	if c.IsSet("vault_client_cert") {
		config.Secrets.Vault.VaultClientCert = c.String("vault_client_cert")
	}
	if c.IsSet("vault_client_key") {
		config.Secrets.Vault.VaultClientKey = c.String("vault_client_key")
	}
	if c.IsSet("vault_skip_verify") {
		config.Secrets.Vault.VaultSkipVerify = c.String("vault_skip_verify")
	}
	if c.IsSet("vault_tls_server_name") {
		config.Secrets.Vault.VaultTlsServerName = c.String("vault_tls_server_name")
	}
	if c.IsSet("vault_base_path") {
		config.Secrets.Vault.VaultBasePath = c.String("vault_base_path")
	}
	return clusterManager.SetClusterConf(config)
}

func showVaultValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets == nil {
		err := errors.New("config.Secrets" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets.Vault == nil {
		err := errors.New("config.Secrets.Vault" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config.Secrets.Vault)
	}
	if c.IsSet("all") || c.IsSet("vault_token") {
		fmt.Println("vault_token:", config.Secrets.Vault.VaultToken)
	}
	if c.IsSet("all") || c.IsSet("vault_addr") {
		fmt.Println("vault_addr:", config.Secrets.Vault.VaultAddr)
	}
	if c.IsSet("all") || c.IsSet("vault_cacert") {
		fmt.Println("vault_cacert:", config.Secrets.Vault.VaultCacert)
	}
	if c.IsSet("all") || c.IsSet("vault_capath") {
		fmt.Println("vault_capath:", config.Secrets.Vault.VaultCapath)
	}
	if c.IsSet("all") || c.IsSet("vault_client_cert") {
		fmt.Println("vault_client_cert:", config.Secrets.Vault.VaultClientCert)
	}
	if c.IsSet("all") || c.IsSet("vault_client_key") {
		fmt.Println("vault_client_key:", config.Secrets.Vault.VaultClientKey)
	}
	if c.IsSet("all") || c.IsSet("vault_skip_verify") {
		fmt.Println("vault_skip_verify:", config.Secrets.Vault.VaultSkipVerify)
	}
	if c.IsSet("all") || c.IsSet("vault_tls_server_name") {
		fmt.Println("vault_tls_server_name:", config.Secrets.Vault.VaultTlsServerName)
	}
	if c.IsSet("all") || c.IsSet("vault_base_path") {
		fmt.Println("vault_base_path:", config.Secrets.Vault.VaultBasePath)
	}
	return nil
}

func setAwsValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets == nil {
		err := errors.New("config.Secrets" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets.Aws == nil {
		err := errors.New("config.Secrets.Aws" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.IsSet("aws_access_key_id") {
		config.Secrets.Aws.AwsAccessKeyId = c.String("aws_access_key_id")
	}
	if c.IsSet("aws_secret_access_key") {
		config.Secrets.Aws.AwsSecretAccessKey = c.String("aws_secret_access_key")
	}
	if c.IsSet("aws_secret_token_key") {
		config.Secrets.Aws.AwsSecretTokenKey = c.String("aws_secret_token_key")
	}
	if c.IsSet("aws_cmk") {
		config.Secrets.Aws.AwsCmk = c.String("aws_cmk")
	}
	if c.IsSet("aws_region") {
		config.Secrets.Aws.AwsRegion = c.String("aws_region")
	}
	return clusterManager.SetClusterConf(config)
}

func showAwsValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if config == nil {
		err := errors.New("config" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets == nil {
		err := errors.New("config.Secrets" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}
	if config.Secrets.Aws == nil {
		err := errors.New("config.Secrets.Aws" + ": no data found, received nil pointer")
		logrus.Error(err)
		return err
	}

	if c.GlobalBool("json") {
		return printJson(config.Secrets.Aws)
	}
	if c.IsSet("all") || c.IsSet("aws_access_key_id") {
		fmt.Println("aws_access_key_id:", config.Secrets.Aws.AwsAccessKeyId)
	}
	if c.IsSet("all") || c.IsSet("aws_secret_access_key") {
		fmt.Println("aws_secret_access_key:", config.Secrets.Aws.AwsSecretAccessKey)
	}
	if c.IsSet("all") || c.IsSet("aws_secret_token_key") {
		fmt.Println("aws_secret_token_key:", config.Secrets.Aws.AwsSecretTokenKey)
	}
	if c.IsSet("all") || c.IsSet("aws_cmk") {
		fmt.Println("aws_cmk:", config.Secrets.Aws.AwsCmk)
	}
	if c.IsSet("all") || c.IsSet("aws_region") {
		fmt.Println("aws_region:", config.Secrets.Aws.AwsRegion)
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
