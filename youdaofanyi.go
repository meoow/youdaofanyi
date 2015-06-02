package youdaofanyi

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Plain = 0
	HTML  = 1
)
const Html_Template = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8"/>
<style>
div.youdao_fanyi {
	background-color: #f9f9f9;
}
p.youdao_translation {
	margin-left: 0.3em;
	margin-right: 0.3em;
}
span.water {
	float:right;
	font-size: 120%;
	font-family: Sans-Serif;
	color:#b9b9b9;
}
</style>
</head>
<body>
<span class="water">有道翻译</span>
<div class="youdao_fanyi">
<h3>{{.word}}</h3>
<hr/>
<p class="youdao_translation">{{.translation}}</p>
</div>
</body>
</html>`

type result struct {
	Type            string `json:"type"`
	ErrorCode       int    `json:"errorCode"`
	ElapsedTime     int    `json:"elapsedTime"`
	TranslateResult [][]struct {
		Src string `json:"src"`
		Tgt string `json:"tgt"`
	} `json:"translateResult"`
}

func Fanyi(str string, style int) (string, error) {

	trans_type := "EN2ZH_CN"
	for _, c := range str {
		if c > '\x7f' {
			trans_type = "ZH_CN2EN"
		}
	}

	uri := "http://fanyi.youdao.com/translate?"
	param := url.Values{"smartresult": {"dict", "rule", "ugc"}}
	query := url.Values{"type": {trans_type},
		"i": {str}, "doctype": {"json"},
		"keyfrom": {"fanyi_web"}, "xmlVersion": {"1.6"},
		"ue": {"UTF-8"}, "typoResult": {"true"},
		"flag": {"false"}}

	resp, err := http.PostForm(uri+param.Encode(), query)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	jsData := new(result)
	json.Unmarshal(body, jsData)

	switch style {
	case Plain:
		return jsData.TranslateResult[0][0].Tgt, nil
	case HTML:
		t, err := template.New("result").Parse(Html_Template)
		if err != nil {
			return "", nil
		}
		buf := new(bytes.Buffer)
		err = t.Execute(buf, map[string]string{"word": jsData.TranslateResult[0][0].Src, "translation": jsData.TranslateResult[0][0].Tgt})
		if err != nil {
			return "", nil
		}
		return buf.String(), nil
	}
	panic("")
}
