package public

import (
	"fmt"
)

// 数据不存在-------------------------------------------------------------------
///----------------------------------------------------------------------------
///----------------------------------------------------------------------------
type ErrNotExist struct {
	Field string
}

func IsErrNotExist(err error) bool {
	_, ok := err.(ErrNotExist)
	return ok
}

func (err ErrNotExist) Error() string {
	return fmt.Sprintf("%s not exists", err.Field)
}

// 数据为空---------------------------------------------------------------------
///----------------------------------------------------------------------------
///----------------------------------------------------------------------------
type ErrDataEmpty struct {
	Field string
}

func IsErrDataEmpty(err error) bool {
	_, ok := err.(ErrNotExist)
	return ok
}

func (err ErrDataEmpty) Error() string {
	return fmt.Sprintf("%s is empty", err.Field)
}

// 数据已经存在------------------------------------------------------------------
///----------------------------------------------------------------------------
///----------------------------------------------------------------------------
type ErrAlreadExist struct {
	Field string
}

func IsErrAlreadExist(err error) bool {
	_, ok := err.(ErrAlreadExist)
	return ok
}

func (err ErrAlreadExist) Error() string {
	return fmt.Sprintf("%s already exists", err.Field)
}
