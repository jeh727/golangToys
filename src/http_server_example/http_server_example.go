
package main

import(
    "fmt"
    "io/ioutil"
    "logger"
    "net/http"
    "strconv"
)

var response = "a"
var mapData = map[string]map[int64]string{}

func main(){


    logger.InitDefaultLogging()

    logger.Info.Println("Starting up")

    http.HandleFunc("/", handler)
    http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request){

    url := r.URL.Path
    values := r.URL.Query()

/*
    if(r.Body != nil) {
        defer r.Body.Close()
    }
*/
    var response string = ""
    switch r.Method {

        case "GET":
            indexString, ok := values["index"]
            if ok {
                index, err := strconv.ParseInt(indexString[0], 10, 64)
                if err != nil {

                    logger.Failure.Printf("Failed to parse index")
                    return
                }

                response = ""
                if mapData[url] != nil {

                    val, ok := mapData[url][index]
                    if ok {
                        //logger.Info.Printf("Returning data for url: %s, index: %d", url, index)
                        response = val
                    } else {
                        logger.Info.Printf("No data for url: %s, index: %d", url, index)
                    }
                } else {
                    logger.Info.Printf("No data for url: %s", url)
                }
            } else {
                logger.Info.Printf("No index provided")
                return
            }

        case "POST":
            body, err := ioutil.ReadAll(r.Body)
            r.Body.Close()
            if err == nil {

                indexString, ok := values["index"]
                if ok {

                    index, err := strconv.ParseInt(indexString[0], 10, 64)
                    if err != nil {

                        logger.Failure.Printf("Failed to parse index")
                        return
                    }

                    _, ok := mapData[url]
                    if !ok {
                       mapData[url] = map[int64]string{}
                    }

                    //logger.Info.Printf("Setting value for url: %s, index: %d", url, index)

                    mapData[url][index] = string(body)
                } else {
                    logger.Failure.Printf("Failed to find index")
                    return
                }
            } else {
                logger.Failure.Printf("Failed to get body")
                return
            }

            response = ""
        default:
            response = fmt.Sprintf("Got unknown Method: %s, URL: %s", r.Method, url)
    }

    fmt.Fprintf(w, response)
}
