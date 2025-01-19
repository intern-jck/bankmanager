package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// var pdfTxtPath = "data/testpdf.txt"
// var txtPath = "data/test.txt"

// Regexes
var dateAtStartRegex = regexp.MustCompile(`^([0-9]{2}\/[0-9]{2})`)
var amountRegex = regexp.MustCompile(`(\d{1,3}(?:,\d{3})*\.\d{2})$`)
var checkIdRegex = regexp.MustCompile(`^([0-9]{1,})`)
var dateRegex = regexp.MustCompile(`([0-9]{2}\/[0-9]{2})`)

// var amountRegex = regexp.MustCompile(`\$(\d{1,3}(?:,\d{3})*(?:\.\d{2})?)|\d{1,3}(?:,\d{3})*(?:\.\d{2})?$|-\d{1,3}(?:,\d{3})*(?:\.\d{2})?`)

var sections = []string{
	"after address message area1",
	"summary",
	"deposits and additions",
	"checks paid section3",
	"atm debit withdrawal",
	"electronic withdrawal",
	"post fees message",
	"dre portrait disclosure message area",
}

type Summary struct {
	Beginning  string
	Ending     string
	Deposits   string
	Checks     string
	Debit      string
	Electronic string
	Fees       string
}

type Deposit struct {
	Date        string
	Amount      string
	Description string
}

type Withdrawal struct {
	Date        string
	Amount      string
	Description string
}

type Check struct {
	ID     string
	Date   string
	Amount string
}

type BankJson struct {
	Summary     Summary
	Deposits    []Deposit
	Withdrawals []Withdrawal
	Checks      []Check
}

