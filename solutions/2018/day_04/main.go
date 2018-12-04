package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/wcharczuk/advent/pkg/fileutil"
	"github.com/wcharczuk/advent/pkg/log"
)

func main() {
	var events []Event
	err := fileutil.ReadByLines("./input", func(line string) error {
		var event Event
		if strings.Contains(line, "#") {
			event = NewEventFromBeginShift(line)
		} else {
			event = NewEventFromAction(line)
		}
		events = append(events, event)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	sort.Sort(Events(events))

	// guard => minute => asleep
	asleepByMinute := map[int]map[int]int{}

	var sleepStart time.Time
	var lastGuardID int
	for _, e := range events {
		log.Context("debug").Print(e)
		if e.Action == "begins shift" {
			lastGuardID = e.GuardID
		} else if e.Action == "falls asleep" {
			sleepStart = e.Timestamp
		} else if e.Action == "wakes up" && !sleepStart.IsZero() {
			minutes := int(e.Timestamp.Sub(sleepStart) / time.Minute)
			var cursor time.Time
			for minute := 0; minute < minutes; minute++ {
				cursor = sleepStart.Add(time.Duration(minute) * time.Minute)
				if _, ok := asleepByMinute[lastGuardID]; !ok {
					asleepByMinute[lastGuardID] = map[int]int{}
				}
				if current, ok := asleepByMinute[lastGuardID][cursor.Minute()]; ok {
					asleepByMinute[lastGuardID][cursor.Minute()] = current + 1
				} else {
					asleepByMinute[lastGuardID][cursor.Minute()] = 1
				}
			}
			sleepStart = time.Time{}
		}
	}

	var worstGuardID, maxTotal int
	for guardID, byMinute := range asleepByMinute {
		var total int
		for _, minuteTotal := range byMinute {
			total = total + minuteTotal
		}
		if maxTotal < total {
			worstGuardID = guardID
			maxTotal = total
		}
	}

	var maxMinuteTotal int
	for _, minuteTotal := range asleepByMinute[worstGuardID] {
		if maxMinuteTotal < minuteTotal {
			maxMinuteTotal = minuteTotal
		}
	}

	var minutes []int
	for minute, minuteTotal := range asleepByMinute[worstGuardID] {
		if minuteTotal == maxMinuteTotal {
			minutes = append(minutes, minute)
		}
	}

	sort.Ints(minutes)

	log.Context("solution").Printf("guard: %d for %d minutes, worst %d", worstGuardID, maxTotal, minutes[0])
}

func NewEventFromBeginShift(line string) Event {
	var event Event
	var err error
	matches := regexExtract(line, `\[(.*)\] Guard #([0-9]*) (.*)`)
	event.Timestamp, err = time.Parse("2006-01-02 15:04", matches[1])
	if err != nil {
		log.Fatal(err)
	}
	event.GuardID, err = strconv.Atoi(matches[2])
	event.Action = matches[3]
	return event
}

func NewEventFromAction(line string) Event {
	var event Event
	var err error
	matches := regexExtract(line, `\[(.*)\] (.*)`)
	event.Timestamp, err = time.Parse("2006-01-02 15:04", matches[1])
	if err != nil {
		log.Fatal(err)
	}
	event.Action = matches[2]
	return event
}

type Events []Event

func (e Events) Len() int {
	return len(e)
}

func (e Events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Events) Less(i, j int) bool {
	return e[i].Timestamp.Before(e[j].Timestamp)
}

// Event is a thing that happens
type Event struct {
	Timestamp time.Time
	GuardID   int
	Action    string
}

func (e Event) String() string {
	return fmt.Sprintf("[%s] #%d %s", e.Timestamp.Format("2006-01-02 15:04"), e.GuardID, e.Action)
}

func regexExtract(corpus, expr string) []string {
	re := regexp.MustCompile(expr)
	allResults := re.FindAllStringSubmatch(corpus, -1)
	results := []string{}
	for _, resultSet := range allResults {
		for _, result := range resultSet {
			results = append(results, result)
		}
	}
	return results
}
