package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"unicode"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	// "golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

var user User

var store = sessions.NewCookieStore([]byte("SESSION_KEY"))

func main() {
	// db, err = sql.Open("{sql-type}", "{username}:{password}@tcp({server:port})/{database-name}")
	db, err = sql.Open("mysql", "root:moaz@tcp(127.0.0.1:3306)/Travelbangla")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("connected to database")

	http.HandleFunc("/", home)
	http.HandleFunc("/service", service)
	http.HandleFunc("/about", about)
	http.HandleFunc("/package", tourPackage)
	http.HandleFunc("/tour-guide", tourGuide)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/signup-handler", signUpHandler)
	http.HandleFunc("/update-user", updateUser)
	http.HandleFunc("/update-user-handler", updateUserHandler)
	http.HandleFunc("/login", logIn)
	http.HandleFunc("/login-handler", logInHandler)

	http.HandleFunc("/my-account", dashboard)
	http.HandleFunc("/my-package", myPackage)
	// http.HandleFunc("/cart", cart)
	// http.HandleFunc("/checkout", checkout)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/admin-resources/", http.StripPrefix("/admin-resources/", http.FileServer(http.Dir("./dashboard/assets"))))

	// http.Handle("/video/", http.StripPrefix("/video/", addHeaders(http.FileServer(http.Dir("video")))))
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// localserver

	fmt.Println("starting server on : 8008")
	http.ListenAndServe(":8008", context.ClearHandler(http.DefaultServeMux))

	// heroku server
	// port := os.Getenv("PORT")
	// log.Print("Listening on :" + port)
	// log.Fatal(http.ListenAndServe(":"+port, context.ClearHandler(http.DefaultServeMux)))

}

// func addHeaders(h http.Handler) http.HandlerFunc {
// 	return func(rw http.ResponseWriter, r *http.Request) {
// 		ext := filepath.Ext(r.RequestURI)
// 		rw.Header().Set("Access-Control-Allow-Origin", "*")
// 		rw.Header().Set("Access-Control-Allow-Headers", "Range")
// 		// rw.Header().Set("MimeType", "application/dash+xml")
// 		if ext == ".mpd" {
// 			rw.Header().Set("MimeType", "application/dash+xml")
// 		} else if ext == ".m3u8" {
// 			rw.Header().Set("MimeType", "application/vnd.apple.mpegurl")
// 		}

// 		h.ServeHTTP(rw, r)
// 	}
// }

func home(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(res, nil)
}

