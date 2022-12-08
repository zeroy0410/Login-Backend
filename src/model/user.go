package model

type User struct {
	ID				uint		`json:"id"`
	Admin			bool 		`json:"admin"`
	Password		string 		`json:"password"`
	NickName		string 		`json:"nick-name"`
	Email			string 		`json:"email"`
}

var Zeroy User

func SetZeroy(){
	Zeroy=User{
		2020150384,
		true,
		"123456",
		"zeroy",
		"zeroy0410@gmail.com",
	}
}