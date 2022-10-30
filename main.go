package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"service-get-applicants/config"
	"strconv"
	"strings"
	"time"

	"github.com/dimchansky/utfbom"
	pq "github.com/pasiol/gopq"
	"github.com/pasiol/serviceLog"
)

var (
	// Version for build
	Version string
	// Build for build
	Build                   string
	jobName                 = "service-get-applicants"
	debugState              = false
	archieveApplicantConfig = ""
	insertApplicantConfig   = ""
)

func readCSVFromString(data string) ([][]string, error) {

	r := csv.NewReader(utfbom.SkipOnly(strings.NewReader(data)))
	r.Comma = ';'
	lines, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func getApplicants() ([][]string, error) {
	query := config.Applicants()
	c := config.GetPrimusConfig()
	query.Host = c.Host
	query.Port = c.Port
	query.User = c.User
	query.Pass = c.Password
	output, _ := pq.ExecuteAndRead(query, 30) // TODO: err
	if output == "" {
		return [][]string{}, nil
	}
	data, err := readCSVFromString(output)
	if err != nil {
		return [][]string{}, err
	}
	return data[1:], nil
}

func getAccounts(accounts map[string]int, service serviceLog.Service) error {
	query := config.Accounts()
	c := config.GetPrimusConfig()
	query.Host = c.Host
	query.Port = c.Port
	query.User = c.User
	query.Pass = c.Password

	output, err := pq.ExecuteAndRead(query, 30)
	if err != nil {
		return err
	}
	if output == "" {
		return errors.New("no account data")
	}
	rows := strings.Fields(output)
	for _, row := range rows[1:] {
		fields := strings.Split(row, ";")
		id, err := strconv.Atoi(fields[0])
		if err != nil {
			return err
		}
		accounts[fields[1]] = id
	}
	return nil
}

func archiveApplicant(id string) (int, error) {
	archieveFile, err := config.ArchieveXML(id)
	if err != nil {
		return 0, err
	}
	c := config.GetPrimusConfig()
	cardID, errorCount, err := pq.ExecuteAtomicImportQuery(archieveFile, c.Host, c.Port, c.User, c.Password, archieveApplicantConfig)
	if err == nil && errorCount == 0 {
		return cardID, nil
	}
	return cardID, errors.New("archiving applicant failed")
}

func insertApplicant(row []string) (string, error) {

	applicantFile, err := config.ApplicantXML(row)
	if err != nil {
		return "", err
	}
	c := config.GetPrimusConfig()
	cardID, errorCount, err := pq.ExecuteAtomicImportQuery(applicantFile, c.Host, c.Port, c.User, c.Password, insertApplicantConfig)
	cardIDstr := strconv.Itoa(cardID)
	if err == nil && errorCount == 0 {
		_, err := archiveApplicant(row[0])
		return cardIDstr, err

	}
	return cardIDstr, errors.New("importing the applicant data to student-registry failed")
}

func userNameAlreadyExists(username string) bool {
	query := config.NewTypeUserAccount(username)
	c := config.GetPrimusConfig()
	query.Host = c.Host
	query.Port = c.Port
	query.User = c.User
	query.Pass = c.Password
	output, _ := pq.ExecuteAndRead(query, 30)

	return strings.Contains(output, username)
}

func alreadyMoved(id string) bool {
	query := config.StudentRegistryApplicantID(id)
	c := config.GetPrimusConfig()
	query.Host = c.Host
	query.Port = c.Port
	query.User = c.User
	query.Pass = c.Password

	output, _ := pq.ExecuteAndRead(query, 30)
	return strings.Contains(output, id)
}

func createNewUserAccount(user []string) (int, error) {
	userAccountFile, err := config.UserAccountXML(user)
	if err != nil {
		return 0, err
	}
	c := config.GetPrimusConfig()
	cardID, _, err := pq.ExecuteAtomicImportQuery(userAccountFile, c.Host, c.Port, c.User, c.Password, "spaikanvastaanotto_uudentyyppinen_wilmatunnus")
	if err == nil && cardID > 0 {
		if cardID == 0 {
			return 0, errors.New("inserting the new user to Wilma user registry failed")
		}
		return cardID, nil
	}
	return 0, errors.New("inserting the new user to Wilma user registry failed")
}

func main() {
	var option string
	if len(os.Args) == 2 {
		if os.Args[1] == "--help" {
			fmt.Printf("Usage: %s [--allowDuplicates]\n", jobName)
			os.Exit(0)
		} else {
			option = string(os.Args[1])
		}
	}

	start := time.Now()
	self, err := serviceLog.SetService(jobName, Version, Build, false)
	if err != nil {
		log.Fatalf("Initializing service failed: %s", err)
	}
	log.Print(serviceLog.GetLogMessage(self, serviceLog.Event{ShortMessage: "Service started.", FullMessage: "", Succesful: true, Severity: "info"}))
	pq.Debug = debugState

	var accounts = map[string]int{}
	applicants, _ := getApplicants() // TODO: err

	if err == nil {
		err = getAccounts(accounts, self)
		if err == nil {
			for _, row := range applicants {
				applicantID := row[0]
				applicantEmail := row[1]
				id := row[0]
				_, founded := accounts[applicantEmail]
				if founded {
					if !alreadyMoved(applicantID) {
						if !userNameAlreadyExists(applicantEmail) || option == "--allowDuplicates" {
							cardID, err := insertApplicant(row)
							if err == nil {
								log.Printf("Inserting new applicant: %s in the student registry: %s", id, cardID)
							} else {
								log.Printf("Inserting new applicant: %s in the student registry failed: %s", id, cardID)
							}
						} else {
							log.Print("Username already exists on the student registry.")
						}
					} else {
						log.Print("Application already exists on the student registry.")
					}
				} else {
					newID, err := createNewUserAccount(row)
					if err == nil {
						log.Print("Inserting the Wilma User account succesfully.")
						accounts[applicantEmail] = newID
						cardID, err := insertApplicant(row)
						if err == nil {
							log.Printf("Inserting new applicant in the student registry: %s", cardID)
						} else {
							log.Printf("Inserting new applicant in the student registry failed: %s", cardID)
						}
					} else {
						log.Print("Inserting the Wilma User account failed.")

					}
				}
			}
		} else {
			log.Print("Getting Wilma accounts failed.")
		}
	} else {
		log.Print("Getting applicants failed.")
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Print("Ending service in a controlled manner.")
	log.Printf("Elapsed processing time %d.", elapsed)

}
