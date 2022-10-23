package main

import (
	"github.com/loft-sh/vcluster-mydeployment-example/syncers"
	"github.com/loft-sh/vcluster-sdk/plugin"
)

func main() {
	ctx := plugin.MustInit()
	plugin.MustRegister(syncers.NewMydeploymentSyncer(ctx))
	plugin.MustStart()
}
