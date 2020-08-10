// Package main provides ...
package main

//DB .....
type DB interface {
	Get(key string) (int, error)
}

//GetFromDB .....
func GetFromDB(db DB, key string) int {
	v, err := db.Get(key)
	if err != nil {
		return -1

	}
	return v
}
