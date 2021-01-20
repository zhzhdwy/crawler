package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strings"
)

var (
	prupleRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08>([^>]*)</div>`)
	pinkRe   = regexp.MustCompile(`<div class="m-btn pink" data-v-ff544c08>([^>]*)</div>`)
	genderRe = regexp.MustCompile(`"genderString":"([^"]+)"`)
)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	//fmt.Printf("%s", contents)
	profile := model.Profile{}
	profile.Name = name

	matchPruple := prupleRe.FindAllSubmatch(contents, -1)
	matchPink := pinkRe.FindAllSubmatch(contents, -1)
	matchGender := genderRe.FindAllSubmatch(contents, -1)

	var match [][][]byte

	match = append(matchPink, matchPruple...)
	match = append(match, matchGender...)

	if len(matchPruple) >= 2 {
		profile.Marriage = string(matchPruple[0][1])
		profile.Occpation = string(matchPruple[len(matchPruple)-2][1])
		profile.Education = string(matchPruple[len(matchPruple)-1][1])
	}

	for _, m := range match {
		op := string(m[1])
		if ok := strings.Contains(op, "籍贯:"); ok {
			profile.Hokou = string(op)
		}
		if ok := strings.Contains(op, "车"); ok {
			profile.Car = string(op)
		}
		if ok := strings.Contains(op, "岁"); ok {
			profile.Age = string(op)
		}
		if ok := strings.Contains(op, "月收入:"); ok {
			profile.Income = string(op)
		}
		if ok := strings.Contains(op, "cm"); ok {
			profile.Height = string(op)
		}
		if ok := strings.Contains(op, "kg"); ok {
			profile.Weight = string(op)
		}

		if ok := strings.Contains(op, "座"); ok {
			profile.Xinzuo = string(op)
		}

		if ok := strings.ContainsAny(op, "家|房"); ok {
			profile.House = string(op)
		}

		if ok := strings.ContainsAny(op, "男|女"); ok {
			profile.Gender = string(op)
		}

	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

var (
	ageRe      = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	heightRe   = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([\d]+)CM</span></td>`)
	weightRe   = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
	marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">离异</div>`)
	//genderRe    = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	incomeRe    = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]*)</td>`)
	educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occpationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	hokouRe     = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	xinzuoRe    = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	houseRe     = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe       = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
)

// 页面被修改了
func oldFind(profile model.Profile, name string, contents []byte) {
	profile.Name = name
	profile.Gender = extactString(contents, genderRe)
	profile.Income = extactString(contents, incomeRe)
	profile.Education = extactString(contents, educationRe)
	profile.Occpation = extactString(contents, occpationRe)
	profile.Hokou = extactString(contents, hokouRe)
	profile.Xinzuo = extactString(contents, xinzuoRe)
	profile.House = extactString(contents, houseRe)
	profile.Car = extactString(contents, carRe)
	profile.Marriage = extactString(contents, marriageRe)
	//height, err := strconv.Atoi(extactString(contents, heightRe))
	//if err == nil {
	//	profile.Height = height
	//}
	//weight, err := strconv.Atoi(extactString(contents, weightRe))
	//if err == nil {
	//	profile.Weight = weight
	//}
	//age, err := strconv.Atoi(extactString(contents, ageRe))
	//if err == nil {
	//	profile.Age = age
	//}
}

func extactString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
