package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"

	fury "github.com/Fury-Labs/fury/app"
)

func main() {
	fury.SetAccountAddressPrefixes()

	root, _ := NewRootCmd()
	if err := servercmd.Execute(root, fury.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
