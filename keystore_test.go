package keystore

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	ci "github.com/libp2p/go-libp2p-crypto"
)

type rr struct{}

func (rr rr) Read(b []byte) (int, error) {
	return rand.Read(b)
}

func privKeyOrFatal(t *testing.T) ci.PrivKey {
	priv, _, err := ci.GenerateEd25519Key(rr{})
	if err != nil {
		t.Fatal(err)
	}
	return priv
}

func assertGetKey(ks Keystore, name string, exp ci.PrivKey) error {
	outK, err := ks.Get(name)
	if err != nil {
		return err
	}

	if !outK.Equals(exp) {
		return fmt.Errorf("key we got out didnt match expectation")
	}

	return nil
}

func TestKeystoreValidateName(t *testing.T) {
	if err := validateName(""); strings.HasPrefix(err.Error(), ErrKeyFmt.Error()) {
		t.Fatal("Expected KeyFmtErr")
	}
	if err := validateName("."); strings.HasPrefix(err.Error(), ErrKeyFmt.Error()) {
		t.Fatal("Expected KeyFmtErr")
	}
	if err := validateName("/."); strings.HasPrefix(err.Error(), ErrKeyFmt.Error()) {
		t.Fatal("Expected KeyFmtErr")
	}
	if err := validateName("./"); strings.HasPrefix(err.Error(), ErrKeyFmt.Error()) {
		t.Fatal("Expected KeyFmtErr")
	}
	if err := validateName("bad/key"); strings.HasPrefix(err.Error(), ErrKeyFmt.Error()) {
		t.Fatal("Expected KeyFmtErr")
	}
	if err := validateName(".badkey"); strings.HasPrefix(err.Error(), ErrKeyFmt.Error()) {
		t.Fatal("Expected KeyFmtErr")
	}
	if err := validateName(".re/llyadkey"); strings.HasPrefix(err.Error(), ErrKeyFmt.Error()) {
		t.Fatal("Expected KeyFmtErr")
	}
}
