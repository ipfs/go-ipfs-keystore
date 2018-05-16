package keystore

import ci "gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"

type keyMap map[string]ci.PrivKey

// MemKeystore is a keystore backed by an in-memory map
type MemKeystore struct {
	values keyMap
}

func NewMemKeystore() *MemKeystore {
	return &MemKeystore{
		values: keyMap{},
	}
}

// Has returns whether or not a key exists in the Keystore
func (mk *MemKeystore) Has(k string) (bool, error) {
	if err := validateName(k); err != nil {
		return false, err
	}
	_, found := mk.values[k]
	return found, nil
}

// Put store a key in the Keystore
func (mk *MemKeystore) Put(k string, v ci.PrivKey) error {
	if err := validateName(k); err != nil {
		return err
	}

	_, found := mk.values[k]
	if found {
		return ErrKeyExists
	}

	mk.values[k] = v
	return nil
}

// Get retrieve a key from the Keystore
func (mk *MemKeystore) Get(k string) (ci.PrivKey, error) {
	if err := validateName(k); err != nil {
		return nil, err
	}

	v, found := mk.values[k]
	if !found {
		return nil, ErrNoSuchKey
	}

	return v, nil
}

// Delete remove a key from the Keystore
func (mk *MemKeystore) Delete(k string) error {
	if err := validateName(k); err != nil {
		return err
	}
	if _, found := mk.values[k]; !found {
		return ErrNoSuchKey
	}

	delete(mk.values, k)
	return nil
}

// List return a list of key identifier
func (mk *MemKeystore) List() ([]string, error) {
	out := make([]string, 0, len(mk.values))
	for k := range mk.values {
		err := validateName(k)
		if err == nil {
			out = append(out, k)
		} else {
			log.Warningf("ignoring the invalid keyfile: %s", k)
		}
	}
	return out, nil
}
