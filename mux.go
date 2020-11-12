package main
import
(
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
type Book struct{
	ID        string     `json:"id"`
	Isbn      string     `json:"isbn"`
	Title     string     `json:"title"`
	Author    *Author    `json:"author"`


}
type author Struct{
	Firstname    string     `json:"firstname"`
	Lastname     string      `json:"Lastname"`

}
var books []Book
func getBooks(w http.Responsewriter,r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(Books)

}
func getBook(w http.Responsewriter,r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,items := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encoder(item)
			return 
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}
func createBook(w http.Responsewriter,r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.body).Decode(&book)
	book.ID = strconv.Inta(rand.Intn(10000000))
	books = append(books,book)
	json.NewEncoder(w).Encode(book)


}
func updateBook(w http.Responsewriter,r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range books{
		if item.ID == params["id"]{
		books = append(books[:index],books[index+1:]...)
		var book Book
	    _ = json.NewDecoder(r.body).Decode(&book)
	    book.ID = strconv.Inta(rand.Intn(10000000))
	    books = append(books,book)
	    json.NewEncoder(w).Encode(book)
		}
	}
	json.NewEncoder(w).Encode(books)

}
func deleteBooks(w http.Responsewriter,r *http.Request)
{
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range books{
		if item.ID == params["id"]{
		books = append(books[:index],books[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(books)


}
func main(){
	r := mux.NewRouter()
	books = append(books, Book{ID : "1",Isbn: "448743",Title:"Book One",Author:&Author{Firstname: "john",Lastname:"Doe"}})
	books = append(books, Book{ID : "2",Isbn: "847564",Title:"Book Two",Author: &Author{Firstname: "steve",Lastname:"Smith"}})
	r.HandleFunc("/api/books",getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("GET")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books",deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))

}