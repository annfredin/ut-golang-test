package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/go-server/handler"
)

//http://localhost:8080/postbook
// params:
// {
//     "content":"In software engineering, software design pattern is commonly occurring problem within given context in software design. It is not a finished design that can be transformed directly into source or machine code"
// }
func main() {
    router := gin.New()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	router.POST("/postbook", func(c *gin.Context) {
       
		var newBook handler.Book
		// Call BindJSON to bind the received JSON to
		// newBook.
		if err := c.BindJSON(&newBook); err != nil {
			c.JSON(400, gin.H{"status": "Invalid Input"})
			return
		}
		
		result, err:= handler.PostBook(&newBook)
		if err != nil{
			c.JSON(500, make(map[string]int,0))
		}

		// Your custom response here
        c.JSON(200, result) 
    })
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

