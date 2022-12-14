/*
API Облачных серверов

Testing NetworkServiceApiService

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

func Test_begetOpenapiVps_NetworkServiceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test NetworkServiceApiService NetworkServiceCreatePrivateNetwork", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.NetworkServiceApi.NetworkServiceCreatePrivateNetwork(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test NetworkServiceApiService NetworkServiceGetNetworkInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.NetworkServiceApi.NetworkServiceGetNetworkInfo(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test NetworkServiceApiService NetworkServiceOrderIpAddress", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.NetworkServiceApi.NetworkServiceOrderIpAddress(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test NetworkServiceApiService NetworkServiceRemoveIpAddress", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ipAddress string

        resp, httpRes, err := apiClient.NetworkServiceApi.NetworkServiceRemoveIpAddress(context.Background(), ipAddress).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test NetworkServiceApiService NetworkServiceSuggestPrivateAddress", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var networkId string

        resp, httpRes, err := apiClient.NetworkServiceApi.NetworkServiceSuggestPrivateAddress(context.Background(), networkId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
