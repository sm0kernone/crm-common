package history

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
)

type (
	mFields struct {
		Before interface{}
		After  interface{}
	}
	M map[string]mFields
)

func GetChangedAttributes(old, new interface{}) M {
	var nForm = reflect.ValueOf(new).Elem().Interface()
	var oForm = reflect.ValueOf(old).Elem().Interface()
	return calcChangedAttrs(oForm, nForm, []string{"created_at", "updated_at", "deleted_at"})
}

func calcChangedAttrs(old, new interface{}, ignoreFields []string) M {
	var memory = M{}
	var memoryMutex = sync.RWMutex{}

	oForm := reflect.Indirect(reflect.ValueOf(&old)).Elem()
	nForm := reflect.Indirect(reflect.ValueOf(&new)).Elem()

	typeOfT := oForm.Type()

	var wg sync.WaitGroup

	for i := 0; i < nForm.NumField(); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					log.Println("RECOVERING FROM FATAL ERROR IN COMPARING CHANGED ATTRIBUTES, ", r)
					return
				}
			}()

			oField := oForm.Field(i)
			nField := nForm.Field(i)
			fieldName := typeOfT.Field(i).Name
			if !InSlice(ignoreFields, fieldName) &&
				oField.Kind() != reflect.Struct &&
				oField.Kind() != reflect.Interface &&
				oField.Kind() != reflect.Func {
				var changed bool

				if oField.Kind() == reflect.Slice {
					changed = sliceDiffers(reflect.ValueOf(oField.Interface()), reflect.ValueOf(nField.Interface()))
				} else if oField.Kind() == reflect.Map {
					changed = mapsDiffers(oField.Interface(), nField.Interface())
				} else {
					changed = nField.Interface() != oField.Interface()
				}
				// if field value is a pointer - compare values by the addresses instead of comparing addresses
				if oField.Kind() == reflect.Ptr {

					var oFieldNil, nFieldNil bool
					nFieldNil = nField.IsNil() || nField.Elem().IsZero()
					oFieldNil = oField.IsNil() || oField.Elem().IsZero()
					// if two elements are nil - they are not changed
					if oFieldNil && nFieldNil {
						changed = false
						return
					}

					// if firstElement is nil and second is not
					if oFieldNil && !nFieldNil {
						memoryMutex.Lock()
						memory[fieldName] = mFields{
							Before: nil,
							After:  nField.Interface(),
						}
						memoryMutex.Unlock()

						changed = false
						return
					}

					// if secondElement is nil and first is not
					if nFieldNil && !oFieldNil {
						memoryMutex.Lock()
						memory[fieldName] = mFields{
							Before: oField.Interface(),
							After:  nil,
						}
						memoryMutex.Unlock()

						changed = false
						return
					}

					if nField.Elem().Interface() == oField.Elem().Interface() {
						changed = false
					}
				}

				if changed {
					memoryMutex.Lock()
					memory[fieldName] = mFields{
						Before: oField.Interface(),
						After:  nField.Interface(),
					}
					memoryMutex.Unlock()
				}
			}

		}(i)
	}

	wg.Wait()

	return memory
}

func sliceDiffers(slice1, slice2 reflect.Value) bool {
	for j := 0; j < slice1.Len(); j++ {
		found := false
		for y := 0; y < slice2.Len(); y++ {
			if slice1.Index(j).Interface() == slice2.Index(y).Interface() {
				found = true
				break
			}
		}
		if !found {
			return true
		}
	}

	return false
}

func mapsDiffers(map1, map2 interface{}) bool {
	switch m1 := map1.(type) {
	case map[string]interface{}:
		return !reflect.DeepEqual(m1, map2.(map[string]interface{}))
	}

	xType := fmt.Sprintf("%T", map1)
	fmt.Println(xType) // "[]int"

	return true
}

func InSlice(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("SliceExists() given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func GetFieldName(model interface{}, fieldName string) string {
	return GetValueByTagFromModel(model, fieldName, "sql")
}

func GetValueByTag(f reflect.StructField, tag string) string {
	return f.Tag.Get(tag)
}

func GetValueByTagFromModel(model interface{}, fieldName, tag string) string {
	if field, exists := reflect.ValueOf(model).Elem().Type().FieldByName(fieldName); exists {
		var fieldName = GetValueByTag(field, tag)
		return strings.Split(fieldName, ",")[0]
	}
	return ""
}
