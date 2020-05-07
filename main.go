package main

import (
        "fmt"
        "log"
        "net/http"
        "time"
        "os"
        "io"
        "io/ioutil"
        "flag"
)

func main() {

        port := flag.String("port", "9001", "Port")
        flag.Parse()

	handler := func(w http.ResponseWriter, r *http.Request) {

                if r.Method != "POST" {
                        w.WriteHeader(http.StatusMethodNotAllowed)
                        w.Write([]byte("Request must be a POST"))
                        return
                }

                var maxReportBytes int64 = 10*1024;
                body, err := ioutil.ReadAll(io.LimitReader(r.Body, maxReportBytes))
                if err != nil || len(body) == 0 {
                        w.WriteHeader(http.StatusBadRequest)
                        w.Write([]byte("Must provide POST body for report"))
                        return
                }

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


                err = ioutil.WriteFile(filename, body, 0644)
                if err != nil {
                        w.WriteHeader(http.StatusInternalServerError)
                        w.Write([]byte("Failed to write file"))
                        return
                }
        }

        log.Println("Starting up")
        http.HandleFunc("/eGJvfRfF300fGpxnB52LmFpD9IIJPzYb", handler);
        log.Fatal(http.ListenAndServe(":"+*port, nil));
}
