package plugin

import (
  "github.com/jmespath-community/go-jmespath"
  "io"
  "net/url"
)

type registerFunc = func(string, modifierFactory, bool, bool)

type modifierFactory = func(Dict) modifierHandler

type modifierHandler = func(interface{}) (interface{}, error)

type Dict = map[string]interface{}

type Config struct {
  TargetRoute jmespath.JMESPath
  QueryTransform jmespath.JMESPath
  HeadersTransform jmespath.JMESPath
  RequestTransform jmespath.JMESPath
  ResponseTransform jmespath.JMESPath
}

type RequestWrapper interface {
  Params() map[string]string
  Headers() map[string][]string
  Body() io.ReadCloser
  Method() string
  URL() *url.URL
  Query() url.Values
  Path() string
}

type ResponseWrapper interface {
  Data() Dict
  Io() io.Reader
  IsComplete() bool
  StatusCode() int
  Headers() map[string][]string
}
