package domain

type Purchase struct {
	ID                  string `db:"id"`
	BranchID            string `db:"branch_id"`
	CustomerID          string `db:"customer_id"`
	PointsPerPurchase   int64  `db:"points_per_purchase"`
	CashbackPerPurchase int64  `db:"cashback_per_purchase"`
	PurchasedAmount     int64  `db:"purchased_amount"`
}
