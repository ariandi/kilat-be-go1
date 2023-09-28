package util

import "database/sql"

func IsSqlValid(sqlStr sql.NullString) string {
	if sqlStr.Valid {
		return sqlStr.String
	}

	return ""
}
