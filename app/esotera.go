package app

import (
	"EsoteraFinal/pkg/sites"
	"EsoteraFinal/pkg/utils"
	"fmt"
	"github.com/corpix/uarand"
	"net"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func Run() {

	Homechoice := utils.HomeMenu()
	if Homechoice == "1. END. Clothing" {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		EndChoice := utils.EndMenu()
		if EndChoice == "1. Entry" {
			Tasks := sites.CreateEndTasksFromCsv("/Users/kacey/go/src/EsoteraFinal/sites/END/END.csv")
			maxRetries := 4

			for tasknum, task := range Tasks {
				userAgent := uarand.GetRandom()
				fmt.Println(utils.EndNeutralPrefix(strconv.Itoa(tasknum+1)) + "Starting task...")
				jar, err := cookiejar.New(nil)
				if err != nil {
					fmt.Println(err)
					return
				}
				var Token string
				var LoginStatus string
				//client := http.Client{Jar: jar}
				dialer := &net.Dialer{
					Timeout: 5 * time.Second,
				}
				client := http.Client{
					Transport: &http.Transport{
						DialContext: dialer.DialContext,
					}, Jar: jar,
				}
				currentRetries := 0
				for currentRetries < maxRetries {
					Token, LoginStatus = sites.EndLogin(client, task.Email, task.Password, userAgent)
					if LoginStatus != "200 OK" && currentRetries < maxRetries {
						fmt.Println(utils.EndFailurePrefix(strconv.Itoa(tasknum+1))+"Failed logging in...", LoginStatus)
						currentRetries++
						time.Sleep(2 * time.Second)

						continue
						//HANDLE FAILED LOGIN WITH RETRIES + RECORD SUCCESSFUL ENTRIES IN JSON BASED ON ENTRY STATUS + MAKE CREATETASKSFROMCSV
						//REMOVE ALL TASKS THAT MATCH JSON

					} else {
						fmt.Println(utils.EndSuccessPrefix(strconv.Itoa(tasknum+1)) + "Successfully logged in!")
						time.Sleep(2 * time.Second)

						break
					}
				}
				time.Sleep(2 * time.Second)
				var CardID int
				var CardStatus string
				for currentRetries < maxRetries {
					CardID, CardStatus = sites.GetPaymentMethod(client, Token, userAgent)
					if CardStatus != "200 OK" && currentRetries < maxRetries {
						fmt.Println(utils.EndFailurePrefix(strconv.Itoa(tasknum+1))+"Info retrieval failure... [1/2]", CardStatus)
						currentRetries++
						time.Sleep(2 * time.Second)

						continue
						//HANDLE FAILED LOGIN WITH RETRIES + RECORD SUCCESSFUL ENTRIES IN JSON BASED ON ENTRY STATUS + MAKE CREATETASKSFROMCSV
						//REMOVE ALL TASKS THAT MATCH JSON

					} else {
						fmt.Println(utils.EndSuccessPrefix(strconv.Itoa(tasknum+1)) + "Successfully loaded info! [1/2]")
						time.Sleep(2 * time.Second)

						break
					}
				}
				time.Sleep(2 * time.Second)
				var EntryData sites.EntryData
				var GetEntryDataStatus string
				for currentRetries < maxRetries {
					EntryData, GetEntryDataStatus = sites.GetEntryData(client, Token, userAgent)
					if GetEntryDataStatus != "202 Accepted" && currentRetries < maxRetries {
						fmt.Println(utils.EndFailurePrefix(strconv.Itoa(tasknum+1))+"Info retrieval failure... [2/2]", GetEntryDataStatus)

						currentRetries++
						time.Sleep(2 * time.Second)

						continue
						//HANDLE FAILED LOGIN WITH RETRIES + RECORD SUCCESSFUL ENTRIES IN JSON BASED ON ENTRY STATUS + MAKE CREATETASKSFROMCSV
						//REMOVE ALL TASKS THAT MATCH JSON

					} else {
						fmt.Println(utils.EndSuccessPrefix(strconv.Itoa(tasknum+1)) + "Successfully loaded info! [2/2]")
						time.Sleep(3 * time.Second)

						break
					}
				}
				time.Sleep(2 * time.Second)
				var EntryStatus string
				for currentRetries < maxRetries {
					EntryStatus = sites.EnterDraw(client, Token, EntryData, CardID, userAgent)
					if string(EntryStatus) == "403 Forbidden" && currentRetries < maxRetries {

						fmt.Println(utils.EndSuccessPrefix(strconv.Itoa(tasknum+1)) + "Raffle already entered!")
						currentRetries++
						time.Sleep(2 * time.Second)

						break
						//HANDLE FAILED LOGIN WITH RETRIES + RECORD SUCCESSFUL ENTRIES IN JSON BASED ON ENTRY STATUS + MAKE CREATETASKSFROMCSV
						//REMOVE ALL TASKS THAT MATCH JSON

					} else if EntryStatus == "201 Created" && currentRetries < maxRetries {
						fmt.Println(utils.EndSuccessPrefix(strconv.Itoa(tasknum+1)) + "Successfully entered raffle!")
						utils.SendSuccessHookEnd(task.Email, string(EntryData.ProductSizeID))
						currentRetries++
						time.Sleep(2 * time.Second)

						break
					} else if currentRetries < maxRetries {

						fmt.Println(utils.EndFailurePrefix(strconv.Itoa(tasknum+1))+"Failed entering raffle...", EntryStatus)
						time.Sleep(2 * time.Second)
						currentRetries++

						continue
					}
				}

			}
			fmt.Println("All tasks completed!")
		} else if EndChoice == "2. Account Gen" {
			userAgent := uarand.GetRandom()
			Tasks := sites.CreateEndTasksFromCsv("/Users/kacey/go/src/EsoteraFinal/sites/END/END.csv")
			jar, err := cookiejar.New(nil)
			client := http.Client{Jar: jar}
			status, guestToken, err := sites.GetGuestID(client, userAgent)
			fmt.Println(status)
			if status == "201 Created" {
				time.Sleep(5 * time.Second)
				//guestToken, _ := sites.GetGuestID(client)
				//fmt.Println(guestToken)
				//time.Sleep(15 * time.Second)
				creationStatus, err := sites.CreateAccount(client, Tasks[0].Email, Tasks[0].FirstName, Tasks[0].LastName, Tasks[0].Password, guestToken, userAgent)
				if err != nil {
					fmt.Println(err)
				}
				if creationStatus == "200 OK" {
					fmt.Println("END account successfully generated.")
				}
			}
			if err != nil {
				fmt.Println(err)
			}
		} else if EndChoice == "3. Address Changer" {
		}

	} else {
		utils.HomeMenu()
	}
}
