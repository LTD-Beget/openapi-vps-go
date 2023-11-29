/*
API Облачных серверов

Testing SoftwareLicenseServiceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package begetOpenapiVps

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_begetOpenapiVps_SoftwareLicenseServiceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SoftwareLicenseServiceApiService SoftwareLicenseServiceChangeLicensePlan", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var vpsId string

        resp, httpRes, err := apiClient.SoftwareLicenseServiceApi.SoftwareLicenseServiceChangeLicensePlan(context.Background(), vpsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SoftwareLicenseServiceApiService SoftwareLicenseServiceGetLicenseInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SoftwareLicenseServiceApi.SoftwareLicenseServiceGetLicenseInfo(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
