package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/libopenstorage/openstorage/osdconfig"
	"github.com/portworx/porx/px/cli/cflags"
	"github.com/sirupsen/logrus"
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

var clusterManager *manager

func main() {
	clusterManager = new(manager)

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
		for i := 0; i < 3; i++ {
			config := new(osdconfig.NodeConfig).Init()
			config.NodeId = "nodeid_" + strconv.FormatInt(int64(i), 10)
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
	app.Commands = []cli.Command{
		{
			Name:        "config",
			Usage:       "Configure cluster",
			Description: "Configure cluster and nodes",
			Hidden:      true,
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
					Name:   "cluster_id",
					Usage:  "(Str)\tCluster ID info",
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
						cli.BoolFlag{
							Name:  "all, a",
							Usage: "(Bool)\tFor all nodes on cluster",
						},
						cli.StringFlag{
							Name:   "node_id",
							Usage:  "(Str)\tID for the node",
							Hidden: false,
						},
						cli.StringFlag{
							Name:   "csi_endpoint",
							Usage:  "(Str)\tCSI endpoint",
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
								cli.BoolFlag{
									Name:   "csi_endpoint",
									Usage:  "(Bool)\tCSI endpoint",
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
									Name:   "mgt_interface",
									Usage:  "(Str)\tManagement interface",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "data_interface",
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
											Name:   "mgt_interface",
											Usage:  "(Bool)\tManagement interface",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "data_interface",
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
								cli.UintFlag{
									Name:   "max_count",
									Usage:  "(Uint)\tMaximum count",
									Hidden: false,
								},
								cli.UintFlag{
									Name:   "max_drive_set_count",
									Usage:  "(Uint)\tMax drive set count",
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
											Name:   "max_count",
											Usage:  "(Bool)\tMaximum count",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "max_drive_set_count",
											Usage:  "(Bool)\tMax drive set count",
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
									},
								},
							},
						},
						{
							Name:        "geo",
							Usage:       "Geographic configuration",
							Description: "Stores geo info for node",
							Hidden:      false,
							Action:      setGeoValues,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:   "rack",
									Usage:  "(Str)\tRack info",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "zone",
									Usage:  "(Str)\tZone info",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "region",
									Usage:  "(Str)\tRegion info",
									Hidden: false,
								},
							},
							Subcommands: []cli.Command{
								{
									Name:        "show",
									Usage:       "Show values",
									Description: "Show values",
									Action:      showGeoValues,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name:   "all, a",
											Usage:  "(Bool)\tShow all data",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "rack",
											Usage:  "(Bool)\tRack info",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "zone",
											Usage:  "(Bool)\tZone info",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "region",
											Usage:  "(Bool)\tRegion info",
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
									Name:   "token",
									Usage:  "(Str)\tVault token",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "address",
									Usage:  "(Str)\tVault address",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "ca_cert",
									Usage:  "(Str)\tVault CA certificate",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "ca_path",
									Usage:  "(Str)\tVault CA path",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "client_cert",
									Usage:  "(Str)\tVault client certificate",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "client_key",
									Usage:  "(Str)\tVault client key",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "skip_verify",
									Usage:  "(Str)\tVault skip verification",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "tls_server_name",
									Usage:  "(Str)\tVault TLS server name",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "base_path",
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
											Name:   "token",
											Usage:  "(Bool)\tVault token",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "address",
											Usage:  "(Bool)\tVault address",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "ca_cert",
											Usage:  "(Bool)\tVault CA certificate",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "ca_path",
											Usage:  "(Bool)\tVault CA path",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "client_cert",
											Usage:  "(Bool)\tVault client certificate",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "client_key",
											Usage:  "(Bool)\tVault client key",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "skip_verify",
											Usage:  "(Bool)\tVault skip verification",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "tls_server_name",
											Usage:  "(Bool)\tVault TLS server name",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "base_path",
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

// setConfigValues sets config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func setConfigValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return SetConfigValues(provider)
}

// SetConfigValues sets config values using cflags provider.
// This func is autogenerated. Please DO NOT EDIT.
func SetConfigValues(provider cflags.Provider) error {
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

	if set, err := provider.IsSet("description"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("description"); err != nil {
				return err
			} else {
				config.Description = value
			}
		}
	}
	if set, err := provider.IsSet("mode"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("mode"); err != nil {
				return err
			} else {
				config.Mode = value
			}
		}
	}
	if set, err := provider.IsSet("version"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("version"); err != nil {
				return err
			} else {
				config.Version = value
			}
		}
	}
	if set, err := provider.IsSet("cluster_id"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("cluster_id"); err != nil {
				return err
			} else {
				config.ClusterId = value
			}
		}
	}
	if set, err := provider.IsSet("domain"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("domain"); err != nil {
				return err
			} else {
				config.Domain = value
			}
		}
	}
	if err := clusterManager.SetClusterConf(config); err != nil {
		logrus.Error("Set config for cluster")
		return err
	}
	logrus.Info("Set config for cluster")
	return nil
}

// showConfigValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func showConfigValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return ShowConfigValues(provider)
}

// ShowConfigValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func ShowConfigValues(provider cflags.Provider) error {
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

	if jsonOut, err := provider.GetGlobalBool("json"); err != nil {
		return err
	} else {
		if jsonOut {
			return printJson(config)
		}
	}
	var set bool
	var setAll bool
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	set, err = provider.IsSet("description")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("description:", config.Description)
	}
	set, err = provider.IsSet("mode")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("mode:", config.Mode)
	}
	set, err = provider.IsSet("version")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("version:", config.Version)
	}
	set, err = provider.IsSet("created")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("created:", config.Created)
	}
	set, err = provider.IsSet("cluster_id")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("cluster_id:", config.ClusterId)
	}
	set, err = provider.IsSet("domain")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("domain:", config.Domain)
	}
	return nil
}

// setNodeValues sets config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func setNodeValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return SetNodeValues(provider)
}

// SetNodeValues sets config values using cflags provider.
// This func is autogenerated. Please DO NOT EDIT.
func SetNodeValues(provider cflags.Provider) error {
	var setNodeID bool
	var setAll bool
	var err error
	setNodeID, err = provider.IsSet("node_id")
	if err != nil {
		return err
	}
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	if !setNodeID && !setAll {
		err := errors.New("--node_id must be provided or --all must be set")
		logrus.Error(err)
		return err
	}
	configs := new(osdconfig.NodesConfig)
	if setAll {
		configs, err = clusterManager.EnumerateNodeConf()
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		if nodeID, err := provider.GetGlobalString("node_id"); err != nil {
			return err
		} else {
			if config, err := clusterManager.GetNodeConf(nodeID); err != nil {
				logrus.Error(err)
				return err
			} else {
				*configs = append(*configs, config)
			}
		}
	}
	for _, config := range *configs {
		config := config
		if config == nil {
			err := errors.New("config" + ": no data found, received nil pointer")
			logrus.Error(err)
			return err
		}

		if set, err := provider.IsSet("node_id"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("node_id"); err != nil {
					return err
				} else {
					config.NodeId = value
				}
			}
		}
		if set, err := provider.IsSet("csi_endpoint"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("csi_endpoint"); err != nil {
					return err
				} else {
					config.CSIEndpoint = value
				}
			}
		}
		if err := clusterManager.SetNodeConf(config); err != nil {
			logrus.Error("Set config for node: ", config.NodeId)
			return err
		}
		logrus.Info("Set config for node: ", config.NodeId)
	}
	return nil
}

// showNodeValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func showNodeValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return ShowNodeValues(provider)
}

// ShowNodeValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func ShowNodeValues(provider cflags.Provider) error {
	var setNodeID bool
	var setAll bool
	var err error
	setNodeID, err = provider.IsSet("node_id")
	if err != nil {
		return err
	}
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	if !setNodeID && !setAll {
		err := errors.New("--node_id must be provided or --all must be set")
		logrus.Error(err)
		return err
	}
	configs := new(osdconfig.NodesConfig)
	if setAll {
		configs, err = clusterManager.EnumerateNodeConf()
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		if nodeID, err := provider.GetGlobalString("node_id"); err != nil {
			return err
		} else {
			if config, err := clusterManager.GetNodeConf(nodeID); err != nil {
				logrus.Error(err)
				return err
			} else {
				*configs = append(*configs, config)
			}
		}
	}
	for _, config := range *configs {
		if config == nil {
			err := errors.New("config" + ": no data found, received nil pointer")
			logrus.Error(err)
			return err
		}

		if jsonOut, err := provider.GetGlobalBool("json"); err != nil {
			return err
		} else {
			if jsonOut {
				if err := printJson(struct {
					NodeId string      `json:"node_id"`
					Config interface{} `json:"config"`
				}{config.NodeId, config}); err != nil {
					return err
				}
			} else {
				fmt.Println("node_id:", config.NodeId)
			}
			var set bool
			var setAll bool
			setAll, err = provider.IsSet("all")
			if err != nil {
				return err
			}
			set, err = provider.IsSet("node_id")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("node_id:", config.NodeId)
			}
			set, err = provider.IsSet("csi_endpoint")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("csi_endpoint:", config.CSIEndpoint)
			}
			fmt.Println()
		}
	}
	return nil
}

// setNetworkValues sets config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func setNetworkValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return SetNetworkValues(provider)
}

// SetNetworkValues sets config values using cflags provider.
// This func is autogenerated. Please DO NOT EDIT.
func SetNetworkValues(provider cflags.Provider) error {
	var setNodeID bool
	var setAll bool
	var err error
	setNodeID, err = provider.IsSet("node_id")
	if err != nil {
		return err
	}
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	if !setNodeID && !setAll {
		err := errors.New("--node_id must be provided or --all must be set")
		logrus.Error(err)
		return err
	}
	configs := new(osdconfig.NodesConfig)
	if setAll {
		configs, err = clusterManager.EnumerateNodeConf()
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		if nodeID, err := provider.GetGlobalString("node_id"); err != nil {
			return err
		} else {
			if config, err := clusterManager.GetNodeConf(nodeID); err != nil {
				logrus.Error(err)
				return err
			} else {
				*configs = append(*configs, config)
			}
		}
	}
	for _, config := range *configs {
		config := config
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

		if set, err := provider.IsSet("mgt_interface"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("mgt_interface"); err != nil {
					return err
				} else {
					config.Network.MgtIface = value
				}
			}
		}
		if set, err := provider.IsSet("data_interface"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("data_interface"); err != nil {
					return err
				} else {
					config.Network.DataIface = value
				}
			}
		}
		if err := clusterManager.SetNodeConf(config); err != nil {
			logrus.Error("Set config for node: ", config.NodeId)
			return err
		}
		logrus.Info("Set config for node: ", config.NodeId)
	}
	return nil
}

// showNetworkValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func showNetworkValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return ShowNetworkValues(provider)
}

// ShowNetworkValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func ShowNetworkValues(provider cflags.Provider) error {
	var setNodeID bool
	var setAll bool
	var err error
	setNodeID, err = provider.IsSet("node_id")
	if err != nil {
		return err
	}
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	if !setNodeID && !setAll {
		err := errors.New("--node_id must be provided or --all must be set")
		logrus.Error(err)
		return err
	}
	configs := new(osdconfig.NodesConfig)
	if setAll {
		configs, err = clusterManager.EnumerateNodeConf()
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		if nodeID, err := provider.GetGlobalString("node_id"); err != nil {
			return err
		} else {
			if config, err := clusterManager.GetNodeConf(nodeID); err != nil {
				logrus.Error(err)
				return err
			} else {
				*configs = append(*configs, config)
			}
		}
	}
	for _, config := range *configs {
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

		if jsonOut, err := provider.GetGlobalBool("json"); err != nil {
			return err
		} else {
			if jsonOut {
				if err := printJson(struct {
					NodeId string      `json:"node_id"`
					Config interface{} `json:"config"`
				}{config.NodeId, config.Network}); err != nil {
					return err
				}
			} else {
				fmt.Println("node_id:", config.NodeId)
			}
			var set bool
			var setAll bool
			setAll, err = provider.IsSet("all")
			if err != nil {
				return err
			}
			set, err = provider.IsSet("mgt_interface")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("mgt_interface:", config.Network.MgtIface)
			}
			set, err = provider.IsSet("data_interface")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("data_interface:", config.Network.DataIface)
			}
			fmt.Println()
		}
	}
	return nil
}

