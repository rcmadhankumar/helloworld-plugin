package main

import (
	"os"
	"fmt"
	"helloworld-plugin/pkg/commands"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
)

var descriptor = plugin.PluginDescriptor{
	Name:        "helloworldplugin",
	Description: "tanzu cli plugin for hello world app.",
	Target:      types.TargetK8s,
	Version:     buildinfo.Version,
	BuildSHA:    buildinfo.SHA,
	Group:       plugin.RunCmdGroup,
}

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "")
	}
	fmt.Println("Hello world application plugin!")
	p.AddCommands(
		commands.TestCmd,
		// commands.testCmd,
	)
	if err := p.Execute(); err != nil {
		os.Exit(0)
	}
}
