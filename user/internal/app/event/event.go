package event

import "time"

type EventBase struct {
	AggregateId string
	Author      string
	Timestamp   time.Time
}

type Event interface {
	GetAggregateId() string
	GetAuthor() string
	GetTimestamp() time.Time
}

func NewEventBase(aggregateId string, author string) EventBase {
	return EventBase{
		AggregateId: aggregateId,
		Author:      author,
		Timestamp:   time.Now(),
	}
}

func (e EventBase) GetAggregateId() string {
	return e.AggregateId
}

func (e EventBase) GetAuthor() string {
	return e.Author
}

func (e EventBase) GetTimestamp() time.Time {
	return e.Timestamp
}
