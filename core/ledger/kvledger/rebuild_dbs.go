/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kvledger

import (
	"github.com/Yunpeng-J/HLF-2.2/common/ledger/blkstorage"
	"github.com/Yunpeng-J/HLF-2.2/common/ledger/util/leveldbhelper"
	"github.com/Yunpeng-J/HLF-2.2/core/ledger"
	"github.com/Yunpeng-J/HLF-2.2/core/ledger/kvledger/txmgmt/statedb/statecouchdb"
	"github.com/pkg/errors"
)

// RebuildDBs drops existing ledger databases.
// Dropped database will be rebuilt upon server restart
func RebuildDBs(config *ledger.Config) error {
	rootFSPath := config.RootFSPath
	fileLockPath := fileLockPath(rootFSPath)
	fileLock := leveldbhelper.NewFileLock(fileLockPath)
	if err := fileLock.Lock(); err != nil {
		return errors.Wrap(err, "as another peer node command is executing,"+
			" wait for that command to complete its execution or terminate it before retrying")
	}
	defer fileLock.Unlock()

	if config.StateDBConfig.StateDatabase == "CouchDB" {
		if err := statecouchdb.DropApplicationDBs(config.StateDBConfig.CouchDB); err != nil {
			return err
		}
	}
	if err := dropDBs(rootFSPath); err != nil {
		return err
	}

	blockstorePath := BlockStorePath(rootFSPath)
	return blkstorage.DeleteBlockStoreIndex(blockstorePath)
}
