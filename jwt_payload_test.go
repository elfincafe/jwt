package jwt

import (
	"fmt"
	"testing"
)

func TestNewPayload(t *testing.T) {
	pld := NewPayload()
	typ := fmt.Sprintf("%T", pld)
	if typ != "*jwt.Payload" {
		t.Errorf(`Fail to initialize *Jwt`)
	}

}

func TestPayloadGet(t *testing.T) {
	pld := NewPayload()
	pld.data["key1"] = "value1"
	pld.data["key2"] = 123
	val1 := pld.Get("key1").(string)
	if val1 != "value1" {
		t.Errorf(`key1 is invalid`)
	}
	val2 := pld.Get("key2").(int)
	if val2 != 123 {
		t.Errorf(`key2 is invalid`)
	}
}

func TestPayloadSet(t *testing.T) {
	pld := NewPayload()
	pld.Set("key1", "value1")
	pld.Set("key2", "value2")
	if pld.data["key1"] != "value1" || pld.data["key2"] != "value2" {
		t.Errorf(`key1 and key2 is invalid`)
	}
	pld.Set("key3", 123)
	if pld.data["key3"] != 123 {
		t.Errorf(`key3 is invalid`)
	}
}
