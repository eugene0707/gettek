package models

import (
	"reflect"
	"github.com/eugene0707/gettek/common"
)

type BaseModel struct {

}

func All(CurrentModel interface{}) (v interface{}, err error){
	modelType := reflect.TypeOf(CurrentModel)
	slice := reflect.MakeSlice(reflect.SliceOf(modelType), 0, 0)

	x := reflect.New(slice.Type())
	x.Elem().Set(slice)
	err = common.CurrentDB().Find(x.Interface()).Error

	return x.Interface(), err
}

func Find(CurrentModel interface{}, id int) (v interface{}, err error){
	modelType := reflect.ValueOf(CurrentModel).Elem()

	x := reflect.New(modelType.Type())
	x.Elem().Set(modelType)

	err = common.CurrentDB().First(x.Interface(), id).Error

	return x.Interface(), err
}

func Destroy(CurrentModel interface{}, id int) (v interface{}, err error){
	modelType := reflect.ValueOf(CurrentModel).Elem()

	x := reflect.New(modelType.Type())
	x.Elem().Set(modelType)

	if err := common.CurrentDB().First(x.Interface(), id).Error; err != nil {
		return nil, err
	}

	err = common.CurrentDB().Where("id = ?", id).Delete(x.Interface()).Error

	return x.Interface(), err
}

func Create(CurrentModel interface{}, val interface{}) (v interface{}, err error){
	modelType := reflect.ValueOf(CurrentModel).Elem()

	x := reflect.New(modelType.Type())
	x.Elem().Set(modelType)

	err = common.CurrentDB().Create(x.Interface()).Error

	return x.Interface(), err
}


