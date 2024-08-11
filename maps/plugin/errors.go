package plugin

import (
  "errors"
  "fmt"
)

var (
	unknownTypeErr = func(what string) error {
    return errors.New(fmt.Sprintf("unknown %s type", what))
  }
  jpathPathNotStringErr = func (rawPath interface{}) error {
    return errors.New(fmt.Sprintf("jpath config is not a string: %v\n", rawPath))
  }

  jpathLoadingErr = func (key string, err error) error {
    return errors.New(fmt.Sprintf("error loading %s: %v\n", key, err))
  }

  jpathCompileErr = func (key string, err error) error {
    return errors.New(fmt.Sprintf("error compiling %s: %v\n", key, err))
  }

  jpathExecutionError = func (key string, err error) error {
    return errors.New(fmt.Sprintf("error executing %s: %v\n", key, err))
  }
)
