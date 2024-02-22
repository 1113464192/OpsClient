package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func IsDir(path string) bool {
	dirStat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return dirStat.IsDir()
}

// 是否包含
func IsSliceContain(slice interface{}, value interface{}) bool {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return false
	}
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(item, value) {
			return true
		}
	}

	return false
}

// 环境变量提取整数
func GetEnvInt(key string, fallback int) int {
	ret := fallback
	value, exists := os.LookupEnv(key)
	if !exists {
		return ret
	}
	if t, err := strconv.Atoi(value); err != nil { //nolint:gosec
		return ret
	} else {
		ret = t
	}
	return ret
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// string转换uint
func StringToUint(idStr *string) (id uint, err error) {
	var oldId uint64
	oldId, err = strconv.ParseUint(*idStr, 10, 0)
	if err != nil {
		return 0, errors.New("uint类型转换失败")
	}
	id = uint(oldId)
	return id, err
}

func IntSliceToStringSlice(intSlice []int) []string {
	stringSlice := make([]string, len(intSlice))
	for i, v := range intSlice {
		stringSlice[i] = strconv.Itoa(v)
	}
	return stringSlice
}

func Float64SliceToStringSlice(floatSlice []float64) []string {
	stringSlice := make([]string, len(floatSlice))
	for i, v := range floatSlice {
		stringSlice[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	return stringSlice
}

// 有最大对应取最大，否则只取[0]
func SplitStringMap(originalMap map[string][]string) []map[string]string {
	maxLength := 0
	for _, values := range originalMap {
		if len(values) > maxLength {
			maxLength = len(values)
		}
	}

	// 创建一个切片用于存储拆分后的map
	splitMaps := make([]map[string]string, maxLength)

	// 遍历原始map
	for key, values := range originalMap {
		for i := 0; i < maxLength; i++ {
			// 如果值的长度大于i，则将值拆分到对应的map中；否则将空字符串放入map中
			if maxLength == len(values) {
				if splitMaps[i] == nil {
					splitMaps[i] = make(map[string]string)
				}
				splitMaps[i][key] = values[i]
			} else {
				if splitMaps[i] == nil {
					splitMaps[i] = make(map[string]string)
				}
				splitMaps[i][key] = values[0]
			}
		}
	}

	return splitMaps
}

func ConvertToJson(params []string) (res string, err error) {
	var extraByte []byte
	var extra = make(map[int]string)
	if len(params) > 0 {
		for i, v := range params {
			extra[i] = v
		}
	}
	extraByte, err = json.Marshal(extra)
	if err != nil {
		return "", err
	}
	return string(extraByte), err
}

// 传x=y切片
func ConvertToJsonPair(params []string) (res string, err error) {
	data := make(map[string][]string)
	for _, param := range params {
		pair := strings.SplitN(param, "=", -1)
		if len(pair) != 2 {
			return "", fmt.Errorf("invalid key-value pair: %s", param)
		}
		key := pair[0]
		value := pair[1]
		data[key] = append(data[key], value)
	}
	var jsonData []byte
	jsonData, err = json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("%s: %v", "转换json报错", err)
	}
	return string(jsonData), err
}

func DeleteUintSlice(s []uint, elem uint) []uint {
	j := 0
	for _, v := range s {
		if v != elem {
			s[j] = v
			j++
		}
	}
	// 如果直接使用 return s，那么返回的切片 s 将包含原始切片中的所有元素，包括指定元素和非指定元素。这是因为切片是引用类型，返回的切片与传入的切片共享相同的底层数组。在这种情况下，虽然在循环中将指定元素跳过并将非指定元素移动到切片的前面，但没有对底层数组进行修改。因此，返回的切片 s 仍然包含了原始切片中的所有元素。
	return s[:j]
}

func DeleteAnySlice(s any, elem any) (any, error) {
	sliceValue := reflect.ValueOf(s)
	if sliceValue.Kind() != reflect.Slice {
		return s, errors.New("传入的首位参数, 类型不是slice")
	}
	j := 0
	for i := 0; i < sliceValue.Len(); i++ {
		v := sliceValue.Index(i).Interface()
		if v != elem {
			sliceValue.Index(j).Set(sliceValue.Index(i))
			j++
		}
	}
	return sliceValue.Slice(0, j).Interface(), nil
}

// 包含字符串切片
func StringSliceContain(substrings []string, target string) bool {
	for _, str := range substrings {
		if strings.Contains(str, target) {
			return true
		}
	}
	// 都不包含返回false
	return false
}

// 获取本机IP
func GetLocalIp(LocalIpCmd string, LocalIpApi string) (ip string, err error) {
	cmd := exec.Command("bash", "-c", LocalIpCmd)
	var output []byte
	output, err = cmd.CombinedOutput()
	// 判断文件是否为空
	if output == nil || err != nil || cmd.ProcessState.ExitCode() != 0 {
		// 文件为空则get获取外网IP
		var response *http.Response
		response, err = http.Get(LocalIpApi)
		if err != nil {
			fmt.Println("查询错误:", err)
			return
		}
		defer response.Body.Close()
		output, err = io.ReadAll(response.Body)
		ip = string(output)
	} else {
		// 文件不为空则直接读取IP
		ip = string(output)
	}
	//	判断字符串是否为正常IP
	if net.ParseIP(ip) == nil {
		return "", fmt.Errorf("%s IP地址不合法", ip)
	}
	return ip, err
}
