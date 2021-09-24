package main

import(
	"fmt"
	"net/http"
	"time"
)

func main(){
	urls := []string{
		"https://www.google.com",
		"https://www.golang.org",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.stackoverflow.com",
		"https://www.urldownmsqkdfjsdmkq.com",
	}
	
	ch1 := make(chan string)
	
	for _, url := range urls{
		go tryUrl(url, ch1)		
	}
	
	
	for u := range ch1{
		go func(url string){
			time.Sleep(3*time.Second)
			tryUrl(url, ch1)
		}(u)
	}
	
	
}

func tryUrl(url string, ch1 chan string){
	time.Sleep(3* time.Second)
	_, err := http.Get(url)
	if err != nil{
		fmt.Println("DOWN:", url)
		ch1 <- url
		return
	}
	fmt.Println("Online:", url)
	ch1 <- url

}
