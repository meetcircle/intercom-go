package intercom

import (
	"errors"
	"testing"
	"time"
)

func TestEventSaveFail(t *testing.T) {
	eventService := EventService{Repository: TestEventAPI{t: t, body: failBody}}
	err := eventService.Save(&Event{})
	if err.Error() != "Missing Identifier" {
		t.Errorf("Error not propagated")
	}
}

func failBody(t *testing.T, event Event) error {
	return errors.New("Missing Identifier")
}

func TestEventSave(t *testing.T) {
	eventService := EventService{Repository: TestEventAPI{t: t, body: successBody}}
	event := Event{}
	event.UserID = "27"
	event.EventName = "govent"
	event.CreatedAt = int64(time.Now().Unix())
	event.Metadata = map[string]interface{}{"is_cool": true}
	eventService.Save(&event)
}

func successBody(t *testing.T, event Event) error {
	if event.UserID != "27" {
		t.Errorf("UserID not set")
	}
	if event.EventName != "govent" {
		t.Errorf("EventName not set")
	}
	if event.Metadata["is_cool"] != true {
		t.Errorf("Metadata not set")
	}
	return nil
}

func TestEventList(t *testing.T) {
	eventList, _ := (&EventService{Repository: TestEventAPI{t: t, eventList: eventListSuccess}}).ListByUserID("1")
	if eventList.Events[0].UserID != "1" {
		t.Errorf("no events for user id found")
	}
	eventList, _ = (&EventService{Repository: TestEventAPI{t: t, eventList: eventListSuccess}}).ListByEmail("2")
	if eventList.Events[0].Email != "2" {
		t.Errorf("no events for email found")
	}
	eventList, _ = (&EventService{Repository: TestEventAPI{t: t, eventList: eventListSuccess}}).ListByID("3")
	if eventList.Events[0].IntercomUserID != "3" {
		t.Errorf("no events for Intercom user id found")
	}

	_, err := (&EventService{Repository: TestEventAPI{t: t, eventList: eventListFail}}).ListByID("3")
	if err.Error() != "event list fail" {
		t.Errorf("error not propagated")
	}
}

func eventListSuccess(params eventParams) (EventList, error) {
	return EventList{Events: []Event{{UserID: params.UserID, IntercomUserID: params.IntercomUserID, Email: params.Email}}}, nil
}

func eventListFail(_ eventParams) (EventList, error) {
	return EventList{}, errors.New("event list fail")
}

func TestEventSummary(t *testing.T) {
	eventSummary, _ := (&EventService{Repository: TestEventAPI{t: t, eventSummary: eventSummarySuccess}}).SummaryByUserID("1")
	if eventSummary.UserID != "1" {
		t.Errorf("no events for user id found")
	}
	eventSummary, _ = (&EventService{Repository: TestEventAPI{t: t, eventSummary: eventSummarySuccess}}).SummaryByEmail("2")
	if eventSummary.Email != "2" {
		t.Errorf("no events for email found")
	}
	eventSummary, _ = (&EventService{Repository: TestEventAPI{t: t, eventSummary: eventSummarySuccess}}).SummaryByID("3")
	if eventSummary.IntercomUserID != "3" {
		t.Errorf("no events for Intercom user id found")
	}

	_, err := (&EventService{Repository: TestEventAPI{t: t, eventSummary: eventSummaryFail}}).SummaryByID("3")
	if err.Error() != "event summary fail" {
		t.Errorf("error not propagated")
	}
}

func eventSummarySuccess(params eventParams) (EventSummaries, error) {
	return EventSummaries{UserID: params.UserID, IntercomUserID: params.IntercomUserID, Email: params.Email}, nil
}

func eventSummaryFail(_ eventParams) (EventSummaries, error) {
	return EventSummaries{}, errors.New("event summary fail")
}

type TestEventAPI struct {
	t            *testing.T
	body         func(*testing.T, Event) error
	eventList    func(eventParams) (EventList, error)
	eventSummary func(eventParams) (EventSummaries, error)
}

func (t TestEventAPI) save(event *Event) error {
	return t.body(t.t, *event)
}

func (t TestEventAPI) list(params eventParams) (EventList, error) {
	return t.eventList(params)
}
func (t TestEventAPI) summary(params eventParams) (EventSummaries, error) {
	return t.eventSummary(params)
}
