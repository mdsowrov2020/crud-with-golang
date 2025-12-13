// Package cmd

package cmd

import (
	"crud/config"
	"crud/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Server(cnf)
}
