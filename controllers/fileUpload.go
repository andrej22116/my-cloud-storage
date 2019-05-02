package controllers

import (
	"io"
	"net/http"
	"os"
)

/*
FileUploadController used for geting files from users
*/
/*
type FileUploadController struct{}
*/
/*
GetUpload return html page for load file
*/
/*
func (controller *FileUploadController) GetUpload() mvc.Result {
	/*now := time.Now().Unix()
	hash := md5.New()
	io.WriteString(hash, strconv.FormatInt(now, 10))

	// render the form with the token for any use you'd like.
	//context.View("upload_form.html", token)
	return mvc.View{
		Name: "upload_form.html",
	}
}
*/
/*
PostUpload upload file from user
*/
/*
func (controller *FileUploadController) PostUpload(context iris.Context) {
	file, info, err := context.FormFile("uploadfile")

	if err != nil {
		context.StatusCode(iris.StatusInternalServerError)
		context.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	defer file.Close()
	fname := info.Filename

	out, err := os.OpenFile("./files/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		context.StatusCode(iris.StatusInternalServerError)
		context.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer out.Close()

	io.Copy(out, file)
}
*/

func UploadFile(responsWriter http.ResponseWriter, request *http.Request) {
	file, header, _ := request.FormFile("uploadfile")

	/*if err != nil {
		request.setS
		request.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}*/

	defer file.Close()

	out, _ := os.OpenFile("./files/"+header.Filename,
		os.O_WRONLY|os.O_CREATE, 0666)

	/*if err != nil {
		request.StatusCode(http.StatusInternalServerError)
		request.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return
	}*/
	defer out.Close()

	io.Copy(out, file)
}