// setStorageValues sets config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func setStorageValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return SetStorageValues(provider)
}

// SetStorageValues sets config values using cflags provider.
// This func is autogenerated. Please DO NOT EDIT.
func SetStorageValues(provider cflags.Provider) error {
	var setNodeID bool
	var setAll bool
	var err error
	setNodeID, err = provider.IsSet("node_id")
	if err != nil {
		return err
	}
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	if !setNodeID && !setAll {
		err := errors.New("--node_id must be provided or --all must be set")
		logrus.Error(err)
		return err
	}
	configs := new(osdconfig.NodesConfig)
	if setAll {
		configs, err = clusterManager.EnumerateNodeConf()
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		if nodeID, err := provider.GetGlobalString("node_id"); err != nil {
			return err
		} else {
			if config, err := clusterManager.GetNodeConf(nodeID); err != nil {
				logrus.Error(err)
				return err
			} else {
				*configs = append(*configs, config)
			}
		}
	}
	for _, config := range *configs {
		config := config
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

		if set, err := provider.IsSet("devices_md"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetStringSlice("devices_md"); err != nil {
					return err
				} else {
					config.Storage.DevicesMd = value
				}
			}
		}
		if set, err := provider.IsSet("devices"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetStringSlice("devices"); err != nil {
					return err
				} else {
					config.Storage.Devices = value
				}
			}
		}
		if set, err := provider.IsSet("max_count"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetUint("max_count"); err != nil {
					return err
				} else {
					config.Storage.MaxCount = uint32(value)
				}
			}
		}
		if set, err := provider.IsSet("max_drive_set_count"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetUint("max_drive_set_count"); err != nil {
					return err
				} else {
					config.Storage.MaxDriveSetCount = uint32(value)
				}
			}
		}
		if set, err := provider.IsSet("raid_level"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("raid_level"); err != nil {
					return err
				} else {
					config.Storage.RaidLevel = value
				}
			}
		}
		if set, err := provider.IsSet("raid_level_md"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("raid_level_md"); err != nil {
					return err
				} else {
					config.Storage.RaidLevelMd = value
				}
			}
		}
		if err := clusterManager.SetNodeConf(config); err != nil {
			logrus.Error("Set config for node: ", config.NodeId)
			return err
		}
		logrus.Info("Set config for node: ", config.NodeId)
	}
	return nil
}

// showStorageValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func showStorageValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return ShowStorageValues(provider)
}

// ShowStorageValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func ShowStorageValues(provider cflags.Provider) error {
	var setNodeID bool
	var setAll bool
	var err error
	setNodeID, err = provider.IsSet("node_id")
	if err != nil {
		return err
	}
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	if !setNodeID && !setAll {
		err := errors.New("--node_id must be provided or --all must be set")
		logrus.Error(err)
		return err
	}
	configs := new(osdconfig.NodesConfig)
	if setAll {
		configs, err = clusterManager.EnumerateNodeConf()
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		if nodeID, err := provider.GetGlobalString("node_id"); err != nil {
			return err
		} else {
			if config, err := clusterManager.GetNodeConf(nodeID); err != nil {
				logrus.Error(err)
				return err
			} else {
				*configs = append(*configs, config)
			}
		}
	}
	for _, config := range *configs {
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

		if jsonOut, err := provider.GetGlobalBool("json"); err != nil {
			return err
		} else {
			if jsonOut {
				if err := printJson(struct {
					NodeId string      `json:"node_id"`
					Config interface{} `json:"config"`
				}{config.NodeId, config.Storage}); err != nil {
					return err
				}
			} else {
				fmt.Println("node_id:", config.NodeId)
			}
			var set bool
			var setAll bool
			setAll, err = provider.IsSet("all")
			if err != nil {
				return err
			}
			set, err = provider.IsSet("devices_md")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("devices_md:", config.Storage.DevicesMd)
			}
			set, err = provider.IsSet("devices")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("devices:", config.Storage.Devices)
			}
			set, err = provider.IsSet("max_count")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("max_count:", config.Storage.MaxCount)
			}
			set, err = provider.IsSet("max_drive_set_count")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("max_drive_set_count:", config.Storage.MaxDriveSetCount)
			}
			set, err = provider.IsSet("raid_level")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("raid_level:", config.Storage.RaidLevel)
			}
			set, err = provider.IsSet("raid_level_md")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("raid_level_md:", config.Storage.RaidLevelMd)
			}
			fmt.Println()
		}
	}
	return nil
}

