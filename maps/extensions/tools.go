package extensions

import (
  "reflect"
  "runtime"
  "strings"
)

func getFunctionName(f interface{}) string {
  name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
  components := strings.Split(name, ".")
  return components[len(components)-1]
}

func getFunctionArguments(f interface{}) []reflect.Type {
  var result []reflect.Type
  for index := 0; index < reflect.TypeOf(f).NumIn(); index++ {
    result = append(result, reflect.TypeOf(f).In(index))
  }
  return result
}

func isVariadic(f interface{}) bool {
  return reflect.TypeOf(f).IsVariadic()
}
