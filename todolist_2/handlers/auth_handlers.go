package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toyProject/todolist_2/models"
	"toyProject/todolist_2/utils"
)

func Sign_upHandler(db *sql.DB) gin.HandlerFunc {
	var requestUser models.User
	return func(c *gin.Context) {
		err := c.ShouldBindJSON(&requestUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "정확한 값을 입력하세요"})
		}

		query := "SELECT * FROM User WHERE username = ?"
		value, err := db.Query(query, requestUser.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "서버 에러, 회원가입을 다시 시도해주세요."})
			return
		}
		if value.Next() {
			c.JSON(http.StatusConflict, gin.H{"msg": "username이 중복됩니다."})
			return
		}
		bcryptPassword := utils.Generate(requestUser.Password)
		if bcryptPassword == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "password 암호화 실패."})
			return
		}
		query = "INSERT INTO User (username, password) VALUES (?,?)"
		result, err := db.Exec(query, requestUser.Username, bcryptPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "서버 에러, 회원가입을 다시 시도해주세요."})
			return
		}

		rowsUpdate, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "업데이트된 행 확인 중 오류"})
			return
		}

		if rowsUpdate > 0 {
			c.JSON(http.StatusOK, gin.H{"msg": "사용자 등록 성공"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "사용자 등록 실패"})
		}
	}
}

func Sign_inHandler(db *sql.DB) gin.HandlerFunc {
	var requestUser models.User
	return func(c *gin.Context) {
		err := c.ShouldBindJSON(&requestUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "정확한 값을 입력하세요"})
		}

		query := "SELECT * FROM User WHERE username = ?"
		value, err := db.Query(query, requestUser.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "1서버 에러, 회원가입을 다시 시도해주세요."})
			return
		}
		if !value.Next() {
			c.JSON(http.StatusConflict, gin.H{"msg": "username 이나 password가 잘못되었습니다."})
			return
		}

		query = "SELECT password FROM User WHERE username = ?"
		row := db.QueryRow(query, requestUser.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "2서버 에러, 회원가입을 다시 시도해주세요."})
			return
		}

		var password string
		scanerr := row.Scan(&password)
		if scanerr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "3서버 에러, 회원가입을 다시 시도해주세요."})
			return
		}
		if !utils.Compare(requestUser.Password, password) {
			c.JSON(http.StatusConflict, gin.H{"msg": "username 이나 password가 잘못되었습니다."})
			return
		}

		token, err := utils.GenerateJWT(requestUser.Username)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "4서버 에러, 회원가입을 다시 시도해주세요."})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":   "로그인 성공",
			"token": token,
		})

	}

}
