package common

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

func encodePrivateKey( *key rsa.PrivateKey ) []byte {
  //base64?
}

func decodePrivateKey( key []byte] ) (rsa.PrivateKey, error) {

}

func encodePublicKey( *key rsa.PublicKey ) []byte {

}

func decodePublicKey( key []byte] ) (rsa.PublicKey, error) {

}

func generateNewKeys() (*rsa.PrivateKey, error){
  // Returns: PrivateKey, error
  // NOTE: privateKey.Public() gives you the public key corresponding to the private one
  bitsize := 4096 //more than enough
  privkey, err := rsa.GenerateKey(rand.Reader, bitsize);
  if err != nil {  return privkey, err }
  err = privkey.Validate()
  return privkey, err
}
