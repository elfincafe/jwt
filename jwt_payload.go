package jwt

type Payload struct {
	Iss  string   `json:"iss,omitempty"`
	Sub  string   `json:"sub,omitempty"`
	Aud  []string `json:"aud,omitempty"`
	Exp  int      `json:"exp,omitempty"`
	Nbf  int      `json:"nbf,omitempty"`
	Iat  int      `json:"iat,omitempty"`
	Jti  string   `json:"jti,omitempty"`
	Typ  string   `json:"typ,omitempty"`
	data map[string]interface{}
}

func NewPayload() *Payload {
	pld := new(Payload)
	pld.Iss = ""
	pld.Sub = ""
	pld.Aud = []string{}
	pld.Exp = 0
	pld.Nbf = 0
	pld.Iat = 0
	pld.Jti = ""
	pld.data = map[string]interface{}{}
	return pld
}

func (pld *Payload) Get(key string) interface{} {
	if _, ok := pld.data[key]; ok {
		return pld.data[key]
	} else {
		return ""
	}
}

func (jwt *Payload) Set(key string, value interface{}) {
	jwt.data[key] = value
}
