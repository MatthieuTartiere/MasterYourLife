package main

import providerCmd "github.com/MasterYourLife/back/activity-provider/cmd/provider"

var (
	BinaryName    string = "activity-provider"
	version       string = "v0.1.0"
	buildDateTime string = "unknown"
)

func main() {
	providerCmd.BinaryName = BinaryName
	providerCmd.BuildDateTime = buildDateTime
	providerCmd.RootCmd.Version = version
	_ = providerCmd.RootCmd.Execute()
}
