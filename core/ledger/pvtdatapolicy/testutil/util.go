/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package testutil

import (
	"github.com/Yunpeng-J/fabric-protos-go/peer"
	"github.com/Yunpeng-J/HLF-2.2/core/ledger/pvtdatapolicy"
	"github.com/Yunpeng-J/HLF-2.2/core/ledger/pvtdatapolicy/mock"
)

// SampleBTLPolicy helps tests create a sample BTLPolicy
// The example input entry is [2]string{ns, coll}:btl
func SampleBTLPolicy(m map[[2]string]uint64) pvtdatapolicy.BTLPolicy {
	ccInfoRetriever := &mock.CollectionInfoProvider{}
	ccInfoRetriever.CollectionInfoStub = func(ccName, collName string) (*peer.StaticCollectionConfig, error) {
		btl := m[[2]string{ccName, collName}]
		return &peer.StaticCollectionConfig{BlockToLive: btl}, nil
	}
	return pvtdatapolicy.ConstructBTLPolicy(ccInfoRetriever)
}
