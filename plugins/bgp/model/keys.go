/*REMEMBER TO USE THIS MODEL WHEN BRINGING OVER THE PLUG-IN*/

package model

import (
	"github.com/contiv/vpp/plugins/ksr/model/ksrkey"
	"github.com/ligato/vpp-agent/pkg/models"
)


const ModuleName = "bgp"
var (
	ModelBgpConf = models.Register(&BgpConf{}, models.Spec{
		Module:  ModuleName,
		Version: "v1",
		Type:    "bgpconf",
	},
	models.WithNameTemplate("{{.Name}}"),
	)
)


// Keyword defines the keyword identifying NodeConfig data.
const Keyword = "bgpconfig"

// KeyPrefix return prefix where all node configs are persisted.
func KeyPrefix() string {
	return ksrkey.KsrK8sPrefix + "/" + Keyword + "/"
}

// Key returns the key for configuration of a given node.
func Key(conf string) string {
	return KeyPrefix() + conf
}