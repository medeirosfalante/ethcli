package tron

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestMnemonic    = "crunch goat helmet buddy metal crowd fish rack entire glare job milk call wish protect"
)

func TestTronAccount_PubKeyFromPath(t *testing.T) {
	addr, err := NewTronAccount(TestMnemonic, "m/44'/195'/0'/0/0")
	assert.NotNil(t, addr)
	assert.NoError(t, err)
	t.Log(addr.PublicKey())
}