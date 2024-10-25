package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// <-chan - canal somente-leitura
func titulos(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

//multiplexar - misturar (mensagem) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		titulos("https://www.cod3r.com.br", "https://www.google.com"),
		titulos("https://app.rocketseat.com.br", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
