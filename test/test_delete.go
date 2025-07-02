package main

import "fmt"

type Book struct{
	ID int
	Title string
	Author string
}
func main() {
	books := []Book{
	{ID: 1, Title: "HarryPoter", Author: "Peter"},
	{ID: 2, Title: "Ironman", Author: "Marvel"},
	{ID: 3, Title: "Focus", Author: "John"},
	{ID: 4, Title: "John Wick", Author: "JJ"},
	}
	
	i := 0
	fmt.Println("ก่อนลบ:", books)

	fmt.Println("ตัวที่โดนลบ",books[i])

	books = append(books[:i],books[i+1:]...)

	fmt.Println("หลังลบ:", books)

}