package pkg

import (
	"html/template"
	"os"
)

type MetaData struct {
	Title       string
	Author      string
	Description string
}

type WebData struct {
	Resume   Resume
	MetaData MetaData
}

func GenerateHTML(resume *Resume) error {
	funcMap := template.FuncMap{
		"lenMinusOne": lenMinusOne,
	}

	tmplHTML := template.New("").Funcs(funcMap)
	tmplHTML, err := tmplHTML.ParseGlob("templates/*.gohtml")
	if err != nil {
		return err
	}

	htmlFile, err := os.Create("output.html")
	if err != nil {
		return err
	}
	defer htmlFile.Close()

	data := WebData{
		Resume: *resume,
		MetaData: MetaData{
			Title:       resume.PageTitle,
			Author:      resume.Name,
			Description: "My resume, a brief summary highlighting my skills and capabilities",
		},
	}

	err = tmplHTML.ExecuteTemplate(htmlFile, "_.gohtml", data)
	if err != nil {
		return err
	}

	return nil
}

func lenMinusOne(s []string) int {
	return len(s) - 1
}
