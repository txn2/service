package service

import (
	"errors"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/go-yaml/yaml"
)

type TokenCfg struct {
	EncKey   string `yaml:"encKey"`
	ExpHours int    `yaml:"expHours"`
}

// tok defines a token generator object
type tok struct {
	encKey []byte
	exp    int
}

// GetToken generated a HS256 token from an object
func (t *tok) GetToken(v interface{}) (string, error) {
	// make a token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	token.Claims = jwt_lib.MapClaims{
		"data": v,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	tokenString, err := token.SignedString(t.encKey)

	return tokenString, err
}

// NewTokenGenerator returns a configured tok used
// for generating tokens
func NewTokenGenerator(cfg Cfg) (*tok, error) {
	tcfg := TokenCfg{}

	// get some configuration
	if cfg["token"] == nil {
		return nil, errors.New("no token configuration found")
	}

	// get yaml
	d, err := yaml.Marshal(cfg["token"])
	if err != nil {
		return nil, err
	}

	// use the yaml to hydrate the configuration
	err = yaml.Unmarshal(d, &tcfg)
	if err != nil {
		return nil, err
	}

	tok := &tok{encKey: []byte(tcfg.EncKey), exp: tcfg.ExpHours}

	return tok, nil
}
