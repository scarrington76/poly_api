package model

type StockFinancials struct {
	Cik         string `json:"cik"`
	CompanyName string `json:"company_name"`
	EndDate     string `json:"end_date"`
	Financials  struct {
		BalanceSheet struct {
			Assets struct {
				Label string `json:"label"`
				Order int    `json:"order"`
				Unit  string `json:"unit"`
				Value int64  `json:"value"`
				Xpath string `json:"xpath"`
			} `json:"assets"`
			Equity struct {
				Formula string `json:"formula"`
				Label   string `json:"label"`
				Order   int    `json:"order"`
				Unit    string `json:"unit"`
				Value   int64  `json:"value"`
			} `json:"equity"`
			Liabilities struct {
				Label string `json:"label"`
				Order int    `json:"order"`
				Unit  string `json:"unit"`
				Value int64  `json:"value"`
				Xpath string `json:"xpath"`
			} `json:"liabilities"`
		} `json:"balance_sheet"`
		CashFlowStatement struct {
			ExchangeGainsLosses struct {
				Label string `json:"label"`
				Order int    `json:"order"`
				Unit  string `json:"unit"`
				Value int    `json:"value"`
				Xpath string `json:"xpath"`
			} `json:"exchange_gains_losses"`
			NetCashFlow struct {
				Formula string `json:"formula"`
				Label   string `json:"label"`
				Order   int    `json:"order"`
				Unit    string `json:"unit"`
				Value   int64  `json:"value"`
			} `json:"net_cash_flow"`
			NetCashFlowFromFinancingActivities struct {
				Label string `json:"label"`
				Order int    `json:"order"`
				Unit  string `json:"unit"`
				Value int64  `json:"value"`
				Xpath string `json:"xpath"`
			} `json:"net_cash_flow_from_financing_activities"`
		} `json:"cash_flow_statement"`
		ComprehensiveIncome struct {
			ComprehensiveIncomeLoss struct {
				Formula string `json:"formula"`
				Label   string `json:"label"`
				Order   int    `json:"order"`
				Unit    string `json:"unit"`
				Value   int64  `json:"value"`
			} `json:"comprehensive_income_loss"`
			ComprehensiveIncomeLossAttributableToParent struct {
				Label string `json:"label"`
				Order int    `json:"order"`
				Unit  string `json:"unit"`
				Value int64  `json:"value"`
				Xpath string `json:"xpath"`
			} `json:"comprehensive_income_loss_attributable_to_parent"`
			OtherComprehensiveIncomeLoss struct {
				Formula string `json:"formula"`
				Label   string `json:"label"`
				Order   int    `json:"order"`
				Unit    string `json:"unit"`
				Value   int64  `json:"value"`
			} `json:"other_comprehensive_income_loss"`
		} `json:"comprehensive_income"`
		IncomeStatement struct {
			BasicEarningsPerShare struct {
				Label string  `json:"label"`
				Order int     `json:"order"`
				Unit  string  `json:"unit"`
				Value float64 `json:"value"`
				Xpath string  `json:"xpath"`
			} `json:"basic_earnings_per_share"`
			CostOfRevenue struct {
				Label string `json:"label"`
				Order int    `json:"order"`
				Unit  string `json:"unit"`
				Value int64  `json:"value"`
				Xpath string `json:"xpath"`
			} `json:"cost_of_revenue"`
			GrossProfit struct {
				Label string `json:"label"`
				Order int    `json:"order"`
				Unit  string `json:"unit"`
				Value int64  `json:"value"`
				Xpath string `json:"xpath"`
			} `json:"gross_profit"`
			OperatingExpenses struct {
				Formula string `json:"formula"`
				Label   string `json:"label"`
				Order   int    `json:"order"`
				Unit    string `json:"unit"`
				Value   int64  `json:"value"`
			} `json:"operating_expenses"`
			Revenues struct {
				Label string `json:"label"`
				Order int    `json:"order"`
				Unit  string `json:"unit"`
				Value int64  `json:"value"`
				Xpath string `json:"xpath"`
			} `json:"revenues"`
		} `json:"income_statement"`
	} `json:"financials"`
	FiscalPeriod        string `json:"fiscal_period"`
	FiscalYear          string `json:"fiscal_year"`
	SourceFilingFileURL string `json:"source_filing_file_url"`
	SourceFilingURL     string `json:"source_filing_url"`
	StartDate           string `json:"start_date"`
}
