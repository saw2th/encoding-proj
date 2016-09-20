package client

import (
     "crypto/aes"
     "crypto/rand"
     "crypto/cipher"
     "encoding/base64"
     "errors"
     "io"
  // "log"
)

// simple in memory key/value store 
// Key - content Id, Value - encrypted Content
type ClientStore struct {
    Storage map[string] []byte
}


// Client provides functionality to interact with the encryption-server
type Client interface {
	// Store accepts an id and a payload in bytes and requests that the
	// encryption-server stores them in its data store
	Store(id, payload []byte) (aesKey []byte, err error)

	// Retrieve accepts an id and an AES key, and requests that the
	// encryption-server retrieves the original (decrypted) bytes stored
	// with the provided id
	Retrieve(id, aesKey []byte) (payload []byte, err error)
}

// Store accepts an id and a payload in bytes and requests that the
// encryption-server stores them in its data store
func (cs ClientStore) Store(id, payload []byte) (aesKey []byte, err error) {
     // create random key for content
     aesKey = make([]byte, 16)
     _, err = rand.Read(aesKey)

     if err != nil {
	return []byte(""), err
     }

     // encrypt
     encrypted_payload, err := encrypt(aesKey, payload)

     // store the encrypted text
     cs.Storage[string(id)] = encrypted_payload
     return aesKey, err
}

// Retrieve accepts an id and an AES key, and requests that the
// encryption-server retrieves the original (decrypted) bytes stored
// with the provided id
func (cs ClientStore) Retrieve(id, aesKey []byte) (payload []byte, err error) {
     stored_enc := cs.Storage[string(id)]
     if stored_enc == nil {
	return []byte(""), errors.New("No encrypted content with supplied Id")
     }

     // create copy to decrypt
     safe_stored_enc := make([]byte, len(stored_enc))
     copy(safe_stored_enc, stored_enc)
     return decrypt(aesKey, safe_stored_enc)
}

// from https://play.golang.org/p/_9zQJ0aWaG
// AES encrypt and decrypt functionality
func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

