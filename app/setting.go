package app

import (
	"html/template"
)

var Setting = func(getKey string, values ...string) interface{} {
	var ret = make(map[string]interface{})
	ret["web_title"] = "Aixgl.艾辛阁"
	ret["web_host"] = "go.aixgl.com"
	ret["asset_path"] = "/tmp/"
	intro := "先专精一个领域，而后涉猎其它 静思之地 <br/> first need to Specialize in a field  Then you can get more . The place of meditation"
	ret["intro"] = (template.HTML(intro))

	if getKey == "all" || getKey == "" {
		return ret
	}
	if len(values) == 0 && getKey != "" {
		return ret[getKey]
	}
	return ret
}
