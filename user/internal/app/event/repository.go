package event

type EventRepository interface {
	Save(streamId string, events []Event) error
	Get(streamId string) []Event
}
