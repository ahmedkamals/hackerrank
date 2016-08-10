package strings

import (
	"testing"
	"hackerrank/strings"
	testingUtil "hackerrank/testing"
)

func TestCompress(t *testing.T) {

	testCases := []*testingUtil.TestCase {
		&testingUtil.TestCase{
			Id: "One character.",
			Input: "a",
			Expected: "a",
		},
		&testingUtil.TestCase{
			Id: "Two different characters.",
			Input: "ab",
			Expected: "ab",
		},
		&testingUtil.TestCase{
			Id: "Four similar characters.",
			Input: "aaaa",
			Expected: "a4",
		},
		&testingUtil.TestCase{
			Id: "Repeated two different characters.",
			Input: "aabb",
			Expected: "a2b2",
		},
		&testingUtil.TestCase{
			Id: "Repeated character, and one character to skip at the end.",
			Input: "aap",
			Expected: "a2p",
		},
		&testingUtil.TestCase{
			Id: "Repeated character, and two similar characters to skip at the end.",
			Input: "aapp",
			Expected: "a2pp",
		},
		&testingUtil.TestCase{
			Id: "Combination of characters, and a group of characters to skip at the end.",
			Input: "abcpqrstuv",
			Expected: "abcpqrstuv",
		},
		&testingUtil.TestCase{
			Id: "Combination of characters, and a group of characters to skip at the middle.",
			Input: "aabcpqrs",
			Expected: "a2bcpqrs",
		},
		&testingUtil.TestCase{
			Id: "Combination of characters, and a group of characters to skip at the front, the middle, and the end.",
			Input: "ppbcpqrstuvaaq",
			Expected: "ppbcpqrstuva2q",
		},
		&testingUtil.TestCase{
			Id: "Combination of characters, and a group of characters to skip at the front, the middle, and non-repated character at the end.",
			Input: "aabcpqrstuvaab",
			Expected: "a2bcpqrstuva2b",
		},
	}

	oStrUtil := strings.NewStringUtil()

	for _, testCase := range testCases {
		got := oStrUtil.Compress(testCase.Input.(string))

		if(testCase.Expected != got) {
			t.Error(
				"For", testCase.Id,
				"expected", testCase.Expected,
				"got", got,
			)
		}
	}
}
