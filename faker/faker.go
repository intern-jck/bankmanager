package faker

import (
	"bankmanager/types"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func Test() {
	fmt.Println("faker go")
}

var cardPurchase = "Card Purchase"
var atmWithdrawal = "ATM Withdrawal"
var nonAtmWithdraw = "Non-Chase ATM Withdraw"
var recurringCardPurchase = "Recurring Card Purchase"
var billPayment = "Bill Payment"
var cardLastFour = "1234"
var cardNumber = "1234 1234 1234 1234"

var months = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// var years = []int{2018, 2019, 2020, 2021, 2022, 2023, 2024}
var years = []int{2018}

type Month struct {
	Name string
	Num  int
	Days int
	Leap bool
}

type Year struct {
	Count  int
	Leap   bool
	Months []Month
}

type Date struct {
	Year  Year
	Month Month
}

type Statement struct {
	Date Date
	Data types.BankJson
}

func CreateStatement(year int, month int) {
	statement := types.BankJson{}

	beginning := 5000.00
	// ending := beginning
	deposits := 0.00
	withdrawals := 0.00

	// create statement for first month
	for day := 1; day <= monthDays[month]; day++ {
		date := fmt.Sprintf("%s/%s", formatIntToString(months[month]), formatIntToString(day))

		w := 0.0
		d := 0.0

		addDeposit := false
		addWithdrawal := false

		description := ""
		amount := ""

		// Add a paycheck
		if day == 20 {
			d = 2500.00
			description = "Job LLC"
			addDeposit = true
		}

		// Add other deposit activity like Venmo
		if day == 22 {
			d = 10.0 + rand.Float64()*(20.0)
			description = "Venmo\tCashout\tPPD ID: 1234567890"
			addDeposit = true
		}

		// Create some standard purchases
		// Add rent
		if day == 1 || day == 15 {
			w = 1500.00
			description = recurringCardPurchase + "\t" + date + " Landlord LLC Card " + cardLastFour
			addWithdrawal = true
		}

		// Cellphone
		if day == 14 {
			w = 100.00
			description = cardPurchase + "\t" + date + " Cellphone Company USA Card " + cardLastFour
			addWithdrawal = true
		}

		// Spotify
		if day == 23 {
			w = 9.99
			description = recurringCardPurchase + "\t" + date + " Spotify USA Card " + cardLastFour
			addWithdrawal = true
		}

		// Netflix
		if day == 14 {
			w = 7.99
			description = cardPurchase + "\t" + date + " Netflix.Com Netflix.Com CA Card " + cardLastFour
			addWithdrawal = true
		}

		// Restaurants
		if day%12 == 0 {
			w = 20.0 + rand.Float64()*(30.0)
			description = cardPurchase + "\t" + date + " Dive Bar Restaurant LLC Card " + cardLastFour
			addWithdrawal = true
		}

		if day%30 == 0 {
			w = 60.0 + rand.Float64()*(20.0)
			description = cardPurchase + "\t" + date + " Gourmet Restaurant LLC Card " + cardLastFour
			addWithdrawal = true
		}

		// Convert amount to string and create purchase

		if addDeposit {
			amount = fmt.Sprintf("%.2f", d)
			deposit := types.Deposit{
				Date:        date,
				Amount:      amount,
				Description: description,
			}
			statement.Deposits = append(statement.Deposits, deposit)
			addDeposit = false
		}

		if addWithdrawal {
			amount = fmt.Sprintf("%.2f", w)
			withdrawal := types.Withdrawal{
				Date:        date,
				Amount:      amount,
				Description: description,
			}
			statement.Withdrawals = append(statement.Withdrawals, withdrawal)
			addWithdrawal = false
		}

		// Update statement
		deposits = deposits + d
		withdrawals = withdrawals + w
	}

	statement.Summary.Beginning = fmt.Sprintf("%.2f", beginning)
	statement.Summary.Ending = fmt.Sprintf("%.2f", beginning-withdrawals+deposits)
	statement.Summary.Deposits = fmt.Sprintf("%.2f", deposits)
	statement.Summary.Withdrawals = fmt.Sprintf("%.2f", withdrawals)

	// fmt.Println(beginning, ending, deposits, withdrawals)
	fmt.Println(statement)

	// Now create json of statement
	jsonPath := fmt.Sprintf("./faker/fakedata/%d%s%s.json", year, formatIntToString(month), formatIntToString(1))
	jsonFile, err := os.Create(jsonPath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	// Save the json data
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	encoder.Encode(statement)
}

func createDeposit(date string, amount string, desc string) (types.Deposit, error) {
	d := types.Deposit{}
	d.Date = date
	d.Amount = amount
	d.Description = desc
	return d, nil
}

func formatIntToString(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}
	return strconv.Itoa(i)
}
