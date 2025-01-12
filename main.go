package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
LABELS
after address message area1
summary
deposits and additions
checks paid section3
atm debit withdrawal
atm debit withdrawal
electronic withdrawal
post fees message
post fees message
dre portrait disclosure message area
*/

/*
regex patterns

amount at end of line
[0-9]{1,3},?[0-9]{1,3},?[0-9]{1,3}.[0-9]{2})$
from chatgpt
(?<=\b|\s)\$(\d{1,3}(?:,\d{3})*(?:\.\d{2})?)|\d{1,3}(?:,\d{3})*(?:\.\d{2})?$

date at start of line
^([0-9]{2}\/[0-9]{2})
*/

var pdfTxtPath = "data/testpdf.txt"
var txtPath = "data/test.txt"

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

func main() {

	file, err := os.Open(pdfTxtPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	txt, err := os.Create(txtPath)
	if err != nil {
		panic(err)
	}
	defer txt.Close()

	scanner := bufio.NewScanner(file)
	readTxt := false
	startTk := "*start*"
	endTk := "*end*"
	label := ""

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
				_, err = txt.WriteString(line + "\n")
				if err != nil {
					panic(err)
				}
			// deposits
			case labels[2], labels[4], labels[5]:
				// case labels[4]:
				dateRegex := regexp.MustCompile(`^([0-9]{2}\/[0-9]{2})`)
				amountRegex := regexp.MustCompile(`\$(\d{1,3}(?:,\d{3})*(?:\.\d{2})?)|\d{1,3}(?:,\d{3})*(?:\.\d{2})?$`)

				var entry []string

				match := dateRegex.FindStringSubmatch(line)
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

				checkIdRegex := regexp.MustCompile(`^([0-9]{1,})`)
				dateRegex := regexp.MustCompile(`([0-9]{2}\/[0-9]{2})`)
				amountRegex := regexp.MustCompile(`\$(\d{1,3}(?:,\d{3})*(?:\.\d{2})?)|\d{1,3}(?:,\d{3})*(?:\.\d{2})?$`)

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
			// // withdrawals
			// case labels[4]:
			// 	_, err = txt.WriteString(line + "\n")
			// 	if err != nil {
			// 		panic(err)
			// 	}

			default:
				continue
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
