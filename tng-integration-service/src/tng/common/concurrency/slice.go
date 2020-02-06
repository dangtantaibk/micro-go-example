package concurrency

import (
	"fmt"
	"strings"
	"sync"
)

type ConcurrentSlice struct {
	sync.RWMutex
	items []interface{}
	max   uint32
}

func NewConcurrentSlice(maxItems uint32) *ConcurrentSlice {
	cs := &ConcurrentSlice{
		items: make([]interface{}, 0),
		max:   maxItems,
	}
	return cs
}

func (cs *ConcurrentSlice) Append(item interface{}) {
	cs.Lock()
	defer cs.Unlock()
	if cs.max == 0 || len(cs.items) < int(cs.max) {
		cs.items = append(cs.items, item)
	}
}

func (cs *ConcurrentSlice) Clear() {
	cs.Lock()
	defer cs.Unlock()
	cs.items = cs.items[:0]
}

func (cs *ConcurrentSlice) ToStringAndClear() string {
	cs.Lock()
	defer cs.Unlock()
	if len(cs.items) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, item := range cs.items {
		sb.WriteString(fmt.Sprintf("%v\n", item))
	}
	cs.items = cs.items[:0]
	return sb.String()
}
