package test

import (
	"context"
	"testing"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestGetProvisioningProfile(t *testing.T) {
	data, _, err := apiClient.ProfilesAPI.ProfilesGetInstance(context.Background(), "M4BW5YXCWX").Include([]string{"bundleId", "certificates"}).Execute()

	assert.Equal(t, "Test Provisioning Profile", data.Data.Attributes.GetName())
	assert.Equal(t, "IOS_APP_STORE", data.Data.Attributes.GetProfileType())

	assert.Equal(t, "Y74BKFQXG8", data.Data.Relationships.GetBundleId().Data.GetId())
	assert.Equal(t, "JVLG7LVRRL", data.Data.Relationships.GetCertificates().Data[0].GetId())

	assert.NoError(t, err)
}

func TestCreateAndDeleteProvisioningProfile(t *testing.T) {
	input := *openapi.NewProfileCreateRequest(
		*openapi.NewProfileCreateRequestData(
			"profiles",
			*openapi.NewProfileCreateRequestDataAttributes("Test Provisioning Profile X", "IOS_APP_STORE"),
			*openapi.NewProfileCreateRequestDataRelationships(
				*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(
					*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("bundleIds", "Y74BKFQXG8"),
				),
				*openapi.NewProfileCreateRequestDataRelationshipsCertificates(
					[]openapi.ProfileRelationshipsCertificatesDataInner{{Type: "certificates", Id: "JVLG7LVRRL"}},
				),
			),
		),
	)

	bundleId := &openapi.ProfileRelationshipsBundleId{
		Data: &openapi.BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData{Type: "bundleIds", Id: "Y74BKFQXG8"},
	}

	certificate := &openapi.ProfileRelationshipsCertificates{
		Data: []openapi.ProfileRelationshipsCertificatesDataInner{{Type: "certificates", Id: "JVLG7LVRRL"}},
	}

	data, _, err := apiClient.ProfilesAPI.ProfilesCreateInstance(context.Background()).ProfileCreateRequest(input).Execute()

	assert.Equal(t, bundleId, data.Data.Relationships.BundleId)
	assert.Equal(t, certificate, data.Data.Relationships.Certificates)

	assert.Equal(t, "Test Provisioning Profile X", data.Data.Attributes.GetName())
	assert.Equal(t, "IOS_APP_STORE", data.Data.Attributes.GetProfileType())

	assert.NoError(t, err)

	_, err = apiClient.ProfilesAPI.ProfilesDeleteInstance(context.Background(), data.Data.Id).Execute()
	assert.NoError(t, err)
}
