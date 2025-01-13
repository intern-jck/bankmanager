package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"
)

var pdfTxtPath = "data/testpdf.txt"
var txtPath = "data/test.txt"

// Regexes
var dateAtStartRegex = regexp.MustCompile(`^([0-9]{2}\/[0-9]{2})`)
var dateRegex = regexp.MustCompile(`([0-9]{2}\/[0-9]{2})`)
var amountRegex = regexp.MustCompile(`\$(\d{1,3}(?:,\d{3})*(?:\.\d{2})?)|\d{1,3}(?:,\d{3})*(?:\.\d{2})?$|-\d{1,3}(?:,\d{3})*(?:\.\d{2})?`)
var checkIdRegex = regexp.MustCompile(`^([0-9]{1,})`)

var labels = []string{
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
	Balance struct {
		Starting string
		Ending   string
	}
	Deposits    string
	Checks      string
	Withdrawals struct {
		Debit      string
		Electronic string
	}
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
	ID     int
	Date   string
	Amount string
}

type BankJson struct {
	CheckingSummary Summary
	Deposits        []Deposit
	Withdrawals     struct {
		Debit      []Withdrawal
		Electronic []Withdrawal
	}
	Checks []Check
}

func main() {

	// Create files
	// Json for statement data
	jsonFile, err := os.Create("test.json")
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
			case labels[2], labels[4], labels[5]:
				var entry []string

				match := dateAtStartRegex.FindStringSubmatch(line)
				trimmedLine := dateRegex.ReplaceAllString(line, "")
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}

				match = amountRegex.FindStringSubmatch(trimmedLine)
				trimmedLine = amountRegex.ReplaceAllString(line, "")
				if match != nil {
					entry = append(entry, match[0])
				} else {
					continue
				}

				entry = append(entry, trimmedLine)
				row := strings.Join(entry, ",")
				_, err = txt.WriteString(row + "\n")
				if err != nil {
					panic(err)
				}

			// checks
			case labels[3]:
				var entry []string

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
