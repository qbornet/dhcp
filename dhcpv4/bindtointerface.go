package dhcpv4

import (
	"github.com/qbornet/dhcp/interfaces"
)

// BindToInterface (deprecated) redirects to interfaces.BindToInterface
func BindToInterface(fd int, ifname string) error {
	return interfaces.BindToInterface(fd, ifname)
}
