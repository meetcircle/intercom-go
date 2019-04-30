package intercom

import (
	"io/ioutil"
	"testing"
	"time"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

func TestEventAPISave(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events"}
	api := EventAPI{httpClient: &http}
	event := Event{UserID: "27", CreatedAt: int64(time.Now().Unix()), EventName: "govent"}
	api.save(&event)
}

func TestEventAPISaveFail(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events", shouldFail: true}
	api := EventAPI{httpClient: &http}
	event := Event{UserID: "444", CreatedAt: int64(time.Now().Unix()), EventName: "govent"}
	err := api.save(&event)
	if herr, ok := err.(interfaces.HTTPError); ok && herr.Code != "not_found" {
		t.Errorf("Error not returned")
	}
}

func TestEventAPIList(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events", fixtureFilename: "fixtures/event_list.json"}
	api := EventAPI{httpClient: &http}
	params := eventParams{UserID: "444"}
	eventList, _ := api.list(params)
	if len(eventList.Events) != 2 {
		t.Errorf("Error retrieving event list")
	}
	if eventList.Events[0].ID != "1234" {
		t.Errorf("Incorrect EventID")
	}
	if eventList.Events[1].ID != "5678" {
		t.Errorf("Incorrect EventID")
	}
}

func TestEventAPIListFail(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events", shouldFail: true}
	api := EventAPI{httpClient: &http}
	params := eventParams{UserID: "444"}
	_, err := api.list(params)
	if herr, ok := err.(interfaces.HTTPError); ok && herr.Code != "not_found" {
		t.Errorf("Error not returned")
	}
}

func TestEventAPISummary(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events", fixtureFilename: "fixtures/event_summaries.json"}
	api := EventAPI{httpClient: &http}
	params := eventParams{UserID: "444"}
	eventList, _ := api.summary(params)
	if len(eventList.Events) != 2 {
		t.Errorf("Error retrieving event list")
	}
	if eventList.Events[0].Count != 1234 {
		t.Errorf("Incorrect Count")
	}
	if eventList.Events[1].Count != 5678 {
		t.Errorf("Incorrect Count")
	}
}

func TestEventAPISummaryFail(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events", shouldFail: true}
	api := EventAPI{httpClient: &http}
	params := eventParams{UserID: "444"}
	_, err := api.summary(params)
	if herr, ok := err.(interfaces.HTTPError); ok && herr.Code != "not_found" {
		t.Errorf("Error not returned")
	}
}

type TestEventHTTPClient struct {
	TestHTTPClient
	t *testing.T

	fixtureFilename string

	expectedURI string
	shouldFail  bool
}

func (t TestEventHTTPClient) Post(uri string, event interface{}) ([]byte, error) {
	if uri != "/events" {
		t.t.Errorf("Wrong endpoint called")
	}
	if t.shouldFail {
		err := interfaces.HTTPError{StatusCode: 404, Code: "not_found", Message: "User Not Found"}
		return nil, err
	}
	return nil, nil
}

func (t TestEventHTTPClient) Get(uri string, event interface{}) ([]byte, error) {
	if uri != "/events" {
		t.t.Errorf("Wrong endpoint called")
	}
	if t.shouldFail {
		err := interfaces.HTTPError{StatusCode: 404, Code: "not_found", Message: "User Not Found"}
		return nil, err
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
