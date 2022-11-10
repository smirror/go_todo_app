// @Title
// @Description
// @Author smirror
// @Update 2022-09-23

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc(
			func(w http.ResponseWriter, request *http.Request) {
				fmt.Fprintf(w, "Hello, %s", request.URL.Path[1:])
			}),
	)
	if err != nil {
		fmt.Printf("faile to terminate server: %v", err)
		os.Exit(1)
	}
}
