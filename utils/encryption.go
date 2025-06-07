package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

type EncryptionService struct {
	EncryptionKey []byte
}

func NewService() *EncryptionService {
	var service EncryptionService
	existingKey := os.Getenv("BYEMBER_ENCRYPTION_KEY")

	if strings.TrimSpace(existingKey) == "" {
		service.EncryptionKey = []byte(existingKey)
	} else {
		service.GenerateKey()
	}
	return &service
}

func (e *EncryptionService) GenerateKey() (bool, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return false, err
	}
	e.EncryptionKey = key
	return true, nil
}

func (e *EncryptionService) Encrypt(plaintext []byte) (string, error) {
	key := []byte(e.EncryptionKey)
	fmt.Println(key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (e *EncryptionService) Decrypt(cipherTextBase64 string) ([]byte, error) {
	key := []byte(e.EncryptionKey)

	ciphertext, err := base64.StdEncoding.DecodeString(cipherTextBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext trop court")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
