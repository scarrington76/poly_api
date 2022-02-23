package model

type MarketHolidaysResponse struct {
	// The market close time on the holiday (if it's not closed).
	Close string `sql:"close"`

	// The date of the holiday.
	Date string `sql:"date"`

	// Which market the record is for.
	Exchange string `sql:"exchange"`

	// The name of the holiday.
	Name string `sql:"name"`

	// The market open time on the holiday (if it's not closed).
	Open string `sql:"open"`

	// The status of the market on the holiday.
	Status string `sql:"status"`
}