// setGeoValues sets config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func setGeoValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return SetGeoValues(provider)
}

// SetGeoValues sets config values using cflags provider.
// This func is autogenerated. Please DO NOT EDIT.
func SetGeoValues(provider cflags.Provider) error {
	var setNodeID bool
	var setAll bool
	var err error
	setNodeID, err = provider.IsSet("node_id")
	if err != nil {
		return err
	}
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	if !setNodeID && !setAll {
		err := errors.New("--node_id must be provided or --all must be set")
		logrus.Error(err)
		return err
	}
	configs := new(osdconfig.NodesConfig)
	if setAll {
		configs, err = clusterManager.EnumerateNodeConf()
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		if nodeID, err := provider.GetGlobalString("node_id"); err != nil {
			return err
		} else {
			if config, err := clusterManager.GetNodeConf(nodeID); err != nil {
				logrus.Error(err)
				return err
			} else {
				*configs = append(*configs, config)
			}
		}
	}
	for _, config := range *configs {
		config := config
		if config == nil {
			err := errors.New("config" + ": no data found, received nil pointer")
			logrus.Error(err)
			return err
		}
		if config.Geo == nil {
			err := errors.New("config.Geo" + ": no data found, received nil pointer")
			logrus.Error(err)
			return err
		}

		if set, err := provider.IsSet("rack"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("rack"); err != nil {
					return err
				} else {
					config.Geo.Rack = value
				}
			}
		}
		if set, err := provider.IsSet("zone"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("zone"); err != nil {
					return err
				} else {
					config.Geo.Zone = value
				}
			}
		}
		if set, err := provider.IsSet("region"); err != nil {
			return err
		} else {
			if set {
				if value, err := provider.GetString("region"); err != nil {
					return err
				} else {
					config.Geo.Region = value
				}
			}
		}
		if err := clusterManager.SetNodeConf(config); err != nil {
			logrus.Error("Set config for node: ", config.NodeId)
			return err
		}
		logrus.Info("Set config for node: ", config.NodeId)
	}
	return nil
}

// showGeoValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func showGeoValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return ShowGeoValues(provider)
}

// ShowGeoValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func ShowGeoValues(provider cflags.Provider) error {
	var setNodeID bool
	var setAll bool
	var err error
	setNodeID, err = provider.IsSet("node_id")
	if err != nil {
		return err
	}
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	if !setNodeID && !setAll {
		err := errors.New("--node_id must be provided or --all must be set")
		logrus.Error(err)
		return err
	}
	configs := new(osdconfig.NodesConfig)
	if setAll {
		configs, err = clusterManager.EnumerateNodeConf()
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		if nodeID, err := provider.GetGlobalString("node_id"); err != nil {
			return err
		} else {
			if config, err := clusterManager.GetNodeConf(nodeID); err != nil {
				logrus.Error(err)
				return err
			} else {
				*configs = append(*configs, config)
			}
		}
	}
	for _, config := range *configs {
		if config == nil {
			err := errors.New("config" + ": no data found, received nil pointer")
			logrus.Error(err)
			return err
		}
		if config.Geo == nil {
			err := errors.New("config.Geo" + ": no data found, received nil pointer")
			logrus.Error(err)
			return err
		}

		if jsonOut, err := provider.GetGlobalBool("json"); err != nil {
			return err
		} else {
			if jsonOut {
				if err := printJson(struct {
					NodeId string      `json:"node_id"`
					Config interface{} `json:"config"`
				}{config.NodeId, config.Geo}); err != nil {
					return err
				}
			} else {
				fmt.Println("node_id:", config.NodeId)
			}
			var set bool
			var setAll bool
			setAll, err = provider.IsSet("all")
			if err != nil {
				return err
			}
			set, err = provider.IsSet("rack")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("rack:", config.Geo.Rack)
			}
			set, err = provider.IsSet("zone")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("zone:", config.Geo.Zone)
			}
			set, err = provider.IsSet("region")
			if err != nil {
				return err
			}
			if set || setAll {
				fmt.Println("region:", config.Geo.Region)
			}
			fmt.Println()
		}
	}
	return nil
}

