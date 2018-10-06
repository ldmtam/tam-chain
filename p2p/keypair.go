package p2p

import (
	"crypto/rand"
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p-crypto"
)

func marshalPrivKey(key crypto.PrivKey) (string, error) {
	bytes, err := crypto.MarshalPrivateKey(key)
	if err != nil {
		return "", err
	}
	return crypto.ConfigEncodeKey(bytes), nil
}

func unmarshalPrivKey(data string) (crypto.PrivKey, error) {
	bytes, err := crypto.ConfigDecodeKey(data)
	if err != nil {
		return nil, err
	}
	return crypto.UnmarshalPrivateKey(bytes)
}

func getKeyFromFile(path string) (crypto.PrivKey, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return unmarshalPrivKey(string(data))
}

func writeKeyToFile(path string, key crypto.PrivKey) error {
	data, err := marshalPrivKey(key)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(data), 0666)
}

func getOrCreateKey(path string) (crypto.PrivKey, error) {
	privKey, err := getKeyFromFile(path)
	if err == nil {
		return privKey, nil
	}
	privKey, _, err = crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}
	if path != "" {
		writeKeyToFile(path, privKey)
	}
	return privKey, nil
}
