package provision

import (
	"fmt"

	"github.com/rancher/machine/libmachine/auth"
	"github.com/rancher/machine/libmachine/drivers"
	"github.com/rancher/machine/libmachine/engine"
	"github.com/rancher/machine/libmachine/log"
	"github.com/rancher/machine/libmachine/mcnutils"
	"github.com/rancher/machine/libmachine/provision/pkgaction"
	"github.com/rancher/machine/libmachine/provision/serviceaction"
	"github.com/rancher/machine/libmachine/swarm"
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
