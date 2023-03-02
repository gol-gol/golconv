package golconv

import (
	"testing"
)

func TestStringToInt(t *testing.T) {
	if result := StringToInt("0", 1); result != 0 {
		t.Fatal("StringToInt failed for 0")
	}
	if result := StringToInt("100", 1); result != 100 {
		t.Fatal("StringToInt failed for 100")
	}
	if result := StringToInt("-1", 1); result != -1 {
		t.Fatal("StringToInt failed for -1")
	}
	if result := StringToInt("-", 1); result != 1 {
		t.Fatal("StringToInt failed for -")
	}
}

func TestStringToUint64(t *testing.T) {
	if result := StringToUint64("0", 1); result != 0 {
		t.Fatal("StringToUint64 failed for 0")
	}
	if result := StringToUint64("100", 1); result != 100 {
		t.Fatal("StringToUint64 failed for 100")
	}
	if result := StringToUint64("-1", 1); result < 0 {
		t.Fatal("StringToUint64 failed for -1")
	}
	if result := StringToUint64("-", 1); result != 1 {
		t.Fatal("StringToUint64 failed for -")
	}
}

func TestStringToBool(t *testing.T) {
	if result := StringToBool("false", true); result != false {
		t.Fatal("StringToBool failed for false")
	}
	if result := StringToBool("true", false); result != true {
		t.Fatal("StringToBool failed for true")
	}
	if result := StringToBool("0", true); result != true {
		t.Fatal("StringToBool failed for 0")
	}
	if result := StringToBool("", true); result != true {
		t.Fatal("StringToBool failed for <empty>")
	}
}
