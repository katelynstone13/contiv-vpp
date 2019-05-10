/*
 * // Copyright (c) 2018 Cisco and/or its affiliates.
 * //
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at:
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

//go:generate protoc -I ./model --gogo_out=plugins=grpc:./model ./model/bgpconfig.proto

package bgpconfig

import (
	"github.com/contiv/vpp/plugins/crd/handler/bgpconfig/model"
	"github.com/contiv/vpp/plugins/crd/pkg/apis/bgpconfig/v1"
	"reflect"
	"sync"

	informers "github.com/contiv/vpp/plugins/crd/pkg/client/informers/externalversions/bgpconfig/v1"
	"github.com/gogo/protobuf/proto"
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/datasync/kvdbsync"
	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/logging"
)

const (
	minResyncTimeout = 100  // minimum timeout between resync attempts, in ms
	maxResyncTimeout = 1000 // maximum timeout between resync attempts, in ms
)

// Handler handler implements Handler interface,
type Handler struct {
	Deps

	// Data store sync status and the mutex that protects access to it
	dsSynced bool
	dsMutex  sync.Mutex

	broker KeyProtoValBroker
	prefix string
	kpc    K8sToProtoConverter

	syncStopCh chan bool
}

// Deps defines dependencies for BgpConfig CRD Handler.
type Deps struct {
	Log                logging.Logger
	Publish            *kvdbsync.Plugin // KeyProtoValWriter does not define Delete
	ControllerInformer informers.BgpConfigInformer
}

// KeyProtoValBroker defines handler's interface to the key-value data store. It
// defines a subset of operations from a generic cn-infra broker interface
// (keyval.ProtoBroker in cn-infra).
type KeyProtoValBroker interface {
	// Put <data> to ETCD or to any other key-value based data source.
	Put(key string, data proto.Message, opts ...datasync.PutOption) error

	// Delete data under the <key> in ETCD or in any other key-value based data
	// source.
	Delete(key string, opts ...datasync.DelOption) (existed bool, err error)

	// GetValue reads a value from etcd stored under the given key.
	GetValue(key string, reqObj proto.Message) (found bool, revision int64, err error)

	// List values stored in etcd under the given prefix.
	ListValues(prefix string) (keyval.ProtoKeyValIterator, error)
}

// K8sToProtoConverter defines the signature for a function converting k8s
// objects to bgp config protobuf objects.
type K8sToProtoConverter func(interface{}) (interface{}, string, bool)

// DsItems defines the structure holding items listed from the data store.
type DsItems map[string]interface{}





// Init initializes handler configuration
// BgpConfig Handler will be taking action on resource CRUD
func (h *Handler) Init() error {
	ksrPrefix := h.Publish.ServiceLabel.GetAgentPrefix()
	h.broker = h.Publish.Deps.KvPlugin.NewBroker(ksrPrefix)
	h.syncStopCh = make(chan bool, 1)
	h.prefix = model.KeyPrefix()

	h.kpc = func(obj interface{}) (interface{}, string, bool) {
		bgpConfig, ok := obj.(*v1.BgpConfig)
		if !ok {
			h.Log.Warn("Failed to cast newly created node-config object")
			return nil, "", false
		}
		return h.bgpConfigToProto(bgpConfig), model.Key(bgpConfig.Name), true
	}
	return nil
}

// ObjectCreated is called when a CRD object is created
func (h *Handler) ObjectCreated(obj interface{}) {
	h.Log.Debugf("Object created with value: %v", obj)
	bgpConfig, ok := obj.(*v1.BgpConfig)
	if !ok {
		h.Log.Warn("Failed to cast newly created bgp-config object")
		return
	}
	globalConfigProto := h.bgpGlobalConfigToProto(bgpConfig.Spec.BGPGlobal)
	err := h.Publish.Put(model.Key(bgpConfig.Name) + "/global", globalConfigProto)
	if err != nil {
		h.Log.Errorf("error publish.put global : %v", err)
	}
	for _, nextPeer := range bgpConfig.Spec.Peers {
		peerProto := h.bgpPeersConfigToProto(nextPeer)
		err := h.Publish.Put(model.Key(bgpConfig.Name) + "/peers/" + nextPeer.Name, peerProto)
		h.Log.Errorf("error publish.put peers : %v" , err)
	}
	/*
	if err != nil {
		h.dsSynced = false
		h.startDataStoreResync()
	}*/


}

