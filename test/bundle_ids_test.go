package test

import (
	"context"
	"testing"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestGetBundleIds(t *testing.T) {
	apiClient.BundleIdsAPI.BundleIdsGetCollection(context.Background()).Execute()
}

func TestGetBundleId(t *testing.T) {
	expectedName := "Integration Test Bundle ID"
	expectedIdentifier := "com.fintreal.test"
	expectedId := "Y74BKFQXG8"

	expectedCapabilities := []openapi.BundleIdRelationshipsBundleIdCapabilitiesDataInner{
		*openapi.NewBundleIdRelationshipsBundleIdCapabilitiesDataInner("bundleIdCapabilities", expectedId+"_IN_APP_PURCHASE"),
		*openapi.NewBundleIdRelationshipsBundleIdCapabilitiesDataInner("bundleIdCapabilities", expectedId+"_APPLE_ID_AUTH"),
	}

	bundleId, _, err := apiClient.BundleIdsAPI.BundleIdsGetInstance(context.Background(), expectedId).Include([]string{"bundleIdCapabilities"}).Execute()
	assert.Equal(t, expectedId, bundleId.Data.GetId())
	assert.Equal(t, expectedIdentifier, bundleId.Data.Attributes.GetIdentifier())
	assert.Equal(t, expectedName, bundleId.Data.Attributes.GetName())
	assert.Equal(t, expectedCapabilities, bundleId.Data.Relationships.BundleIdCapabilities.Data)
	assert.NoError(t, err)
}

func TestCreateAndDeleteBundleId(t *testing.T) {
	identifier := "com.fintreal.test." + GenerateRandomString()
	name := "Integration Test Bundle ID"
	input := *openapi.NewBundleIdCreateRequest(
		*openapi.NewBundleIdCreateRequestData(
			"bundleIds",
			*openapi.NewBundleIdCreateRequestDataAttributes(name, "IOS", identifier),
		),
	)

	data, _, err := apiClient.BundleIdsAPI.BundleIdsCreateInstance(context.Background()).BundleIdCreateRequest(input).Execute()

	assert.Equal(t, identifier, data.Data.Attributes.GetIdentifier())
	assert.Equal(t, name, data.Data.Attributes.GetName())
	assert.NoError(t, err)

	_, err = apiClient.BundleIdsAPI.BundleIdsDeleteInstance(context.Background(), data.Data.GetId()).Execute()
	assert.NoError(t, err)
}
