package test

import (
	"context"
	"fmt"
	"testing"
)

func TestGetCertificate(t *testing.T) {
	res, _, _ := apiClient.CertificatesAPI.CertificatesGetCollection(context.Background()).Execute()
	fmt.Printf("Certificates: %v\n", res.Data)
}
