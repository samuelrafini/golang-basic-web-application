package views

import (
	"fmt"
	"html/template"
	"path/filepath"
)

var (
	LayoutDir string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout string
}

func NewView(layout string, files ...string) * View {
	files = append(files, layoutFiles()...)
	fmt.Println(files)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout: layout,
	}
}

//return slice of strings for all the layouts
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}