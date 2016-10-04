
package main

import(
    "fmt"
    "io/ioutil"
    "logger"
    "net/http"
    "os"
    "strings"
    "time"
)

func main(){

    logger.InitDefaultLogging()

    logger.Info.Println("Starting up")

    startTime := time.Now()
    for loop := int64(0); loop < 100000; loop++ {
        StoreData("/topic", loop, fmt.Sprintf("Payload %d", loop))
    }
    logger.Info.Printf("Time: %s", time.Since(startTime))

    logger.Info.Printf("Got Back: %s", LoadData("", 0))
}

func StoreData(topic string, index int64, payload string) {

    url := fmt.Sprintf("http://localhost:8000%s?index=%d", topic, index)

    Post(url, payload)
}

func LoadData(topic string, index int64) string {

    url := fmt.Sprintf("http://localhost:8000%s?index=%d", topic, index)

    return Get(url)
}
func Get(url string) string {

    resp, err := http.Get(url)
    if err != nil {
        logger.Error.Printf("Failed to GET %s\n", url)
        return ""
    }

    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {

        logger.Error.Printf("Failed to read body")
        return ""
    }

    logger.Info.Printf("%s", body)

    return string(body)
}

func Post(url string, payload string){

    resp, err := http.Post(url, "text/plain; charset=UTF-8", strings.NewReader(payload))
    if err != nil {
        logger.Error.Printf("Failed to POST %s\n", url)
        os.Exit(1)
    }

    _, err = ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {

        logger.Error.Printf("Failed to read body")
        os.Exit(1)
    }

    //logger.Info.Printf("Response: %s", body)
}
