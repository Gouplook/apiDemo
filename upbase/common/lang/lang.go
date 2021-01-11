/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  lang
 * @Version: 1.0.0
 * @Date: 2021/1/10 21:21
 */
package lang

import (
	"apiDemo/upgin"
	"apiDemo/upgin/i18n"
	"apiDemo/upgin/utils"
	"log"
	"path/filepath"
)


//定义变量
var(
	loadLangs map[string]*Load
	defaultLang string
)

type Load struct {
	i18n *i18n.Locale  // 描述本地化信息
}

//获取lang
func GetLang(str string, langs ...string) string{
	for _,lang := range langs {
		if msg := getLang(str, lang);msg != "" {
			return msg
		}
	}
	if msg := getLang(str, defaultLang);msg != "" {
		return msg
	}
	return ""
}
//返回值
func getLang(str string, lang string) string{
	if loadLang,ok := loadLangs[lang]; ok == true {
		return loadLang.i18n.Tr(str)
	}
	return ""
}
func init(){
	//加载语言包
	langs := upgin.AppConfig.Strings("lang")
	lang_path := upgin.AppConfig.String("lang.path")
	if lang_path == "" {
		lang_path = "lang"
	}
	loadLangs = make(map[string]*Load)
	defaultLang = upgin.AppConfig.String("lang.default")
	basePath := ""
	if utils.FileExists(filepath.Join(upgin.WorkPath,lang_path)) {
		basePath = upgin.WorkPath
	} else {
		basePath = upgin.AppPath
	}
	for _, lang := range langs {
		langfile := filepath.Join(basePath,lang_path,upgin.AppConfig.String("lang."+lang))
		err := i18n.SetMessage(lang,langfile)
		if err != nil {
			log.Fatalf("load % Error:%v", langfile, err)
		}
		loadLangs[lang] = &Load{
			i18n: &i18n.Locale{Lang: lang},
		}
	}
}