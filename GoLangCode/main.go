package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
    "JSMPJ_go_db_vue/GoLangCode/Models"
)
func handleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is Home page")
}
func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	methods := handlers.AllowedHeaders([]string{"GET","POST","PUT","DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	myRouter.HandleFunc("/",handleHome).Methods("GET")
	myRouter.HandleFunc("/user",user.Alluser).Methods("GET")
	myRouter.HandleFunc("/user/{email}/{password}",user.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{email}/{title}/{message}",user.PostUser).Methods("POST")
	myRouter.HandleFunc("/user/{email}",user.GetUserPost).Methods("GET")
	myRouter.HandleFunc("/user/{id}",user.DeleteUserPost).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",handlers.CORS(methods,origins)(myRouter)))
}
func main(){
	fmt.Println("JSMPJ Corporation")
	user.InitialMigration()
	handleRequests()

}