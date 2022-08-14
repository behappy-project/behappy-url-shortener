package util

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/pkg/errors"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

const (
	DATETIME_FORMAT = "2006/01/02 15:04:05"
	DATE_FORMAT     = "2006/01/02"
	TIME_FORMART    = "15:04:05"
	ASIA_LOC        = "Asia/Shanghai"
)

// CTime 这个方法就是把加减时间值和时间整合到了一起
func CTime(t time.Time, timeStr string) time.Time {
	timePart, err := time.ParseDuration(timeStr)
	if err != nil {
		return t
	}
	return t.Add(timePart)
}

// Substring 字符串截取(中文乱码处理)
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if end > length {
		return string(r[start:length])
	}
	return string(r[start:end])
}

// 合并两个数组并去重
func MergeDuplicateIntArray(slice []int, elems []int) []int {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []int
	for i := range t.Iterator().C {
		result = append(result, i.(int))
	}
	return result
}

// 数组去重
func DuplicateIntArray(m []int) []int {
	s := make([]int, 0)
	smap := make(map[int]int)
	for _, value := range m {
		//计算map长度
		length := len(smap)
		smap[value] = 1
		//比较map长度, 如果map长度不相等， 说明key不存在
		if len(smap) != length {
			s = append(s, value)
		}
	}
	return s
}

// 数组取出不同元素 放入结果
// sourceList中的元素不在sourceList2中 则取到result中
func GetDifferentIntArray(sourceList, sourceList2 []int) (result []int) {
	for _, src := range sourceList {
		var find bool
		for _, target := range sourceList2 {
			if src == target {
				find = true
				continue
			}
		}
		if !find {
			result = append(result, src)
		}
	}
	return
}

// 数组存在某个数字
func ExistIntArray(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// 字符串数组存在某个字符串
func ExistStringArray(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// 数组取出不同元素 放入结果
// sourceList中的元素不在sourceList2中 则取到result中
func GetDifferentStringArray(sourceList, sourceList2 []string) (result []string) {
	for _, src := range sourceList {
		var find bool
		for _, target := range sourceList2 {
			if src == target {
				find = true
				continue
			}
		}
		if !find {
			result = append(result, src)
		}
	}
	return
}

// 合并两个字符串数组并去重
func MergeDuplicateStringArray(slice []string, elems []string) []string {
	listPId := append(slice, elems...)
	t := mapset.NewSet()
	for _, i := range listPId {
		t.Add(i)
	}
	var result []string
	for i := range t.Iterator().C {
		result = append(result, i.(string))
	}
	return result
}

// 字符串数组去重
func DuplicateStringArray(m []string) []string {
	s := make([]string, 0)
	smap := make(map[string]string)
	for _, value := range m {
		//计算map长度
		length := len(smap)
		smap[value] = value
		//比较map长度, 如果map长度不相等， 说明key不存在
		if len(smap) != length {
			s = append(s, value)
		}
	}
	return s
}

// 时间转换 将1993-12-26 10:30:00转换为time
func ParseTimeByTimeStr(str, errPrefix string) (time.Time, error) {
	p := strings.TrimSpace(str)
	if p == "" {
		return time.Time{}, errors.Errorf("%s不能为空", errPrefix)
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return time.Time{}, errors.Errorf("%s格式错误", errPrefix)
	}

	return t, nil
}

// 获取int64 当前时间戳/输入time时间戳
func ParseTimeToInt64(t ...time.Time) int64 {
	if len(t) == 0 {
		return time.Now().UnixNano() / 1e6
	} else {
		return t[0].UnixNano() / 1e6
	}
}

// 获取int64 秒
func ParseSecondTimeToInt64() int64 {
	return time.Now().Unix()
}

// 获取int64 小时
func ParseHourTimeToInt64() int64 {
	return time.Now().Unix() / 3600 * 3600
}

// 捕获异常 error
func Catch(err error) {
	if err != nil {
		panic(err)
	}
}

// 获取最近的周一
func ParseCurrentMonday(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart
}

// 返回某一天的当地时区0点
func ParseMorningTime(t time.Time) time.Time {
	s := t.Format("19931226")
	result, _ := time.ParseInLocation("19931226", s, time.Local)
	return result
}

// 当月第一天0点
func ParseFirstDayOfMonthMorning(t time.Time) time.Time {
	if t.IsZero() {
		return t
	}
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// 将指定time按时区格式化为字符串
func TimeFormat(format, location string, t time.Time) time.Time {
	// 加载时间
	loc, _ := time.LoadLocation(location)
	value := t.Format(format)
	formatTime, _ := time.ParseInLocation(format, value, loc)
	return formatTime
}

// 获取传入时间前一天的时间，不传默认是昨天
func ParseYesterdayTime(t ...time.Time) time.Time {
	if len(t) == 0 {
		return time.Now().AddDate(0, 0, -1)
	} else {
		return t[0].AddDate(0, 0, -1)
	}
}

// 把int64转换成1993-12-26 10:30:00
func ParseTimeToTimeStr(intTime int64, strfmt ...string) string {
	t := time.Unix(intTime/1e3, 0)
	defaultFmt := "2006-01-02 15:04:05"
	if len(strfmt) > 0 {
		defaultFmt = strfmt[0]
	}
	return t.Format(defaultFmt)
}

// int64 to time
func Int64ConvertToTime(intTime int64) time.Time {
	return time.Unix(intTime/1e3, 0)
}

func GetRandomInt(min, max float64) float64 {
	return math.Floor(rand.Float64()*(max-min)) + min
}

// GroupByMap 切片按key分组转map
func GroupByMap[T interface{}](values []*T, key string) map[string][]*T {
	mapObj := make(map[string][]*T)
	// 先去重
	for _, groupKey := range removeRepeatedElement(values, key) {
		// 再过滤
		arr := filter(values, key, groupKey)
		mapObj[groupKey] = arr
	}
	return mapObj
}

func filter[T interface{}](arr []*T, key, keyAsVal string) (newArr []*T) {
	newArr = make([]*T, 0)
	for _, item := range arr {
		itemVal := reflect.ValueOf(item).Elem().FieldByName(key).String()
		if itemVal == keyAsVal {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

func removeRepeatedElement[T interface{}](arr []*T, key string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		// key对应的val
		ikeyAsVal := reflect.ValueOf(arr[i]).Elem().FieldByName(key).String()
		for j := i + 1; j < len(arr); j++ {
			jkeyAsVal := reflect.ValueOf(arr[j]).Elem().FieldByName(key).String()
			if ikeyAsVal == jkeyAsVal {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, ikeyAsVal)
		}
	}
	return newArr
}
