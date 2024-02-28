package libs

import "encoding/json"

func JsonStringify[T any](o T, def string) string {
	b, err := json.Marshal(o)

	if err != nil {
		return def
	}

	return string(b)
}

func JsonDecode[T any](j string, def T) T {
	var res T

	err := json.Unmarshal([]byte(j), &res)
	if err != nil {
		return def
	}

	return res
}
