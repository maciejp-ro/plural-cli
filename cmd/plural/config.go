package main

import (
	"github.com/pluralsh/plural/pkg/config"
	"github.com/urfave/cli"
	"github.com/olekukonko/tablewriter"
	"os"
)

func configCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "amend",
			Usage:     "modify config",
			ArgsUsage: "[key] [value]",
			Action:    handleAmend,
		},
		{
			Name:      "read",
			Usage:     "dumps config",
			ArgsUsage: "",
			Action:    handleRead,
		},
	}
}

func profileCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "use",
			Usage:     "moves the config in PROFILE to the current config",
			ArgsUsage: "PROFILE",
			Action:    handleUseProfile,
		},
		{
			Name:      "save",
			Usage:     "saves the current config as PROFILE",
			ArgsUsage: "PROFILE",
			Action:    handleSaveProfile,
		},
		{
			Name:      "list",
			Usage:     "lists all saved profiles",
			ArgsUsage: "",
			Action:    listProfiles,
		},
	}
}

func handleAmend(c *cli.Context) error {
	return config.Amend(c.Args().Get(0), c.Args().Get(1))
}

func handleRead(c *cli.Context) error {
	conf := config.Read()
	d, err := conf.Marshal()
	if err != nil {
		return err
	}

	os.Stdout.Write(d)
	return nil
}

func handleUseProfile(c *cli.Context) error {
	return config.Profile(c.Args().Get(0))
}

func handleSaveProfile(c *cli.Context) error {
	conf := config.Read()
	return conf.SaveProfile(c.Args().Get(0))
}

func listProfiles(c *cli.Context) error {
	profiles, err := config.Profiles()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Email", "Endpoint"})
	for _, profile := range profiles {
		table.Append([]string{profile.Metadata.Name, profile.Spec.Email, profile.Spec.BaseUrl()})
	}

	table.Render()
	return nil
}