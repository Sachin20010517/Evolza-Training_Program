package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type MonthlyReport struct {
	Month     string
	Year      int
	Employees []string
	Products  []string
}

// Using the init function to make sure the template is only parsed once in the program
func main() {
	// Execute person into the template and print to Stdout
	report := MonthlyReport{"June", 2024, []string{"Oluwatomisin", "Oluwaseun", "Oluwasegun"}, []string{"Books", "Pens", "Pencils"}}
	filepath := "sample.template"
	outLocation := "generated.txt"

	tFunc := template.FuncMap{
		"add": add,
	}

	err := generateFromTemplate(report, filepath, tFunc, outLocation, filepath)
	if err != nil {
		log.Println(err)
	}
}

// generateFromTemplate parses the template and executes it with the provided data
func generateFromTemplate(data interface{}, fileName string, function template.FuncMap, outFileLocation string, tmplLocation ...string) error {
	tmpl, err := template.New(fileName).Funcs(function).ParseFiles(tmplLocation...)
	if err != nil {
		return err
	}

	buf := new(strings.Builder)
	err = tmpl.Execute(buf, data)
	if err != nil {
		return err
	}

	_, err = os.Create(outFileLocation)
	if err != nil {
		return err
	}
	buff := []byte(buf.String())
	err = os.WriteFile(outFileLocation, buff, 0644)
	if err != nil {
		return err
	}
	return nil
}

func add(x, y int) int {
	return x + y
}
