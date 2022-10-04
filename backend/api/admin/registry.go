package admin

import (
	"fmt"
	"reflect"
	"strings"

	"models"
)

var typeRegistry = make(map[string]reflect.Type)

func init() {
	myTypes := []interface{}{
		models.Damage{},
		models.Request{},
	}
	for _, v := range myTypes {
		typeRegistry[strings.ToLower(fmt.Sprintf("%T", v))] = reflect.TypeOf(v)
	}

}

func makeInstance(name string) (interface{}, error) {
	n, ok := typeRegistry[name]
	if ok {
		return reflect.New(n).Interface(), nil
	}

	return nil, fmt.Errorf("unhandled registry %s", name)
}

func getEntityFromReq(name string, isList bool) (val interface{}, err error) {
	if isList {
		return makeInstance("models." + name + "slice")
	}
	return makeInstance("models." + name)
}
