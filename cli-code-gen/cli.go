package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
func (m *manager) GetNodeConf() (*osdconfig.NodeConfig, error) {
	config := new(osdconfig.NodeConfig)
	if b, err := ioutil.ReadFile("/tmp/nodeConfig.json"); err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal(b, config); err != nil {
			return nil, err
		}
	}
	return config, nil
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
	if b, err := json.Marshal(config); err != nil {
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
			Usage:       "usage",
			Description: "description",
			Hidden:      false,
			Action:      setConfigValues,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "description",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "mode",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "version",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "created",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "cluster_id",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "logging_url",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "alerting_url",
					Usage:  "(Str)\tusage to be added",
					Hidden: false,
				},
				cli.StringFlag{
					Name:   "scheduler",
					Usage:  "(Str)\tusage to be added",
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
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "mode",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "version",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "created",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "cluster_id",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "logging_url",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "alerting_url",
							Usage:  "(Bool)\tusage to be added",
							Hidden: false,
						},
						cli.BoolFlag{
							Name:   "scheduler",
							Usage:  "(Bool)\tusage to be added",
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
							Usage:  "(Str)\tusage to be added",
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
									Usage:  "(Bool)\tusage to be added",
									Hidden: false,
								},
							},
						},
						{
							Name:        "network",
							Usage:       "usage to be added",
							Description: "description to be added",
							Hidden:      false,
							Action:      setNetworkValues,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:   "mgt_iface",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "data_iface",
									Usage:  "(Str)\tusage to be added",
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
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "data_iface",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
									},
								},
							},
						},
						{
							Name:        "storage",
							Usage:       "usage to be added",
							Description: "description to be added",
							Hidden:      false,
							Action:      setStorageValues,
							Flags: []cli.Flag{
								cli.StringSliceFlag{
									Name:   "devices_md",
									Usage:  "(Str...)\tusage to be added",
									Hidden: false,
								},
								cli.StringSliceFlag{
									Name:   "devices",
									Usage:  "(Str...)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "raid_level",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "raid_level_md",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "async_io",
									Usage:  "(Bool)\tusage to be added",
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
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "devices",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "raid_level",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "raid_level_md",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "async_io",
											Usage:  "(Bool)\tusage to be added",
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
							Usage:  "(Str)\tusage to be added",
							Hidden: false,
						},
						cli.StringFlag{
							Name:   "cluster_secret_key",
							Usage:  "(Str)\tusage to be added",
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
									Usage:  "(Bool)\tusage to be added",
									Hidden: false,
								},
								cli.BoolFlag{
									Name:   "cluster_secret_key",
									Usage:  "(Bool)\tusage to be added",
									Hidden: false,
								},
							},
						},
						{
							Name:        "vault",
							Usage:       "usage to be added",
							Description: "none yet",
							Hidden:      false,
							Action:      setVaultValues,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:   "vault_token",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_addr",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_cacert",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_capath",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_client_cert",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_client_key",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_skip_verify",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_tls_server_name",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "vault_base_path",
									Usage:  "(Str)\tusage to be added",
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
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_addr",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_cacert",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_capath",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_client_cert",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_client_key",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_skip_verify",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_tls_server_name",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "vault_base_path",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
									},
								},
							},
						},
						{
							Name:        "aws",
							Usage:       "usage to be added",
							Description: "none yet",
							Hidden:      false,
							Action:      setAwsValues,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:   "aws_access_key_id",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "aws_secret_access_key",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "aws_secret_token_key",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "aws_cmk",
									Usage:  "(Str)\tusage to be added",
									Hidden: false,
								},
								cli.StringFlag{
									Name:   "aws_region",
									Usage:  "(Str)\tusage to be added",
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
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_secret_access_key",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_secret_token_key",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_cmk",
											Usage:  "(Bool)\tusage to be added",
											Hidden: false,
										},
										cli.BoolFlag{
											Name:   "aws_region",
											Usage:  "(Bool)\tusage to be added",
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
		return err
	}
	if c.IsSet("description") {
		description := c.String("description")
		config.Description = description
	}
	if c.IsSet("mode") {
		mode := c.String("mode")
		config.Mode = mode
	}
	if c.IsSet("version") {
		version := c.String("version")
		config.Version = version
	}
	if c.IsSet("created") {
		created := c.String("created")
		config.Created = created
	}
	if c.IsSet("cluster_id") {
		clusterId := c.String("cluster_id")
		config.ClusterId = clusterId
	}
	if c.IsSet("logging_url") {
		loggingUrl := c.String("logging_url")
		config.LoggingUrl = loggingUrl
	}
	if c.IsSet("alerting_url") {
		alertingUrl := c.String("alerting_url")
		config.AlertingUrl = alertingUrl
	}
	if c.IsSet("scheduler") {
		scheduler := c.String("scheduler")
		config.Scheduler = scheduler
	}
	if c.IsSet("multicontainer") {
		multicontainer := c.Bool("multicontainer")
		config.Multicontainer = multicontainer
	}
	if c.IsSet("nolh") {
		nolh := c.Bool("nolh")
		config.Nolh = nolh
	}
	if c.IsSet("callhome") {
		callhome := c.Bool("callhome")
		config.Callhome = callhome
	}
	if c.IsSet("bootstrap") {
		bootstrap := c.Bool("bootstrap")
		config.Bootstrap = bootstrap
	}
	if c.IsSet("tunnel_end_point") {
		tunnelEndPoint := c.String("tunnel_end_point")
		config.TunnelEndPoint = tunnelEndPoint
	}
	if c.IsSet("tunnel_certs") {
		tunnelCerts := c.StringSlice("tunnel_certs")
		config.TunnelCerts = tunnelCerts
	}
	if c.IsSet("driver") {
		driver := c.String("driver")
		config.Driver = driver
	}
	if c.IsSet("debug_level") {
		debugLevel := c.String("debug_level")
		config.DebugLevel = debugLevel
	}
	if c.IsSet("domain") {
		domain := c.String("domain")
		config.Domain = domain
	}
	return clusterManager.SetClusterConf(config)
}

func showConfigValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		return err
	}
	if c.GlobalBool("json") {
		return printJson(config)
	}
	if c.IsSet("all") || c.IsSet("description") {
		fmt.Println("Description:", config.Description)
	}
	if c.IsSet("all") || c.IsSet("mode") {
		fmt.Println("Mode:", config.Mode)
	}
	if c.IsSet("all") || c.IsSet("version") {
		fmt.Println("Version:", config.Version)
	}
	if c.IsSet("all") || c.IsSet("created") {
		fmt.Println("Created:", config.Created)
	}
	if c.IsSet("all") || c.IsSet("cluster_id") {
		fmt.Println("ClusterId:", config.ClusterId)
	}
	if c.IsSet("all") || c.IsSet("logging_url") {
		fmt.Println("LoggingUrl:", config.LoggingUrl)
	}
	if c.IsSet("all") || c.IsSet("alerting_url") {
		fmt.Println("AlertingUrl:", config.AlertingUrl)
	}
	if c.IsSet("all") || c.IsSet("scheduler") {
		fmt.Println("Scheduler:", config.Scheduler)
	}
	if c.IsSet("all") || c.IsSet("multicontainer") {
		fmt.Println("Multicontainer:", config.Multicontainer)
	}
	if c.IsSet("all") || c.IsSet("nolh") {
		fmt.Println("Nolh:", config.Nolh)
	}
	if c.IsSet("all") || c.IsSet("callhome") {
		fmt.Println("Callhome:", config.Callhome)
	}
	if c.IsSet("all") || c.IsSet("bootstrap") {
		fmt.Println("Bootstrap:", config.Bootstrap)
	}
	if c.IsSet("all") || c.IsSet("tunnel_end_point") {
		fmt.Println("TunnelEndPoint:", config.TunnelEndPoint)
	}
	if c.IsSet("all") || c.IsSet("tunnel_certs") {
		fmt.Println("TunnelCerts:", config.TunnelCerts)
	}
	if c.IsSet("all") || c.IsSet("driver") {
		fmt.Println("Driver:", config.Driver)
	}
	if c.IsSet("all") || c.IsSet("debug_level") {
		fmt.Println("DebugLevel:", config.DebugLevel)
	}
	if c.IsSet("all") || c.IsSet("domain") {
		fmt.Println("Domain:", config.Domain)
	}
	return nil
}

func setNodeValues(c *cli.Context) error {
	config, err := clusterManager.GetNodeConf()
	if err != nil {
		return err
	}
	if c.IsSet("node_id") {
		nodeId := c.String("node_id")
		config.NodeId = nodeId
	}
	return clusterManager.SetNodeConf(config)
}

func showNodeValues(c *cli.Context) error {
	config, err := clusterManager.GetNodeConf()
	if err != nil {
		return err
	}
	if c.GlobalBool("json") {
		return printJson(config)
	}
	if c.IsSet("all") || c.IsSet("node_id") {
		fmt.Println("NodeId:", config.NodeId)
	}
	return nil
}

