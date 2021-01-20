package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "抱墨堂")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element, but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:      "抱墨堂",
		Gender:    "女士",
		Age:       "38岁",
		Height:    "167cm",
		Weight:    "55kg",
		Income:    "月收入:3千以下",
		Marriage:  "离异",
		Education: "中专",
		Occpation: "计算机/互联网",
		Hokou:     "籍贯:河南开封",
		Xinzuo:    "天秤座(09.23-10.22)",
		House:     "和家人同住",
		Car:       "未买车",
	}

	if profile != expected {
		t.Errorf("\nexpected %v\nbut was  %v.", expected, profile)
	}

}
