package plural

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pkg/browser"
	"github.com/urfave/cli"

	"github.com/pluralsh/plural/pkg/api"
	"github.com/pluralsh/plural/pkg/config"
	"github.com/pluralsh/plural/pkg/crypto"
	"github.com/pluralsh/plural/pkg/manifest"
	"github.com/pluralsh/plural/pkg/provider"
	"github.com/pluralsh/plural/pkg/scm"
	"github.com/pluralsh/plural/pkg/server"
	"github.com/pluralsh/plural/pkg/utils"
	"github.com/pluralsh/plural/pkg/utils/git"
	"github.com/pluralsh/plural/pkg/utils/pathing"
	"github.com/pluralsh/plural/pkg/wkspace"
)

const DemoingErrorMsg = "You're currently running a gcp demo cluster. Spin that down by deleting you shell at https://app.plural.sh/shell before beginning a local installation"

func (p *Plural) handleInit(c *cli.Context) error {
	gitCreated := false
	repo := ""

	git, err := wkspace.Preflight()
	if err != nil && git {
		return err
	}

	if err := handleLogin(c); err != nil {
		return err
	}
	p.InitPluralClient()

	me, err := p.Me()
	if err != nil {
		return api.GetErrorResponse(err, "Me")
	}
	if me.Demoing {
		return fmt.Errorf(DemoingErrorMsg)
	}

	if _, err := os.Stat(manifest.ProjectManifestPath()); err == nil && git && !affirm("This repository's workspace.yaml already exists. Would you like to use it?", "PLURAL_INIT_AFFIRM_CURRENT_REPO") {
		fmt.Println("Run `plural init` from empty repository or outside any in order to start from scratch.")
		return nil
	}

	prov, err := runPreflights()
	if err != nil && !c.Bool("ignore-preflights") {
		return err
	}

	if !git && affirm("you're attempting to setup plural outside a git repository. would you like us to set one up for you here?", "PLURAL_INIT_AFFIRM_SETUP_REPO") {
		repo, err = scm.Setup()
		if err != nil {
			return err
		}
		gitCreated = true
	}
	if !git && !gitCreated {
		return fmt.Errorf("you're not in a git repository, either clone one directly or let us set it up for you by rerunning `plural init`")
	}

	// create workspace.yaml when git repository is ready
	if err := prov.Flush(); err != nil {
		return err
	}
	if err := cryptoInit(c); err != nil {
		return err
	}
	_ = wkspace.DownloadReadme()

	if affirm(backupMsg, "PLURAL_INIT_AFFIRM_BACKUP_KEY") {
		if err := crypto.BackupKey(p.Client); err != nil {
			return api.GetErrorResponse(err, "BackupKey")
		}
	}

	if err := crypto.CreateKeyFingerprintFile(); err != nil {
		return err
	}

	utils.Success("Workspace is properly configured!\n")
	if gitCreated {
		utils.Highlight("Be sure to `cd %s` to use your configured git repo\n", repo)
	}
	return nil
}

func preflights(c *cli.Context) error {
	_, err := runPreflights()
	return err
}

func runPreflights() (provider.Provider, error) {
	prov, err := provider.GetProvider()
	if err != nil {
		return prov, err
	}

	for _, pre := range prov.Preflights() {
		if err := pre.Validate(); err != nil {
			return prov, err
		}
	}

	return prov, nil
}

func handleLogin(c *cli.Context) error {
	conf := &config.Config{}
	conf.Token = ""
	conf.Endpoint = c.String("endpoint")
	client := api.FromConfig(conf)

	if config.Exists() {
		conf := config.Read()
		if affirm(fmt.Sprintf("It looks like your current Plural user is %s, use this profile?", conf.Email), "PLURAL_LOGIN_AFFIRM_CURRENT_USER") {
			client = api.FromConfig(&conf)
			return postLogin(&conf, client, c)
		}
	}

	device, err := client.DeviceLogin()
	if err != nil {
		return api.GetErrorResponse(err, "DeviceLogin")
	}

	fmt.Printf("logging into Plural at %s\n", device.LoginUrl)
	if err := browser.OpenURL(device.LoginUrl); err != nil {
		fmt.Printf("Open %s in your browser to proceed\n", device.LoginUrl)
	}

	var jwt string
	for {
		result, err := client.PollLoginToken(device.DeviceToken)
		if err == nil {
			jwt = result
			break
		}

		time.Sleep(2 * time.Second)
	}

	conf.Token = jwt
	conf.ReportErrors = affirm("Would you be willing to report any errors to Plural to help with debugging?", "PLURAL_LOGIN_AFFIRM_REPORT_ERRORS")
	client = api.FromConfig(conf)
	return postLogin(conf, client, c)
}

func handleClone(c *cli.Context) error {
	url := c.Args().Get(0)
	cmd := exec.Command("git", "clone", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	repo := git.RepoName(url)
	_ = os.Chdir(repo)
	if err := cryptoInit(c); err != nil {
		return err
	}

	if err := handleUnlock(c); err != nil {
		return err
	}

	utils.Success("Your repo has been cloned and decrypted, cd %s to start working\n", repo)
	return nil
}

func downloadReadme(c *cli.Context) error {
	return wkspace.DownloadReadme()
}

func postLogin(conf *config.Config, client api.Client, c *cli.Context) error {
	me, err := client.Me()
	if err != nil {
		return api.GetErrorResponse(err, "Me")
	}

	conf.Email = me.Email
	fmt.Printf("\nlogged in as %s!\n", me.Email)

	saEmail := c.String("service-account")
	if saEmail != "" {
		jwt, email, err := client.ImpersonateServiceAccount(saEmail)
		if err != nil {
			return api.GetErrorResponse(err, "ImpersonateServiceAccount")
		}

		conf.Email = email
		conf.Token = jwt
		fmt.Printf("Assumed service account %s\n", saEmail)
		config.SetConfig(conf)
		return nil
	}

	accessToken, err := client.GrabAccessToken()
	if err != nil {
		return api.GetErrorResponse(err, "GrabAccessToken")
	}

	conf.Token = accessToken
	return conf.Flush()
}

func handleImport(c *cli.Context) error {
	dir, err := filepath.Abs(c.Args().Get(0))
	if err != nil {
		return err
	}

	conf := config.Import(pathing.SanitizeFilepath(filepath.Join(dir, "config.yml")))
	if err := conf.Flush(); err != nil {
		return err
	}

	if err := cryptoInit(c); err != nil {
		return err
	}

	data, err := os.ReadFile(pathing.SanitizeFilepath(filepath.Join(dir, "key")))
	if err != nil {
		return err
	}

	key, err := crypto.Import(data)
	if err != nil {
		return err
	}
	if err := key.Flush(); err != nil {
		return err
	}

	utils.Success("Workspace properly imported\n")
	return nil
}

func handleServe(c *cli.Context) error {
	return server.Run()
}
