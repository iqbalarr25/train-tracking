package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func EncryptAES(plainText string, key string) (string, error) {
	keyByte, _ := stringTo32ByteArray(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encrypted := aead.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func DecryptAES(encryptedText string, key string) (string, error) {
	// Dekode dari base64
	keyByte, _ := stringTo32ByteArray(key)
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Ambil nonce dari data terenkripsi
	nonceSize := aead.NonceSize()
	if len(encryptedData) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	nonce, cipherText := encryptedData[:nonceSize], encryptedData[nonceSize:]

	// Dekripsi data
	decrypted, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func stringTo32ByteArray(input string) ([]byte, error) {
	if len(input) > 32 {
		return nil, fmt.Errorf("string terlalu panjang, maksimal 32 byte")
	}

	byteArray := make([]byte, 32)
	copy(byteArray, input) // Menyalin isi string ke byte array

	return byteArray, nil
}
