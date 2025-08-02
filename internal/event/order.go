package event

import "time"

// OrderCreated struct, JSON (de)serialization

// OrderCreated defines the payload for an order creation event
type OrderCreated struct {
	ID        string    `json:"id"`        // unique order ID
	Status    string    `json:"status"`    // e.g. "created"
	TimeStamp time.Time `json:"timestamp"` // when the event occurred
}
