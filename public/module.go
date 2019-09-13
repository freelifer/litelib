package public

import (
	"fmt"
)

type BaseErr struct {
	Module string
}

type ErrNameAlreadExist struct {
	Module string
	Name   string
}

func IsErrNameAlreadExist(err error) bool {
	_, ok := err.(ErrNameAlreadExist)
	return ok
}

func (err ErrNameAlreadExist) Error() string {
	return fmt.Sprintf("[%s] [name:%s] already exists", err.Module, err.Name)
}

type ErrUUIDNotExist struct {
	Module string
	Uuid   string
}

func IsErrUUIDNotExist(err error) bool {
	_, ok := err.(ErrUUIDNotExist)
	return ok
}

func (err ErrUUIDNotExist) Error() string {
	return fmt.Sprintf("[%s] [uuid:%s] not exists", err.Module, err.Uuid)
}

type ErrCodeAlreadExist struct {
	Module string
	Code   string
}

func IsErrCodeAlreadExist(err error) bool {
	_, ok := err.(ErrCodeAlreadExist)
	return ok
}

func (err ErrCodeAlreadExist) Error() string {
	return fmt.Sprintf("[%s] [code:%s] already exists", err.Module, err.Code)
}

// Error 名称重复
type ErrNameRepeat struct {
	BaseErr
	Name string
}

func IsErrNameRepeat(err error) bool {
	_, ok := err.(ErrNameRepeat)
	return ok
}

func (err ErrNameRepeat) Error() string {
	return fmt.Sprintf("[%s] [name:%s] repeat", err.Module, err.Name)
}

// Error 名称不存在
type ErrNameNotExist struct {
	Module string
	Name   string
}

func IsErrNameNotExist(err error) bool {
	_, ok := err.(ErrNameNotExist)
	return ok
}

func (err ErrNameNotExist) Error() string {
	return fmt.Sprintf("[%s] [name:%s] not found", err.Module, err.Name)
}

type ErrIdNotExist struct {
	Module string
	Id     int64
}

func IsErrIdNotExist(err error) bool {
	_, ok := err.(ErrIdNotExist)
	return ok
}

func (err ErrIdNotExist) Error() string {
	return fmt.Sprintf("[%s] [id:%d]  not found", err.Module, err.Id)
}
