package main

import (
	_ "embed"

	"github.com/lerax-chain/lerax-chain/command/root"
	"github.com/lerax-chain/lerax-chain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
