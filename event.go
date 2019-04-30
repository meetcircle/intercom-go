package intercom

import (
	"fmt"
	"time"
)

// EventService handles interactions with the API through an EventRepository.
type EventService struct {
	Repository EventRepository
}

// An Event represents a new event that happens to a User.
type Event struct {
	ID             string                 `json:"id,omitempty"`
	Email          string                 `json:"email,omitempty"`
	UserID         string                 `json:"user_id,omitempty"`
	IntercomUserID string                 `json:"intercom_user_id,omitempty"`
	EventName      string                 `json:"event_name,omitempty"`
	CreatedAt      int64                  `json:"created_at,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

type EventPages struct {
	Next string `json:"next,omitempty"`
}

type EventList struct {
	Pages  EventPages `json:"pages,omitempty"`
	Events []Event
}

type EventSummaries struct {
	Email          string `json:"email,omitempty"`
	UserID         string `json:"user_id,omitempty"`
	IntercomUserID string `json:"intercom_user_id,omitempty"`
	Events         []EventSummary
}

type EventSummary struct {
	Name        string    `json:"name,omitempty"`
	First       time.Time `json:"first,omitempty"`
	Last        time.Time `json:"last,omitempty"`
	Count       int       `json:"count,omitempty"`
	Description string    `json:"description,omitempty"`
}

type eventParams struct {
	Type           string `url:"type,omitempty"`
	UserID         string `url:"user_id,omitempty"`
	Email          string `url:"email,omitempty"`
	IntercomUserID string `url:"intercom_user_id,omitempty"`
	Summary        bool   `url:"summary,omitempty"`
}

// Save a new Event
func (e *EventService) Save(event *Event) error {
	return e.Repository.save(event)
}

func (e *EventService) ListByID(id string) (EventList, error) {
	return e.Repository.list(eventParams{IntercomUserID: id})
}

func (e *EventService) ListByUserID(userID string) (EventList, error) {
	return e.Repository.list(eventParams{UserID: userID})
}

func (e *EventService) ListByEmail(email string) (EventList, error) {
	return e.Repository.list(eventParams{Email: email})
}

func (e *EventService) SummaryByID(id string) (EventSummaries, error) {
	return e.Repository.summary(eventParams{IntercomUserID: id, Summary: true})
}

func (e *EventService) SummaryByUserID(userID string) (EventSummaries, error) {
	return e.Repository.summary(eventParams{UserID: userID, Summary: true})
}

func (e *EventService) SummaryByEmail(email string) (EventSummaries, error) {
	return e.Repository.summary(eventParams{Email: email, Summary: true})
}

func (e Event) String() string {
	return fmt.Sprintf("[intercom] event { name: %s, user_id: %s, email: %s }", e.EventName, e.UserID, e.Email)
}
