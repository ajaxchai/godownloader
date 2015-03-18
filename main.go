// httpclient project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.17500.cn/getData/ssq.TXT")
	if err != nil {
		log.Fatal(err)
	}

	ssqData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("%s", ssqData)

}
