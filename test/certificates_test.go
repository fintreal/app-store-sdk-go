package test

import (
	"context"
	"testing"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestGetCertificate(t *testing.T) {
	data, _, err := apiClient.CertificatesAPI.CertificatesGetInstance(context.Background(), "ZXGDLKU59H").Execute()

	assert.NoError(t, err)
	assert.Equal(t, "ZXGDLKU59H", data.Data.Id)
	assert.Equal(t, "certificates", data.Data.Type)
	assert.Equal(t, "Apple Development: Adam Kovacs", data.Data.Attributes.Name)
	assert.Equal(t, openapi.CertificateType("DEVELOPMENT"), data.Data.Attributes.CertificateType)
	assert.Equal(t, "3B41B0552665AF0B4BE18D0F747C19A1", data.Data.Attributes.SerialNumber)
}
