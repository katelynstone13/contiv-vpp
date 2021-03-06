// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package govppmux

import (
	"git.fd.io/govpp.git/adapter"
	govppapi "git.fd.io/govpp.git/api"
	"github.com/ligato/cn-infra/logging/measure/model/apitrace"
)

// TraceAPI is extended API with ability to get traced VPP binary API calls
type TraceAPI interface {
	API

	// GetTrace serves to obtain measured binary API calls
	GetTrace() *apitrace.Trace
}

// StatsAPI is extended API with ability to get VPP stats data
type StatsAPI interface {
	API

	// ListStats returns all stats names present on the VPP. Patterns can be used as a prefix
	// to filter the output
	ListStats(patterns ...string) ([]string, error)

	// ListStats returns all stats names, types and values from the VPP. Patterns can be used as a prefix
	// to filter the output. Stats are divided between workers. Example:
	//
	// stats values: {{0, 20, 30}{0, 0, 10}}
	//
	// It means there are three interfaces on two workers (inner arrays, array index == sw_if_index),
	// and statistics are like following:
	//
	// 0 for sw_if_index 0
	// 20 for sw_if_index 1
	// 40 for sw_if_index 2 (sum of stats from all workers)
	//
	DumpStats(patterns ...string) ([]*adapter.StatEntry, error)

	// GetSystemStats retrieves system statistics of the connected VPP instance like Vector rate, Input rate, etc.
	GetSystemStats() (*govppapi.SystemStats, error)

	// GetNodeStats retrieves a list of Node VPP counters (vectors, clocks, ...)
	GetNodeStats() (*govppapi.NodeStats, error)

	// GetInterfaceStats retrieves all counters related to the VPP interfaces
	GetInterfaceStats() (*govppapi.InterfaceStats, error)

	// GetErrorStats retrieves VPP error counters
	GetErrorStats(names ...string) (*govppapi.ErrorStats, error)
}

// API for other plugins to get connectivity to VPP.
type API interface {
	// NewAPIChannel returns a new API channel for communication with VPP via govpp core.
	// It uses default buffer sizes for the request and reply Go channels.
	//
	// Example of binary API call from some plugin using GOVPP:
	//      ch, _ := govpp_mux.NewAPIChannel()
	//      ch.SendRequest(req).ReceiveReply
	NewAPIChannel() (govppapi.Channel, error)

	// NewAPIChannelBuffered returns a new API channel for communication with VPP via govpp core.
	// It allows to specify custom buffer sizes for the request and reply Go channels.
	//
	// Example of binary API call from some plugin using GOVPP:
	//      ch, _ := govpp_mux.NewAPIChannelBuffered(100, 100)
	//      ch.SendRequest(req).ReceiveReply
	NewAPIChannelBuffered(reqChanBufSize, replyChanBufSize int) (govppapi.Channel, error)
}
