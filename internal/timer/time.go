package timer

import "time"

func GetNowTime() time.Time {
	localtion ,_ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(localtion)
}

// GetNowTime 方法是对标准库time的now方法进行封装，用于返回当前本地时间的time对象，此处封装的主要是为了便于后续对time对象做进一步的统一处理

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.Add(duration), nil
}

// GetCalculateTime 方法是从传入的字符串中解析出duration（持续时间），而在add方法中，我没可以传入获得的duration，这样就是能获取到当前时间+duration时间后的时间
// 使用time.ParseDuration 方法的主要原因是，对用户输入的字符串进行格式化处理，因为并不知道用户输入的字符串是否符合要求