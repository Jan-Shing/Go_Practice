package main

import(
	"testing"
	"regexp"
)

type Testcase struct{
	input string
	want string
}

func normalize(phone string) string{
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

func TestNormalize(t *testing.T){
	testCase := []Testcase{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"(123)456-7892", "1234567892"},
	}

	for _, tc := range testCase{
		t.Run(tc.input, func(t *testing.T){
			actual := normalize(tc.input)
			if actual != tc.want{
				t.Errorf("got %s; want %s", actual, tc.want)
			}
		})
	}
}