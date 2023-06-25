package cmd

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/user/:userID", func(c echo.Context) error {
		userID := c.Param("userID")

		// データベースから userID に対応するユーザー情報を取得する
		user, err := getUserFromDatabase(userID)
		if err != nil {
			// エラーハンドリング
			return c.String(http.StatusInternalServerError, "Error: "+err.Error())
		}

		// ユーザー情報を利用してレスポンスを生成する
		response := "User ID: " + user.ID + ", Name: " + user.Name
		return c.String(http.StatusOK, response)
	})

	return e
}

type User struct {
	ID   string
	Name string
}

func getUserFromDatabase(userID string) (*User, error) {
	// MySQL データベースへの接続情報
	db, err := sql.Open("mysql", "root:your_password@tcp(localhost:3307)/oneday")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// ユーザー情報を取得するクエリ
	query := "SELECT id, name FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)

	// ユーザー情報を格納する変数
	user := User{}

	// クエリの結果を取得して変数にセットする
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		// ユーザーが見つからなかった場合の処理
		if err == sql.ErrNoRows {
			// デフォルト値を設定するなどの処理を行う
			user.ID = userID
			user.Name = "Unknown"
			return &user, nil
		}
		return nil, err
	}

	return &user, nil
}
