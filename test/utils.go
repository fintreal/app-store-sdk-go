package test

import (
	"math/rand"
	"os"

	openapi "github.com/fintreal/app-store-sdk-go"
)

var apiClient = GetClient()

func GetClient() *openapi.APIClient {
	keyData := os.Getenv("KEY_DATA")
	keyId := os.Getenv("KEY_ID")
	issuerId := os.Getenv("ISSUER_ID")
	var token, err = openapi.GenerateToken(keyData, keyId, issuerId)
	if err != nil {
		panic(err)
	}
	var configuration = openapi.NewConfiguration(*token)
	var apiClient = openapi.NewAPIClient(configuration)

	return apiClient
}

const charset = "abcdefghijklmnopqrstuvwxyz"

func GenerateRandomString() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
