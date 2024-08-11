package main

import (
  maps_plugin "pheno_ui/maps/plugin"
)

func main() {}

func init() {
  ModifierRegisterer.Init()
}

// ModifierRegisterer is the symbol the plugin loader will be looking for. It must
// implement the plugin.Registerer interface
// https://github.com/luraproject/lura/blob/master/proxy/plugin/modifier.go#L71
var ModifierRegisterer  = maps_plugin.MakePlugin("pheno_ui_maps")

//
//type registerer string
//
//// RegisterModifiers is the function the plugin loader will call to register the
//// modifier(s) contained in the plugin using the function passed as argument.
//// f will register the factoryFunc under the name and mark it as a request
//// and/or response modifier.
//func (r registerer) RegisterModifiers(f func(
//	name string,
//	factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
//	appliesToRequest bool,
//	appliesToResponse bool,
//)) {
//	f(string(r)+"-request", r.requestDump, true, false)
//	f(string(r)+"-response", r.responseDump, false, true)
//	fmt.Println(string(r), "registered nope!!!")
//}
//
//// RequestWrapper is an interface for passing proxy request between the krakend pipe
//// and the loaded plugins
//type RequestWrapper interface {
//	Params() map[string]string
//	Headers() map[string][]string
//	Body() io.ReadCloser
//	Method() string
//	URL() *url.URL
//	Query() url.Values
//	Path() string
//}
//
//// ResponseWrapper is an interface for passing proxy response between the krakend pipe
//// and the loaded plugins
//type ResponseWrapper interface {
//	Data() map[string]interface{}
//	Io() io.Reader
//	IsComplete() bool
//	StatusCode() int
//	Headers() map[string][]string
//}
//
//type ModifiedResponseWrapper struct {
//  original ResponseWrapper
//  data map[string]interface{}
//}
//
//func (r ModifiedResponseWrapper) Data() map[string]interface{} {
//  return r.data
//}
//
//func (r ModifiedResponseWrapper) Io() io.Reader {
//  return r.original.Io()
//}
//
//func (r ModifiedResponseWrapper) IsComplete() bool {
//  return r.original.IsComplete()
//}
//
//func (r ModifiedResponseWrapper) StatusCode() int {
//  return r.original.StatusCode()
//}
//
//func (r ModifiedResponseWrapper) Headers() map[string][]string {
//  return r.original.Headers()
//}
//
//var unknownTypeErr = errors.New("unknown request type")
//
//func (r registerer) requestDump(
//	cfg map[string]interface{},
//) func(interface{}) (interface{}, error) {
//	// check the cfg. If the modifier requires some configuration,
//	// it should be under the name of the plugin.
//	// ex: if this modifier required some A and B config params
//	/*
//	   "extra_config":{
//	       "plugin/req-resp-modifier":{
//	           "name":["krakend-debugger"],
//	           "krakend-debugger":{
//	               "A":"foo",
//	               "B":42
//	           }
//	       }
//	   }
//	*/
//
//	// return the modifier
//	fmt.Println("request dumper injected!!!")
//	return func(input interface{}) (interface{}, error) {
//		req, ok := input.(RequestWrapper)
//		if !ok {
//			return nil, unknownTypeErr
//		}
//
//    fmt.Println("Request Dumper... it works yo!!!")
//		fmt.Println("params:", req.Params())
//		fmt.Println("headers:", req.Headers())
//		fmt.Println("method:", req.Method())
//		fmt.Println("url:", req.URL())
//		fmt.Println("query:", req.Query())
//		fmt.Println("path:", req.Path())
//
//		return input, nil
//	}
//}
//
//func (r registerer) responseDump(
//	cfg map[string]interface{},
//) func(interface{}) (interface{}, error) {
//	// check the cfg. If the modifier requires some configuration,
//	// it should be under the name of the plugin.
//	// ex: if this modifier required some A and B config params
//	/*
//	   "extra_config":{
//	       "plugin/req-resp-modifier":{
//	           "name":["krakend-debugger"],
//	           "krakend-debugger":{
//	               "A":"foo",
//	               "B":42
//	           }
//	       }
//	   }
//	*/
//
//	// return the modifier
//	fmt.Println("response dumper injected!!!")
//  var f []functions.FunctionEntry
//  f = append(f, functions.GetDefaultFunctions()...)
//  f = append(f, extensions.GetExtensions()...)
//  caller := interpreter.NewFunctionCaller(f...)
//  opt := interpreter.WithFunctionCaller(caller)
//
// var responseTransform jmespath.JMESPath = nil
// responseTransformPath, ok := cfg["response_transform"]
// if ok {
//   fmt.Println("loading response_transform:", responseTransformPath)
//   responseTransformBuffer, err := os.ReadFile(responseTransformPath.(string))
//   if err != nil {
//     fmt.Println("error loading response_transform:", err)
//   } else {
//     responseTransformString := string(responseTransformBuffer)
//     fmt.Println("response_transform loaded:", responseTransformString)
//     responseTransform, err = jmespath.Compile(responseTransformString)
//     if err != nil {
//       fmt.Println("error compiling response_transform:", err)
//     }
//   }
// }
//
//	return func(input interface{}) (interface{}, error) {
//		resp, ok := input.(ResponseWrapper)
//		if !ok {
//			return nil, unknownTypeErr
//		}
//
//   if responseTransform != nil {
//     context := map[string]interface{}{
//       "query":   nil,
//       "headers": resp.Headers(),
//       "body":    resp.Data(),
//     }
//
//     result, err := responseTransform.Search(context, opt)
//     if err != nil {
//       fmt.Println("error executing response_transform:", err)
//     } else {
//       fmt.Println("response_transform result 0:", result)
//       return ModifiedResponseWrapper{
//         original: resp,
//         data:     result.(map[string]interface{}),
//       }, nil
//     }
//   }
//
//		return input, nil
//	}
//}



