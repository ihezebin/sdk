package stringer

import "testing"

func TestStringer(t *testing.T) {
	t.Log(IsEmpty(""))
	t.Log(CharAt("你好，world", 1))
	t.Log(Equals("korbin", "korbin"))
	t.Log(StartWithPrefix("hello, world", "hello"))
	t.Log(EndWithSuffix("hello, world", "world"))
	t.Log(SubstringWithBegin("hello, world", 7))
	t.Log(SubstringWithRange("hello, world", 7, len("hello, world")))
	t.Log(Contact("hello", ", world"))
}
