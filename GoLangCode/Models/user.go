package user
import(
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gorilla/mux"

)

var db *gorm.DB
var err error 
type User struct{
	gorm.Model
	Email string
	Password string
}
type Posts struct{
	gorm.Model
	Email string
	Title string
	Message string
}

func InitialMigration(){
	db,err := gorm.Open("sqlite3","test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&User{},&Posts{})
}
func Alluser(w http.ResponseWriter, r *http.Request){
	db,err:= gorm.Open("sqlite3","test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func NewUser(w http.ResponseWriter, r *http.Request){
	db,err:= gorm.Open("sqlite3","test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	email:=vars["email"]
	password:= vars["password"]
	db.Create(&User{Email:email,Password:password})
	fmt.Fprintf(w,"New User Created")
}

func PostUser(w http.ResponseWriter, r *http.Request){
	db,err:= gorm.Open("sqlite3","test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	email:=vars["email"]
	title:= vars["title"]
	message:=vars["message"]
	db.Create(&Posts{Email:email,Title:title,Message:message})
	fmt.Fprintf(w,"New POST Created")
}

func GetUserPost(w http.ResponseWriter, r *http.Request){
	db,err:= gorm.Open("sqlite3","test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	var posts []Posts
	db.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}