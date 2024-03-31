package common

type Results struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}

func StatusOk() *Results {
	results := new(Results)
	results.Code = "200"
	results.Data = make(map[string]any)
	return results
}

func StatusErr() *Results {
	results := new(Results)
	results.Code = "400"
	results.Data = make(map[string]any)
	return results
}

func SetCode(code string) *Results {
	results := new(Results)
	results.Code = code
	results.Data = make(map[string]any)
	return results
}

func (r *Results) AddData(key string, value any) *Results {
	r.Data[key] = value
	return r
}
