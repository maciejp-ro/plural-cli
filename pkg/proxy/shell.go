package proxy

import (
	"github.com/pluralsh/plural-operator/apis/platform/v1alpha1"
	"github.com/pluralsh/plural/pkg/kubernetes/exec"
)

func execShell(namespace string, proxy *v1alpha1.Proxy) error {
	shell := proxy.Spec.ShConfig
	var rest []string
	if len(shell.Command) > 0 {
		rest = append([]string{shell.Command}, shell.Args...)
	} else {
		rest = []string{"/bin/sh"}
	}

	return exec.Exec(namespace, proxy.Spec.Target, rest)
}
