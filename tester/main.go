package main

import (
  "encoding/json"
  "fmt"
  "github.com/jmespath-community/go-jmespath"
  "github.com/jmespath-community/go-jmespath/pkg/functions"
  "github.com/jmespath-community/go-jmespath/pkg/interpreter"
  "pheno_ui/maps/extensions"
)

func main() {
  fmt.Println("tester loaded!!!")
  data := map[string]interface{}{
    "query":   "hello world!",
    "nested":  map[string]interface{}{"key": "value", "array": []interface{}{1, 2, 3}, "nested": map[string]interface{}{"key": "value"}},
  }
  query := `{
    core: 'control string',
    isoDateNow: isoDateNow(),
    encodeJson: encodeJson(nested),
    decodeJson: decodeJson('{"key":"value","array":[1,2,3],"nested":{"key":"value"}}'),
    randomInt: ` + "randomInt(`1`, `10`)," + `
    byKeyNoDefault: ` + "byKey('yolo', `{ \"yolo\": 666, \"loyo\": 999 }`)," + `
    byKeyNoDefaultErr: ` + "byKey('dario', `{ \"yolo\": 666, \"loyo\": 999 }`)," + `
    byKeyDefault: ` + "byKey('loyo', `{ \"yolo\": 666, \"loyo\": 999 }`, `888`)," + `
    byKeyDefaultErr: ` + "byKey('dario', `{ \"yolo\": 666, \"loyo\": 999 }`, `777`)," + `
    split: split('hello,world', ','),
    envDefault: env('ENV_VAR_TEST', 'not found'),
    envDefaultErr: env('ENV_VAR_TEST_ERR', 'not found'),
    envNoDefault: env('ENV_VAR_TEST'),
    envNoDefaultErr: env('ENV_VAR_TEST_ERR'),
    concat: concat('SELECT * FROM ', 'hello_world')
  }`

  compiled, err := jmespath.Compile(query)
  if err != nil {
    fmt.Println("error compiling query:", err)
    return
  }

  var f []functions.FunctionEntry
  f = append(f, functions.GetDefaultFunctions()...)
  f = append(f, extensions.GetExtensions()...)
  caller := interpreter.NewFunctionCaller(f...)
  opt := interpreter.WithFunctionCaller(caller)

  result, err := compiled.Search(data, opt)

  if err != nil {
    fmt.Println("error executing response_transform:", err)
  } else {
    pretty, _ := json.MarshalIndent(result, "", "  ")
    fmt.Println("transform result:", string(pretty))
  }
}