func setNetworkValues(c *cli.Context) error {
	config, err := clusterManager.GetNodeConf()
	if err != nil {
		return err
	}
	if c.IsSet("mgt_iface") {
		mgtIface := c.String("mgt_iface")
		config.Network.MgtIface = mgtIface
	}
	if c.IsSet("data_iface") {
		dataIface := c.String("data_iface")
		config.Network.DataIface = dataIface
	}
	return clusterManager.SetNodeConf(config)
}

func showNetworkValues(c *cli.Context) error {
	config, err := clusterManager.GetNodeConf()
	if err != nil {
		return err
	}
	if c.GlobalBool("json") {
		return printJson(config.Network)
	}
	if c.IsSet("all") || c.IsSet("mgt_iface") {
		fmt.Println("MgtIface:", config.Network.MgtIface)
	}
	if c.IsSet("all") || c.IsSet("data_iface") {
		fmt.Println("DataIface:", config.Network.DataIface)
	}
	return nil
}

func setStorageValues(c *cli.Context) error {
	config, err := clusterManager.GetNodeConf()
	if err != nil {
		return err
	}
	if c.IsSet("devices_md") {
		devicesMd := c.StringSlice("devices_md")
		config.Storage.DevicesMd = devicesMd
	}
	if c.IsSet("devices") {
		devices := c.StringSlice("devices")
		config.Storage.Devices = devices
	}
	if c.IsSet("raid_level") {
		raidLevel := c.String("raid_level")
		config.Storage.RaidLevel = raidLevel
	}
	if c.IsSet("raid_level_md") {
		raidLevelMd := c.String("raid_level_md")
		config.Storage.RaidLevelMd = raidLevelMd
	}
	if c.IsSet("async_io") {
		asyncIo := c.Bool("async_io")
		config.Storage.AsyncIo = asyncIo
	}
	return clusterManager.SetNodeConf(config)
}

func showStorageValues(c *cli.Context) error {
	config, err := clusterManager.GetNodeConf()
	if err != nil {
		return err
	}
	if c.GlobalBool("json") {
		return printJson(config.Storage)
	}
	if c.IsSet("all") || c.IsSet("devices_md") {
		fmt.Println("DevicesMd:", config.Storage.DevicesMd)
	}
	if c.IsSet("all") || c.IsSet("devices") {
		fmt.Println("Devices:", config.Storage.Devices)
	}
	if c.IsSet("all") || c.IsSet("raid_level") {
		fmt.Println("RaidLevel:", config.Storage.RaidLevel)
	}
	if c.IsSet("all") || c.IsSet("raid_level_md") {
		fmt.Println("RaidLevelMd:", config.Storage.RaidLevelMd)
	}
	if c.IsSet("all") || c.IsSet("async_io") {
		fmt.Println("AsyncIo:", config.Storage.AsyncIo)
	}
	return nil
}

func setSecretsValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		return err
	}
	if c.IsSet("secret_type") {
		secretType := c.String("secret_type")
		config.Secrets.SecretType = secretType
	}
	if c.IsSet("cluster_secret_key") {
		clusterSecretKey := c.String("cluster_secret_key")
		config.Secrets.ClusterSecretKey = clusterSecretKey
	}
	return clusterManager.SetClusterConf(config)
}

func showSecretsValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		return err
	}
	if c.GlobalBool("json") {
		return printJson(config.Secrets)
	}
	if c.IsSet("all") || c.IsSet("secret_type") {
		fmt.Println("SecretType:", config.Secrets.SecretType)
	}
	if c.IsSet("all") || c.IsSet("cluster_secret_key") {
		fmt.Println("ClusterSecretKey:", config.Secrets.ClusterSecretKey)
	}
	return nil
}

func setVaultValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		return err
	}
	if c.IsSet("vault_token") {
		vaultToken := c.String("vault_token")
		config.Secrets.Vault.VaultToken = vaultToken
	}
	if c.IsSet("vault_addr") {
		vaultAddr := c.String("vault_addr")
		config.Secrets.Vault.VaultAddr = vaultAddr
	}
	if c.IsSet("vault_cacert") {
		vaultCacert := c.String("vault_cacert")
		config.Secrets.Vault.VaultCacert = vaultCacert
	}
	if c.IsSet("vault_capath") {
		vaultCapath := c.String("vault_capath")
		config.Secrets.Vault.VaultCapath = vaultCapath
	}
	if c.IsSet("vault_client_cert") {
		vaultClientCert := c.String("vault_client_cert")
		config.Secrets.Vault.VaultClientCert = vaultClientCert
	}
	if c.IsSet("vault_client_key") {
		vaultClientKey := c.String("vault_client_key")
		config.Secrets.Vault.VaultClientKey = vaultClientKey
	}
	if c.IsSet("vault_skip_verify") {
		vaultSkipVerify := c.String("vault_skip_verify")
		config.Secrets.Vault.VaultSkipVerify = vaultSkipVerify
	}
	if c.IsSet("vault_tls_server_name") {
		vaultTlsServerName := c.String("vault_tls_server_name")
		config.Secrets.Vault.VaultTlsServerName = vaultTlsServerName
	}
	if c.IsSet("vault_base_path") {
		vaultBasePath := c.String("vault_base_path")
		config.Secrets.Vault.VaultBasePath = vaultBasePath
	}
	return clusterManager.SetClusterConf(config)
}

func showVaultValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		return err
	}
	if c.GlobalBool("json") {
		return printJson(config.Secrets.Vault)
	}
	if c.IsSet("all") || c.IsSet("vault_token") {
		fmt.Println("VaultToken:", config.Secrets.Vault.VaultToken)
	}
	if c.IsSet("all") || c.IsSet("vault_addr") {
		fmt.Println("VaultAddr:", config.Secrets.Vault.VaultAddr)
	}
	if c.IsSet("all") || c.IsSet("vault_cacert") {
		fmt.Println("VaultCacert:", config.Secrets.Vault.VaultCacert)
	}
	if c.IsSet("all") || c.IsSet("vault_capath") {
		fmt.Println("VaultCapath:", config.Secrets.Vault.VaultCapath)
	}
	if c.IsSet("all") || c.IsSet("vault_client_cert") {
		fmt.Println("VaultClientCert:", config.Secrets.Vault.VaultClientCert)
	}
	if c.IsSet("all") || c.IsSet("vault_client_key") {
		fmt.Println("VaultClientKey:", config.Secrets.Vault.VaultClientKey)
	}
	if c.IsSet("all") || c.IsSet("vault_skip_verify") {
		fmt.Println("VaultSkipVerify:", config.Secrets.Vault.VaultSkipVerify)
	}
	if c.IsSet("all") || c.IsSet("vault_tls_server_name") {
		fmt.Println("VaultTlsServerName:", config.Secrets.Vault.VaultTlsServerName)
	}
	if c.IsSet("all") || c.IsSet("vault_base_path") {
		fmt.Println("VaultBasePath:", config.Secrets.Vault.VaultBasePath)
	}
	return nil
}

func setAwsValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		return err
	}
	if c.IsSet("aws_access_key_id") {
		awsAccessKeyId := c.String("aws_access_key_id")
		config.Secrets.Aws.AwsAccessKeyId = awsAccessKeyId
	}
	if c.IsSet("aws_secret_access_key") {
		awsSecretAccessKey := c.String("aws_secret_access_key")
		config.Secrets.Aws.AwsSecretAccessKey = awsSecretAccessKey
	}
	if c.IsSet("aws_secret_token_key") {
		awsSecretTokenKey := c.String("aws_secret_token_key")
		config.Secrets.Aws.AwsSecretTokenKey = awsSecretTokenKey
	}
	if c.IsSet("aws_cmk") {
		awsCmk := c.String("aws_cmk")
		config.Secrets.Aws.AwsCmk = awsCmk
	}
	if c.IsSet("aws_region") {
		awsRegion := c.String("aws_region")
		config.Secrets.Aws.AwsRegion = awsRegion
	}
	return clusterManager.SetClusterConf(config)
}

func showAwsValues(c *cli.Context) error {
	config, err := clusterManager.GetClusterConf()
	if err != nil {
		return err
	}
	if c.GlobalBool("json") {
		return printJson(config.Secrets.Aws)
	}
	if c.IsSet("all") || c.IsSet("aws_access_key_id") {
		fmt.Println("AwsAccessKeyId:", config.Secrets.Aws.AwsAccessKeyId)
	}
	if c.IsSet("all") || c.IsSet("aws_secret_access_key") {
		fmt.Println("AwsSecretAccessKey:", config.Secrets.Aws.AwsSecretAccessKey)
	}
	if c.IsSet("all") || c.IsSet("aws_secret_token_key") {
		fmt.Println("AwsSecretTokenKey:", config.Secrets.Aws.AwsSecretTokenKey)
	}
	if c.IsSet("all") || c.IsSet("aws_cmk") {
		fmt.Println("AwsCmk:", config.Secrets.Aws.AwsCmk)
	}
	if c.IsSet("all") || c.IsSet("aws_region") {
		fmt.Println("AwsRegion:", config.Secrets.Aws.AwsRegion)
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
