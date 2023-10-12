package jwt

import "testing"

func TestNewHeader(t *testing.T) {

}

func TestHeaderGet(t *testing.T) {
	hdr := NewHeader()
	hdr.data["key1"] = "value1"
	hdr.data["key2"] = 123
	val1 := hdr.Get("key1").(string)
	if val1 != "value1" {
		t.Errorf(`key1 is invalid`)
	}
	val2 := hdr.Get("key2").(int)
	if val2 != 123 {
		t.Errorf(`key2 is invalid`)
	}
}

func TestHeaderSet(t *testing.T) {
	hdr := NewHeader()
	hdr.Set("key1", "value1")
	hdr.Set("key2", "value2")
	if hdr.data["key1"] != "value1" || hdr.data["key2"] != "value2" {
		t.Errorf(`key1 and key2 is invalid`)
	}
	hdr.Set("key3", 123)
	if hdr.data["key3"] != 123 {
		t.Errorf(`key3 is invalid`)
	}
}
