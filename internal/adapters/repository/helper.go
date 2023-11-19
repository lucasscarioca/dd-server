package repository

import (
	"database/sql"
)

// nullString converts a string to sql.NullString for empty string check
func nullString(value string) sql.NullString {
	if value == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

// nullBytes converts a byte array to a sql.NullByte array for empty []byte check
func nullBytes(value []byte) []sql.NullByte {
	if value == nil {
		return []sql.NullByte{}
	}

	var data []sql.NullByte
	for _, b := range value {
		data = append(data, sql.NullByte{
			Byte:  b,
			Valid: true,
		})
	}
	return data
}

// func nullJson(value map[string]interface{}) sql.NullString {
// 	d, err := json.Marshal(value)
// 	if err != nil {
// 		return sql.NullString{}
// 	}

// 	return sql.NullString{
// 		String: string(d),
// 		Valid:  true,
// 	}
// }

// // nullUint64 converts an uint64 to sql.NullInt64 for empty uint64 check
// func nullUint64(value uint64) sql.NullInt64 {
// 	if value == 0 {
// 		return sql.NullInt64{}
// 	}

// 	valueInt64 := int64(value)

// 	return sql.NullInt64{
// 		Int64: valueInt64,
// 		Valid: true,
// 	}
// }

// // nullInt64 converts an int64 to sql.NullInt64 for empty int64 check
// func nullInt64(value int64) sql.NullInt64 {
// 	if value == 0 {
// 		return sql.NullInt64{}
// 	}

// 	return sql.NullInt64{
// 		Int64: value,
// 		Valid: true,
// 	}
// }

// // nullFloat64 converts a float64 to sql.NullFloat64 for empty float64 check
// func nullFloat64(value float64) sql.NullFloat64 {
// 	if value == 0 {
// 		return sql.NullFloat64{}
// 	}

// 	return sql.NullFloat64{
// 		Float64: value,
// 		Valid:   true,
// 	}
// }
