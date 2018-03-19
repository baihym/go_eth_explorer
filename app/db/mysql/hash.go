package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

func GetHashIdByHash(hash string) (int64, error) {
	var id int64

	if err := DBQueryRow(
		"SELECT `id` FROM `hashes` WHERE `hash`=?",
		hash,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func InsertHash(hash string) (int64, error) {
	result, err := DBExec("INSERT INTO `hashes` (`hash`) VALUES (?)", hash)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetHashIdByHashForce(hash string) (int64, error) {
	id, err := GetHashIdByHash(hash)
	if err == sql.ErrNoRows {
		id, err = InsertHash(hash)
	}
	return id, err
}

func GetHashIdByHashForceWithPanic(hash string) int64 {
	id, err := GetHashIdByHashForce(hash)
	if err == nil && id <= 0 {
		err = errors.New(fmt.Sprintf("Invalid id=%d", id))
	}
	if err != nil {
		panic(fmt.Sprintf("Get hash err, hash: %s, err: %s", hash, err.Error()))
	}
	return id
}
