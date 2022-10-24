package common

import (
	"fmt"
	"reflect"
)

// func main() {

// 	// source := testClass{
// 	// 	Id:   1,
// 	// 	Name: "mohammad1",
// 	// }

// 	source := struct {
// 		Children []testClass
// 	}{}

// 	source.Children = []testClass{
// 		{Id: 1, Name: "m1"},
// 		{Id: 2, Name: "m2"},
// 	}

// 	des := Map[struct{ Children []testClass2 }](source)
// 	fmt.Println(des)

// }

func Map[T any](source any) T {

	var dest *T = new(T)

	var sourceVal = reflect.ValueOf(source)
	var destVal = reflect.ValueOf(dest).Elem()
	mapValues(sourceVal, destVal)

	return *dest

}

func mapValues(sourceVal, destVal reflect.Value) {
	sourceType := sourceVal.Type()
	if sourceType.Kind() == reflect.Struct {
		if sourceVal.Type().Kind() == reflect.Ptr {
			if sourceVal.IsNil() {
				// If source is nil, it maps to an empty struct
				sourceVal = reflect.New(sourceVal.Type().Elem())
			}
			sourceVal = sourceVal.Elem()
		}
		for i := 0; i < sourceType.NumField(); i++ {
			mapField(sourceVal, destVal, i)
		}
	} else if sourceType.Kind() == reflect.Slice {
		mapSlice(sourceVal, destVal)
	} else if sourceType == sourceVal.Type() {
		destVal.Set(sourceVal)
	} else if sourceType.Kind() == reflect.Ptr {
		if valueIsNil(sourceVal) {
			return
		}
		val := reflect.New(sourceType.Elem())
		mapValues(sourceVal, val.Elem())
		destVal.Set(val)
	} else {
		panic("Currently not supported")
	}

}

func mapSlice(sourceVal, destVal reflect.Value) {
	destType := destVal.Type()
	length := sourceVal.Len()
	target := reflect.MakeSlice(destType, length, length)
	for j := 0; j < length; j++ {
		val := reflect.New(destType.Elem()).Elem()
		mapValues(sourceVal.Index(j), val)
		target.Index(j).Set(val)
	}

	if length == 0 {
		verifyArrayTypesAreCompatible(sourceVal, destVal)
	}
	destVal.Set(target)
}

func verifyArrayTypesAreCompatible(sourceVal, destVal reflect.Value) {
	dummyDest := reflect.New(reflect.PtrTo(destVal.Type()))
	dummySource := reflect.MakeSlice(sourceVal.Type(), 1, 1)
	mapValues(dummySource, dummyDest.Elem())
}

func mapField(source, destVal reflect.Value, i int) {
	destType := destVal.Type()
	fieldName := destType.Field(i).Name

	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Sprintf("Error mapping field: %s. DestType: %v. SourceType: %v. Error: %v", fieldName, destType, source.Type(), r))
		}
	}()

	destField := destVal.Field(i)
	if destType.Field(i).Anonymous {
		mapValues(source, destField)
	} else {
		if valueIsContainedInNilEmbeddedType(source, fieldName) {
			return
		}
		sourceField := source.FieldByName(fieldName)
		if (sourceField == reflect.Value{}) {

			if destField.Kind() == reflect.Struct {
				mapValues(source, destField)
				return
			} else {
				for i := 0; i < source.NumField(); i++ {
					if source.Field(i).Kind() != reflect.Struct {
						continue
					}
					if sourceField = source.Field(i).FieldByName(fieldName); (sourceField != reflect.Value{}) {
						break
					}
				}
			}
		}
		mapValues(sourceField, destField)
	}
}

func valueIsContainedInNilEmbeddedType(source reflect.Value, fieldName string) bool {
	structField, _ := source.Type().FieldByName(fieldName)
	ix := structField.Index
	if len(structField.Index) > 1 {
		parentField := source.FieldByIndex(ix[:len(ix)-1])
		if valueIsNil(parentField) {
			return true
		}
	}
	return false
}

func valueIsNil(value reflect.Value) bool {
	return value.Type().Kind() == reflect.Ptr && value.IsNil()
}
