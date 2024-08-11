package extensions

import (
  "encoding/json"
  "fmt"
  "github.com/jmespath-community/go-jmespath/pkg/functions"
  "math/rand"
  "os"
  "strings"
  "time"
)

func GetExtensions() []functions.FunctionEntry {
  return []functions.FunctionEntry{
    MakeFunctionEntry(isoDateNow),
    MakeFunctionEntry(encodeJson),
    MakeFunctionEntry(decodeJson),
    MakeFunctionEntry(randomInt),
    MakeFunctionEntry(byKey),
    MakeFunctionEntry(split),
    MakeFunctionEntry(env),
    MakeFunctionEntry(concat),
  }
}

func isoDateNow() string {
  return time.Now().UTC().Format(time.RFC3339)
}

func encodeJson(obj map[string]interface{}) (string, error) {
  marshaled, err := json.Marshal(obj)
  if err != nil {
    return "", err
  }
  return string(marshaled), nil
}

func decodeJson(str string) (map[string]interface{}, error) {
  var obj map[string]interface{}
  err := json.Unmarshal([]byte(str), &obj)
  if err != nil {
    return nil, err
  }
  return obj, nil
}

func randomInt(min float64, max float64) int {
  iMin := int(min)
  iMax := int(max)
  return iMin + rand.Intn(iMax-iMin)
}

func byKey(key string, obj map[string]interface{}, def ...interface{}) interface{} {
  val, ok := obj[key]
  if !ok && len(def) > 0 {
    return def[0]
  }
  return val
}

func split(obj interface{}, sep string) []string {
  str, ok := obj.(string)
  if !ok {
    return nil
  }
  return strings.Split(str, sep)
}

func env(key string, def ...string) string {
  val, ok := os.LookupEnv(key)
  if !ok && len(def) > 0 {
    return def[0]
  }
  return val
}

func concat(entries ...interface{}) string {
  var b strings.Builder
  b.Grow(len(entries))
  for _, entry := range entries {
    b.WriteString(fmt.Sprint(entry))
  }
  return b.String()
}
