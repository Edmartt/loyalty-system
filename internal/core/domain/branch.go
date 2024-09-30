package domain

type Branch struct {
	ID         string `db:"id"`
	CommerceID string `db:"commerce_id"`
	Address    string `db:"address"`
}
