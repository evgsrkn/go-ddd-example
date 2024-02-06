package event

type UserBecameAdmin struct {
	EventBase
}

func NewUserBecameAdminEvent(aggregateId string, authorId string) Event {
	return UserBecameAdmin{
		NewEventBase(aggregateId, authorId),
	}
}
