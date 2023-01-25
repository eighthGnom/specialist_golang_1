package main

import (
	"fmt"
	"sort"
	"strings"
)

type Tracker struct {
	shedulle map[string][]string
	onOff    bool
}

func New() *Tracker {
	m := make(map[string][]string)
	return &Tracker{shedulle: m, onOff: false}
}

func (t *Tracker) StartApp() bool {
	t.onOff = true
	return t.onOff
}

func (t *Tracker) Quit() bool {
	t.onOff = false
	return t.onOff
}

func (t *Tracker) Add(date, newEvent string) {
	eventList, ok := t.shedulle[date]
	if !ok {
		eventList = make([]string, 0, 4)
		t.shedulle[date] = append(eventList, newEvent)
		return
	}
	for _, event := range eventList {
		if event == newEvent {
			return
		}
	}
	t.shedulle[date] = append(eventList, newEvent)
	return

}

func (t *Tracker) DelAll(date string) string {
	eventList, ok := t.shedulle[date]
	l := len(eventList)
	if !ok || l == 0 {
		return fmt.Sprintf("Deleted %d events", l)
	}
	delete(t.shedulle, date)
	return fmt.Sprintf("Deleted %d events", l)
}

func (t *Tracker) DelOne(date, delEvent string) string {
	eventList, ok := t.shedulle[date]
	l := len(eventList)
	if !ok || l == 0 {
		return fmt.Sprintf("Event not found")
	}
	for id, event := range eventList {
		if event == delEvent {
			t.shedulle[date][id] = t.shedulle[date][l-1]
			t.shedulle[date] = t.shedulle[date][:l-1]
			return fmt.Sprintf("Deleted successfully")
		}
	}
	return fmt.Sprintf("Event not found")
}

func (t *Tracker) Find(date string) string {
	eventList, ok := t.shedulle[date]
	if !ok {
		return fmt.Sprintf("Event not found")
	}
	sort.Strings(eventList)
	return strings.Join(eventList, "\n")
}

func (t *Tracker) Print() {
	l := len(t.shedulle)
	keys := make([]string, 0, l)
	for date, _ := range t.shedulle {
		sort.Strings(t.shedulle[date])
		keys = append(keys, date)
	}
	sort.Strings(keys)
	for _, date := range keys {
		for _, event := range t.shedulle[date] {
			fmt.Printf("%s %s\n", printDate(date), event)
		}
	}

}
