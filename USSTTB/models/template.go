package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateTB struct {
	*template.Template
}

type HtmlTemplate struct {
	Index        TemplateTB
	Login        TemplateTB
	UserRegister TemplateTB
	ErShou       TemplateTB
	QiuZhu       TemplateTB
}

func (t *TemplateTB) WriteError(w io.Writer, err error) {
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func InitTemplate(templateDir string) HtmlTemplate {
	//返回

	tp := readTemplate(
		[]string{"login", "index", "userRegister"},
		templateDir,
	)
	var htmlTempate HtmlTemplate
	htmlTempate.Index = tp[1]
	htmlTempate.Login = tp[0]
	htmlTempate.UserRegister = tp[2]
	return htmlTempate

}

func readTemplate(templates []string, templateDir string) []TemplateTB {
	var tbs = []TemplateTB{}
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//解析，找：1、拿到当前路径
		//访问博客首页，有多个模板嵌套，因此需要在解析时将所有涉及模板都进行解析
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		t.Funcs(template.FuncMap{"isODD": IsODD})
		t.Funcs(template.FuncMap{"getNextName": GetNextName, "date": Date, "dateDay": Date})
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, post, pagination)
		//fmt.Println(templateDir + viewName)
		if err != nil {
			log.Println("解析模板出错", err)
		}
		var tb TemplateTB
		tb.Template = t
		tbs = append(tbs, tb)
	}

	return tbs
}

func (t *TemplateTB) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
