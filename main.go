package main
import (
	    "encoding/json"
	        "fmt"
		    "log"
		        "net/http"
		)
		var greeting = map[string]string{
			    "hello": "world",
		    }
		    func homePage(w http.ResponseWriter, r *http.Request) {
			        JsoninBytes, err := json.Marshal(greeting)
				    if err != nil {
					            fmt.Println(err)
					    }
					    fmt.Fprintf(w, string(JsoninBytes))
					        fmt.Println("Endpoint Hit: Homepage")
					}
					func main() {
						    http.HandleFunc("/", homePage)
						        log.Fatal(http.ListenAndServe(":8080", nil))
						}

