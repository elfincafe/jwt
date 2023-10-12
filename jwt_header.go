package jwt

type Header struct {
	Typ  string `json:"typ"`
	Alg  string `json:"alg"`
	data map[string]interface{}
}

func NewHeader() *Header {
	hdr := new(Header)
	hdr.Typ = "JWT"
	hdr.Alg = ALGO_NONE
	hdr.data = map[string]interface{}{}
	return hdr
}

func (hdr *Header) Get(key string) interface{} {
	if _, ok := hdr.data[key]; ok {
		return hdr.data[key]
	} else {
		return ""
	}
}

func (hdr *Header) Set(key string, value interface{}) {
	hdr.data[key] = value
}
