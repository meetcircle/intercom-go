package intercom

import (
	"encoding/json"
	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// EventRepository defines the interface for working with Events through the API.
type EventRepository interface {
	save(*Event) error
	list(params eventParams) (EventList, error)
	summary(params eventParams) (EventSummaries, error)
}

// EventAPI implements EventRepository
type EventAPI struct {
	httpClient interfaces.HTTPClient
}

func (api EventAPI) save(event *Event) error {
	_, err := api.httpClient.Post("/events", event)
	return err
}

func (api EventAPI) list(params eventParams) (EventList, error) {
	eventList := EventList{}
	data, err := api.listEvents(params)
	if err != nil {
		return eventList, err
	}
	err = json.Unmarshal(data, &eventList)
	return eventList, err
}

func (api EventAPI) summary(params eventParams) (EventSummaries, error) {
	eventSummaries := EventSummaries{}
	data, err := api.listEvents(params)
	if err != nil {
		return eventSummaries, err
	}
	err = json.Unmarshal(data, &eventSummaries)
	return eventSummaries, err
}

func (api EventAPI) listEvents(params eventParams) ([]byte, error) {
	params.Type = "user"
	data, err := api.httpClient.Get("/events", params)
	if err != nil {
		return nil, err
	}
	return data, nil
}
