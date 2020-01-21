package commands

import (
	"testing"

	"github.com/hsartoris-bard/machine/commands/commandstest"
	"github.com/hsartoris-bard/machine/drivers/fakedriver"
	"github.com/hsartoris-bard/machine/libmachine"
	"github.com/hsartoris-bard/machine/libmachine/auth"
	"github.com/hsartoris-bard/machine/libmachine/engine"
	"github.com/hsartoris-bard/machine/libmachine/host"
	"github.com/hsartoris-bard/machine/libmachine/libmachinetest"
	"github.com/hsartoris-bard/machine/libmachine/provision"
	"github.com/hsartoris-bard/machine/libmachine/swarm"
	"github.com/stretchr/testify/assert"
)

func TestCmdProvision(t *testing.T) {
	testCases := []struct {
		commandLine CommandLine
		api         libmachine.API
		expectedErr error
	}{
		{
			commandLine: &commandstest.FakeCommandLine{
				CliArgs: []string{"foo", "bar"},
			},
			api: &libmachinetest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name:   "foo",
						Driver: &fakedriver.Driver{},
						HostOptions: &host.Options{
							EngineOptions: &engine.Options{},
							AuthOptions:   &auth.Options{},
							SwarmOptions:  &swarm.Options{},
						},
					},
					{
						Name:   "bar",
						Driver: &fakedriver.Driver{},
						HostOptions: &host.Options{
							EngineOptions: &engine.Options{},
							AuthOptions:   &auth.Options{},
							SwarmOptions:  &swarm.Options{},
						},
					},
				},
			},
			expectedErr: nil,
		},
	}

	provision.SetDetector(&provision.FakeDetector{
		Provisioner: provision.NewFakeProvisioner(nil),
	})

	// fakeprovisioner always returns "true" for compatible host, so we
	// just need to register it.
	provision.Register("fakeprovisioner", &provision.RegisteredProvisioner{
		New: provision.NewFakeProvisioner,
	})

	for _, tc := range testCases {
		assert.Equal(t, tc.expectedErr, cmdProvision(tc.commandLine, tc.api))
	}
}
