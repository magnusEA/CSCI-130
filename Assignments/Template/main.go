package main

import (
	"log"
	"os"
	"text/template"
)
type Meal struct{
	Breakfast string
	Lunch string
	Dinner string
}

func main() {

	food := Meal{
		Breakfast: "Eggs",
		Lunch: "Meat",
		Dinner: "Something Else",
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, food)
	if err != nil {
		log.Fatalln(err)
	}
}
