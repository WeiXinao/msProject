package formatx

import "time"

func ToDateTimeString(milli int64) string {
	return time.UnixMilli(milli).Format(time.DateTime)
}

func ParseDateTimeString(str string) (int64, error) {
	time, err := time.Parse("2006-01-02 15:04", str)
	if err != nil {
		return 0, err
	}
	return time.UnixMilli(), nil
}