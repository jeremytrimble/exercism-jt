package clock

import "fmt"

const testVersion = 4

// Clock represents a time-of-day with a range of 24 hours and a precision of 1
// minute.
type Clock struct {
    minutes int
}

func normalize(minutes int) int {

    if minutes >= 60*24 {
        days := minutes / (60*24)
        minutes -= days * 60*24
    }

    if minutes < 0 {
        days := -minutes / (60*24)
        days += 1
        minutes += days * 60*24
    }

    return minutes
}


// New creates a new Clock correpsonding to given time of day.
func New(hour, minute int) Clock {
    return Clock{ normalize( hour * 60 + minute) };
}

// String returns the Clock's time as a string in 24-hour hh:mm form.
func (c Clock) String() string {
    hour   := c.minutes / 60
    minute := c.minutes % 60
    return fmt.Sprintf("%02d:%02d", hour, minute);
}

// Add returns a new Clock, advanced by the given number of minutes.
func (c Clock) Add(minutes int) Clock {
    return Clock{ normalize(c.minutes + minutes) }
}

