package parser

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"
)

// Regexes
var dateAtStartRegex = regexp.MustCompile(`^([0-9]{2}\/[0-9]{2})`)
var amountRegex = regexp.MustCompile(`(\d{1,3}(?:,\d{3})*\.\d{2})$`)
var checkIdRegex = regexp.MustCompile(`^([0-9]{1,})`)
var dateRegex = regexp.MustCompile(`([0-9]{2}\/[0-9]{2})`)

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

func CreateJson(path string) {

	paths := strings.Split(path, "/")

	txtDir := paths[0]
	txtFile := paths[1]

	txtPath := "data/textpdf/" + txtDir + "/" + txtFile

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
