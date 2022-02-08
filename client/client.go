package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"powClient/pow"
	"powClient/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	ip := "157.234.240.223"
	date := "060408"
	url := "http://127.0.0.1:9797/req"
	tick := time.Tick(2 * time.Second)
	var resp *http.Response

	for {
		select {
		case <-tick:
			//this data just for test, in real case it should be collected

			randStr := rand.String(4)

			//TODO perhaps not needed to put into struct
			proof := pow.NewProofOfWork(pow.HashCash{
				IP:        ip,
				Date:      date,
				RandomStr: randStr,
			})
			counter := proof.Run()

			b := strings.Builder{}
			b.WriteString(ip)
			b.WriteString(":")
			b.WriteString(date)
			b.WriteString(":")
			b.WriteString(strconv.Itoa(counter))
			b.WriteString(":")
			b.WriteString(randStr)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatal()
			}
			client := &http.Client{}
			req.Header.Add("X-Hashcash", b.String())
			resp, err = client.Do(req)
			if err != nil {
				log.Fatalf("hello error '%s'\n", err)
			}

			fmt.Println(resp.Status)
			if resp.StatusCode == http.StatusOK {
				bodyBytes, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				bodyString := string(bodyBytes)
				fmt.Println(bodyString)
				fmt.Println()
			}
			resp.Body.Close()
		}
	}
	defer resp.Body.Close()
}
