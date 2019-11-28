package main

import ( 
	"fmt"
	"log"
	"net/http"
	"github.com/mikasd/youtube-stats/youtube"
	"github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func setupRoutes(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main(){
	fmt.Println("youtube")

	item, err := youtube.GetSubscribers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", item)


	// setupRoutes()
}