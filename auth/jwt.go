package auth

import (
	_ "embed"
)

//go:embed cert/key
var rawPrivkey []byte

//go:embed cert/key.pub
var rawPubkey []byte
