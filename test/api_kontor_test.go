/*
SpedV API

Testing KontorAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/NiklasSchmitt/spedv-go-sdk"
)

func Test_openapi_KontorAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KontorAPIService V1KontorGameJobsAvailableGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var game FPHSpedVAPIEnumsGameEnum

		resp, httpRes, err := apiClient.KontorAPI.V1KontorGameJobsAvailableGet(context.Background(), game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KontorAPIService V1KontorGameJobsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var game FPHSpedVAPIEnumsGameEnum

		resp, httpRes, err := apiClient.KontorAPI.V1KontorGameJobsGet(context.Background(), game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KontorAPIService V1KontorGameJobsJobidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var game FPHSpedVAPIEnumsGameEnum
		var jobid int32

		resp, httpRes, err := apiClient.KontorAPI.V1KontorGameJobsJobidGet(context.Background(), game, jobid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KontorAPIService V1KontorGamePartsAvailableGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var game FPHSpedVAPIEnumsGameEnum

		resp, httpRes, err := apiClient.KontorAPI.V1KontorGamePartsAvailableGet(context.Background(), game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KontorAPIService V1KontorGamePartsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var game FPHSpedVAPIEnumsGameEnum

		resp, httpRes, err := apiClient.KontorAPI.V1KontorGamePartsGet(context.Background(), game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KontorAPIService V1KontorGamePartsPartidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var game FPHSpedVAPIEnumsGameEnum
		var partid int32

		resp, httpRes, err := apiClient.KontorAPI.V1KontorGamePartsPartidGet(context.Background(), game, partid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KontorAPIService V1KontorGamePartsPartidJobsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var game FPHSpedVAPIEnumsGameEnum
		var partid int32

		resp, httpRes, err := apiClient.KontorAPI.V1KontorGamePartsPartidJobsGet(context.Background(), game, partid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KontorAPIService V1KontorGameTrailersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var game FPHSpedVAPIEnumsGameEnum

		resp, httpRes, err := apiClient.KontorAPI.V1KontorGameTrailersGet(context.Background(), game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}