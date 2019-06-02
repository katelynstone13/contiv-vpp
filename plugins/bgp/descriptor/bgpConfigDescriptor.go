package descriptor

import (
	"context"
	"github.com/contiv/vpp/plugins/bgp/descriptor/adapter"
	"github.com/contiv/vpp/plugins/bgp/model"
	"github.com/ligato/cn-infra/logging"
	kvs "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	bgpapi "github.com/osrg/gobgp/api"
	gobgp "github.com/osrg/gobgp/pkg/server"
	"net"
)

const (
	bgpDescriptorName = "bgp-conf"
)

//our bgp descriptor
type BgpDescriptor struct {
	log       logging.Logger
	server    *gobgp.BgpServer
	hasConfig chan bool
}

// NewGlobalConfDescriptor creates a new instance of the descriptor.
func NewBgpConfDescriptor(log logging.PluginLogger, server *gobgp.BgpServer, hasConfig chan bool) *kvs.KVDescriptor {
	d := &BgpDescriptor{log: log, server: server, hasConfig: hasConfig}

	// Set plugin descriptor init values
	gcd := &adapter.BgpConfDescriptor{
		Name:        bgpDescriptorName,
		NBKeyPrefix: model.ModelBgpConf.KeyPrefix(),
		ValueTypeName: model.ModelBgpConf.ProtoName(),
		KeySelector:   model.ModelBgpConf.IsKeyValid,
		KeyLabel:      model.ModelBgpConf.StripKeyPrefix,
		Create: d.Create,
		Delete: d.Delete,
		UpdateWithRecreate: func(key string, oldValue, newValue *model.BgpConf, metadata interface{}) bool {
			// Modify always performed via re-creation
			return true
		},
	}
	return adapter.NewBgpConfDescriptor(gcd)
}

// Create creates new value.
func (d *BgpDescriptor) Create(key string, value *model.BgpConf) (metadata interface{}, err error) {
	globalConf := value.Global
	peersConf := value.Peers

	/*GLOBAL DESCRIPTOR*/
	//Checks if RouterId is a valid IP Address
	if net.ParseIP(globalConf.RouterId) == nil {
		d.log.Errorf("Invalid IP Address for RouterId = %s", globalConf.RouterId)
		return nil, err
	}
	d.log.Info("parsed IP global descriptor")
	//Checks if AS is above 64512 and below 65536
	if globalConf.As < 64512 || globalConf.As > 65536 {
		d.log.Errorf("Invalid AS Number = %d. AS Number should be above 64512 and below 65536", globalConf.As)
		return nil, err
	}
	d.log.Info("parsed AS global descriptor")
	//Checks if ListenPort is -1
	if globalConf.ListenPort != -1 {
		d.log.Errorf("Invalid ListenPort = %d. ListenPort should be -1", globalConf.ListenPort)
		return nil, err
	}
	d.log.Info("parsed ListenPort global descriptor")
	d.log.Infof("Creating GlobalConf As = %d, RouterId = %s, ListenPort = %d",
		globalConf.As, globalConf.RouterId, globalConf.ListenPort)
	err = d.server.StartBgp(context.Background(), &bgpapi.StartBgpRequest{
		Global: &bgpapi.Global{
			As:         globalConf.As,
			RouterId:   globalConf.RouterId,
			ListenPort: globalConf.ListenPort,
		},
	})

	/*PEER DESCRIPTOR*/
	for _, nextPeer := range peersConf {
		if net.ParseIP(nextPeer.NeighborAddress) == nil {
			d.log.Errorf("Invalid IP Address for NeighborAddress = %s", nextPeer.NeighborAddress)
			return nil, err
		}
		d.log.Debug("parsed IP peer descriptor")
		//Checks if AS is above 64512 and below 65536
		if nextPeer.PeerAs < 64512 || nextPeer.PeerAs > 65536 {
			d.log.Errorf("Invalid AS Number = %d. AS Number should be above 64512 and below 65536", nextPeer.PeerAs)
			return nil, err
		}
		d.log.Debug("parsed AS peer descriptor")
		n := &bgpapi.Peer{
			Conf: &bgpapi.PeerConf{
				NeighborAddress: nextPeer.NeighborAddress,
				PeerAs:          nextPeer.PeerAs,
			},
		}
		err = d.server.AddPeer(context.Background(), &bgpapi.AddPeerRequest{
			Peer: n,
		})
	}

	if err != nil {
		d.log.Errorf("Error creating BgpConf = %s", err)
		return nil, err
	}
	d.hasConfig <- true
	return nil, nil
}

// Delete removes an existing value.
func (d *BgpDescriptor) Delete(key string, value *model.BgpConf, metadata interface{}) error {
	globalConf := value.Global
	peersConf := value.Peers

	/*DELETING GLOBAL*/
	d.log.Infof("Deleting GlobalConf As = %d, RouterId = %s, ListenPort = %d",
		globalConf.As, globalConf.RouterId, globalConf.ListenPort)
	err := d.server.StopBgp(context.Background(), &bgpapi.StopBgpRequest{})

	/*DELETING PEERS*/
	for _, nextPeer := range peersConf {
		//d.log.Infof("Deleting Peer %s", nextPeer.Name)
		err = d.server.DeletePeer(context.Background(), &bgpapi.DeletePeerRequest{
			Address: nextPeer.NeighborAddress,
		})
	}

	if err != nil {
		d.log.Errorf("Error deleting BgpConf = %s", err)
		return err
	}
	return nil
}
