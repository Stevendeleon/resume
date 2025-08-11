package pkg

import (
	"fmt"
	"html/template"
	"os"
	"strings"
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
		"phoneToText": formatPhone,
		"formatUrl":   formatUrl,
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

func formatPhone(p string) string {
	parts := strings.Split(p, "-")
	if len(parts) != 3 {
		return p
	}

	return fmt.Sprintf("+1(%s) %s-%s", parts[0], parts[1], parts[2])
}

func formatUrl(u string) string {
	u = strings.TrimPrefix(u, "https://")
	u = strings.TrimPrefix(u, "http://")
	u = strings.TrimPrefix(u, "www.")
	return u
}
