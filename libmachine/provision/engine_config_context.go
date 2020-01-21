package provision

import (
	"github.com/hsartoris-bard/machine/libmachine/auth"
	"github.com/hsartoris-bard/machine/libmachine/engine"
)

type EngineConfigContext struct {
	DockerPort       int
	AuthOptions      auth.Options
	EngineOptions    engine.Options
	DockerOptionsDir string
}
