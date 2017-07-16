package weblru

type Request struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (r Request) Valid() bool {
	if r.Key == "" || r.Value == "" {
		return false
	}
	return true
}
