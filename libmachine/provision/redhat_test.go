package provision

import (
	"testing"

	"github.com/hsartoris-bard/machine/drivers/fakedriver"
	"github.com/hsartoris-bard/machine/libmachine/auth"
	"github.com/hsartoris-bard/machine/libmachine/engine"
	"github.com/hsartoris-bard/machine/libmachine/provision/provisiontest"
	"github.com/hsartoris-bard/machine/libmachine/swarm"
)

func TestRedHatDefaultStorageDriver(t *testing.T) {
	p := NewRedHatProvisioner("", &fakedriver.Driver{})
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "devicemapper" {
		t.Fatal("Default storage driver should be devicemapper")
	}
}
