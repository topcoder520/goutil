package goutil

import (
	"time"
)

//Parse_datetime_to_timestr 把 datetime 转换成 时间字符串
//t: datetime 时间，比如：2019-09-17 09:45:42.5962359 +0800 CST m=+0.003989201
//flag: 标识位，决定输出的时间字符串的格式
func Parse_datetime_to_timestr(t time.Time, flag int) string {
	var time_str string
	if flag == 1 {
		time_str = t.Format("2006-01-02 15:04:05")
	} else if flag == 2 {
		time_str = t.Format("2006-01-02 15:04")
	} else if flag == 3 {
		time_str = t.Format("2006-01-02")
	} else if flag == 4 {
		time_str = t.Format("2006.01.02 15:04:05")
	} else if flag == 6 {
		time_str = t.Format("2006.01.02 15:04")
	} else {
		time_str = t.Format("2006.01.02")
	}
	return time_str
}

//Parse_datetime_to_timestamp  把 datetime 转换成时间戳
//t: datetime 时间
func Parse_datetime_to_timestamp(t time.Time) int64 {
	return t.Unix()
}

//Parse_timestr_to_datetime 把时间字符串转换成 datetime 时间
//time_str: 时间字符串，比如：2019-09-17 15:04:05
//flag: 标识位，决定输入的时间字符串的格式
func Parse_timestr_to_datetime(time_str string, flag int) time.Time {
	if flag == 1 {
		t, error1 := time.Parse("2006-01-02 15:04:05", time_str)
		if error1 != nil {
			panic(error1)
		}
		return t
	} else if flag == 2 {
		t, error2 := time.Parse("2006-01-02 15:04", time_str)
		if error2 != nil {
			panic(error2)
		}
		return t
	} else if flag == 3 {
		t, error3 := time.Parse("2006-01-02", time_str)
		if error3 != nil {
			panic(error3)
		}
		return t
	} else if flag == 4 {
		t, error4 := time.Parse("2006.01.02 15:04:05", time_str)
		if error4 != nil {
			panic(error4)
		}
		return t
	} else if flag == 5 {
		t, error5 := time.Parse("2006.01.02 15:04", time_str)
		if error5 != nil {
			panic(error5)
		}
		return t
	} else {
		t, err := time.Parse("2006.01.02", time_str)
		if err != nil {
			panic(err)
		}
		return t
	}
}

//Parse_timestr_to_timestamp 把时间字符串转换成时间戳
//time_str: 时间字符串，比如：2019-09-17 09:45:42
//flag: 标识位，决定传入的时间字符串的格式
func Parse_timestr_to_timestamp(time_str string, flag int) int64 {
	var t int64
	loc, _ := time.LoadLocation("Local")
	if flag == 1 {
		t1, _ := time.ParseInLocation("2006.01.02 15:04:05", time_str, loc)
		t = t1.Unix()
	} else if flag == 2 {
		t1, _ := time.ParseInLocation("2006-01-02 15:04", time_str, loc)
		t = t1.Unix()
	} else if flag == 3 {
		t1, _ := time.ParseInLocation("2006-01-02", time_str, loc)
		t = t1.Unix()
	} else if flag == 4 {
		t1, _ := time.ParseInLocation("2006.01.02", time_str, loc)
		t = t1.Unix()
	} else {
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, loc)
		t = t1.Unix()
	}
	return t
}

//Parse_timestamp_to_timestr 把时间戳转换成时间字符串
//stamp: 时间戳，比如：1568685105，调用此方法时需先声明为 int64 类型
//flag: 标识位，决定输入的时间字符串的格式
func Parse_timestamp_to_timestr(stamp int64, flag int) string {
	var time_str string
	if flag == 1 {
		time_str = time.Unix(stamp, 0).Format("2006-01-02")
	} else if flag == 2 {
		time_str = time.Unix(stamp, 0).Format("2006-01-02 15:04:05")
	} else if flag == 3 {
		time_str = time.Unix(stamp, 0).Format("2006-01-02 15:04")
	} else if flag == 4 {
		time_str = time.Unix(stamp, 0).Format("2006.01.02 15:04:05")
	} else if flag == 5 {
		time_str = time.Unix(stamp, 0).Format("2006.01.02 15:04")
	} else {
		time_str = time.Unix(stamp, 0).Format("2006.01.02")
	}
	return time_str
}

//Parse_timestamp_to_datetime 时间戳转换成 datetime 时间
func Parse_timestamp_to_datetime(t int64) time.Time {
	return time.Unix(t, 0)
}

//Get_after_day 获取多少天,多少月或者多少年之前或之后的时间
//day_range: 间隔的天数，月数或者年份数
//flag: 决定是取天数，月数还是年数  1:年 2:月 3:日
func Get_after_day(day_range, flag int) time.Time {
	now := time.Now()
	var tmp_day time.Time
	if flag == 1 {
		tmp_day = now.AddDate(day_range, 0, 0)
	} else if flag == 2 {
		tmp_day = now.AddDate(0, day_range, 0)
	} else {
		tmp_day = now.AddDate(0, 0, day_range)
	}
	return tmp_day
}

//get_after_time  获取多少小时，分钟及秒之前或之后的时间
//time_range: 时间差，比如：
//		10h     获取10小时之后的时间
//		-10h     获取10小时之前的时间
//		10m        获取10分钟之后的时间
//		-10m    获取10分钟之后的时间
//		10s        获取10秒之后的时间
//		-10s    获取10秒之后的时间
func Get_after_time(time_range string) time.Time {
	m, _ := time.ParseDuration(time_range)
	tmp := time.Now().Add(m)
	return tmp
}
