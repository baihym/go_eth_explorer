package mysql

var DBTokens = make(map[string]int64)

func InitDBTokens() error {

	rows, err := DBQuery(
		"SELECT id,contract FROM tokens WHERE id != 1",
	)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		var contract string
		if err := rows.Scan(&id, &contract); err != nil {
			return err
		}
		DBTokens[contract] = id
	}

	return nil
}
