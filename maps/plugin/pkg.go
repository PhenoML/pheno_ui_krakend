package plugin

import (
  "fmt"
  "github.com/jmespath-community/go-jmespath/pkg/functions"
  "github.com/jmespath-community/go-jmespath/pkg/interpreter"
  "pheno_ui/maps/extensions"
  "reflect"
)

type Plugin string

func MakePlugin(name string) Plugin {
  return Plugin(name)
}

var pExtensions interpreter.Option

func (p *Plugin) Init() {
  var f []functions.FunctionEntry
  f = append(f, functions.GetDefaultFunctions()...)
  f = append(f, extensions.GetExtensions()...)
  caller := interpreter.NewFunctionCaller(f...)
  option := interpreter.WithFunctionCaller(caller)
  pExtensions = option

  fmt.Println(string(*p), "has been initialized")
}

func (p *Plugin) RegisterModifiers(register registerFunc) {
  // ... krakend cannot register a single handler for both request and response
  register(string(*p) + "_request", p.makeHandler, true, false)
  register(string(*p) + "_response", p.makeHandler, false, true)
  fmt.Println(string(*p), "registered")
}

func (p *Plugin) makeHandler(config Dict) modifierHandler {
  pConfig := p.initConfig(config)
  return func(input interface{}) (interface{}, error) {
    req, ok := input.(RequestWrapper)
    if ok {
      return p.handleRequest(req, pConfig)
    }

    res, ok := input.(ResponseWrapper)
    if ok {
      return p.handleResponse(res, pConfig)
    }

    fmt.Println("unknown type", reflect.TypeOf(input))

    return nil, unknownTypeErr("modifier input")
  }
}

func (p *Plugin) initConfig(config Dict) *Config {
  targetRoute, err := compileJpathConfig(config, "target_route")
  if err != nil {
    fmt.Println(err)
  }
  return &Config{
    TargetRoute: targetRoute,
    QueryTransform: unsafeLoadJpathConfig(config, "query_transform"),
    HeadersTransform: unsafeLoadJpathConfig(config, "headers_transform"),
    RequestTransform: unsafeLoadJpathConfig(config, "request_transform"),
    ResponseTransform: unsafeLoadJpathConfig(config, "response_transform"),
  }
}
