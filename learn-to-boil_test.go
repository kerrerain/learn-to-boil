package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestExtractLines(t *testing.T) {
	testCases := []struct {
		FilePath       string
		ExpectedOutput []string
	}{
		{
			"case1.txt",
			[]string{
				"1  pâte feuilletée",
				"1/2 chèvre en buche",
				"3 cuillères à soupe de pesto",
				"4  oeufs",
				"20 cl de crème fraîche",
				"75 g de  gruyère  rapé",
			},
		},
		{
			"case2.txt",
			[]string{
				"600 g d' épinards  hachés (peuvent être  surgelés )",
				"3  oeufs",
				"1 petite faisselle de  fromage  blanc",
				"1 boursin (ou autre fromage  ail  et fines herbes)",
				"4 belles tranches de  saumon fumé",
			},
		},
	}

	for _, testCase := range testCases {
		dat, err := ioutil.ReadFile("./test_files/" + testCase.FilePath)
		check(err)

		lines := ExtractLines(string(dat))

		if !reflect.DeepEqual(testCase.ExpectedOutput, lines) {
			t.Errorf("Expected %v but was %v", testCase.ExpectedOutput, lines)
		}
	}
}