// ObjectDeleted is called when a CRD object is deleted
func (h *Handler) ObjectDeleted(obj interface{}) {
	bgpConfig, ok := obj.(*v1.BgpConfig)
	if !ok {
		h.Log.Warn("Failed to cast newly created bgp-config object")
		return
	}
	_, err := h.Publish.Delete(model.Key("global"))
	if err != nil {
		h.Log.Errorf("error publish.put global : %v", err)
	}
	for _, nextPeer := range bgpConfig.Spec.Peers {
		_, err := h.Publish.Delete(model.Key(bgpConfig.Name) + "/peers/" + nextPeer.Name)
		h.Log.Errorf("error publish.put peer : %v" , err)
	}
}

// ObjectUpdated is called when a CRD object is updated
func (h *Handler) ObjectUpdated(oldObj, newObj interface{}) {
	/*h.Log.Debugf("Object updated with value: %v", newObj)
	if !reflect.DeepEqual(oldObj, newObj) {

		h.Log.Debugf("bgp config updating item in data store, %v", newObj)
		bgpConfig, ok := newObj.(*v1.BgpConfig)
		if !ok {
			h.Log.Warn("Failed to cast delete event")
			return
		}

		nodeConfigProto := h.nodeConfigToProto(nodeConfig)
		err := h.Publish.Put(model.Key(nodeConfig.GetName()), nodeConfigProto)
		if err != nil {
			h.Log.WithField("rwErr", err).
				Warnf("node config failed to update item in data store %v", nodeConfigProto)
			h.dsSynced = false
			h.startDataStoreResync()
			return
		}
	}*/
}
// bgpConfigToProto converts bgp-config data from the Contiv's own CRD representation
// into the corresponding protobuf-modelled data format.
func (h *Handler) bgpConfigToProto(bgpConfig *v1.BgpConfig) *model.BgpConf {
	bgpConfigProto := &model.BgpConf{}

	bgpConfigProto.Global = h.bgpGlobalConfigToProto(bgpConfig.Spec.BGPGlobal)

	for _, nextPeer := range bgpConfig.Spec.Peers {
		bgpConfigProto.Peers = append(bgpConfigProto.Peers,
			h.bgpPeersConfigToProto(nextPeer))
	}

	return bgpConfigProto
}

func (h *Handler) bgpPeersConfigToProto(bgpPeersConfig v1.PeerConf) *model.PeerConf {
	bgpPeersConfigProto := &model.PeerConf{}
	bgpPeersConfigProto.Name = bgpPeersConfig.Name
	bgpPeersConfigProto.AuthPassword = bgpPeersConfig.AuthPassword
	bgpPeersConfigProto.Description = bgpPeersConfig.Description
	bgpPeersConfigProto.LocalAs = bgpPeersConfig.LocalAs
	bgpPeersConfigProto.NeighborAddress = bgpPeersConfig.NeighborAddress
	bgpPeersConfigProto.PeerAs = bgpPeersConfig.PeerAs
	bgpPeersConfigProto.PeerGroup = bgpPeersConfig.PeerGroup
	bgpPeersConfigProto.PeerType = bgpPeersConfig.PeerType
	bgpPeersConfigProto.RemovePrivateAs = model.PeerConf_RemovePrivateAs(bgpPeersConfig.RemovePrivateAs)
	bgpPeersConfigProto.RouteFlapDamping = bgpPeersConfig.RouteFlapDamping
	bgpPeersConfigProto.SendCommunity = bgpPeersConfig.SendCommunity
	bgpPeersConfigProto.NeighborInterface = bgpPeersConfig.NeighborInterface
	bgpPeersConfigProto.Vrf = bgpPeersConfig.Vrf
	bgpPeersConfigProto.AllowOwnAs = bgpPeersConfig.AllowOwnAs
	bgpPeersConfigProto.ReplacePeerAs = bgpPeersConfig.ReplacePeerAs
	bgpPeersConfigProto.AdminDown = bgpPeersConfig.AdminDown

	return bgpPeersConfigProto
}

func (h *Handler) bgpGlobalConfigToProto(bgpGlobalConfig v1.GlobalConf) *model.GlobalConf {
	bgpGlobalConfigProto := &model.GlobalConf{}
	bgpGlobalConfigProto.As = bgpGlobalConfig.As
	bgpGlobalConfigProto.Families = bgpGlobalConfig.Families
	bgpGlobalConfigProto.ListenAddresses = bgpGlobalConfig.ListenAddresses
	bgpGlobalConfigProto.RouterId = bgpGlobalConfig.RouterId
	bgpGlobalConfigProto.UseMultiplePaths = bgpGlobalConfig.UseMultiplePaths

	return bgpGlobalConfigProto
}
