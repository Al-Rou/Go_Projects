package CheckPassword

/*
In absence of a database, these two variables are to be used to authenticate the entered
username and password.
*/
var username string = "Al"
var password string = "1234"

//This function is to check whether entered username and password are correct
func CheckPass(inputUsername string, inputPassword string) bool {
	if inputUsername == username && inputPassword == password {
		return true
	}
	return false
}
