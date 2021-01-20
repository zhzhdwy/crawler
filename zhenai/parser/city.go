package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^>]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)">上海离异征婚</a>`)
)

func ParseCity(contents []byte) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		// 这里range会为m创建副本，使用的都是一个内存块。
		// 所以往slice中添加时，都会指向那一块内存。
		// 遍历完后这块内存自然是最后一个数字，所以range中地址类赋值都要考虑到这个问题
		name := string(m[2])
		// result.Items = append(result.Items, "User "+name)
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, url+" "+name)
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
