package database

import (
	"SimpleService/pkg/model"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type SqlStorage struct {
	base *sql.DB
}

func New(options string) (*SqlStorage, error) {
	database, err := sql.Open("postgres", options)
	if err != nil {
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successful connect")
	return &SqlStorage{
		base: database,
	}, nil
}

func (s *SqlStorage) GetTotalCount() (int, error) {
	var res int

	rows, err := s.base.Query("SELECT COUNT(*) FROM article")
	if err != nil {
		return 0, err
	}

	if rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			return 0, err
		}
	}

	return res, nil
}

func (s *SqlStorage) Search(limit, offset int) ([]model.Tribe, error) {

	query := fmt.Sprintf("SELECT * FROM article ORDER BY id LIMIT %d OFFSET (%d) * %d", limit, offset, limit)
	rows, err := s.base.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data model.Tribe
	var result []model.Tribe

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Name, &data.Description, &data.Content)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}

	return result, nil
}

func (s *SqlStorage) SearchById(id int) (model.Tribe, error) {
	var res model.Tribe
	query := fmt.Sprintf("SELECT * FROM article WHERE id=%d", id)
	rows, err := s.base.Query(query)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&res.Id, &res.Name, &res.Description, &res.Content)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (s *SqlStorage) PublishArticle(data []string) error {

	if len(data) < 3 {
		err := errors.New("incorrect create article")
		return err
	}
	rows, err := s.base.Query("SELECT COUNT(*) FROM article")
	if err != nil {
		return err
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
	}

	query := fmt.Sprintf("INSERT INTO article (id, name, description, content) VALUES (%d, '%s', '%s', '%s')", id+1, data[0], data[1], data[2])
	_, err = s.base.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *SqlStorage) Authorization(login, password string) (string, error) {
	var token string
	if login == "" || password == "" {
		err := errors.New("empty login or password")
		return token, err
	}

	rows, err := s.base.Query("SELECT login, password FROM users WHERE root")
	if err != nil {
		return token, err
	}
	defer rows.Close()

	result := make([]string, 2)
	if rows.Next() {
		err := rows.Scan(&result[0], &result[1])
		if err != nil {
			return token, err
		}
	}

	if login != result[0] || password != result[1] {
		err := errors.New("incorrect login or password")
		return token, err
	}
	
	token, err = GenerateToken()
	if err != nil {
		return token, err
	}
	return token, nil
}

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	var Key []byte
	tokenString, err := token.SignedString(Key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
