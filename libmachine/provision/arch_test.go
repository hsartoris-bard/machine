package provision

import (
	"testing"

	"github.com/hsartoris-bard/machine/drivers/fakedriver"
	"github.com/hsartoris-bard/machine/libmachine/auth"
	"github.com/hsartoris-bard/machine/libmachine/engine"
	"github.com/hsartoris-bard/machine/libmachine/provision/provisiontest"
	"github.com/hsartoris-bard/machine/libmachine/swarm"
)

func TestArchDefaultStorageDriver(t *testing.T) {
	p := NewArchProvisioner(&fakedriver.Driver{}).(*ArchProvisioner)
	p.SSHCommander = provisiontest.NewFakeSSHCommander(provisiontest.FakeSSHCommanderOptions{})
	p.Provision(swarm.Options{}, auth.Options{}, engine.Options{})
	if p.EngineOptions.StorageDriver != "overlay" {
		t.Fatal("Default storage driver should be overlay")
	}
}
