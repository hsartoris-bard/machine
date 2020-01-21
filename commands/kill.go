package commands

import "github.com/hsartoris-bard/machine/libmachine"

func cmdKill(c CommandLine, api libmachine.API) error {
	return runAction("kill", c, api)
}
