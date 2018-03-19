package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

func GetAddressIdByAddress(address string) (int64, error) {
	var id int64

	if err := DBQueryRow(
		"SELECT `id` FROM `addresses` WHERE `address`=?",
		address,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func InsertAddress(address string) (int64, error) {
	result, err := DBExec("INSERT INTO `addresses` (`address`) VALUES (?)", address)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetAddressIdByAddressForce(address string) (int64, error) {
	id, err := GetAddressIdByAddress(address)
	if err == sql.ErrNoRows {
		id, err = InsertAddress(address)
	}
	return id, err
}

func GetAddressIdByAddressForceWithPanic(address string) int64 {
	id, err := GetAddressIdByAddressForce(address)
	if err == nil && id <= 0 {
		err = errors.New(fmt.Sprintf("Invalid id=%d", id))
	}
	if err != nil {
		panic(fmt.Sprintf("Get address err, address: %s, err: %s", address, err.Error()))
	}
	return id
}
