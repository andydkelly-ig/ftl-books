package main

import "fmt"

//Book represents info about a book
type Book struct {
	Title string
	Author string
	Copies int
	Series bool
	SeriesId int
	PriceCents int
	DiscountPercent int
}

func main () {
	var myBook = Book {
		Title: "The Rainmaker",
		//Author: "John Grisham",
		Copies: 12,
		Series: true,
		//SeriesId: 1,
		PriceCents: 900,
		DiscountPercent: 10,
	}
	netPrice := myBook.NetPriceCalculate(myBook.PriceCents)
	//Integrer value returned from NetPriceCalculate function
	fmt.Printf("Net Price = $%v\n", netPrice)
	//Testing what an int reference returns when not configured. Comment out SeriesId in myBook
	fmt.Printf("Series ID = %v\n", myBook.SeriesId)
	//Testing what a string reference returns when not configured. Comment out Author in myBook	
	fmt.Printf("Author = %v\n", myBook.Author)
}

func (b Book) NetPriceCalculate(originalPrice int) int {
	return originalPrice - (originalPrice / 100 * b.DiscountPercent) 
}