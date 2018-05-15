package keystore

import (
	"fmt"

	logging "github.com/ipfs/go-log"
	ci "github.com/libp2p/go-libp2p-crypto"
)

var log = logging.Logger("keystore")

// Keystore provides a key management interface
type Keystore interface {
	// Has returns whether or not a key exist in the Keystore
	Has(string) (bool, error)
	// Put stores a key in the Keystore, if a key with the same name already exists, returns ErrKeyExists
	Put(string, ci.PrivKey) error
	// Get retrieves a key from the Keystore if it exists, and returns ErrNoSuchKey
	// otherwise.
	Get(string) (ci.PrivKey, error)
	// Delete removes a key from the Keystore
	Delete(string) error
	// List returns a list of key identifier
	List() ([]string, error)
}

var ErrNoSuchKey = fmt.Errorf("no key by the given name was found")
var ErrKeyExists = fmt.Errorf("key by that name already exists, refusing to overwrite")
