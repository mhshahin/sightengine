package utility

import "net/url"

// MapToURLValues ...
func MapToURLValues(params ...map[string]string) url.Values {
	var vals url.Values

	if params == nil {
		return nil
	}

	for _, p := range params {
		for k, v := range p {
			vals.Set(k, v)
		}
	}

	return vals
}
