package avl

import (
	"fmt"
)

type DuplicateError struct {
	Key     interface{}
	Message string
}

func NewDuplicateError(k interface{}) *DuplicateError {
	return &DuplicateError{
		Key:     k,
		Message: "DUPLICATE ERROR: ",
	}
}

func (e *DuplicateError) Error() string {
	switch e.Key.(type) {
	case string:
		val := e.Key.(string)
		return fmt.Sprintf(e.Message + "Key = " + val + " already exists in the tree. " +
			"Please use Update() if you wish to update the value of the node.")
	case int8:
		val := int64(e.Key.(int8))
		return fmt.Sprintf(e.Message+"Key = %d"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case int16:
		val := int64(e.Key.(int16))
		return fmt.Sprintf(e.Message+"Key = %d"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case int32:
		val := int64(e.Key.(int32))
		return fmt.Sprintf(e.Message+"Key = %c (rune representation, if key is a rune) // "+
			"%d (int32 representation)"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val, val)
	case int64:
		val := int64(e.Key.(int64))
		return fmt.Sprintf(e.Message+"Key = %d"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case uint8:
		val := uint64(e.Key.(uint8))
		return fmt.Sprintf(e.Message+"Key = %d"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case uint16:
		val := uint64(e.Key.(uint16))
		return fmt.Sprintf(e.Message+"Key = %d"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case uint32:
		val := uint64(e.Key.(uint32))
		return fmt.Sprintf(e.Message+"Key = %d"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case uint64:
		val := e.Key.(uint64)
		return fmt.Sprintf(e.Message+"Key = %d"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case float32:
		val := float64(e.Key.(float32))
		return fmt.Sprintf(e.Message+"Key = %d"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case float64:
		val := e.Key.(float64)
		return fmt.Sprintf(e.Message+"Key = %g"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	case bool:
		val := e.Key.(bool)
		return fmt.Sprintf(e.Message+"Key = %t"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", val)
	default:
		return fmt.Sprintf(e.Message+"Key = %+v"+" already exists in the tree. "+
			"Please use Update() if you wish to update the value of the node.", e.Key)
	}
}

type NilNodeError struct {
	Key     interface{}
	Message string
}

func NewNilNodeError(k interface{}) *NilNodeError {
	return &NilNodeError{
		Key:     k,
		Message: "NIL NODE ERROR: ",
	}
}

func (e *NilNodeError) Error() string {
	switch e.Key.(type) {
	case string:
		val := e.Key.(string)
		return fmt.Sprintf(e.Message + "Key = " + val + " does not exist in the tree.")
	case int8:
		val := int64(e.Key.(int8))
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case int16:
		val := int64(e.Key.(int16))
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case int32:
		val := int64(e.Key.(int32))
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case int64:
		val := int64(e.Key.(int64))
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case uint8:
		val := uint64(e.Key.(uint8))
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case uint16:
		val := uint64(e.Key.(uint16))
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case uint32:
		val := uint64(e.Key.(uint32))
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case uint64:
		val := e.Key.(uint64)
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case float32:
		val := float64(e.Key.(float32))
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case float64:
		val := e.Key.(float64)
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	case bool:
		val := e.Key.(bool)
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", val)
	default:
		return fmt.Sprintf(e.Message+"Key = %+v"+" does not exist in the tree.", e.Key)
	}
}
