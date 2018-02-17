package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{

		{
			Name: "config",
			Usage: "To be added",
			Description: "To be added",
			Subcommands: []cli.Command{
				{
					Name: "Secrets",
					Usage: "To be added",
					Description: "To be added",
					Subcommands: []cli.Command{
						{
							Name: "Vault",
							Usage: "To be added",
							Description: "To be added",
							Subcommands: []cli.Command{
								{
									Name: "get",
									Usage: "Get values",
									Description: "Get values",
									Action: get_Vault_values,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name: "VaultToken",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "VaultAddr",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "VaultCacert",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "VaultCapath",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "VaultClientCert",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "VaultClientKey",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "VaultSkipVerify",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "VaultTlsServerName",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "VaultBasePath",
											Usage: "(Bool)\tTo be added",
										},
									},
								},
								{
									Name: "set",
									Usage: "Set values",
									Description: "Set values",
									Action: set_Vault_values,
									Flags: []cli.Flag{
										cli.StringFlag{
											Name: "VaultToken",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "VaultAddr",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "VaultCacert",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "VaultCapath",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "VaultClientCert",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "VaultClientKey",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "VaultSkipVerify",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "VaultTlsServerName",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "VaultBasePath",
											Usage: "(Str)\tTo be added",
										},
									},
								},
							},
						},
						{
							Name: "Aws",
							Usage: "To be added",
							Description: "To be added",
							Subcommands: []cli.Command{
								{
									Name: "get",
									Usage: "Get values",
									Description: "Get values",
									Action: get_Aws_values,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name: "AwsAccessKeyId",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "AwsSecretAccessKey",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "AwsSecretTokenKey",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "AwsCmk",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "AwsRegion",
											Usage: "(Bool)\tTo be added",
										},
									},
								},
								{
									Name: "set",
									Usage: "Set values",
									Description: "Set values",
									Action: set_Aws_values,
									Flags: []cli.Flag{
										cli.StringFlag{
											Name: "AwsAccessKeyId",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "AwsSecretAccessKey",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "AwsSecretTokenKey",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "AwsCmk",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "AwsRegion",
											Usage: "(Str)\tTo be added",
										},
									},
								},
							},
						},
						{
							Name: "get",
							Usage: "Get values",
							Description: "Get values",
							Action: get_Secrets_values,
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name: "SecretType",
									Usage: "(Bool)\tTo be added",
								},
								cli.BoolFlag{
									Name: "ClusterSecretKey",
									Usage: "(Bool)\tTo be added",
								},
							},
						},
						{
							Name: "set",
							Usage: "Set values",
							Description: "Set values",
							Action: set_Secrets_values,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name: "SecretType",
									Usage: "(Str)\tTo be added",
								},
								cli.StringFlag{
									Name: "ClusterSecretKey",
									Usage: "(Str)\tTo be added",
								},
							},
						},
					},
				},
				{
					Name: "Kvdb",
					Usage: "To be added",
					Description: "To be added",
					Subcommands: []cli.Command{
						{
							Name: "get",
							Usage: "Get values",
							Description: "Get values",
							Action: get_Kvdb_values,
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name: "Username",
									Usage: "(Bool)\tTo be added",
								},
								cli.BoolFlag{
									Name: "Password",
									Usage: "(Bool)\tTo be added",
								},
								cli.BoolFlag{
									Name: "CaFile",
									Usage: "(Bool)\tTo be added",
								},
								cli.BoolFlag{
									Name: "CertFile",
									Usage: "(Bool)\tTo be added",
								},
								cli.BoolFlag{
									Name: "TrustedCaFile",
									Usage: "(Bool)\tTo be added",
								},
								cli.BoolFlag{
									Name: "ClientCertAuth",
									Usage: "(Bool)\tTo be added",
								},
								cli.BoolFlag{
									Name: "AclToken",
									Usage: "(Bool)\tTo be added",
								},
								cli.BoolFlag{
									Name: "KvdbAddr",
									Usage: "(Bool)\tTo be added",
								},
							},
						},
						{
							Name: "set",
							Usage: "Set values",
							Description: "Set values",
							Action: set_Kvdb_values,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name: "Username",
									Usage: "(Str)\tTo be added",
								},
								cli.StringFlag{
									Name: "Password",
									Usage: "(Str)\tTo be added",
								},
								cli.StringFlag{
									Name: "CaFile",
									Usage: "(Str)\tTo be added",
								},
								cli.StringFlag{
									Name: "CertFile",
									Usage: "(Str)\tTo be added",
								},
								cli.StringFlag{
									Name: "TrustedCaFile",
									Usage: "(Str)\tTo be added",
								},
								cli.StringFlag{
									Name: "ClientCertAuth",
									Usage: "(Str)\tTo be added",
								},
								cli.StringFlag{
									Name: "AclToken",
									Usage: "(Str)\tTo be added",
								},
								cli.StringSliceFlag{
									Name: "KvdbAddr",
									Usage: "(Str...)\tTo be added",
								},
							},
						},
					},
				},
				{
					Name: "node",
					Usage: "To be added",
					Description: "To be added",
					Subcommands: []cli.Command{
						{
							Name: "Network",
							Usage: "To be added",
							Description: "To be added",
							Subcommands: []cli.Command{
								{
									Name: "get",
									Usage: "Get values",
									Description: "Get values",
									Action: get_Network_values,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name: "MgtIface",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "DataIface",
											Usage: "(Bool)\tTo be added",
										},
									},
								},
								{
									Name: "set",
									Usage: "Set values",
									Description: "Set values",
									Action: set_Network_values,
									Flags: []cli.Flag{
										cli.StringFlag{
											Name: "MgtIface",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "DataIface",
											Usage: "(Str)\tTo be added",
										},
									},
								},
							},
						},
						{
							Name: "Storage",
							Usage: "To be added",
							Description: "To be added",
							Subcommands: []cli.Command{
								{
									Name: "get",
									Usage: "Get values",
									Description: "Get values",
									Action: get_Storage_values,
									Flags: []cli.Flag{
										cli.BoolFlag{
											Name: "DevicesMd",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "Devices",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "RaidLevel",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "RaidLevelMd",
											Usage: "(Bool)\tTo be added",
										},
										cli.BoolFlag{
											Name: "AsyncIo",
											Usage: "(Bool)\tTo be added",
										},
									},
								},
								{
									Name: "set",
									Usage: "Set values",
									Description: "Set values",
									Action: set_Storage_values,
									Flags: []cli.Flag{
										cli.StringSliceFlag{
											Name: "DevicesMd",
											Usage: "(Str...)\tTo be added",
										},
										cli.StringSliceFlag{
											Name: "Devices",
											Usage: "(Str...)\tTo be added",
										},
										cli.StringFlag{
											Name: "RaidLevel",
											Usage: "(Str)\tTo be added",
										},
										cli.StringFlag{
											Name: "RaidLevelMd",
											Usage: "(Str)\tTo be added",
										},
										cli.BoolFlag{
											Name: "AsyncIo",
											Usage: "(Bool)\tTo be added",
										},
									},
								},
							},
						},
						{
							Name: "get",
							Usage: "Get values",
							Description: "Get values",
							Action: get_node_values,
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name: "NodeId",
									Usage: "(Bool)\tTo be added",
								},
							},
						},
						{
							Name: "set",
							Usage: "Set values",
							Description: "Set values",
							Action: set_node_values,
							Flags: []cli.Flag{
								cli.StringFlag{
									Name: "NodeId",
									Usage: "(Str)\tTo be added",
								},
							},
						},
					},
					Flags: []cli.Flag{
						cli.StringSliceFlag{
							Name: "id",
							Usage: "(Str...)\tNode id",
						},
					},
				},
				{
					Name: "get",
					Usage: "Get values",
					Description: "Get values",
					Action: get_config_values,
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "Description",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Mode",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Version",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Created",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "ClusterId",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "LoggingUrl",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "AlertingUrl",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Scheduler",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Multicontainer",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Nolh",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Callhome",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Bootstrap",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "TunnelEndPoint",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "TunnelCerts",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Driver",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "DebugLevel",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Domain",
							Usage: "(Bool)\tTo be added",
						},
					},
				},
				{
					Name: "set",
					Usage: "Set values",
					Description: "Set values",
					Action: set_config_values,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "Description",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "Mode",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "Version",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "Created",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "ClusterId",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "LoggingUrl",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "AlertingUrl",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "Scheduler",
							Usage: "(Str)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Multicontainer",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Nolh",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Callhome",
							Usage: "(Bool)\tTo be added",
						},
						cli.BoolFlag{
							Name: "Bootstrap",
							Usage: "(Bool)\tTo be added",
						},
						cli.StringFlag{
							Name: "TunnelEndPoint",
							Usage: "(Str)\tTo be added",
						},
						cli.StringSliceFlag{
							Name: "TunnelCerts",
							Usage: "(Str...)\tTo be added",
						},
						cli.StringFlag{
							Name: "Driver",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "DebugLevel",
							Usage: "(Str)\tTo be added",
						},
						cli.StringFlag{
							Name: "Domain",
							Usage: "(Str)\tTo be added",
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
func get_Secrets_values (c *cli.Context) {}
func set_Secrets_values (c *cli.Context) {}
func get_Kvdb_values (c *cli.Context) {}
func set_Kvdb_values (c *cli.Context) {}
func get_Network_values (c *cli.Context) {}
func set_Network_values (c *cli.Context) {}
func get_Storage_values (c *cli.Context) {}
func set_Storage_values (c *cli.Context) {}
func get_node_values (c *cli.Context) {}
func set_node_values (c *cli.Context) {}
func get_config_values (c *cli.Context) {}
func set_config_values (c *cli.Context) {}
func get_Vault_values (c *cli.Context) {}
func set_Vault_values (c *cli.Context) {}
func get_Aws_values (c *cli.Context) {}
func set_Aws_values (c *cli.Context) {}