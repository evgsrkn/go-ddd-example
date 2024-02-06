package event

type UserActivated struct {
	EventBase
}

func NewUserActivatedEvent(aggregateId string, author string) UserActivated {
	return UserActivated{
		NewEventBase(aggregateId, author),
	}
}
