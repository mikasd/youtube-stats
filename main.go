package main

import ( 
	"fmt"
	"log"
	"net/http"
	"github.com/mikasd/youtube-stats/youtube"
	"github.com/joho/godotenv"
	"github.com/mikasd/youtube-stats/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func stats(w http.ResponseWriter, r *http.Request){
	ws, err := websocket.Upgrade(w,r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
}

func setupRoutes(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
	fmt.Println("youtube subscriber monitor")

	err1 := godotenv.Load()
  	if err1 != nil {
    	log.Fatal("Error loading .env file")
  	}

	item, err := youtube.GetSubscribers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", item)



	setupRoutes()
}