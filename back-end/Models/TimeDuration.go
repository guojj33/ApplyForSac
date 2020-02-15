package Models

import (
	"encoding/json"
	"time"
)

type TimeDuration struct {
	StartTime time.Time
	EndTime   time.Time
}

func (timeDuration *TimeDuration) String() string {
	b, err := json.Marshal(*timeDuration)
	if err == nil {
		return string(b)
	} else {
		return ""
	}
}

func isTimeDurationValid(timeDuration TimeDuration) bool {
	if timeDuration.StartTime.Before(timeDuration.EndTime) {
		return true
	}
	return false
}

func (timeDuration TimeDuration) compareWith(timeDuration2 TimeDuration) TimeDurCompType {
	if timeDuration.EndTime.Before(timeDuration2.StartTime) || timeDuration.EndTime.Equal(timeDuration2.StartTime) {
		return TimeDurComp_Earlier
	} else if timeDuration.StartTime.After(timeDuration2.EndTime) || timeDuration.StartTime.Equal(timeDuration2.EndTime) {
		return TimeDurComp_Later
	} else {
		return TimeDurComp_Overlapping
	}
}
