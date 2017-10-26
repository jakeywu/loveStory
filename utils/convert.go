package utils

import (
	"github.com/jinzhu/copier"
)

func StructToStruct(toValue interface{}, fromValue interface{}) error {
	return copier.Copy(toValue, fromValue)
}
