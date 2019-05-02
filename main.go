package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"./controllers"
)

func main() {
	//app := iris.Default()
	/*
		app.Get("/ping", func(context iris.Context) {
			context.JSON(iris.Map{
				"message": "Pong!",
			})
		})
	*/
	//app.Get("/ping", new(controllers.TestController))

	//app.Use(recover.New())

	/*
		app.RegisterView(iris.HTML("./views", ".html"))

		appMvc := mvc.New(app)

		appMvc.Handle(new(controllers.IndexController))
		appMvc.Handle(new(controllers.HelpController))
		appMvc.Handle(new(controllers.FileUploadController))
		appMvc.Handle(new(controllers.FileSendController))

		app.Run(iris.Addr(":8080"))
	*/

	router := mux.NewRouter()

	//router.HandleFunc("/load", controllers.SendFile)
	router.HandleFunc("/load", controllers.SendPrivateFile).Methods("POST")
	router.HandleFunc("/files", controllers.GetAllFiles).Methods("GET")

	http.ListenAndServe(":8080", router)
}
