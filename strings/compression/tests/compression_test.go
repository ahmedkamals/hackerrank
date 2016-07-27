package compression

import (
	"testing"
	"../"
)

func TestCompress(t *testing.T) {

	tests := map[string]string {
		"a" : "a",
		"ab" : "ab",
		"aaaa" : "a4",
		"aabb" : "a2b2",
		"aap" : "a2p",
		"aapp" : "a2pp",
		"abcpqrstuv" : "abcpqrstuv",
		"aabcpqrs" : "a2bcpqrs",
		"aabcpqrstuvaa" : "a2bcpqrstuva2",
		"aabcpqrstuvaab" : "a2bcpqrstuva2b",
	}

	oStrUtil := compression.NewStringUtil()

	for index, expected := range tests {
		got := oStrUtil.Compress(index)

		if(expected != got) {
			t.Error(
				"For", index,
				"expected", expected,
				"got", got,
			)
		}
	}
}
