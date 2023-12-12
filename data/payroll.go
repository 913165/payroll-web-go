package data

type Payroll struct {
	PayrollID  int     `json:"payroll_id"`
	EmployeeID int     `json:"employee_id"`
	Salary     float64 `json:"salary"`
	PayDate    string  `json:"pay_date"`
	Taxes      float64 `json:"taxes"`
	Bonuses    float64 `json:"bonuses"`
	NetPay     float64 `json:"net_pay"`
}
