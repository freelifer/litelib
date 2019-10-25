package dao

import (
	. "github.com/freelifer/litelib/public"
)

func IsEmptyByField(value string, field string) error {
	if len(value) == 0 {
		return ErrDataEmpty{Field: field}
	}
	return nil
}
func IsAlreadExistByField(bean interface{}, field string) error {
	exist, err := DB.Engine.Get(bean)
	if exist {
		return ErrAlreadExist{Field: field}
	}
	return err
}

func GetBeanById(id int64, bean interface{}, field string) error {
	exist, err := DB.Engine.Id(id).Get(bean)
	if err != nil {
		return err
	} else if !exist {
		return ErrNotExist{Field: field}
	}
	return nil
}

func GetBeanByOtherField(bean interface{}, field string) error {
	exist, err := DB.Engine.Get(bean)
	if err != nil {
		return err
	} else if !exist {
		return ErrNotExist{Field: field}
	}
	return nil
}
