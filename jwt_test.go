package jwt

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	jwt := New()
	typ := fmt.Sprintf("%T", jwt)
	if typ != "*jwt.Jwt" {
		t.Errorf(`Fail to initialize *Jwt`)
	}
}

func TestToken(t *testing.T) {

}
