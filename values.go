package xhttp

type Values map[string]string

//get value
//if key exists
func (v Values) Get(key string) string {
	if v == nil {
		return ""
	}
	return v[key]
}

//set key value
func (v Values) Set(key, value string) {
	v[key] = value
}

//
func (v Values) Has(key string) bool {
	return v[key] != ""
}