// setSecretsValues sets config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func setSecretsValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return SetSecretsValues(provider)
}

// SetSecretsValues sets config values using cflags provider.
// This func is autogenerated. Please DO NOT EDIT.
func SetSecretsValues(provider cflags.Provider) error {
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

	if set, err := provider.IsSet("secret_type"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("secret_type"); err != nil {
				return err
			} else {
				config.Secrets.SecretType = value
			}
		}
	}
	if set, err := provider.IsSet("cluster_secret_key"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("cluster_secret_key"); err != nil {
				return err
			} else {
				config.Secrets.ClusterSecretKey = value
			}
		}
	}
	if err := clusterManager.SetClusterConf(config); err != nil {
		logrus.Error("Set config for cluster")
		return err
	}
	logrus.Info("Set config for cluster")
	return nil
}

// showSecretsValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func showSecretsValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return ShowSecretsValues(provider)
}

// ShowSecretsValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func ShowSecretsValues(provider cflags.Provider) error {
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

	if jsonOut, err := provider.GetGlobalBool("json"); err != nil {
		return err
	} else {
		if jsonOut {
			return printJson(config.Secrets)
		}
	}
	var set bool
	var setAll bool
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	set, err = provider.IsSet("secret_type")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("secret_type:", config.Secrets.SecretType)
	}
	set, err = provider.IsSet("cluster_secret_key")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("cluster_secret_key:", config.Secrets.ClusterSecretKey)
	}
	return nil
}

// setVaultValues sets config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func setVaultValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return SetVaultValues(provider)
}

// SetVaultValues sets config values using cflags provider.
// This func is autogenerated. Please DO NOT EDIT.
func SetVaultValues(provider cflags.Provider) error {
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

	if set, err := provider.IsSet("token"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("token"); err != nil {
				return err
			} else {
				config.Secrets.Vault.Token = value
			}
		}
	}
	if set, err := provider.IsSet("address"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("address"); err != nil {
				return err
			} else {
				config.Secrets.Vault.Address = value
			}
		}
	}
	if set, err := provider.IsSet("ca_cert"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("ca_cert"); err != nil {
				return err
			} else {
				config.Secrets.Vault.CACert = value
			}
		}
	}
	if set, err := provider.IsSet("ca_path"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("ca_path"); err != nil {
				return err
			} else {
				config.Secrets.Vault.CAPath = value
			}
		}
	}
	if set, err := provider.IsSet("client_cert"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("client_cert"); err != nil {
				return err
			} else {
				config.Secrets.Vault.ClientCert = value
			}
		}
	}
	if set, err := provider.IsSet("client_key"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("client_key"); err != nil {
				return err
			} else {
				config.Secrets.Vault.ClientKey = value
			}
		}
	}
	if set, err := provider.IsSet("skip_verify"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("skip_verify"); err != nil {
				return err
			} else {
				config.Secrets.Vault.TLSSkipVerify = value
			}
		}
	}
	if set, err := provider.IsSet("tls_server_name"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("tls_server_name"); err != nil {
				return err
			} else {
				config.Secrets.Vault.TLSServerName = value
			}
		}
	}
	if set, err := provider.IsSet("base_path"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("base_path"); err != nil {
				return err
			} else {
				config.Secrets.Vault.BasePath = value
			}
		}
	}
	if err := clusterManager.SetClusterConf(config); err != nil {
		logrus.Error("Set config for cluster")
		return err
	}
	logrus.Info("Set config for cluster")
	return nil
}

// showVaultValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func showVaultValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return ShowVaultValues(provider)
}

// ShowVaultValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func ShowVaultValues(provider cflags.Provider) error {
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

	if jsonOut, err := provider.GetGlobalBool("json"); err != nil {
		return err
	} else {
		if jsonOut {
			return printJson(config.Secrets.Vault)
		}
	}
	var set bool
	var setAll bool
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	set, err = provider.IsSet("token")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("token:", config.Secrets.Vault.Token)
	}
	set, err = provider.IsSet("address")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("address:", config.Secrets.Vault.Address)
	}
	set, err = provider.IsSet("ca_cert")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("ca_cert:", config.Secrets.Vault.CACert)
	}
	set, err = provider.IsSet("ca_path")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("ca_path:", config.Secrets.Vault.CAPath)
	}
	set, err = provider.IsSet("client_cert")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("client_cert:", config.Secrets.Vault.ClientCert)
	}
	set, err = provider.IsSet("client_key")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("client_key:", config.Secrets.Vault.ClientKey)
	}
	set, err = provider.IsSet("skip_verify")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("skip_verify:", config.Secrets.Vault.TLSSkipVerify)
	}
	set, err = provider.IsSet("tls_server_name")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("tls_server_name:", config.Secrets.Vault.TLSServerName)
	}
	set, err = provider.IsSet("base_path")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("base_path:", config.Secrets.Vault.BasePath)
	}
	return nil
}

// setAwsValues sets config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func setAwsValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return SetAwsValues(provider)
}

// SetAwsValues sets config values using cflags provider.
// This func is autogenerated. Please DO NOT EDIT.
func SetAwsValues(provider cflags.Provider) error {
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

	if set, err := provider.IsSet("aws_access_key_id"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("aws_access_key_id"); err != nil {
				return err
			} else {
				config.Secrets.Aws.AccessKeyId = value
			}
		}
	}
	if set, err := provider.IsSet("aws_secret_access_key"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("aws_secret_access_key"); err != nil {
				return err
			} else {
				config.Secrets.Aws.SecretAccessKey = value
			}
		}
	}
	if set, err := provider.IsSet("aws_secret_token_key"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("aws_secret_token_key"); err != nil {
				return err
			} else {
				config.Secrets.Aws.SecretTokenKey = value
			}
		}
	}
	if set, err := provider.IsSet("aws_cmk"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("aws_cmk"); err != nil {
				return err
			} else {
				config.Secrets.Aws.Cmk = value
			}
		}
	}
	if set, err := provider.IsSet("aws_region"); err != nil {
		return err
	} else {
		if set {
			if value, err := provider.GetString("aws_region"); err != nil {
				return err
			} else {
				config.Secrets.Aws.Region = value
			}
		}
	}
	if err := clusterManager.SetClusterConf(config); err != nil {
		logrus.Error("Set config for cluster")
		return err
	}
	logrus.Info("Set config for cluster")
	return nil
}

// showAwsValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func showAwsValues(c *cli.Context) error {
	provider, err := cflags.NewCodeGangstaProvider(c)
	if err != nil {
		return err
	}
	return ShowAwsValues(provider)
}

// ShowAwsValues retrieves config values using codegangsta cli context.
// This func is autogenerated. Please DO NOT EDIT.
func ShowAwsValues(provider cflags.Provider) error {
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

	if jsonOut, err := provider.GetGlobalBool("json"); err != nil {
		return err
	} else {
		if jsonOut {
			return printJson(config.Secrets.Aws)
		}
	}
	var set bool
	var setAll bool
	setAll, err = provider.IsSet("all")
	if err != nil {
		return err
	}
	set, err = provider.IsSet("aws_access_key_id")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("aws_access_key_id:", config.Secrets.Aws.AccessKeyId)
	}
	set, err = provider.IsSet("aws_secret_access_key")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("aws_secret_access_key:", config.Secrets.Aws.SecretAccessKey)
	}
	set, err = provider.IsSet("aws_secret_token_key")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("aws_secret_token_key:", config.Secrets.Aws.SecretTokenKey)
	}
	set, err = provider.IsSet("aws_cmk")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("aws_cmk:", config.Secrets.Aws.Cmk)
	}
	set, err = provider.IsSet("aws_region")
	if err != nil {
		return err
	}
	if set || setAll {
		fmt.Println("aws_region:", config.Secrets.Aws.Region)
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
