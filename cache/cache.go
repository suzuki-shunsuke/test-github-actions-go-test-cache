package cache

import (
	"time"

	"github.com/suzuki-shunsuke/go-ci-env/cienv"
)

func Sleep(a time.Duration) {
	time.Sleep(a * time.Second)
}

func Foo() string {
	platform := cienv.Get()
	if platform == nil {
		return ""
	}
	return platform.Branch()
}
