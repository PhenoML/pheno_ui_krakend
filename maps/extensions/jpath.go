package extensions

import (
  "github.com/jmespath-community/go-jmespath/pkg/functions"
  "reflect"
)

// Future Dario: this should be implemented as code generation in the future,
// reflection is expensive and could cause slowdowns in heavy usage scenarios

func MakeFunctionEntry(f interface{}) functions.FunctionEntry {
  name := getFunctionName(f)
  args := getFunctionArguments(f)
  variadic := isVariadic(f)
  jpArgs := makeJpArgs(args, variadic)

  fValue := reflect.ValueOf(f)
  fType := fValue.Type()
  fArgCount := fType.NumIn()
  handler := func (arguments []interface{}) (interface{}, error) {
    var args []reflect.Value
    for i, arg := range arguments {
      if arg == nil {
        expectedType := fType.In(min(i, fArgCount - 1))
        if i >= fArgCount && variadic {
          args = append(args, reflect.Zero(expectedType.Elem()))
        } else {
          args = append(args, reflect.Zero(expectedType))
        }
      } else {
        args = append(args, reflect.ValueOf(arg))
      }
    }
    result := fValue.Call(args)
    var err error = nil
    if len(result) > 1 && !result[1].IsNil() {
      err = result[1].Interface().(error)
    }
    return result[0].Interface(), err
  }

  return functions.FunctionEntry{
    Name:      name,
    Handler:   handler,
    Arguments: jpArgs,
  }
}
func argsToJpType(args []reflect.Type) []functions.JpType {
  var result []functions.JpType
  for _, arg := range args {
    result = append(result, argToJpType(arg))
  }
  return result
}

func makeJpArgs(args []reflect.Type, variadic bool) []functions.ArgSpec {
  jpTypes := argsToJpType(args)
  if variadic {
    jpTypes = jpTypes[:len(jpTypes)-1]
  }

  var result []functions.ArgSpec
  for _, jpType := range jpTypes {
    result = append(result, functions.ArgSpec{
      Types: []functions.JpType{jpType},
    })
  }

  if variadic {
    result = append(result, functions.ArgSpec{
      Types: []functions.JpType{argToJpType(args[len(args)-1].Elem())},
      Variadic: true,
      Optional: true,
    })
  }

  return result
}

var expRefType = reflect.TypeOf((*functions.ExpRef)(nil)).Elem()

func argToJpType(arg reflect.Type) functions.JpType {
  if arg.AssignableTo(expRefType) {
    return functions.JpExpref
  }
  switch arg.Kind() {
  case reflect.String:
    return functions.JpString

  case reflect.Int:
    fallthrough
  case reflect.Float64:
    return functions.JpNumber

  case reflect.Slice:
    fallthrough
  case reflect.Array:
    return listArgToJpType(arg)

  case reflect.Map:
    return functions.JpObject

  default:
    return functions.JpAny
  }
}

func listArgToJpType(arg reflect.Type) functions.JpType {
  switch arg.Elem().Kind() {
  case reflect.String:
    return functions.JpArrayString

  case reflect.Int:
    fallthrough
  case reflect.Float64:
    return functions.JpArrayNumber

  case reflect.Array:
    return functions.JpArrayArray

  default:
    return functions.JpArray
  }
}


