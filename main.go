package main

import (
        "fmt"
        "log"
        "net/http"
        "time"
        "os"
        "io"
        "flag"
)

func main() {

        port := flag.String("port", "9001", "Port")
        flag.Parse()

	handler := func(w http.ResponseWriter, r *http.Request) {

                w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

                timestamp := time.Now().UTC().Format(time.RFC3339)

                filename := timestamp
                index := 1

                // loop until we find an unused filename
                for {
                        _, err := os.Stat(filename)
                        if os.IsNotExist(err) {
                                break
                        } else {
                                filename = fmt.Sprintf("%s_%d", timestamp, index)
                                index += 1
                        }
                }


                f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
                if err != nil {
                        log.Fatal(err)
                }

                defer f.Close()

                io.Copy(f, r.Body)
        }

        log.Println("Starting up")
        log.Fatal(http.ListenAndServe(":"+*port, http.HandlerFunc(handler)));
}
