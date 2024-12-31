package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func main() {
	// Initialize AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		panic("Error creating AWS session: " + err.Error())
	}

	// Create KMS client
	svc := kms.New(sess)

	// Generate a new data key using KMS
	dataKeyInput := &kms.GenerateDataKeyInput{
		KeyId:   aws.String("alias/your-kms-key-alias"),
		KeySpec: aws.String("AES_256"),
		EncryptionContext: map[string]*string{
			"dataType": aws.String("PII"),
		},
	}

	resp, err := svc.GenerateDataKey(dataKeyInput)
	if err != nil {
		panic("Error generating data key: " + err.Error())
	}

	// Encrypt the data slice using the generated data key
	plaintextData := []byte("Your sensitive data slice here")
	encryptedData, err := encryptData(plaintextData, resp.Plaintext)
	if err != nil {
		panic("Error encrypting data: " + err.Error())
	}

	fmt.Println("Encrypted data:", encryptedData)

	// Decrypt the data slice using the data key
	decryptedData, err := decryptData(encryptedData, resp.Plaintext)
	if err != nil {
		panic("Error decrypting data: " + err.Error())
	}

	fmt.Println("Decrypted data:", string(decryptedData))
}

// Encrypt data using AES-256-GCM
func encryptData(plaintext []byte, dataKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(dataKey)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return aesGCM.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt data using AES-256-GCM
func decryptData(ciphertext []byte, dataKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(dataKey)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return aesGCM.Open(nil, nonce, ciphertext, nil)
}