func CreateJson(path string) {

	paths := strings.Split(path, "/")

	txtDir := paths[0]
	txtFile := paths[1]

	fmt.Println(txtDir, txtFile)

	txtPath := "data/textpdf/" + txtDir + "/" + txtFile
	fmt.Println("Converting: ", txtPath)

	// PDF txt file to parse
	txtPdf, err := os.Open(txtPath)
	if err != nil {
		panic(err)
	}
	defer txtPdf.Close()

	// Tokens to look for in each line to find sections
	startTk := "*start*"
	endTk := "*end*"
	section := ""

	// Flag to start parsing text
	readTxt := false

	jsonData := BankJson{}

	// Parse the txt file
	scanner := bufio.NewScanner(txtPdf)
	for scanner.Scan() {

		// Get each line
		line := scanner.Text()

		// Clean each line
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Check if line contains start or end token
		if strings.Contains(line, startTk) {
			section = line[7:]
			readTxt = true
			continue
		} else if strings.Contains(line, endTk) {
			readTxt = false
			continue
		}

		// Parse text

		// var sections = []string{
		// 	"after address message area1",
		// 	"summary",
		// 	"deposits and additions",
		// 	"checks paid section3",
		// 	"atm debit withdrawal",
		// 	"electronic withdrawal",
		// 	"post fees message",
		// 	"dre portrait disclosure message area",
		// }

		if readTxt {
			switch section {

			// "summary"
			case sections[1]:

				match := amountRegex.FindStringSubmatch(line)
				amountStr := ""

				if match != nil {
					amountStr = regexp.MustCompile(`[\$-]`).ReplaceAllString(match[0], "")
					amountStr = strings.Replace(amountStr, ",", "", -1)
				}

				description := amountRegex.ReplaceAllString(line, "")
				description = strings.TrimSpace(description)

				switch description {
				case "Beginning Balance":
					jsonData.Summary.Beginning = amountStr
				case "Ending":
					jsonData.Summary.Ending = amountStr
				case "Deposits and Additions":
					jsonData.Summary.Deposits = amountStr
				case "Checks Paid":
					jsonData.Summary.Checks = amountStr
				case "ATM & Debit Card Withdrawals":
					jsonData.Summary.Debit = amountStr
				case "Electronic Withdrawals":
					jsonData.Summary.Electronic = amountStr
				}

			// "deposits and additions"
			case sections[2]:

				var deposit Deposit

				amountStr := ""

				match := dateAtStartRegex.FindStringSubmatch(line)
				line = dateAtStartRegex.ReplaceAllString(line, "")
				if match != nil {
					deposit.Date = match[0]
				} else {
					continue
				}

				match = amountRegex.FindStringSubmatch(line)
				line = amountRegex.ReplaceAllString(line, "")
				if match != nil {
					amountStr = regexp.MustCompile(`[\$-]`).ReplaceAllString(match[0], "")
					amountStr = strings.Replace(amountStr, ",", "", -1)
					deposit.Amount = amountStr
				} else {
					continue
				}

				line = strings.TrimSpace(line)
				deposit.Description = line

				jsonData.Deposits = append(jsonData.Deposits, deposit)

			// "checks paid section3"
			case sections[3]:

				var check Check
				amountStr := ""

				match := checkIdRegex.FindStringSubmatch(line)
				if match != nil {
					check.ID = match[0]
				} else {
					continue
				}

				match = dateRegex.FindStringSubmatch(line)
				if match != nil {
					check.Date = match[0]
				} else {
					continue
				}

				match = amountRegex.FindStringSubmatch(line)
				if match != nil {
					amountStr = regexp.MustCompile(`[\$-]`).ReplaceAllString(match[0], "")
					amountStr = strings.Replace(amountStr, ",", "", -1)
					check.Amount = amountStr
				} else {
					continue
				}

				fmt.Println(check)
				jsonData.Checks = append(jsonData.Checks, check)

			// "atm debit withdrawal"
			// "electronic withdrawal"
			case sections[4], sections[5]:

				var withdrawal Withdrawal

				amountStr := ""

				match := dateAtStartRegex.FindStringSubmatch(line)
				line = dateAtStartRegex.ReplaceAllString(line, "")
				if match != nil {
					withdrawal.Date = match[0]
				} else {
					continue
				}

				match = amountRegex.FindStringSubmatch(line)
				line = amountRegex.ReplaceAllString(line, "")
				if match != nil {
					amountStr = regexp.MustCompile(`[\$-]`).ReplaceAllString(match[0], "")
					amountStr = strings.Replace(amountStr, ",", "", -1)
					withdrawal.Amount = amountStr
				} else {
					continue
				}

				line = strings.TrimSpace(line)
				withdrawal.Description = line

				jsonData.Withdrawals = append(jsonData.Withdrawals, withdrawal)

			// "post fees message"
			case sections[6]:
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Json for statement data
	jsonPath := "data/json/" + txtDir + "/" + txtFile
	jsonPath = strings.Replace(jsonPath, ".txt", ".json", -1)
	fmt.Println(jsonPath)

	jsonFile, err := os.Create(jsonPath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	// Save the json data
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	encoder.Encode(jsonData)

}

/*
	// Save the json data
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	encoder.Encode(jsonData)

	// label := ""
	// jsonData := BankJson{}
	// // Create files

	// // Json for statement data
	// jsonFile, err := os.Create("data/json/2018/20180104.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer jsonFile.Close()

	// 	if readTxt {

	// 		switch label {

	// 		// summary
	// 		case labels[1]:
	// 			var entry []string

	// 			match := amountRegex.FindStringSubmatch(line)
	// 			amount := ""
	// 			if match != nil {
	// 				amount = regexp.MustCompile(`[\$-]`).ReplaceAllString(match[0], "")
	// 				entry = append(entry, amount)
	// 			}

	// 			description := amountRegex.ReplaceAllString(line, "")
	// 			description = strings.TrimSpace(description)

	// 			entry = append(entry, description)
	// 			row := strings.Join(entry, ",")

	// 			_, err = txt.WriteString(row + "\n")
	// 			if err != nil {
	// 				panic(err)
	// 			}

	// 			switch description {
	// 			case "Beginning Balance":
	// 				jsonData.CheckingSummary.Balance.Starting = entry[0]
	// 			case "Ending Balance":
	// 				jsonData.CheckingSummary.Balance.Ending = entry[0]
	// 			case "Deposits and Additions":
	// 				jsonData.CheckingSummary.Deposits = entry[0]
	// 			case "Checks Paid":
	// 				jsonData.CheckingSummary.Checks = entry[0]
	// 			case "ATM & Debit Card Withdrawals":
	// 				jsonData.CheckingSummary.Withdrawals.Debit = entry[0]
	// 			case "Electronic Withdrawals":
	// 				jsonData.CheckingSummary.Withdrawals.Electronic = entry[0]
	// 			}

	// 		// deposits
	// 		case labels[2]:
	// 			var entry []string
	// 			var deposit Deposit

	// 			match := dateAtStartRegex.FindStringSubmatch(line)
	// 			if match != nil {
	// 				entry = append(entry, match[0])
	// 			} else {
	// 				continue
	// 			}

	// 			trimmedLine := dateAtStartRegex.ReplaceAllString(line, "")

	// 			match = amountRegex.FindStringSubmatch(trimmedLine)
	// 			description := amountRegex.ReplaceAllString(line, "")
	// 			description = strings.TrimSpace(description)
	// 			if match != nil {
	// 				entry = append(entry, match[0])
	// 			} else {
	// 				continue
	// 			}

	// 			entry = append(entry, description)
	// 			row := strings.Join(entry, ",")
	// 			_, err = txt.WriteString(row + "\n")
	// 			if err != nil {
	// 				panic(err)
	// 			}

	// 			deposit.Date = entry[0]
	// 			deposit.Amount = entry[1]
	// 			deposit.Description = entry[2]

	// 			jsonData.Deposits = append(jsonData.Deposits, deposit)

	// 		// Withdrawals
	// 		case labels[4], labels[5]:
	// 			var entry []string
	// 			var withdrawal Withdrawal

	// 			match := dateAtStartRegex.FindStringSubmatch(line)
	// 			if match != nil {
	// 				entry = append(entry, match[0])
	// 			} else {
	// 				continue
	// 			}
	// 			trimmedLine := dateAtStartRegex.ReplaceAllString(line, "")

	// 			match = amountRegex.FindStringSubmatch(trimmedLine)
	// 			description := amountRegex.ReplaceAllString(line, "")
	// 			description = strings.TrimSpace(description)
	// 			if match != nil {
	// 				entry = append(entry, match[0])
	// 			} else {
	// 				continue
	// 			}

	// 			entry = append(entry, description)
	// 			row := strings.Join(entry, ",")
	// 			_, err = txt.WriteString(row + "\n")
	// 			if err != nil {
	// 				panic(err)
	// 			}

	// 			withdrawal.Date = entry[0]
	// 			withdrawal.Amount = entry[1]
	// 			withdrawal.Description = entry[2]

	// 			if label == labels[4] {
	// 				jsonData.Withdrawals.Debit = append(jsonData.Withdrawals.Debit, withdrawal)
	// 			} else if label == labels[5] {
	// 				jsonData.Withdrawals.Electronic = append(jsonData.Withdrawals.Electronic, withdrawal)
	// 			}

	// 		// checks
	// 		case labels[3]:
	// 			var entry []string
	// 			var check Check

	// 			match := checkIdRegex.FindStringSubmatch(line)
	// 			if match != nil {
	// 				entry = append(entry, match[0])
	// 			} else {
	// 				continue
	// 			}

	// 			match = dateRegex.FindStringSubmatch(line)
	// 			if match != nil {
	// 				entry = append(entry, match[0])
	// 			} else {
	// 				continue
	// 			}

	// 			match = amountRegex.FindStringSubmatch(line)
	// 			if match != nil {
	// 				entry = append(entry, match[0])
	// 			} else {
	// 				continue
	// 			}

	// 			row := strings.Join(entry, ",")
	// 			_, err = txt.WriteString(row + "\n")
	// 			if err != nil {
	// 				panic(err)
	// 			}

	// 			check.ID = entry[0]
	// 			check.Date = entry[1]
	// 			check.Amount = entry[2]

	// 			jsonData.Checks = append(jsonData.Checks, check)

	// 		default:
	// 			continue
	// 		}

	// 	}
func ParsePdf() {

	// Create files
	// Json for statement data
	jsonFile, err := os.Create("json/2018/20180104.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	// PDF to parse
	file, err := os.Open(pdfTxtPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Text file to save cleaned PDF
	txt, err := os.Create(txtPath)
	if err != nil {
		panic(err)
	}
	defer txt.Close()

	// To parse PDF txt file
	scanner := bufio.NewScanner(file)
	readTxt := false
	startTk := "*start*"
	endTk := "*end*"
	label := ""
	jsonData := BankJson{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Contains(line, startTk) {
			label = line[7:]
			_, err = txt.WriteString("<<" + strings.ToUpper(label) + ">>\n")
			if err != nil {
				panic(err)
			}

			readTxt = true
			continue
		} else if strings.Contains(line, endTk) {
			_, err = txt.WriteString("<<" + strings.ToUpper(label) + ">>\n")
			if err != nil {
				panic(err)
			}

			readTxt = false
			continue
		}

		if readTxt {

			switch label {

			// summary
			case labels[1]:
				var entry []string

				match := amountRegex.FindStringSubmatch(line)
				amount := ""
				if match != nil {
					amount = regexp.MustCompile(`[\$-]`).ReplaceAllString(match[0], "")
					entry = append(entry, amount)
				}

				description := amountRegex.ReplaceAllString(line, "")
				description = strings.TrimSpace(description)

				entry = append(entry, description)
				row := strings.Join(entry, ",")

				_, err = txt.WriteString(row + "\n")
				if err != nil {
					panic(err)
				}

				switch description {
				case "Beginning Balance":
					jsonData.CheckingSummary.Balance.Starting = entry[0]
				case "Ending Balance":
					jsonData.CheckingSummary.Balance.Ending = entry[0]
				case "Deposits and Additions":
					jsonData.CheckingSummary.Deposits = entry[0]
				case "Checks Paid":
					jsonData.CheckingSummary.Checks = entry[0]
				case "ATM & Debit Card Withdrawals":
					jsonData.CheckingSummary.Withdrawals.Debit = entry[0]
				case "Electronic Withdrawals":
					jsonData.CheckingSummary.Withdrawals.Electronic = entry[0]
				}

			// deposits
			case labels[2]:
				var entry []string
				var deposit Deposit

				match := dateAtStartRegex.FindStringSubmatch(line)
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}

				trimmedLine := dateAtStartRegex.ReplaceAllString(line, "")

				match = amountRegex.FindStringSubmatch(trimmedLine)
				description := amountRegex.ReplaceAllString(line, "")
				description = strings.TrimSpace(description)
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}

				entry = append(entry, description)
				row := strings.Join(entry, ",")
				_, err = txt.WriteString(row + "\n")
				if err != nil {
					panic(err)
				}

				deposit.Date = entry[0]
				deposit.Amount = entry[1]
				deposit.Description = entry[2]

				jsonData.Deposits = append(jsonData.Deposits, deposit)

			// Withdrawals
			case labels[4], labels[5]:
				var entry []string
				var withdrawal Withdrawal

				match := dateAtStartRegex.FindStringSubmatch(line)
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}
				trimmedLine := dateAtStartRegex.ReplaceAllString(line, "")

				match = amountRegex.FindStringSubmatch(trimmedLine)
				description := amountRegex.ReplaceAllString(line, "")
				description = strings.TrimSpace(description)
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}

				entry = append(entry, description)
				row := strings.Join(entry, ",")
				_, err = txt.WriteString(row + "\n")
				if err != nil {
					panic(err)
				}

				withdrawal.Date = entry[0]
				withdrawal.Amount = entry[1]
				withdrawal.Description = entry[2]

				if label == labels[4] {
					jsonData.Withdrawals.Debit = append(jsonData.Withdrawals.Debit, withdrawal)
				} else if label == labels[5] {
					jsonData.Withdrawals.Electronic = append(jsonData.Withdrawals.Electronic, withdrawal)
				}

			// checks
			case labels[3]:
				var entry []string
				var check Check

				match := checkIdRegex.FindStringSubmatch(line)
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}

				match = dateRegex.FindStringSubmatch(line)
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}

				match = amountRegex.FindStringSubmatch(line)
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}

				row := strings.Join(entry, ",")
				_, err = txt.WriteString(row + "\n")
				if err != nil {
					panic(err)
				}

				check.ID = entry[0]
				check.Date = entry[1]
				check.Amount = entry[2]

				jsonData.Checks = append(jsonData.Checks, check)

			default:
				continue
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Save the json data
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	encoder.Encode(jsonData)

}


*/
