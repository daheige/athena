package prioritymq

import (
	"fmt"
)

// Level 队列优先级
type Level int64

const (
	// Normal 普通优先级
	Normal Level = iota

	// Low 低优先级
	Low

	// Medium 中优先级
	Medium

	// High 高优先级
	High
)

var suffixLevel = map[Level]string{
	High:   "high",
	Medium: "medium",
	Low:    "low",
	Normal: "normal",
}

// NewTopicName 创建优先级topic名字
// 格式: xxx-high,xxx-medium,xxx-low,xxx-normal
func NewTopicName(name string, l Level) string {
	suffix, exist := suffixLevel[l]
	if !exist {
		l = Normal
	}

	topic := fmt.Sprintf("%s-%s", name, suffix)
	return topic
}
