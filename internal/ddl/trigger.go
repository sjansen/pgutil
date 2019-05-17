package ddl

import "strings"

func (t *Trigger) addEvent(event string) {
	event = strings.ToUpper(event)
	if t.Events == nil {
		t.Events = []*TriggerEvent{
			{Event: event},
		}
	} else {
		t.Events = append(
			t.Events,
			&TriggerEvent{Event: event},
		)
	}
}

func (t *Trigger) setCalled(called string) {
	called = strings.ToUpper(
		collapseWhitespace(called),
	)
	t.Called = called
}

func (t *Trigger) setForEach(forEach string) {
	t.ForEach = strings.ToUpper(forEach)
}

func (t *Trigger) setName(name string) {
	t.Name = name
}

func (t *Trigger) setTable(table string) {
	t.Table = table
}

func (t *Trigger) setTiming(timing string) {
	timing = strings.ToUpper(
		collapseWhitespace(timing),
	)
	t.Timing = timing
}
