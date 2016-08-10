package strings

import (
	"testing"
	"hackerrank/strings"
	testingUtil "hackerrank/testing"
)

func TestGetCost(t *testing.T) {

	testCases := []*testingUtil.TestCase {
		&testingUtil.TestCase{
			Id: "Group of different characters.",
			Input: "abcd",
			Expected: 4,
		},
		&testingUtil.TestCase{
			Id: "Group of different characters, repeated in order.",
			Input: "abab",
			Expected: 2,
		},
		&testingUtil.TestCase{
			Id: "Group of two reated different characters, with non-repeated character in the middle.",
			Input: "abcab",
			Expected: 3,
		},
		&testingUtil.TestCase{
			Id: "Group of three repated different characters, with non-repeated character in the middle.",
			Input: "abdcabd",
			Expected: 4,
		},
	}

	oStrUtil := strings.NewStringUtil()

	for _, testCase := range testCases {
		got := oStrUtil.GetCost(testCase.Input.(string))

		if(testCase.Expected != got) {
			t.Error(
				"For", testCase.Id,
				"expected", testCase.Expected,
				"got", got,
			)
		}
	}
}

