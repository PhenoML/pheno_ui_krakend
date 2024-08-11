package plugin

func (p *Plugin) handleResponse(res ResponseWrapper, config *Config) (interface{}, error) {
  responseTransform := config.ResponseTransform
  if responseTransform != nil {
    context := map[string]interface{}{
      "query":   nil,
      "headers": res.Headers(),
      "body":    res.Data(),
    }

    result, err := responseTransform.Search(context, pExtensions)
    if err != nil {
      return nil, jpathExecutionError("response_transform", err)
    }
    return ModifiedResponseWrapper{
      res,
      result.(Dict),
    }, nil
  }

  return res, nil
}

type ModifiedResponseWrapper struct {
  ResponseWrapper
  data Dict
}

func (r ModifiedResponseWrapper) Data() Dict {
  return r.data
}