func service(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp, err = ptmp.ParseFiles("wpage/service.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(res, nil)
}

func about(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp, err = ptmp.ParseFiles("wpage/about.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(res, nil)
}

func tourPackage(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp, err = ptmp.ParseFiles("wpage/package.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(res, nil)
}

func tourGuide(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp, err = ptmp.ParseFiles("wpage/guide.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(res, nil)
}

func contact(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp, err = ptmp.ParseFiles("wpage/contact.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(res, nil)
}

func signUp(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp, err = ptmp.ParseFiles("wpage/signup.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(res, nil)
}

func signUpHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****registerAuthHandler running*****")
	req.ParseForm()
	username := req.FormValue("name-input")
	email := req.FormValue("email-input")
	phone := req.FormValue("phone-input")
	password := req.FormValue("pass-input")

	// Validate username
	var nameAlphaNumeric = true
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			nameAlphaNumeric = false
			break
		}
	}
	if len(username) < 4 || len(username) > 50 || !nameAlphaNumeric {
		displaySignupError(res, "Invalid username")
		return
	}

	// Validate password
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			pswdLowercase = true
		case unicode.IsUpper(char):
			pswdUppercase = true
		case unicode.IsNumber(char):
			pswdNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		case unicode.IsSpace(char):
			pswdNoSpaces = false
		}
	}
	if len(password) < 8 || !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdNoSpaces {
		displaySignupError(res, "Invalid password")
		return
	}

	// Check if username already exists
	var uID string
	err := db.QueryRow("SELECT u_name FROM users WHERE u_name=?", username).Scan(&uID)
	if err != sql.ErrNoRows {
		displaySignupError(res, "Username already taken")
		return
	}

	// Check if email already used
	var eID string
	err = db.QueryRow("SELECT u_email FROM users WHERE u_email=?", email).Scan(&eID)
	if err != sql.ErrNoRows {
		displaySignupError(res, "Email already used")
		return
	}

	// Insert new user into the database
	insertStmt, err := db.Prepare("INSERT INTO users(u_name, u_email, u_password, u_phone) VALUES(?, ?, ?, ?);")
	if err != nil {
		displaySignupError(res, "There was a problem registering account")
		return
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(username, email, password, phone)
	if err != nil {
		displaySignupError(res, "There was a problem registering account")
		return
	}

	fmt.Println("user created")
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}

func displaySignupError(res http.ResponseWriter, message string) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ptmp, err = ptmp.ParseFiles("wpage/signup.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ptmp.Execute(res, message)
}

func logIn(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp, err = ptmp.ParseFiles("wpage/login.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(res, nil)
}

func logInHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("*****loginAuthHandler running*****")
	req.ParseForm()
	username := req.FormValue("name-input")
	password := req.FormValue("pass-input")
	fmt.Println("username:", username, "password:", password)

	var userID, dbPass string
	err := db.QueryRow("SELECT u_id, u_password FROM users WHERE u_name=?", username).Scan(&userID, &dbPass)
	if err != nil {
		displayLoginError(res, "Check username and password")
		return
	}

	if dbPass == password {
		session, _ := store.Get(req, "session")
		session.Values["userID"] = userID
		session.Save(req, res)
		http.Redirect(res, req, "/my-account", http.StatusSeeOther)
		return
	}

	displayLoginError(res, "Check password")
}

func displayLoginError(res http.ResponseWriter, message string) {
	ptmp, err := template.ParseFiles("template/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ptmp, err = ptmp.ParseFiles("wpage/login.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ptmp.Execute(res, message)
}

// Handler to show the update profile form
func updateUser(res http.ResponseWriter, req *http.Request) {
	// Fetch the logged-in user's data from session or database
	// Assuming session has the userID
	session, _ := store.Get(req, "session")
	userID, ok := session.Values["userID"]
	if !ok || userID == nil {
		http.Error(res, "User not logged in", http.StatusUnauthorized)
		return
	}

	var user User
	err := db.QueryRow("SELECT u_name, u_email, u_phone FROM users WHERE u_id=?", userID).Scan(&user.UserName, &user.UserEmail, &user.UserPhone)
	if err != nil {
		http.Error(res, "User not found", http.StatusInternalServerError)
		return
	}

	// Render the update profile form with user's data
	ptmp, err := template.ParseFiles("template/dashboard.html")
	if err != nil {
		http.Error(res, "Template not found", http.StatusInternalServerError)
		return
	}

	ptmp.Execute(res, user)
}

// Handler to process form data and update the user's profile
func updateUserHandler(res http.ResponseWriter, req *http.Request) {
	// Check for POST method
	if req.Method != http.MethodPost {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := req.ParseForm()
	if err != nil {
		http.Error(res, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	// Get session to fetch the userID
	session, _ := store.Get(req, "session")
	userID := session.Values["userID"]

	// Extract form values
	username := req.FormValue("name-input")
	email := req.FormValue("email-input")
	phone := req.FormValue("phone-input")
	password := req.FormValue("password-input")

	fmt.Println("Form values:", userID, username, email, phone, password)

	// Optional: You can revalidate the inputs if needed (similar to signupHandler)

	// Prepare the SQL query for updating user details
	updateQuery := `UPDATE users SET u_name=?, u_email=?, u_phone=? WHERE u_id=?`

	// If the password field is provided, update the password too
	if password != "" {
		updateQuery = `UPDATE users SET u_name=?, u_email=?, u_phone=?, u_password=? WHERE u_id=?`
		_, err = db.Exec(updateQuery, username, email, phone, password, userID)
	} else {
		updateQuery = `UPDATE users SET u_name=?, u_email=?, u_phone=? WHERE u_id=?`
		_, err = db.Exec(updateQuery, username, email, phone, userID)
	}

	if err != nil {
		// log.Println("SQL Error:", err)  // Add logging for better debugging
		http.Error(res, "Unable to update profile", http.StatusInternalServerError)
		return
	}

	// Handle SQL errors
	if err != nil {
		http.Error(res, "Unable to update profile", http.StatusInternalServerError)
		return
	}

	// Redirect to the profile or dashboard after update
	http.Redirect(res, req, "/my-account", http.StatusSeeOther)
}

func dashboard(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	userID, ok := session.Values["userID"]
	if !ok {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	var sessionUserName, sessionUserEmail, sessionUserPhone string
	db.QueryRow("SELECT u_name, u_email, u_phone FROM users WHERE u_id=?", userID).Scan(&sessionUserName, &sessionUserEmail, &sessionUserPhone)

	user := User{
		UserID:    fmt.Sprint(userID),
		UserName:  sessionUserName,
		UserEmail: sessionUserEmail,
		UserPhone: sessionUserPhone,
	}

	ptmp, err := template.ParseFiles("dashboard/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ptmp.Execute(res, user)
}

func myPackage(res http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("dashboard/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp, err = ptmp.ParseFiles("dashboard/wpage/my-package.html")
	if err != nil {
		fmt.Println(err.Error())

	}

	ptmp.Execute(res, nil)
}
