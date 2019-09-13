// this acts as a simple in memory db to handle users and refresh tokens.
// replace these functions with actual calls to your db

package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/proximax-storage/xpx-catapult-faucet"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
	"os"
	"strings"
	"time"
)

var (
	discardRatio float64

	gcInterval time.Duration
)

var DbStorage = new(BadgerDB)

type (
	// DB defines an embedded key/value store database interface.
	DB interface {
		Get(key []byte) (value []byte, err error)
		GetAddress(sign string) error
		AddIp(sign string) error
		getAll() (value *map[string]string, err error)
		Close() error
	}

	// BadgerDB is a wrapper around a BadgerDB backend database that implements
	// the DB interface.
	BadgerDB struct {
		db         *badger.DB
		ctx        context.Context
		cancelFunc context.CancelFunc
		logger     badger.Logger
	}
)

func InitDB(dataDir string) error {
	utils.Logger(0, "initializing database")

	if err := os.MkdirAll(dataDir, 0774); err != nil {
		return err
	}

	opts := badger.DefaultOptions(dataDir)
	opts.Logger = nil

	badgerDB, err := badger.Open(opts)
	if err != nil {
		return err
	}

	bdb := &BadgerDB{
		db: badgerDB,
		//logger: logger.With("module", "db"),
	}
	bdb.ctx, bdb.cancelFunc = context.WithCancel(context.Background())

	//go bdb.runGC()

	DbStorage = bdb

	discardRatio = Faucet.Config.DbStorage.DiscardRatio

	gcInterval = Faucet.Config.DbStorage.GcInterval * time.Minute

	utils.Logger(0, "initialized database")

	return nil
}

func (bdb *BadgerDB) get(key []byte) bool {
	if err := bdb.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return false
	}
	return true
}

func (bdb *BadgerDB) set(prefix byte, table, hash string) error {
	err := bdb.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(setPrefix(table, []byte(hash)), []byte("true")).WithTTL(24 + time.Hour + time.Minute).WithMeta(prefix)
		return txn.SetEntry(e)
	})

	if err != nil {
		bdb.logger.Debugf("failed to set hash %s: %v", hash, err)
		return err
	}

	return nil
}

// runGC triggers the garbage collection for the BadgerDB backend database. It
// should be run in a goroutine.
func (bdb *BadgerDB) runGC() {
	ticker := time.NewTicker(gcInterval)
	for {
		select {
		case <-ticker.C:
			err := bdb.db.RunValueLogGC(discardRatio)
			if err != nil {
				// don't report error when GC didn't result in any cleanup
				if err == badger.ErrNoRewrite {
					bdb.logger.Debugf("no BadgerDB GC occurred: %v", err)
				} else {
					bdb.logger.Errorf("failed to GC BadgerDB: %v", err)
				}
			}
		case <-bdb.ctx.Done():
			return
		}
	}
}

func (bdb *BadgerDB) GetAddress(sigHash string) bool {
	hash := setPrefix("address", []byte(sigHash))
	return bdb.get(hash)
}

func (bdb *BadgerDB) GetIp(sigHash string) bool {
	hash := setPrefix("ip", []byte(sigHash))
	return bdb.get(hash)
}

func (bdb *BadgerDB) AddAddress(hash string) error {
	if hash == "" {
		return errors.New("empty hash")
	}
	return bdb.set(1, "address", hash)
}

func (bdb *BadgerDB) AddIp(hash string) error {
	if hash == "" {
		return errors.New("empty hash")
	}
	return bdb.set(2, "ip", hash)
}

// Close implements the DB interface. It closes the connection to the underlying
// BadgerDB database as well as invoking the context's cancel function.
func (bdb *BadgerDB) Close() error {
	bdb.cancelFunc()
	return bdb.db.Close()
}

func setPrefix(table string, hash []byte) []byte {
	prefix := []byte(fmt.Sprintf("%s/", table))
	return append(prefix, hash...)
}

func getHash(key string) string {
	return strings.Split(key, "/")[1]
}

func getStatus(key string) string {
	return strings.Split(key, "/")[1]
}
