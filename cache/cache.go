package cache

import "time"

func Sleep(a time.Duration) {
	time.Sleep(a * time.Second)
}
