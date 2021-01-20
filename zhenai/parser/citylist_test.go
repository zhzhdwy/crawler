package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	resutl := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if len(resutl.Requests) != resultSize {
		t.Errorf("resutl should have %d requsets; but had %d", resultSize, len(resutl.Requests))

	}

	for i, url := range expectedUrls {
		if resutl.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, resutl.Requests[i].Url)
		}
	}
	if len(resutl.Items) != resultSize {
		t.Errorf("resutl should have %d requsets; but had %d", resultSize, len(resutl.Items))

	}

	for i, city := range expectedCities {
		if resutl.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s", i, city, resutl.Items[i].(string))
		}
	}

}
