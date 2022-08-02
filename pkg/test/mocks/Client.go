// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	gqlclient "github.com/pluralsh/gqlclient"
	api "github.com/pluralsh/plural/pkg/api"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// AcquireLock provides a mock function with given fields: repo
func (_m *Client) AcquireLock(repo string) (*api.ApplyLock, error) {
	ret := _m.Called(repo)

	var r0 *api.ApplyLock
	if rf, ok := ret.Get(0).(func(string) *api.ApplyLock); ok {
		r0 = rf(repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ApplyLock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAccessToken provides a mock function with given fields:
func (_m *Client) CreateAccessToken() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateArtifact provides a mock function with given fields: repo, attrs
func (_m *Client) CreateArtifact(repo string, attrs api.ArtifactAttributes) (api.Artifact, error) {
	ret := _m.Called(repo, attrs)

	var r0 api.Artifact
	if rf, ok := ret.Get(0).(func(string, api.ArtifactAttributes) api.Artifact); ok {
		r0 = rf(repo, attrs)
	} else {
		r0 = ret.Get(0).(api.Artifact)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, api.ArtifactAttributes) error); ok {
		r1 = rf(repo, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCrd provides a mock function with given fields: repo, chart, file
func (_m *Client) CreateCrd(repo string, chart string, file string) error {
	ret := _m.Called(repo, chart, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(repo, chart, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDomain provides a mock function with given fields: name
func (_m *Client) CreateDomain(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateEvent provides a mock function with given fields: event
func (_m *Client) CreateEvent(event *api.UserEventAttributes) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(*api.UserEventAttributes) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateKey provides a mock function with given fields: name, content
func (_m *Client) CreateKey(name string, content string) error {
	ret := _m.Called(name, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateRecipe provides a mock function with given fields: repoName, attrs
func (_m *Client) CreateRecipe(repoName string, attrs gqlclient.RecipeAttributes) (string, error) {
	ret := _m.Called(repoName, attrs)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, gqlclient.RecipeAttributes) string); ok {
		r0 = rf(repoName, attrs)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, gqlclient.RecipeAttributes) error); ok {
		r1 = rf(repoName, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRepository provides a mock function with given fields: name, publisher, input
func (_m *Client) CreateRepository(name string, publisher string, input *gqlclient.RepositoryAttributes) error {
	ret := _m.Called(name, publisher, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *gqlclient.RepositoryAttributes) error); ok {
		r0 = rf(name, publisher, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteEabCredential provides a mock function with given fields: cluster, provider
func (_m *Client) DeleteEabCredential(cluster string, provider string) error {
	ret := _m.Called(cluster, provider)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(cluster, provider)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteShell provides a mock function with given fields:
func (_m *Client) DeleteShell() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeviceLogin provides a mock function with given fields:
func (_m *Client) DeviceLogin() (*api.DeviceLogin, error) {
	ret := _m.Called()

	var r0 *api.DeviceLogin
	if rf, ok := ret.Get(0).(func() *api.DeviceLogin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.DeviceLogin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChartInstallations provides a mock function with given fields: repoId
func (_m *Client) GetChartInstallations(repoId string) ([]*api.ChartInstallation, error) {
	ret := _m.Called(repoId)

	var r0 []*api.ChartInstallation
	if rf, ok := ret.Get(0).(func(string) []*api.ChartInstallation); ok {
		r0 = rf(repoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.ChartInstallation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCharts provides a mock function with given fields: repoId
func (_m *Client) GetCharts(repoId string) ([]*api.Chart, error) {
	ret := _m.Called(repoId)

	var r0 []*api.Chart
	if rf, ok := ret.Get(0).(func(string) []*api.Chart); ok {
		r0 = rf(repoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.Chart)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEabCredential provides a mock function with given fields: cluster, provider
func (_m *Client) GetEabCredential(cluster string, provider string) (*api.EabCredential, error) {
	ret := _m.Called(cluster, provider)

	var r0 *api.EabCredential
	if rf, ok := ret.Get(0).(func(string, string) *api.EabCredential); ok {
		r0 = rf(cluster, provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.EabCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(cluster, provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstallation provides a mock function with given fields: name
func (_m *Client) GetInstallation(name string) (*api.Installation, error) {
	ret := _m.Called(name)

	var r0 *api.Installation
	if rf, ok := ret.Get(0).(func(string) *api.Installation); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Installation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstallationById provides a mock function with given fields: id
func (_m *Client) GetInstallationById(id string) (*api.Installation, error) {
	ret := _m.Called(id)

	var r0 *api.Installation
	if rf, ok := ret.Get(0).(func(string) *api.Installation); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Installation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstallations provides a mock function with given fields:
func (_m *Client) GetInstallations() ([]*api.Installation, error) {
	ret := _m.Called()

	var r0 []*api.Installation
	if rf, ok := ret.Get(0).(func() []*api.Installation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.Installation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPackageInstallations provides a mock function with given fields: repoId
func (_m *Client) GetPackageInstallations(repoId string) ([]*api.ChartInstallation, []*api.TerraformInstallation, error) {
	ret := _m.Called(repoId)

	var r0 []*api.ChartInstallation
	if rf, ok := ret.Get(0).(func(string) []*api.ChartInstallation); ok {
		r0 = rf(repoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.ChartInstallation)
		}
	}

	var r1 []*api.TerraformInstallation
	if rf, ok := ret.Get(1).(func(string) []*api.TerraformInstallation); ok {
		r1 = rf(repoId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*api.TerraformInstallation)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(repoId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetRecipe provides a mock function with given fields: repo, name
func (_m *Client) GetRecipe(repo string, name string) (*api.Recipe, error) {
	ret := _m.Called(repo, name)

	var r0 *api.Recipe
	if rf, ok := ret.Get(0).(func(string, string) *api.Recipe); ok {
		r0 = rf(repo, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Recipe)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(repo, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepository provides a mock function with given fields: repo
func (_m *Client) GetRepository(repo string) (*api.Repository, error) {
	ret := _m.Called(repo)

	var r0 *api.Repository
	if rf, ok := ret.Get(0).(func(string) *api.Repository); ok {
		r0 = rf(repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShell provides a mock function with given fields:
func (_m *Client) GetShell() (api.CloudShell, error) {
	ret := _m.Called()

	var r0 api.CloudShell
	if rf, ok := ret.Get(0).(func() api.CloudShell); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(api.CloudShell)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTerraformInstallations provides a mock function with given fields: repoId
func (_m *Client) GetTerraformInstallations(repoId string) ([]*api.TerraformInstallation, error) {
	ret := _m.Called(repoId)

	var r0 []*api.TerraformInstallation
	if rf, ok := ret.Get(0).(func(string) []*api.TerraformInstallation); ok {
		r0 = rf(repoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.TerraformInstallation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTerraforma provides a mock function with given fields: repoId
func (_m *Client) GetTerraforma(repoId string) ([]*api.Terraform, error) {
	ret := _m.Called(repoId)

	var r0 []*api.Terraform
	if rf, ok := ret.Get(0).(func(string) []*api.Terraform); ok {
		r0 = rf(repoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.Terraform)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTfProviderScaffold provides a mock function with given fields: name, version
func (_m *Client) GetTfProviderScaffold(name string, version string) (string, error) {
	ret := _m.Called(name, version)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(name, version)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTfProviders provides a mock function with given fields:
func (_m *Client) GetTfProviders() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVersions provides a mock function with given fields: chartId
func (_m *Client) GetVersions(chartId string) ([]*api.Version, error) {
	ret := _m.Called(chartId)

	var r0 []*api.Version
	if rf, ok := ret.Get(0).(func(string) []*api.Version); ok {
		r0 = rf(chartId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(chartId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GrabAccessToken provides a mock function with given fields:
func (_m *Client) GrabAccessToken() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImpersonateServiceAccount provides a mock function with given fields: email
func (_m *Client) ImpersonateServiceAccount(email string) (string, string, error) {
	ret := _m.Called(email)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InstallRecipe provides a mock function with given fields: id
func (_m *Client) InstallRecipe(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListArtifacts provides a mock function with given fields: repo
func (_m *Client) ListArtifacts(repo string) ([]api.Artifact, error) {
	ret := _m.Called(repo)

	var r0 []api.Artifact
	if rf, ok := ret.Get(0).(func(string) []api.Artifact); ok {
		r0 = rf(repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.Artifact)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListKeys provides a mock function with given fields: emails
func (_m *Client) ListKeys(emails []string) ([]*api.PublicKey, error) {
	ret := _m.Called(emails)

	var r0 []*api.PublicKey
	if rf, ok := ret.Get(0).(func([]string) []*api.PublicKey); ok {
		r0 = rf(emails)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.PublicKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(emails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRecipes provides a mock function with given fields: repo, provider
func (_m *Client) ListRecipes(repo string, provider string) ([]*api.Recipe, error) {
	ret := _m.Called(repo, provider)

	var r0 []*api.Recipe
	if rf, ok := ret.Get(0).(func(string, string) []*api.Recipe); ok {
		r0 = rf(repo, provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.Recipe)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(repo, provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositories provides a mock function with given fields: query
func (_m *Client) ListRepositories(query string) ([]*api.Repository, error) {
	ret := _m.Called(query)

	var r0 []*api.Repository
	if rf, ok := ret.Get(0).(func(string) []*api.Repository); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, pwd
func (_m *Client) Login(email string, pwd string) (string, error) {
	ret := _m.Called(email, pwd)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, pwd)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, pwd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginMethod provides a mock function with given fields: email
func (_m *Client) LoginMethod(email string) (*api.LoginMethod, error) {
	ret := _m.Called(email)

	var r0 *api.LoginMethod
	if rf, ok := ret.Get(0).(func(string) *api.LoginMethod); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.LoginMethod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Me provides a mock function with given fields:
func (_m *Client) Me() (*api.Me, error) {
	ret := _m.Called()

	var r0 *api.Me
	if rf, ok := ret.Get(0).(func() *api.Me); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Me)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OIDCProvider provides a mock function with given fields: id, attributes
func (_m *Client) OIDCProvider(id string, attributes *api.OidcProviderAttributes) error {
	ret := _m.Called(id, attributes)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *api.OidcProviderAttributes) error); ok {
		r0 = rf(id, attributes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PollLoginToken provides a mock function with given fields: token
func (_m *Client) PollLoginToken(token string) (string, error) {
	ret := _m.Called(token)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseLock provides a mock function with given fields: repo, lock
func (_m *Client) ReleaseLock(repo string, lock string) (*api.ApplyLock, error) {
	ret := _m.Called(repo, lock)

	var r0 *api.ApplyLock
	if rf, ok := ret.Get(0).(func(string, string) *api.ApplyLock); ok {
		r0 = rf(repo, lock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ApplyLock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(repo, lock)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetInstallations provides a mock function with given fields:
func (_m *Client) ResetInstallations() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scaffolds provides a mock function with given fields: in
func (_m *Client) Scaffolds(in *api.ScaffoldInputs) ([]*api.ScaffoldFile, error) {
	ret := _m.Called(in)

	var r0 []*api.ScaffoldFile
	if rf, ok := ret.Get(0).(func(*api.ScaffoldInputs) []*api.ScaffoldFile); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*api.ScaffoldFile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*api.ScaffoldInputs) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnlockRepository provides a mock function with given fields: name
func (_m *Client) UnlockRepository(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateVersion provides a mock function with given fields: spec, tags
func (_m *Client) UpdateVersion(spec *api.VersionSpec, tags []string) error {
	ret := _m.Called(spec, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(*api.VersionSpec, []string) error); ok {
		r0 = rf(spec, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadTerraform provides a mock function with given fields: dir, repoName
func (_m *Client) UploadTerraform(dir string, repoName string) (api.Terraform, error) {
	ret := _m.Called(dir, repoName)

	var r0 api.Terraform
	if rf, ok := ret.Get(0).(func(string, string) api.Terraform); ok {
		r0 = rf(dir, repoName)
	} else {
		r0 = ret.Get(0).(api.Terraform)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dir, repoName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}