package commands

import (
	"testing"

	"github.com/hsartoris-bard/machine/commands/commandstest"
	"github.com/hsartoris-bard/machine/drivers/fakedriver"
	"github.com/hsartoris-bard/machine/libmachine/host"
	"github.com/hsartoris-bard/machine/libmachine/libmachinetest"
	"github.com/hsartoris-bard/machine/libmachine/state"
	"github.com/stretchr/testify/assert"
)

func TestCmdURLMissingMachineName(t *testing.T) {
	commandLine := &commandstest.FakeCommandLine{}
	api := &libmachinetest.FakeAPI{}

	err := cmdURL(commandLine, api)

	assert.Equal(t, ErrNoDefault, err)
}

func TestCmdURLTooManyNames(t *testing.T) {
	commandLine := &commandstest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1", "machineToRemove2"},
	}
	api := &libmachinetest.FakeAPI{}

	err := cmdURL(commandLine, api)

	assert.EqualError(t, err, "Error: Expected one machine name as an argument")
}

func TestCmdURL(t *testing.T) {
	commandLine := &commandstest.FakeCommandLine{
		CliArgs: []string{"machine"},
	}
	api := &libmachinetest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name: "machine",
				Driver: &fakedriver.Driver{
					MockState: state.Running,
					MockIP:    "120.0.0.1",
				},
			},
		},
	}

	stdoutGetter := commandstest.NewStdoutGetter()
	defer stdoutGetter.Stop()

	err := cmdURL(commandLine, api)

	assert.NoError(t, err)
	assert.Equal(t, "tcp://120.0.0.1:2376\n", stdoutGetter.Output())
}
