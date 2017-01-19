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
			"case1.html",
			[]string{
				"450 g de beurre",
				"3 citrons (ou 3 oranges)",
				"90 g de beurre doux",
				"3 oeufs",
				"100 g de sucre",
				"zestes de citron (ou d' orange)", // FIXME @magleff Try to remove the whitespace
			},
		},
		{
			"case2.html",
			[]string{
				"1,2 kg de gras double",
				"3 carottes",
				"2 beaux oignons",
				"100 g beurre persillé",
				"1 petit bocal de sauce tomate cuisinée",
				"3/4 d'une bouteille de vin blanc sec",
				"thym, laurier",
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
