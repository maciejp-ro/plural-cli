package bundle_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pluralsh/plural/pkg/api"
	"github.com/pluralsh/plural/pkg/bundle"
	"github.com/pluralsh/plural/pkg/manifest"
	"github.com/stretchr/testify/assert"

	"gopkg.in/yaml.v2"
)

func TestConfigureEnvVariables(t *testing.T) {
	tests := []struct {
		name          string
		ctx           map[string]interface{}
		item          *api.ConfigurationItem
		context       *manifest.Context
		repo          string
		expectedError string
		expectedValue string
		envVars       map[string]string
	}{
		{
			name: "test integer item",
			item: &api.ConfigurationItem{
				Name:    "test_item",
				Default: "123",
				Type:    bundle.Int,
			},
			context:       &manifest.Context{},
			ctx:           map[string]interface{}{},
			repo:          "test",
			envVars:       map[string]string{"PLURAL_TEST_TEST_ITEM": "95"},
			expectedValue: "95",
		},
		{
			name: "test bool item",
			item: &api.ConfigurationItem{
				Name:    "test_item",
				Default: "false",
				Type:    bundle.Bool,
			},
			context:       &manifest.Context{},
			ctx:           map[string]interface{}{},
			repo:          "test",
			envVars:       map[string]string{"PLURAL_TEST_TEST_ITEM": "true"},
			expectedValue: "true",
		},
		{
			name: "test domain item",
			item: &api.ConfigurationItem{
				Name:    "test_item",
				Default: "false",
				Type:    bundle.Domain,
			},
			context:       &manifest.Context{},
			ctx:           map[string]interface{}{},
			repo:          "test",
			envVars:       map[string]string{"PLURAL_TEST_TEST_ITEM": "https://test.com"},
			expectedValue: "https://test.com",
		},
		{
			name: "test string item",
			item: &api.ConfigurationItem{
				Name:    "test_item",
				Default: "abc",
				Type:    bundle.String,
			},
			context:       &manifest.Context{},
			ctx:           map[string]interface{}{},
			repo:          "test",
			envVars:       map[string]string{"PLURAL_TEST_TEST_ITEM": "test"},
			expectedValue: "test",
		},
		{
			name: "test password item",
			item: &api.ConfigurationItem{
				Name:    "test_item",
				Default: "secret",
				Type:    bundle.Password,
			},
			context:       &manifest.Context{},
			ctx:           map[string]interface{}{},
			repo:          "test",
			envVars:       map[string]string{"PLURAL_TEST_TEST_ITEM": "secret-123"},
			expectedValue: "secret-123",
		},
		{
			name: "test bucket item",
			item: &api.ConfigurationItem{
				Name:    "test_item",
				Default: "test",
				Type:    bundle.Bucket,
			},
			context:       &manifest.Context{Buckets: []string{}},
			ctx:           map[string]interface{}{},
			repo:          "test",
			envVars:       map[string]string{"PLURAL_TEST_TEST_ITEM": "abc"},
			expectedValue: "abc",
		}, {
			name: "test file item",
			item: &api.ConfigurationItem{
				Name:    "test_item",
				Default: "context.yaml",
				Type:    bundle.File,
			},
			context:       &manifest.Context{Buckets: []string{}},
			ctx:           map[string]interface{}{},
			repo:          "test",
			envVars:       map[string]string{"PLURAL_TEST_TEST_ITEM": "workspace.yaml"},
			expectedValue: "apiVersion: \"\"\nkind: \"\"\nmetadata: null\nspec:\n  cluster: \"\"\n  bucket: \"\"\n  project: test\n  provider: \"\"\n  region: \"\"\n  owner: null\n  network: null\n  bucketPrefix: \"\"\n  context: {}\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			for k, v := range test.envVars {
				os.Setenv(k, v)
			}
			defer func(envVars map[string]string) {
				for k := range envVars {
					os.Unsetenv(k)
				}
			}(test.envVars)

			dir, err := ioutil.TempDir("", "config")
			assert.NoError(t, err)
			defer os.RemoveAll(dir)

			err = os.Chdir(dir)
			assert.NoError(t, err)

			project := manifest.VersionedProjectManifest{
				Spec: &manifest.ProjectManifest{
					Project: test.repo,
				},
			}
			out, err := yaml.Marshal(project)
			assert.NoError(t, err)
			err = os.WriteFile(filepath.Join(dir, "workspace.yaml"), out, 0644)
			assert.NoError(t, err)

			err = bundle.Configure(test.ctx, test.item, test.context, test.repo)
			if test.expectedError != "" {
				assert.Equal(t, err.Error(), test.expectedError)
			} else {
				assert.NoError(t, err)
			}

			val := test.ctx[test.item.Name]
			assert.Equal(t, test.expectedValue, fmt.Sprint(val))

		})
	}
}