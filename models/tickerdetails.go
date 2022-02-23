package model

// A single ticker response that will have detailed information about the ticker and the company behind it.
type TickerDetails struct {

	//Ticker with details.
	Results []TickerDetailsResults `sql:"adjusted"`

	// A request id assigned by the server.
	Request_id string `sql:"adjusted"`

	// The total number of results for this request.
	Count int64 `sql:"adjusted"`

	// The status of this request's response.
	Status string `sql:"adjusted"`
}

type TickerDetailsResults struct {

	// Whether or not the asset is actively traded. False means the asset has been delisted.
	Active bool `sql:"active"`

	// Address of the company (object)
	Address []Address `sql:"adjusted"`

	// Links to the company's branding/logos (object)
	Branding []Branding `sql:"adjusted"`

	// The CIK number for this ticker.
	Cik string `sql:"adjusted"`

	// The composite OpenFIGI number for this ticker.
	Composite_figi string `sql:"adjusted"`

	// The name of the currency that this asset is traded with.
	Currency_name string `sql:"adjusted"`

	// The last date that the asset was traded.
	Delisted_utc string `sql:"adjusted"`

	// A description of the company and what they do/offer.
	Description string `sql:"adjusted"`

	// The URL of the company's website homepage.
	Homepage_url string `sql:"adjusted"`

	// The date that the symbol was first publicly listed in the format YYYY-MM-DD.
	List_date string `sql:"adjusted"`

	// The locale of the asset.
	Locale string `sql:"adjusted"`

	// The market type of the asset.
	Market string `sql:"adjusted"`

	// The most recent close price of the ticker multiplied by weighted outstanding shares.
	Market_cap int64 `sql:"adjusted"`

	// The name of the asset. For stocks/equities this will be the companies registered name. For crypto/fx this will be the name of the currency or coin pair.
	Name string `sql:"adjusted"`

	// The phone number for the company behind this ticker.
	Phone_number string `sql:"adjusted"`

	// The ISO code of the primary listing exchange for this asset.
	Primary_exchange string `sql:"adjusted"`

	// The share Class OpenFIGI number for this ticker.
	Share_class_figi string `sql:"adjusted"`

	// The recorded number of outstanding shares for this particular share class.
	Share_class_shares_outstanding int64 `sql:"adjusted"`

	// The standard industrial classification code for this ticker. For a list of SIC Codes, see the SEC's SIC Code List.
	Sic_code string `sql:"adjusted"`

	// A description of this ticker's SIC code.
	Sic_description string `sql:"adjusted"`

	// The exchange symbol that this item is traded under.
	Ticker string `sql:"adjusted"`

	// The approximate number of employees for the company.
	Total_employees int64 `sql:"adjusted"`

	// The type of the asset. Find the types that we support via our Ticker Types API.
	Type string `sql:"adjusted"`

	// The shares outstanding calculated assuming all shares of other share classes are converted to this share class.
	Weighted_shares_outstanding int64 `sql:"adjusted"`
}

type Address struct {

	// The first line of the company's headquarters address.
	Address1 string `sql:"adjusted"`

	// The city of the company's headquarters address.
	City string `sql:"adjusted"`

	// The state of the company's headquarters address.
	State string `sql:"adjusted"`
}

type Branding struct {

	// A link to this ticker's company's icon. Icon's are generally smaller, square images that represent the company at a glance.
	Icon_url string `sql:"adjusted"`

	// A link to this ticker's company's logo.
	Logo_url string `sql:"adjusted"`
}
