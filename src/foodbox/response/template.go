package response

import (
	"html/template"
	"fmt"
	"net/http"
)

type FoodboxTemplate struct {
	template *template.Template
}

type empty struct {}

func RenderHtml(htmlFile string, writer http.ResponseWriter) {
	t := newTemplate(htmlFile)
	t.execute(writer, empty{})
}

func RenderTemplate(templateFile string, writer http.ResponseWriter, data interface{}) {
	t := newTemplate(templateFile)
	t.execute(writer, data)
}

func newTemplate(templateFile string) *FoodboxTemplate {
	templatePath := fmt.Sprintf("html/%s", templateFile)
	t := template.Must(template.New(templateFile).Delims("{%", "%}").ParseFiles(templatePath))
	return &FoodboxTemplate{t}
}

func (t *FoodboxTemplate) execute(writer http.ResponseWriter, data interface{}) {
	err := t.template.Execute(writer, data)
	if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
    }
}
