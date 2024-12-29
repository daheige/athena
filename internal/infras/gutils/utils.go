// Package gutils 用于基础函数的封装
package gutils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Uuid 生成 version4 的uuid
// 返回格式:eba1e8cd0460491049c644bdf3cf024d
func Uuid() string {
	u, err := uuid.NewRandom()
	if err != nil {
		return strings.Replace(RndUuid(), "-", "", -1)
	}

	return strings.Replace(u.String(), "-", "", -1)
}

// RndUuid 基于时间ns和随机数实现唯一的uuid
// 在单台机器上是不会出现重复的uuid
// 如果要在分布式的架构上生成不重复的uuid
// 只需要在rndStr的前面加一些自定义的字符串
// 返回格式:eba1e8cd-0460-4910-49c6-44bdf3cf024d
func RndUuid() string {
	ns := time.Now().UnixNano()
	rndStr := strings.Join([]string{
		strconv.FormatInt(ns, 10), strconv.FormatInt(RandInt64(1000, 9999), 10),
	}, "")
	s := Md5(rndStr)
	return fmt.Sprintf("%s-%s-%s-%s-%s", s[:8], s[8:12], s[12:16], s[16:20], s[20:])
}

// RandInt64 生成m-n之间的 int64 随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}

	return rand.Int63n(max-min) + min
}

// Md5 md5 string
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
