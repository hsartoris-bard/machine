package provision

import (
	"github.com/hsartoris-bard/machine/libmachine/auth"
	"github.com/hsartoris-bard/machine/libmachine/drivers"
	"github.com/hsartoris-bard/machine/libmachine/engine"
	"github.com/hsartoris-bard/machine/libmachine/log"
	"github.com/hsartoris-bard/machine/libmachine/provision/pkgaction"
	"github.com/hsartoris-bard/machine/libmachine/swarm"
)

func init() {
	Register("PhotonOS", &RegisteredProvisioner{
		New: NewPhotonOSProvisioner,
	})
}

func NewPhotonOSProvisioner(d drivers.Driver) Provisioner {
	return &PhotonOSProvisioner{
		NewSystemdProvisioner("photon", d),
	}
}

type PhotonOSProvisioner struct {
	SystemdProvisioner
}

func (provisioner *PhotonOSProvisioner) String() string {
	return "photon"
}

func (provisioner *PhotonOSProvisioner) CompatibleWithHost() bool {
	return provisioner.OsReleaseInfo.ID == provisioner.OsReleaseID || provisioner.OsReleaseInfo.IDLike == provisioner.OsReleaseID
}

func (provisioner *PhotonOSProvisioner) Package(name string, action pkgaction.PackageAction) error {
	return nil
}

func (provisioner *PhotonOSProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions

	if err := provisioner.SetHostname(provisioner.Driver.GetMachineName()); err != nil {
		return err
	}

	if err := makeDockerOptionsDir(provisioner); err != nil {
		return err
	}

	log.Debugf("Preparing certificates")
	provisioner.AuthOptions = setRemoteAuthOptions(provisioner)

	log.Debugf("Setting up certificates")
	if err := ConfigureAuth(provisioner); err != nil {
		return err
	}

	log.Debug("Configuring swarm")
	err := configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions)
	return err
}
