package main

import (
	"os"
    "fmt"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
)

var descriptor = plugin.PluginDescriptor{
	Name:        "helloworld-plugin",
	Description: "tanzu cli plugin for hello world app.",
	Target:      types.TargetGlobal, // <<<FIXME! set the Target of the plugin to one of {TargetGlobal,TargetK8s,TargetTMC}
	Version:     buildinfo.Version,
	BuildSHA:    buildinfo.SHA,
	Group:       plugin.ManageCmdGroup, // set group
}

func main() {
	p, err := plugin.NewPlugin(&descriptor)
	if err != nil {
		log.Fatal(err, "")
	}
	// fmt.Println("Hello world application plugin!")
	p.AddCommands(
	)
	if err := p.Execute(); err != nil {
		os.Exit(1)
	}
}
