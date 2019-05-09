package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	/*var token struct {
		Token string `json:"token"`
	}
	err := json.NewDecoder(r.Body).Decode(&token)
	if err == nil {
		fmt.Println(token.Token)
	} else {
		fmt.Println("Huy tebe w rot!!!!!!!!!!!!!!!!")
	}*/

	r.ParseForm()
	fmt.Println(r.PostForm)

	file, header, _ := r.FormFile("file")

	defer file.Close()

	out, _ := os.OpenFile("./files/"+header.Filename,
		os.O_WRONLY|os.O_CREATE, 0666)

	defer out.Close()

	io.Copy(out, file)
}
