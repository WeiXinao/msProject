package formatx

import "time"

func ToDateTimeString(milli int64) string {
	return time.UnixMilli(milli).Format(time.DateTime)
}