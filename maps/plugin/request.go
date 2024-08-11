package plugin

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io"
  "net/url"
)

func (p *Plugin) handleRequest(req RequestWrapper, config *Config) (interface{}, error) {
  targetRoute := config.TargetRoute
  queryTransform := config.QueryTransform
  headersTransform := config.HeadersTransform
  requestTransform := config.RequestTransform

  if targetRoute != nil || queryTransform != nil || headersTransform != nil || requestTransform != nil {
    rawBody, err := io.ReadAll(req.Body())
    var body map[string]interface{} = nil
    if err != nil {
      fmt.Println("error reading body:", err)
    } else {
      err = json.Unmarshal(rawBody, &body)
      if err != nil {
        fmt.Println("error unmarshalling body:", err)
      }
    }

    context := map[string]interface{}{
      "query":   toObject(req.Query()),
      "headers": toObject(req.Headers()),
      "body":    body,
    }


    // Future Dario: Modifying the request path in the plugin seems to not be working
    path := ""
    if targetRoute != nil {
      route, err := targetRoute.Search(context, pExtensions)
      if err != nil {
        return nil, jpathExecutionError("target_route", err)
      }
      routeURL, err := url.Parse(route.(string))
      if err != nil {
        return nil, fmt.Errorf("error parsing target_route: %w", err)
      }
      path = routeURL.Path
    }

    var query url.Values = nil
    if queryTransform != nil {
      result, err := queryTransform.Search(context, pExtensions)
      if err != nil {
        return nil, jpathExecutionError("query_transform", err)
      }
      query = toValue(result.(map[string]interface{}))
    }

    var headers map[string][]string = nil
    if headersTransform != nil {
      result, err := headersTransform.Search(context, pExtensions)
      if err != nil {
        return nil, jpathExecutionError("headers_transform", err)
      }
      headers = toValue(result.(map[string]interface{}))
    }

    var requestBody io.ReadCloser = nil
    if requestTransform != nil {
      result, err := requestTransform.Search(context, pExtensions)
      if err != nil {
        return nil, jpathExecutionError("request_transform", err)
      }
      newBody, err := json.Marshal(result)
      if err != nil {
        return nil, fmt.Errorf("error marshalling request body: %w", err)
      }
      requestBody = io.NopCloser(bytes.NewReader(newBody))
    }

    return ModifiedRequestWrapper{
      req,
      path,
      query,
      headers,
      requestBody,
    }, nil
  }

  return req, nil
}

type ModifiedRequestWrapper struct {
  RequestWrapper
  path string
  query url.Values
  headers map[string][]string
  body io.ReadCloser
}

// future Dario: Overriding this method seems to no have effect in the final request

func (r ModifiedRequestWrapper) Path() string {
  if r.path != "" {
    return r.path
  }
  return r.RequestWrapper.Path()
}

func (r ModifiedRequestWrapper) Query() url.Values {
  if r.query != nil {
    return r.query
  }
  return r.RequestWrapper.Query()
}

func (r ModifiedRequestWrapper) Headers() map[string][]string {
  if r.headers != nil {
    return r.headers
  }
  return r.RequestWrapper.Headers()
}

func (r ModifiedRequestWrapper) Body() io.ReadCloser {
  if r.body != nil {
    return r.body
  }
  return r.RequestWrapper.Body()
}
