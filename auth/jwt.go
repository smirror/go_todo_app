package auth

import (
	_ "embed"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"golang.org/x/net/context"
	"todo_app/clock"
	"todo_app/entity"
)

//go:embed cert/key
var rawPrivKey []byte

//go:embed cert/key.pub
var rawPubKey []byte

type JWTer struct {
	PrivateKey, PublixKey jwk.Key
	Store                 Store
	Clocker               clock.Clocker
}

//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID entity.UserID) error
	Load(ctx context.Context, key string) (entity.UserID, error)
}

func NewJWTer(s Store) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	j.PrivateKey = privkey
	j.PublixKey = pubkey
	j.Clocker = clock.RealClocker{}
	return j, nil
}

func parse(rawkey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawkey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse key: %w", err)
	}

	return key, nil
}
