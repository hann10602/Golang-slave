package handler

import (
	"gorm_demo/db"
	"gorm_demo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	GetUsersMapConditionsParams struct {
		Params map[string] interface{}
	}

	GetUsersSliceConditionsParams struct {
		Ids []int64 `json:"ids"`
	}

	GetUsersOrConditionsParams struct {
		Username string `json:"username"`
		User models.User
	}

	GetUsersDTOParams struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UpdateUserParams struct {
		Username *string `json:"username, omitempty"`
		Password *string `json:"password, omitempty"`
	}

	DeleteUserParams struct {
		Ids []int64 `json:"ids"`
	}
)

func GetUsers(c echo.Context) error {
	var Db = db.DbManager()

	users := []models.User{}
	result := Db.Find(&users)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUsersWhereConditions(c echo.Context) error {
	var Db = db.DbManager()

	users := []models.User{}

	username := c.QueryParam("username")
	password := c.QueryParam("password")

	result := Db.Where("username = ? AND password = ?", username, password).Find(&users)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUsersStructConditions(c echo.Context) error {
	var Db = db.DbManager()

	username := c.QueryParam("username")
	password := c.QueryParam("password")

	users := []models.User{}
	result := Db.Where(&models.User{Username: username, Password: password }).Find(&users)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUsersMapConditions(c echo.Context) error {
	var Db = db.DbManager()

	users := []models.User{}
	var params GetUsersMapConditionsParams

	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := Db.Where(params.Params).Find(&users)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUsersSliceConditions(c echo.Context) error {
	var Db = db.DbManager()

	users := []models.User{}
	var params GetUsersSliceConditionsParams

    if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := Db.Where(params.Ids).Find(&users)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUsersOrConditions(c echo.Context) error {
	var Db = db.DbManager()

	users := []models.User{}
	
	var params *GetUsersOrConditionsParams

	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := Db.Where("username = ?", params.Username).Or(params.User).Find(&users) 

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUsersSelectFields(c echo.Context) error {
	var Db = db.DbManager()

	users := []models.User{}

	result := Db.Select("ID", "username").Find(&users) 

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUsersDTO(c echo.Context) error {
	var Db = db.DbManager()

	users := []GetUsersDTOParams{}

	result := Db.Model(&models.User{}).Find(&users) 

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	var Db = db.DbManager()

	id := c.Param("id")
	user := models.User{}

	result := Db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  "User does not exist")
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	var Db = db.DbManager()

	user := &models.User{}

    if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := Db.Create(&user)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  result.Error.Error())
	}

	return c.JSON(http.StatusOK, "Create successfully")
}

func UpdateUser(c echo.Context) error {
	Db := db.DbManager()

	id := c.Param("id")
	
	userDto := UpdateUserParams{}
	user := models.User{}

	if getResult := Db.Where("id = ?", id).First(&user); getResult.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	if err := c.Bind(&userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if saveResult := Db.Model(&user).Updates(&userDto); saveResult.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  saveResult.Error.Error())
	}

	return c.JSON(http.StatusOK, "Update successfully")
}

func UpdateUserBySave(c echo.Context) error {
	Db := db.DbManager()

	id := c.Param("id")
	
	userDto := UpdateUserParams{}
	user := models.User{}

	if getResult := Db.Where("id = ?", id).First(&user); getResult.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User does not exist")
	}

	if err := c.Bind(&userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if userDto.Username != nil {
		user.Username = *userDto.Username
	}

	if userDto.Password != nil {
		user.Password = *userDto.Password
	}

	if saveResult := Db.Save(&user); saveResult.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest,  saveResult.Error.Error())
	}

	return c.JSON(http.StatusOK, "Update successfully")
}

func DeleteUsers(c echo.Context) error {
	Db := db.DbManager()

	ids := DeleteUserParams{}

	if err := c.Bind(&ids); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	Db.Delete(&models.User{}, ids.Ids)

	return c.JSON(http.StatusOK, "Delete successfully")
}