package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	fmt.Println(Now())
	now := time.Now()
	fmt.Println(Format(now))
	fmt.Println(Format2Milli(now))
	tt, err := Parse("2021-12-12 12:12:12")
	if err != nil {
		return
	}
	fmt.Println(tt.Unix() == 1639282332)
	fmt.Println(Format(Unix2Time(1639282332)))
}
