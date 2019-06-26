package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
    "JSMPJ_go_db_vue/GoLangCode/Models"
)
func handleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is Home page")
}
func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",handleHome).Methods("GET")
	myRouter.HandleFunc("/user",user.Alluser).Methods("GET")
	myRouter.HandleFunc("/user/{email}/{password}",user.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{email}",user.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{email}/{password}",user.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000",myRouter))
}
func main(){
	fmt.Println("JSMPJ Corporation")
	user.InitialMigration()
	handleRequests()

}