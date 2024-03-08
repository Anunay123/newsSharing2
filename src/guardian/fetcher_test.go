package guardian

import (
	"net/http"
	"testing"
	"time"
)

/*
	Similar test cases could be written for nyTimes fetcher as well.
*/

// MockNetworkService mocks the network.CallRestService function
type MockNetworkService struct{}

func (m *MockNetworkService) CallRestService(requestMethod, httpEndPoint, requestType, responseType string,
	requestHeaders, queryParams map[string]string, requestCookies []*http.Cookie, body interface{}, client http.Client) ([]byte, error, int) {
	// Simulating a successful response
	responseData := []byte(`{
  "response": {
    "status": "ok",
    "total": 30072,
    "startIndex": 1,
    "results": [
      {
        "id": "world/2024/jan/24/france-debates-plan-to-enshrine-abortion-as-constitutional-right",
        "type": "article",
        "sectionId": "world",
        "sectionName": "World news",
        "webPublicationDate": "2024-01-24T18:25:53Z",
        "webTitle": "France debates plan to enshrine abortion as constitutional right",
        "webUrl": "https://www.theguardian.com/world/2024/jan/24/france-debates-plan-to-enshrine-abortion-as-constitutional-right",
        "apiUrl": "https://content.guardianapis.com/world/2024/jan/24/france-debates-plan-to-enshrine-abortion-as-constitutional-right",
        "isHosted": false,
        "pillarId": "pillar/news",
        "pillarName": "News"
      }
    ]
  }
}`)
	return responseData, nil, http.StatusOK
}

type MockNetworkService2 struct{}

func (m *MockNetworkService2) CallRestService(requestMethod, httpEndPoint, requestType, responseType string,
	requestHeaders, queryParams map[string]string, requestCookies []*http.Cookie, body interface{}, client http.Client) ([]byte, error, int) {
	// Simulating a successful response
	time.Sleep(5 * time.Second)
	return nil, nil, http.StatusBadRequest
}

func TestGetNewsData_Positive(t *testing.T) {
	query := "debates"
	page := "1"
	newsDataChannel := make(chan *NewsData)

	// Mocking network call
	networkService := &MockNetworkService{}
	//responseBytes, _, _ := networkService.CallRestService("", "", "", "", nil, nil, nil, nil, http.Client{})

	// Invoke function under test
	go GetNewsData(query, page, newsDataChannel, networkService)

	// Verify the result
	select {
	case newsData := <-newsDataChannel:
		if newsData.GetResponse().Status != "ok" {
			t.Errorf("status not ok")
		}
		if len(newsData.GetResponse().GetResults()) <= 0 {
			t.Errorf("no content from the API")
		}
		if newsData.GetResponse().GetResults()[0].GetWebTitle() != "France debates plan to enshrine abortion as constitutional right" {
			t.Errorf("Expected title 'France debates plan to enshrine abortion as constitutional right', got %s", newsData.GetResponse().GetResults()[0].GetWebTitle())
		}
	case <-time.After(2 * time.Second):
		t.Errorf("Timeout waiting for news data")
	}
}

func TestGetNewsData_NetworkError(t *testing.T) {
	query := "debates"
	page := "1"
	newsDataChannel := make(chan *NewsData)

	// Mocking network call to return an error
	networkService := &MockNetworkService2{}
	//_, expectedError, _ := networkService.CallRestService("", "", "", "", nil, nil, nil, nil, http.Client{})

	// Invoke function under test
	go GetNewsData(query, page, newsDataChannel, networkService)

	// Verify the result
	select {
	case <-newsDataChannel:
		t.Error("Expected no data to be received due to network error")
	case <-time.After(time.Second):
		// No data received within the timeout, this is expected
	}
}
