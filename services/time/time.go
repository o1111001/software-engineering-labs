package time

import "time"

func GetTime() string {
	return time.Now().Format(time.RFC3339)
}