package model

type MarketStatus struct {

	// The market close time on the holiday (if it's not closed).
	Close string `sql:"close"`

	// The date of the holiday.
	Date string `sql:"close"`

	// Which market the record is for.
	Exchange string `sql:"close"`

	// The name of the holiday.
	Name string `sql:"close"`

	// The market open time on the holiday (if it's not closed).
	Open string `sql:"close"`

	// The status of the market on the holiday.
	Status string `sql:"close"`

	// Whether or not the market is in post-market hours.
	AfterHours bool `sql:"close"`

	Currencies []Currencies `sql:"close"`

	// Whether or not the market is in pre-market hours.
	EarlyHours bool `sql:"close"`

	Exchanges []Exchanges `sql:"close"`

	// The status of the market as a whole.
	Market string `sql:"close"`

	// The current time of the server.
	ServerTime string `sql:"close"`
}

type Currencies struct {
	// The status of the Nasdaq market.
	Nasdaq string `sql:"close"`

	// The status of the NYSE market.
	Nyse string `sql:"close"`

	// The status of the OTC market.
	Otc string `sql:"close"`
}

type Exchanges struct {
	// The status of the crypto market.
	Crypto string `sql:"close"`

	// The status of the forex market.
	Fx string `sql:"close"`
}
