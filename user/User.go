package user

import (
	"Echo1/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func Create(ctx echo.Context) error {

	db := database.Connection()

	defer db.Close()

	//name := ctx.FormValue("name")
	//surname := ctx.FormValue("surname")
	//age := ctx.FormValue("age")

	user := User{}

	err := ctx.Bind(&user)

	if err != nil {
		panic(err.Error())
	}

	insert, err := db.Query("INSERT INTO users (name,surname,age) values (?,?,?)", user.Name, user.Surname, user.Age)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	return ctx.JSON(http.StatusOK, user)
}

func Info(ctx echo.Context) error {

	db := database.Connection()

	defer db.Close()

	var users []User

	userList, err := db.Query("SELECT * from users")

	if err != nil {
		panic(err.Error())
	}

	for userList.Next() {
		user := User{}

		err := userList.Scan(&user.Id, &user.Name, &user.Surname, &user.Age)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)

	}

	return ctx.JSON(http.StatusOK, users)
}

func Detail(ctx echo.Context) error {
	id := ctx.Param("id")

	db := database.Connection()

	user := User{}

	err := db.QueryRow("SELECT * from  users where id=?", id).Scan(&user.Id, &user.Name, &user.Surname, &user.Age)

	if err != nil {
		panic(err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func Delete(ctx echo.Context) error {
	id := ctx.FormValue("id")

	db := database.Connection()

	_, err := db.Query("DELETE from users where id = ?", id)

	if err != nil {
		panic(err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"response": "Kullanıcı silindi.",
	})
}
