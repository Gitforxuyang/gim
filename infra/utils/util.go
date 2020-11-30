package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

func Must(e error) {
	if e != nil {
		panic(e)
	}
}

func MustNil(e error) {

}
func StructToJson(object interface{}) (string, error) {
	str, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func StructToJsonOrError(object interface{}) (string) {
	str, err := json.Marshal(object)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

func ErrStr(err error) string {
	if err == nil {
		return "nil"
	}
	return err.Error()
}

func PrintStrcut(obj interface{}) {
	str, _ := json.Marshal(obj)
	fmt.Println(string(str))
}
func NowSecond() int32 {
	return int32(time.Now().Unix())
}
func NowMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}
