package debugger

import (
	"fmt"
	"sort"
	"strings"
)

type any interface{}

type Debugger map[string]any

func NewDebugger() *Debugger {
	return &Debugger{}
}

func (d *Debugger) Set(key string, value any) {
	(*d)[key] = value
}

func (d *Debugger) Get(key string) any {
	return (*d)[key]
}

func (d *Debugger) Delete(key string) {
	delete(*d, key)
}

func (d *Debugger) Clear() {
	*d = Debugger{}
}

func (d *Debugger) Inc(key string) error {
	v := (*d)[key]
	switch v.(type) {
	case int:
		(*d)[key] = v.(int) + 1
	case int8:
		(*d)[key] = v.(int8) + 1
	case int16:
		(*d)[key] = v.(int16) + 1
	case int32:
		(*d)[key] = v.(int32) + 1
	case int64:
		(*d)[key] = v.(int64) + 1
	case uint:
		(*d)[key] = v.(uint) + 1
	case uint8:
		(*d)[key] = v.(uint8) + 1
	case uint16:
		(*d)[key] = v.(uint16) + 1
	case uint32:
		(*d)[key] = v.(uint32) + 1
	case uint64:
		(*d)[key] = v.(uint64) + 1
	default:
		return fmt.Errorf("Debugger.Inc: %s is not a number", key)
	}
	return nil
}

func (d *Debugger) ToString() string {
	s := make([]string, len(*d))
	i := 0
	for k, v := range *d {
		s[i] = fmt.Sprintf("%s: %v", k, v)
		i++
	}
	sort.Strings(s)
	return strings.Join(s, "\n")
}
