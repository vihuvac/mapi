package helpers

import (
	"reflect"
)

// CopyIdenticalFields will copy identical fields between two interfaces, e.g: fields sent within the payload and a model struct.
func CopyIdenticalFields(a, b interface{}) {
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b).Elem()

	at := av.Type()
	for i := 0; i < at.NumField(); i++ {
		name := at.Field(i).Name

		bf := bv.FieldByName(name)
		if bf.IsValid() {
			bf.Set(av.Field(i))
		}
	}
}
