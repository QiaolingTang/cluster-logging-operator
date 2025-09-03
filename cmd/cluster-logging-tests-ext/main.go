package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/openshift-eng/openshift-tests-extension/pkg/cmd"
	e "github.com/openshift-eng/openshift-tests-extension/pkg/extension"
	g "github.com/openshift-eng/openshift-tests-extension/pkg/ginkgo"

	// If using ginkgo, import your tests here
	_ "github.com/openshift/cluster-logging-operator/test/ext"
)

func main() {
	// Extension registry
	registry := e.NewRegistry()

	// You can declare multiple extensions, but most people will probably only need to create one.
	ext := e.NewExtension("openshift-logging", "non-payload", "cluster-logging-operator")

	ext.AddSuite(e.Suite{
		Name: "logging/fast",
		Qualifiers: []string{
			`!labels.exists(l, l=="Slow")`,
		},
	})

	ext.AddSuite(e.Suite{
		Name: "logging/slow",
		Qualifiers: []string{
			`labels.exists(l, l=="Slow")`,
		},
	})

	ext.AddSuite(e.Suite{
		Name: "logging/serial",
		Qualifiers: []string{
			`labels.exists(l, l=="Serial")`,
		},
	})

	ext.AddSuite(e.Suite{
		Name: "logging/parallel",
		Qualifiers: []string{
			`!labels.exists(l, l=="Serial")`,
		},
	})

	specs, err := g.BuildExtensionTestSpecsFromOpenShiftGinkgoSuite()
	if err != nil {
		panic(fmt.Sprintf("couldn't build extension test specs from ginkgo: %+v", err.Error()))
	}

	// specs = specs.MustFilter([]string{`name.contains("sig-openshift-logging")`}) //This works

	/*
		specs, err = specs.MustSelect(et.NameContains("sig-openshift-logging"))
		if err != nil {
			panic(fmt.Sprintf("no specs found: %v", err))
		}
	*/
	ext.AddSpecs(specs)

	registry.Register(ext)

	root := &cobra.Command{
		Long: "OpenShift Logging extended tests",
	}

	root.AddCommand(cmd.DefaultExtensionCommands(registry)...)

	if err := func() error {
		return root.Execute()
	}(); err != nil {
		os.Exit(1)
	}
}
