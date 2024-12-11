package service_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreparelho/star-wars-challenge-golang/planet/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const endpoint string = "https://swapi.dev/api/planets/?page="

type HTTPClientMock struct {
	mock.Mock
}

func (mock *HTTPClientMock) Get(url string) (*http.Response, error) {
	args := mock.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetAllPlanets(test *testing.T) {
	mockClient := &HTTPClientMock{}

	planetService := &service.GetAllPlanetsService{
		Client: mockClient,
	}

	test.Run("Success - Valid response", func(test *testing.T) {
		mockResponseBody := `{"results": [{"name": "Earth"}, {"name": "Mars"}]}`

		mockResponse := &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(mockResponseBody)),
		}
		mockClient.On("Get", endpoint).Return(mockResponse, nil)

		context, _ := gin.CreateTestContext(httptest.NewRecorder())
		response, err := planetService.GetAllPlanets(context)
		if err != nil {
			test.Errorf("error test")
		}

		assert.NotEmpty(test, response)
	})
}
