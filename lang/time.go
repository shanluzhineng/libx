package lang

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// convert time.Now() to *time.Time
func NowToPtr() *time.Time {
	return TimeToPtr(time.Now())
}

// convert time to *time.Time
func TimeToPtr(t time.Time) *time.Time {
	return &t
}

// parse time string to time using China timeZone
func ParseTimeToChinaTimezone(layoutList []string, timeString string) (*time.Time, error) {
	if len(layoutList) <= 0 {
		return nil, errors.New("layoutList不能为空,或者长度不能小于0")
	}
	for _, eachLayout := range layoutList {
		timeValue, err := time.ParseInLocation(eachLayout, timeString, ChinaTimezone)
		if err == nil {
			return &timeValue, nil
		}
	}
	return nil, fmt.Errorf("无效的时间,时间格式必须是%s", strings.Join(layoutList, ","))
}
