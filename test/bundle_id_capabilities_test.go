package test

import (
	"context"
	"testing"

	openapi "github.com/fintreal/app-store-sdk-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteBundleIdCapability(t *testing.T) {
	key := "APPLE_ID_AUTH_APP_CONSENT"
	optionKey := "PRIMARY_APP_CONSENT"
	capabilityType := "APPLE_ID_AUTH"
	bundleId := "S65568LTPR"
	input := *openapi.NewBundleIdCapabilityCreateRequest(
		*openapi.NewBundleIdCapabilityCreateRequestData(
			"bundleIdCapabilities",
			*openapi.NewBundleIdCapabilityCreateRequestDataAttributes(
				openapi.CapabilityType(capabilityType),
				[]openapi.CapabilitySetting{{
					Key: &key,
					Options: []openapi.CapabilityOption{{
						Key: &optionKey,
					}},
				}},
			),
			*openapi.NewBundleIdCapabilityCreateRequestDataRelationships(
				*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(
					*openapi.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("bundleIds", bundleId),
				),
			),
		),
	)
	data, _, err := apiClient.BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance(context.Background()).BundleIdCapabilityCreateRequest(input).Execute()
	assert.Equal(t, capabilityType, string(data.Data.Attributes.GetCapabilityType()))
	assert.Equal(t, bundleId+"_"+capabilityType, data.Data.GetId())

	_, err = apiClient.BundleIdCapabilitiesAPI.BundleIdCapabilitiesDeleteInstance(context.Background(), data.Data.GetId()).Execute()
	assert.NoError(t, err)
}
