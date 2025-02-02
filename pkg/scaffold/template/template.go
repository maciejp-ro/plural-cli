package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pluralsh/gqlclient"
	"github.com/pluralsh/plural/pkg/config"
	"github.com/pluralsh/plural/pkg/output"
	"github.com/pluralsh/plural/pkg/template"
	"github.com/pluralsh/plural/pkg/utils"
	"github.com/pluralsh/plural/pkg/utils/pathing"
	"github.com/pluralsh/plural/pkg/wkspace"
	"gopkg.in/yaml.v2"
)

func templateInfo(path string) (t gqlclient.TemplateType, contents string, err error) {
	gopath := pathing.SanitizeFilepath(filepath.Join(path, "values.yaml.tpl"))
	if utils.Exists(gopath) {
		contents, err = utils.ReadFile(gopath)
		t = gqlclient.TemplateTypeGotemplate
		return
	}

	luapath := pathing.SanitizeFilepath(filepath.Join(path, "values.yaml.lua"))
	if utils.Exists(gopath) {
		contents, err = utils.ReadFile(luapath)
		t = gqlclient.TemplateTypeLua
		return
	}

	err = fmt.Errorf("could not find values.yaml.tpl or values.yaml.lua in directory, perhaps your link is to the wrong folder?")
	return
}

func BuildValuesFromTemplate(vals map[string]interface{}, w *wkspace.Workspace) (map[string]map[string]interface{}, error) {
	globals := map[string]interface{}{}
	output := make(map[string]map[string]interface{})
	for _, chartInst := range w.Charts {
		chartName := chartInst.Chart.Name
		tplate := chartInst.Version.ValuesTemplate
		templateType := chartInst.Version.TemplateType

		if w.Links != nil {
			if path, ok := w.Links.Helm[chartName]; ok {
				var err error
				templateType, tplate, err = templateInfo(path)
				if err != nil {
					return nil, err
				}
			}
		}

		switch templateType {
		case gqlclient.TemplateTypeGotemplate:
			if err := FromGoTemplate(vals, globals, output, chartName, tplate); err != nil {
				return nil, err
			}
		case gqlclient.TemplateTypeLua:
			if err := FromLuaTemplate(vals, globals, output, chartName, tplate); err != nil {
				return nil, err
			}
		}
	}

	if len(globals) > 0 {
		output["global"] = globals
	}

	output["plrl"] = map[string]interface{}{
		"license": w.Installation.LicenseKey,
	}
	return output, nil
}

func TmpValuesFile(path string) (f *os.File, err error) {
	conf := config.Read()
	if strings.HasSuffix(path, "lua") {
		return luaTmpValuesFile(path, &conf)
	}

	return goTmpValuesFile(path, &conf)

}

func luaTmpValuesFile(path string, conf *config.Config) (f *os.File, err error) {
	valuesTmpl, err := utils.ReadFile(path)
	if err != nil {
		return
	}
	f, err = os.CreateTemp("", "values.yaml")
	if err != nil {
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	vals := genDefaultValues(conf)

	output, err := ExecuteLua(vals, valuesTmpl)
	if err != nil {
		return nil, err
	}

	io, err := yaml.Marshal(output)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(io))
	_, err = f.Write(io)
	if err != nil {
		return nil, err
	}
	return
}

func goTmpValuesFile(path string, conf *config.Config) (f *os.File, err error) {
	valuesTmpl, err := utils.ReadFile(path)
	if err != nil {
		return
	}
	tmpl, err := template.MakeTemplate(valuesTmpl)
	if err != nil {
		return
	}

	vals := genDefaultValues(conf)
	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, vals); err != nil {
		return
	}

	f, err = os.CreateTemp("", "values.yaml")
	if err != nil {
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	fmt.Println(buf.String())
	err = wkspace.FormatValues(f, buf.String(), output.New())
	return
}

func genDefaultValues(conf *config.Config) map[string]interface{} {
	return map[string]interface{}{
		"Values":        map[string]interface{}{},
		"Configuration": map[string]map[string]interface{}{},
		"License":       "example-license",
		"Region":        "region",
		"Project":       "example",
		"Cluster":       "cluster",
		"Provider":      "provider",
		"Config":        conf,
		"Context":       map[string]interface{}{},
	}
}
