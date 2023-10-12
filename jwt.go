package jwt

const (
	ALGO_NONE  = "none"
	ALGO_HS256 = "HS256"
	ALGO_RS256 = "RS256"
)

type Jwt struct {
	Header  *Header
	Payload *Payload
}

func New() *Jwt {
	jwt := new(Jwt)
	jwt.Header = NewHeader()
	jwt.Payload = NewPayload()
	return jwt
}

func (jwt *Jwt) Token() string {
	token := ""
	return token
}
