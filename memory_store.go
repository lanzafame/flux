package cqrs

//InMemory implementation of the event store
type InMemoryEventStore struct {
	events     []Event
	aggregates map[string][]Event
}

func (store *InMemoryEventStore) GetEvents(aggregateId string) []Event {
	return store.aggregates[aggregateId]
}

func (store *InMemoryEventStore) GetEventsFrom(aggregateId, eventId string, count int) []Event {
	if events, ok := store.aggregates[aggregateId]; ok {
		for i, e := range events {
			if e.Id == eventId {
				if i+count > len(events) {
					return events[i:]
				} else {
					return events[i : i+count]
				}
			}
		}
	}
	return []Event{}
}

func (store *InMemoryEventStore) GetAllEventsFrom(eventId string, count int) []Event {
	for i, e := range store.events {
		if e.Id == eventId {
			if i+count > len(store.events) {
				return store.events[i:]
			} else {
				return store.events[i : i+count]
			}

		}
	}
	return []Event{}
}

func (store *InMemoryEventStore) SaveEvents(aggregateId string, events []Event) error {
	if _, ok := store.aggregates[aggregateId]; !ok {
		store.aggregates[aggregateId] = make([]Event, 0)
	}
	store.aggregates[aggregateId] = append(store.aggregates[aggregateId], events...)
	store.events = append(store.events, events...)
	return nil
}

func (store *InMemoryEventStore) GetAllEvents() []Event {
	events := make([]Event, len(store.events))
	copy(events, store.events)
	return events
}

func (store *InMemoryEventStore) GetEvent(id string) Event {
	var event Event
	for _, e := range store.events {
		if e.Id == id {
			return e
		}
	}
	return event
}

func NewEventStore() EventStore {
	return &InMemoryEventStore{make([]Event, 0), make(map[string][]Event)}
}
