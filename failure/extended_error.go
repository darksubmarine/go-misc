package failure

import (
	"errors"
	"fmt"
)

type ExtendedError struct {
	err  error
	meta map[string]interface{}
}

func NewExtendedError(err error) *ExtendedError {
	return &ExtendedError{err: err, meta: map[string]interface{}{}}
}

func (e *ExtendedError) Error() string {
	return e.String()
}

func (e *ExtendedError) String() string {
	return fmt.Sprintf("%s meta=%v", e.err.Error(), e.meta)
}

func (e *ExtendedError) IsErr(b *ExtendedError) bool {
	if b == nil {
		return false
	}
	return errors.Is(e.err, b.err)
}

func (e *ExtendedError) Is(b error) bool {
	return errors.Is(e.err, b)
}

//-- Metadata

func (e *ExtendedError) with(key string, val interface{}) *ExtendedError {
	e.meta[key] = val
	return e
}

func (e *ExtendedError) WithString(key string, val string) *ExtendedError {
	return e.with(key, val)
}

func (e *ExtendedError) WithInt(key string, val int) *ExtendedError {
	return e.with(key, val)
}

func (e *ExtendedError) WithFloat(key string, val float64) *ExtendedError {
	return e.with(key, val)
}

func (e *ExtendedError) WithBool(key string, val bool) *ExtendedError {
	return e.with(key, val)
}

func (e *ExtendedError) HasMetaKey(key string) bool {
	if _, ok := e.meta[key]; ok {
		return true
	}
	return false
}

func (e *ExtendedError) metaVal(key string) interface{} {
	if v, ok := e.meta[key]; ok {
		return v
	}
	return nil
}

func (e *ExtendedError) MetaRawVal(key string) interface{} {
	return e.metaVal(key)
}

func (e *ExtendedError) MetaString(key string) *string {
	if v := e.metaVal(key); v != nil {
		if val, ok := v.(string); ok {
			return &val
		}
	}
	return nil
}

func (e *ExtendedError) MetaInt(key string) *int {
	if v := e.metaVal(key); v != nil {
		if val, ok := v.(int); ok {
			return &val
		}
	}
	return nil
}

func (e *ExtendedError) MetaFloat(key string) *float64 {
	if v := e.metaVal(key); v != nil {
		if val, ok := v.(float64); ok {
			return &val
		}
	}
	return nil
}

func (e *ExtendedError) MetaBool(key string) *bool {
	if v := e.metaVal(key); v != nil {
		if val, ok := v.(bool); ok {
			return &val
		}
	}
	return nil
}

func (e *ExtendedError) Meta() map[string]interface{} {
	ret := make(map[string]interface{}, len(e.meta))
	for k, v := range e.meta {
		ret[k] = v
	}

	return ret
}
