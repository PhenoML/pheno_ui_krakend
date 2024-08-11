package plugin

import (
  "fmt"
  "github.com/jmespath-community/go-jmespath"
  "os"
)

func loadJpathConfig(config Dict, key string) (jmespath.JMESPath, error) {
  path, err := getConfigString(config, key)
  if path == "" {
    return nil, err
  }

  fmt.Printf("loading %s: %s\n", key, path)
  buffer, err := os.ReadFile(path)
  if err != nil {
    return nil, jpathLoadingErr(key, err)
  }

  return compileJpathExpression(string(buffer), key)
}

func unsafeLoadJpathConfig(config Dict, key string) jmespath.JMESPath {
  result, err := loadJpathConfig(config, key)
  if err != nil {
    fmt.Println(err)
  }
  return result
}

func getConfigString(config Dict, key string) (string, error) {
  rawValue, ok := config[key]
  if !ok {
    return "", nil // no error, just no config
  }

  value, ok := rawValue.(string)
  if !ok {
    return "", jpathPathNotStringErr(rawValue)
  }

  return value, nil
}

func compileJpathConfig(config Dict, key string) (jmespath.JMESPath, error) {
  value, err := getConfigString(config, key)
  if value == "" {
    return nil, err
  }
  return compileJpathExpression(value, key)
}

func compileJpathExpression(exp string, key string) (jmespath.JMESPath, error) {
  fmt.Printf("compiling %s...\n", key)
  result, err := jmespath.Compile(exp)
  if err != nil {
    return nil, jpathCompileErr(key, err)
  }
  return result, nil
}

func toObject[T any](src map[string]T) Dict {
  result := make(Dict)
  for k, v := range src {
    result[k] = v
  }
  return result
}

func toValue(src Dict) map[string][]string {
  result := make(map[string][]string)
  for k, v := range src {
    result[k] = toStringArray(v)
  }
  return result
}

func toStringArray(src interface{}) []string {
  strArr, ok := src.([]string)
  if ok {
    return strArr
  }

  interfaceArr, ok := src.([]interface{})
  if !ok {
    return []string{fmt.Sprint(src)}
  }

  var result []string
  for _, v := range interfaceArr {
    result = append(result, fmt.Sprint(v))
  }
  return result
}
