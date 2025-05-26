package test

import (
	"context"
	"testing"
)

func TestGetCertificate(t *testing.T) {
	apiClient.CertificatesAPI.CertificatesGetCollection(context.Background()).Execute()
}
