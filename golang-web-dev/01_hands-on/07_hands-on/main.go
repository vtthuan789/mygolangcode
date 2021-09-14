package main

import (
	"log"
	"os"
	"text/template"
)

type MenuItem struct {
	Name        string
	Price       float32
	Description string
}

type Menu struct {
	Meal  string
	Items []MenuItem
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := []struct {
		Name  string
		Menus []Menu
	}{
		{
			Name: "California I",
			Menus: []Menu{
				{
					Meal: "Breakfast",
					Items: []MenuItem{
						{
							Name:        "MAC & CHEESE BALLS",
							Price:       7.00,
							Description: "Spicy Tomato Sauce, Basil Purée, Pecorino Cheese",
						},
						{
							Name:        "VEGAN CHOP CHOP",
							Price:       8.00,
							Description: "Vegetables, Quinoa, Puffed Grains, Mint, Parsley, Spicy Tahini Dressing, Beet Hummus, Pea Shoots",
						},
						{
							Name:        "MANGO & BRIE QUESADILLA",
							Price:       10.00,
							Description: "Pickled Jalapeño, Scallions, Aged Cheddar, Cumin Cilantro Crema, Tempura Bits, Fresh Lime",
						},
					},
				},
				{
					Meal: "Lunch",
					Items: []MenuItem{
						{
							Name:        "CHICKEN CAESAR LETTUCE WRAPS",
							Price:       10.00,
							Description: "Harissa Smoked Chicken, Iceberg, Lettuce, Grilled Corn Tomato Salsa, Jalapeño Caesar Aioli, Grana Padano, Chicken Crackling",
						},
						{
							Name:        "BIBIMBAP",
							Price:       12.00,
							Description: "Kalbi Beef, Coconut Rice, Kimchi Slaw, Sous Vide Egg, Gochujang Sauce, Thai,Basil, Smoked Peanuts, Crispy Shallots",
						},
						{
							Name:        "BANQUET BURGER",
							Price:       12.00,
							Description: "Prime Beef Patty, Glazed Pork Belly, American Cheese, Chili Relish, Mustard Aioli, Arugula, Sesame Brioche, Frizzled Onions",
						},
					},
				},
				{
					Meal: "Dinner",
					Items: []MenuItem{
						{
							Name:        "REUBEN SANDWICH",
							Price:       12.00,
							Description: "Smoked Meat, Swiss Cheese, Spicy Russian Dressing, Pickled Slaw, Caraway Brioche, Dill Pickles",
						},
						{
							Name:        "FISH & CHIPS",
							Price:       12.00,
							Description: "Beer Battered Icelandic Cod, Herb Potato Fries, Smoked Tartar Sauce, Pickled Slaw, Charred Lemon",
						},
						{
							Name:        "FD POUTINE",
							Price:       12.00,
							Description: "Shoestring Fries dusted in Rosemary Salt, Bone Marrow Gravy, Cheese Curds, Shredded Mozzarella, Smoked Tomato Ketchup",
						},
					},
				},
			},
		},
		{
			Name: "California II",
			Menus: []Menu{
				{
					Meal: "Breakfast",
					Items: []MenuItem{
						{
							Name:        "MAC & CHEESE BALLS",
							Price:       7.00,
							Description: "Spicy Tomato Sauce, Basil Purée, Pecorino Cheese",
						},
						{
							Name:        "VEGAN CHOP CHOP",
							Price:       8.00,
							Description: "Vegetables, Quinoa, Puffed Grains, Mint, Parsley, Spicy Tahini Dressing, Beet Hummus, Pea Shoots",
						},
						{
							Name:        "MANGO & BRIE QUESADILLA",
							Price:       10.00,
							Description: "Pickled Jalapeño, Scallions, Aged Cheddar, Cumin Cilantro Crema, Tempura Bits, Fresh Lime",
						},
					},
				},
				{
					Meal: "Lunch",
					Items: []MenuItem{
						{
							Name:        "CHICKEN CAESAR LETTUCE WRAPS",
							Price:       10.00,
							Description: "Harissa Smoked Chicken, Iceberg, Lettuce, Grilled Corn Tomato Salsa, Jalapeño Caesar Aioli, Grana Padano, Chicken Crackling",
						},
						{
							Name:        "BIBIMBAP",
							Price:       12.00,
							Description: "Kalbi Beef, Coconut Rice, Kimchi Slaw, Sous Vide Egg, Gochujang Sauce, Thai,Basil, Smoked Peanuts, Crispy Shallots",
						},
						{
							Name:        "BANQUET BURGER",
							Price:       12.00,
							Description: "Prime Beef Patty, Glazed Pork Belly, American Cheese, Chili Relish, Mustard Aioli, Arugula, Sesame Brioche, Frizzled Onions",
						},
					},
				},
				{
					Meal: "Dinner",
					Items: []MenuItem{
						{
							Name:        "REUBEN SANDWICH",
							Price:       12.00,
							Description: "Smoked Meat, Swiss Cheese, Spicy Russian Dressing, Pickled Slaw, Caraway Brioche, Dill Pickles",
						},
						{
							Name:        "FISH & CHIPS",
							Price:       12.00,
							Description: "Beer Battered Icelandic Cod, Herb Potato Fries, Smoked Tartar Sauce, Pickled Slaw, Charred Lemon",
						},
						{
							Name:        "FD POUTINE",
							Price:       12.00,
							Description: "Shoestring Fries dusted in Rosemary Salt, Bone Marrow Gravy, Cheese Curds, Shredded Mozzarella, Smoked Tomato Ketchup",
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
