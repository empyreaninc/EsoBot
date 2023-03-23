package sites

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"
	"time"
)

type ShapeHeaders struct {
	Shapeheadera  string `json:"shapeheadera"`
	Shapeheadera0 string `json:"shapeheadera0"`
	Shapeheaderb  string `json:"shapeheaderb"`
	Shapeheaderc  string `json:"shapeheaderc"`
	Shapeheaderd  string `json:"shapeheaderd"`
	Shapeheaderf  string `json:"shapeheaderf"`
	Shapeheaderz  string `json:"shapeheaderz"`
}

type Enabled struct {
	enabled bool `json:"enabled"`
}

type GuestID struct {
	Body struct {
		GuestID string `json:"guest_id"`
	} `json:"body"`
	Message struct {
		Body        string `json:"body"`
		Translation struct {
			ID        interface{}   `json:"id"`
			Variables []interface{} `json:"variables"`
		} `json:"translation"`
	} `json:"message"`
}

type EmailAvailabilityResponse struct {
	Body struct {
		Available bool `json:"available"`
	} `json:"body"`
	Message struct {
		Body        string `json:"body"`
		Translation struct {
			ID        interface{}   `json:"id"`
			Variables []interface{} `json:"variables"`
		} `json:"translation"`
	} `json:"message"`
}
type BearerFromLogin struct {
	Body struct {
		Token      string `json:"token"`
		CustomerID string `json:"customer_id"`
	} `json:"body"`
	Message struct {
		Body        string `json:"body"`
		Translation struct {
			ID        interface{}   `json:"id"`
			Variables []interface{} `json:"variables"`
		} `json:"translation"`
	} `json:"message"`
}

//type EntryResponse struct {
//	CustomerID         int    `json:"customer_id"`
//	ProductSizeID      int    `json:"product_size_id"`
//	ShippingAddressID  int    `json:"shipping_address_id"`
//	BillingAddressID   int    `json:"billing_address_id"`
//	PaymentMethodID    int    `json:"payment_method_id"`
//	ShippingMethodID   int    `json:"shipping_method_id"`
//	WebsiteID          int    `json:"website_id"`
//	Postcode           string `json:"postcode"`
//	Street             string `json:"street"`
//	SubscriptionOrigin string `json:"subscription_origin"`
//}

type AccountData struct {
	Body struct {
		ID                  int    `json:"id"`
		Email               string `json:"email"`
		FirstName           string `json:"first_name"`
		LastName            string `json:"last_name"`
		StoreID             int    `json:"store_id"`
		WebsiteID           int    `json:"website_id"`
		CustomerID          string `json:"customer_id"`
		ExtensionAttributes struct {
			StorecreditBalance int `json:"storecredit_balance"`
		} `json:"extension_attributes"`
		Addresses []struct {
			ID                  int      `json:"id"`
			CountryID           string   `json:"country_id"`
			FirstName           string   `json:"first_name"`
			LastName            string   `json:"last_name"`
			Street              []string `json:"street"`
			City                string   `json:"city"`
			Postcode            string   `json:"postcode"`
			Telephone           string   `json:"telephone"`
			DefaultBilling      bool     `json:"default_billing"`
			DefaultShipping     bool     `json:"default_shipping"`
			ExtensionAttributes struct {
			} `json:"extension_attributes"`
			CustomAttributes []interface{} `json:"custom_attributes"`
			Region           struct {
				RegionCode interface{} `json:"region_code"`
				Region     string      `json:"region"`
				RegionID   int         `json:"region_id"`
			} `json:"region"`
		} `json:"addresses"`
	} `json:"body"`
	Message struct {
		Body        string `json:"body"`
		Translation struct {
			ID        interface{}   `json:"id"`
			Variables []interface{} `json:"variables"`
		} `json:"translation"`
	} `json:"message"`
}

type PaymentVaultData struct {
	EntityID          int    `json:"entity_id"`
	PublicHash        string `json:"public_hash"`
	VaultMethodCode   string `json:"vault_method_code"`
	PaymentMethodCode string `json:"payment_method_code"`
	Type              string `json:"type"`
	TypeDetail        string `json:"type_detail"`
	ExpiresAt         string `json:"expires_at"`
	GatewayToken      string `json:"gateway_token"`
	IsActive          bool   `json:"is_active"`
	IsDefault         bool   `json:"is_default"`
	IsVisible         bool   `json:"is_visible"`
	IsNew             bool   `json:"is_new"`
	UniqueIdentifier  string `json:"unique_identifier"`
	Label             string `json:"label"`
}

type EntryData struct {
	CustomerID         int    `json:"customer_id"`
	ProductSizeID      int    `json:"product_size_id"`
	ShippingAddressID  int    `json:"shipping_address_id"`
	BillingAddressID   int    `json:"billing_address_id"`
	PaymentMethodID    int    `json:"payment_method_id"`
	ShippingMethodID   int    `json:"shipping_method_id"`
	WebsiteID          int    `json:"website_id"`
	Postcode           string `json:"postcode"`
	Street             string `json:"street"`
	SubscriptionOrigin string `json:"subscription_origin"`
}

var EmailAvailabilityBody EmailAvailabilityResponse

type EndGenTask struct {
	Email      string `csv:"Email"`
	Password   string `csv:"Password"`
	FirstName  string `csv:"FirstName"`
	LastName   string `csv:"LastName"`
	Address1   string `csv:"Address1"`
	Address2   string `csv:"Address2"`
	County     string `csv:"County"`
	Phone      int    `csv:"Phone"`
	Country    string `csv:"Country"`
	CardNumber int    `csv:"CardNumber"`
	CardExpiry string `csv:"CardExpiry"`
	CardCvC    int    `csv:"CardCvC"`
	TaskNumber int
}

func getShapeHeaders() ShapeHeaders {
	for getVpnStatus() != "200 OK" {
		fmt.Println("checking vpn")
		time.Sleep(10 * time.Second)
		//status := getVpnStatus()
	}
	//	if status != "200 OK" {
	//		break
	//	} else {
	//		continue
	//	}
	//}
	jar, err := cookiejar.New(nil)
	client := http.Client{Jar: jar}
	var returnHeaders ShapeHeaders
	getHeadersReq, err := http.NewRequest("GET", "http://localhost:3000", nil)
	if err != nil {
		log.Println(err)
	}
	getHeaders, err := client.Do(getHeadersReq)
	if err != nil {
		fmt.Println(err)
	}
	if getHeaders.Body != nil {
		defer getHeaders.Body.Close()
	}
	headerText, err := ioutil.ReadAll(getHeaders.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Get Header Status: ", getHeaders.Status)
	unMarshallingerr := json.Unmarshal([]byte(headerText), &returnHeaders)
	if unMarshallingerr != nil {
		return ShapeHeaders{}
	}
	//fmt.Println(returnHeaders)
	if getHeaders.Status == "500" {
	}
	recover()
	return returnHeaders

}

func getVpnStatus() string {

	jar, err := cookiejar.New(nil)
	client := http.Client{Jar: jar}

	vpnReq, err := http.NewRequest("GET", "http://localhost:3000/connection", nil)
	if err != nil {
		log.Println(err)
	}
	vpnStatusCheck, err := client.Do(vpnReq)
	if err != nil {
		fmt.Println(err)
	}
	//defer vpnStatusCheck.Body.Close()
	//vpnText, err := ioutil.ReadAll(vpnStatusCheck.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	////fmt.Println("Get Header Status: ", getHeaders.Status)
	//json.Unmarshal([]byte(vpnText), &vpnStatus)
	////fmt.Println(returnHeaders)
	////if vpnStatusCheck.Status == "500" {
	////}
	////recover()
	//fmt.Println(vpnStatus.enabled)
	fmt.Println(vpnStatusCheck.Status)
	return vpnStatusCheck.Status

}

func CreateEndTasksFromCsv(fileName string) []*EndGenTask {
	f, err := os.Open("/Users/kacey/go/src/EsoteraFinal/sites/END/END.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		// handle error
	}
	var EndGenTasks []*EndGenTask
	// UnmarshalBytes parses the CSV from the bytes in the interface.
	err = gocsv.UnmarshalBytes(fileBytes, &EndGenTasks)

	var TaskCount int = 1
	for _, task := range EndGenTasks {
		task.TaskNumber = TaskCount
		TaskCount++

	}
	fmt.Println("Finished Loading ", TaskCount-1, " tasks.")

	return EndGenTasks
}

func EmailCheck(client http.Client, email string) (string, bool, error) {

	headers := getShapeHeaders()
	//jar, err := cookiejar.New(nil)
	////client = http.Client{Jar: jar}
	//client.Jar = jar
	var PostData string = `{"email": "` + email + `"}`
	var data = strings.NewReader(PostData)
	req, err := http.NewRequest("POST", "https://api.endclothing.com/customer/rest/v2/gb/email-availability", data)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authority", "api.endclothing.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-GB,en;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("exj5WzXnUF-a", headers.Shapeheadera)
	req.Header.Set("exj5WzXnUF-b", headers.Shapeheaderb)
	req.Header.Set("exj5WzXnUF-c", headers.Shapeheaderc)
	req.Header.Set("exj5WzXnUF-d", headers.Shapeheaderd)
	req.Header.Set("exj5WzXnUF-f", headers.Shapeheaderf)
	req.Header.Set("exj5WzXnUF-z", headers.Shapeheaderz)
	req.Header.Set("Origin", "https://www.endclothing.com")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://www.endclothing.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Gpc", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")

	fmt.Println("Checking E-mail...")

	checkEmailResp, err := client.Do(req)
	if checkEmailResp.Body != nil {
		defer checkEmailResp.Body.Close()
	}

	status := checkEmailResp.Status

	bodyText, err := ioutil.ReadAll(checkEmailResp.Body)

	unMarshallingerr := json.Unmarshal([]byte(bodyText), &EmailAvailabilityBody)
	if unMarshallingerr != nil {
		return "", false, unMarshallingerr
	}
	exists := EmailAvailabilityBody.Body.Available

	checkEmailResp.Body.Close()
	return status, exists, err
}
func GetGuestID(client http.Client, useragent string) (string, string, error) {
	headers := getShapeHeaders()
	getIDReq, err := http.NewRequest("GET", "https://api.endclothing.com/customer/rest/v2/gb/customer/guest", nil)
	if err != nil {
		fmt.Println(err)
	}
	getIDReq.Header.Set("Authority", "api.endclothing.com")
	getIDReq.Header.Set("Accept", "application/json, text/plain, */*")
	getIDReq.Header.Set("Accept-Language", "en-GB,en;q=0.9")
	getIDReq.Header.Set("Cache-Control", "no-cache")
	getIDReq.Header.Set("Content-Type", "application/json")
	getIDReq.Header.Set("exj5WzXnUF-a", headers.Shapeheadera)
	getIDReq.Header.Set("exj5WzXnUF-b", headers.Shapeheaderb)
	getIDReq.Header.Set("exj5WzXnUF-c", headers.Shapeheaderc)
	getIDReq.Header.Set("exj5WzXnUF-d", headers.Shapeheaderd)
	getIDReq.Header.Set("exj5WzXnUF-f", headers.Shapeheaderf)
	getIDReq.Header.Set("exj5WzXnUF-z", headers.Shapeheaderz)
	getIDReq.Header.Set("Origin", "https://www.endclothing.com")
	getIDReq.Header.Set("Pragma", "no-cache")
	getIDReq.Header.Set("Referer", "https://www.endclothing.com/")
	getIDReq.Header.Set("Sec-Fetch-Dest", "empty")
	getIDReq.Header.Set("Sec-Fetch-Mode", "cors")
	getIDReq.Header.Set("Sec-Fetch-Site", "same-site")
	getIDReq.Header.Set("Sec-Gpc", "1")
	getIDReq.Header.Set("User-Agent", useragent)
	getIDResp, err := client.Do(getIDReq)
	if err != nil {
		fmt.Println(err)
	}
	if getIDResp.Body != nil {
		defer getIDResp.Body.Close()
	}

	headerText, err := ioutil.ReadAll(getIDResp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var ReturnedID GuestID

	unMarshallingerr := json.Unmarshal([]byte(headerText), &ReturnedID)
	if unMarshallingerr != nil {
		return "", "", unMarshallingerr
	}
	fmt.Println("ID: ", ReturnedID)
	return getIDResp.Status, ReturnedID.Body.GuestID, err
}
func CreateAccount(client http.Client, email string, firstName string, lastName string, password string, guestid string, useragent string) (string, error) {
	headers := getShapeHeaders()
	//type CreateAccountData struct {
	//	Email      string `json:"email"`
	//	First_name string `json:"email"`
	//	Firstname  string `json:"firstname"`
	//	Guest_id   string `json:"guest_id"`
	//	Last_name  string `json:"last_name"`
	//	Lastname   string `json:"lastname"`
	//	Password   string `json:"password"`
	//	Store_id   int    `json:"store_id"`
	//	Website_id int    `json:"website_id"`
	//}
	//
	//accountData := &CreateAccountData{
	//	Email:      email,
	//	First_name: firstName,
	//	Firstname:  firstName,
	//	Guest_id:   guestid,
	//	Last_name:  lastName,
	//	Lastname:   lastName,
	//	Password:   password,
	//	Store_id:   1,
	//	Website_id: 1,
	//}
	//marshalledData, err := json.Marshal(accountData)
	//if err != nil {
	//	fmt.Println(err)
	//
	//}
	//var postData string = string(marshalledData)
	//fmt.Println(postData)
	//var jsonStr = []byte(postData)

	postData := "{\"email\":\"" + email + "\",\"first_name\":\"" + firstName + "\",\"firstname\":\"" + firstName + "\",\"last_name\":\"" + lastName + "\",\"lastname\":\"" + lastName + "\",\"password\":\"" + password + "\",\"website_id\":1,\"store_id\":1,\"guest_id\":\"" + guestid + "\"}"
	//fmt.Println(postData)
	//var genData = strings.NewReader(postData)
	genReq, err := http.NewRequest("POST", "https://api.endclothing.com/customer/rest/v2/gb/customer", strings.NewReader(postData))
	genReq.Header.Set("Authority", "api.endclothing.com")
	genReq.Header.Set("method", "POST")
	genReq.Header.Set("path", "/customer/rest/v2/gb/customer")
	genReq.Header.Set("scheme", "https")
	genReq.Header.Set("Accept", "application/json, text/plain, */*")
	genReq.Header.Set("accept-encoding", "gzip, deflate, br")
	genReq.Header.Set("Accept-Language", "en-GB,en;q=0.9")
	genReq.Header.Set("Cache-Control", "no-cache")
	genReq.Header.Set("content-length", string(len(postData)))
	genReq.Header.Set("Content-Type", "application/json")
	//genReq.Header.Set("Cookie", "end_country=GB; end_store=gb; 4tagIi0FLf=A0tad9iDAQAAPyEHkoYST1jxlq009zJI4LG84yNcXY6d0jlQNVMEgrSGUQM4ATO2U2yucm46wH9eCOfvosJeCA|1|1|98f3cd7e40faad7f9a1b1171f8a1b23b357050ae; _abck=2BEE959FD3F53B49720B7AFA0E22ABEB~-1~YAAQbHAyuMuQt8+DAQAABWl32AiF/DFDcxz/Z1/vW+ebue0PolHngc9e7AxD0w6Xbh7j0eUGB8dNP2MlbCOE8zlBdhKGnLFZ0M2ezj5M7oLXyMbH7tcPIiAnsf9C6zVMyqHW07lpS67m/+/lNCPgUp2K8mTDTZYKFa9DhQj+wjSYZ46LcQga6zFNICNvxyLNh007g5Brav6ANMs0aag4fB8InSoBYDR0YR71ARenX7qjvevptCFkNemQTIZVUpg9CvHeh4/Zbr0ZLxv2bFQtXme9ialsO7RvpJ5P4yGFNfBB7+1O8G1V+J8TflHrwj4NQzfjTph9eknnywAeisGzZBRU0jv7kY/HkFQmlIiggCpaSMSdD6gwcXDSaztLTBMnWg==~-1~-1~-1; bm_sz=2273337B526F8BA32FF14F7A691EFCA1~YAAQbHAyuM2Qt8+DAQAABWl32BGnp6Si4L2a+utawx6MPggmzYJXBgW453FAbPWSMP05Veema9fedE+div+cDI29uSM6+VCpbCIGtjWA/Oqk4RwAxsM9Wzp5orM09NG49I2AttTRv2H502O39I94JNc+qI6q3S3sGBEN3aEwmW+HCxGPHiQ01P+6tpUQJ/Qh8GqwrztlRn7y78piCJBTff8PgWIJG2kwnPJZ38XjlZvVQ/LvsHviBnRaXx9TP1p+FLpFMkbKuiVHKrxK0ogdfe1p+KmllmnPIn8EFTAWadbDRyIg5KFSIw==~4405044~3686724; guest_tracking_id=97808ea5-727a-460a-b5aa-f02137c30a76; __zlcmid=1CRlIhmshISXMck; AKA_A2=A; end_root_menu=2; ak_bmsc=ED77B84D3885FF0DFA8E0D6B6BCDDDC0~000000000000000000000000000000~YAAQiO1lX519adCDAQAAzZfl2BEK4FLtR8G8FP5/cHPZ9XCe6siJuT0yHnJqLJAxGjAeHDNC9qZqlig0dNvKDf+IyDGvMLh7rUZU6va7oy+h01gxl8PPXjVagUyE1EkVHi1Ao0gMgpvN3pHqekzJbwjgLAAENSS4JuMWtHhi3Jzkt+geQmWu1Pl6fF6V/1r02NBLyXbDtPYEYlmW2GJFa+TW6yjZW9T40II0tpBGf518Xr+Z5jg5DokFG64Ez9mW/krYVhw6ww72ObAy5HzAFPS6YdFKvaHK5fGG4xidfynrMAhdVvdLNZoIx//peztxz7xTHA+6x97A4rrXHQMM9065vIEfSsnsLxkyID43z9+g3QTAzktwRqjAkIjjBFYbHGjMFPqN1ZUsNpyjFAF3")
	genReq.Header.Set("Exj5wzxnuf-A", headers.Shapeheadera)
	genReq.Header.Set("Exj5wzxnuf-A0", headers.Shapeheadera0)
	genReq.Header.Set("Exj5wzxnuf-B", headers.Shapeheaderb)
	genReq.Header.Set("Exj5wzxnuf-C", headers.Shapeheaderc)
	genReq.Header.Set("Exj5wzxnuf-D", headers.Shapeheaderd)
	genReq.Header.Set("Exj5wzxnuf-F", headers.Shapeheaderf)
	genReq.Header.Set("Exj5wzxnuf-Z", headers.Shapeheaderz)
	genReq.Header.Set("Origin", "https://www.endclothing.com")
	//genReq.Header.Set("pragma", "no-cache")
	genReq.Header.Set("Referer", "https://www.endclothing.com/")
	genReq.Header.Set("Sec-Fetch-Dest", "empty")
	genReq.Header.Set("Sec-Fetch-Mode", "cors")
	genReq.Header.Set("Sec-Fetch-Site", "same-site")
	genReq.Header.Set("Sec-Gpc", "1")
	genReq.Header.Set("User-Agent", useragent)
	genResp, genReqErr := client.Do(genReq)
	if genReqErr != nil {
		fmt.Println(err)
	}
	if genResp.Body != nil {
		defer genResp.Body.Close()
	}
	genStatus := genResp.Status
	fmt.Println("Account creation status: ", genStatus)
	return genStatus, genReqErr
}

func EndLogin(client http.Client, email string, password string, useragent string) (string, string) {
	maxRetries := 3
	retries := 0
	for retries < maxRetries {
		headers := getShapeHeaders()

		type loginPayload struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		data := loginPayload{
			email, password,
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			// handle err
		}
		body := bytes.NewReader(payloadBytes)

		req, err := http.NewRequest("POST", "https://api.endclothing.com/customer/rest/v2/gb/customer/token", body)
		if err != nil {
			// handle err
		}
		req.Header.Set("Authority", "api.endclothing.com")
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Exj5wzxnuf-A", headers.Shapeheadera)
		req.Header.Set("Exj5wzxnuf-B", headers.Shapeheaderb)
		req.Header.Set("Exj5wzxnuf-C", headers.Shapeheaderc)
		req.Header.Set("Exj5wzxnuf-D", headers.Shapeheaderd)
		req.Header.Set("Exj5wzxnuf-F", headers.Shapeheaderf)
		req.Header.Set("Exj5wzxnuf-Z", headers.Shapeheaderz)
		req.Header.Set("Origin", "https://www.endclothing.com")
		req.Header.Set("Referer", "https://www.endclothing.com/")
		req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-site")
		req.Header.Set("User-Agent", useragent)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			//retries = retries + 1
			//continue
		}
		fmt.Println(resp.Status)
		if resp != nil {

			defer resp.Body.Close()
		} else {
			retries = retries + 1
			continue
		}
		getBearer, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var unmarshalledAcctData BearerFromLogin
		UnmarshallingErr := json.Unmarshal([]byte(getBearer), &unmarshalledAcctData)
		if UnmarshallingErr != nil {
			return "", ""
		}
		fmt.Println(unmarshalledAcctData.Body.Token)
		return unmarshalledAcctData.Body.Token, resp.Status

	}

	return "", "test"
	//getBearer, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var unmarshalledAcctData BearerFromLogin
	//UnmarshallingErr := json.Unmarshal([]byte(getBearer), &unmarshalledAcctData)
	//if UnmarshallingErr != nil {
	//	return "", ""
	//}
	//return unmarshalledAcctData.Body.Token, resp.Status
	{
		return "nil", "nil"
	}
}

func GetEntryData(client http.Client, token string, useragent string) (EntryData, string) {
	maxRetries := 3
	retries := 0
	for retries < maxRetries {
		headers := getShapeHeaders()

		req, err := http.NewRequest("GET", "https://api.endclothing.com/customer/rest/v2/gb/account/me", nil)
		if err != nil {
			// handle err
		}
		req.Header.Set("Authority", "api.endclothing.com")
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Exj5wzxnuf-A", headers.Shapeheadera)
		req.Header.Set("Exj5wzxnuf-B", headers.Shapeheaderb)
		req.Header.Set("Exj5wzxnuf-C", headers.Shapeheaderc)
		req.Header.Set("Exj5wzxnuf-D", headers.Shapeheaderd)
		req.Header.Set("Exj5wzxnuf-F", headers.Shapeheaderf)
		req.Header.Set("Exj5wzxnuf-Z", headers.Shapeheaderz)
		req.Header.Set("Origin", "https://launches.endclothing.com")
		req.Header.Set("Referer", "https://launches.endclothing.com/")
		req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-site")
		req.Header.Set("User-Agent", useragent)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			retries = retries + 1
			continue
		}
		if resp != nil {
			defer resp.Body.Close()
		} else {
			retries = retries + 1
			continue
		}
		accountData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var unmarshalledAcctData AccountData
		unMarshallingerr := json.Unmarshal([]byte(accountData), &unmarshalledAcctData)
		if unMarshallingerr != nil {
			//return EntryData{}, ""
			fmt.Println(unMarshallingerr)
		}
		var createdEntryData EntryData
		createdEntryData.CustomerID, _ = strconv.Atoi(unmarshalledAcctData.Body.CustomerID)
		createdEntryData.WebsiteID = unmarshalledAcctData.Body.WebsiteID
		createdEntryData.CustomerID = unmarshalledAcctData.Body.ID

		for _, element := range unmarshalledAcctData.Body.Addresses {
			if element.DefaultShipping == true {
				createdEntryData.ShippingAddressID = element.ID
				createdEntryData.Postcode = element.Postcode
				createdEntryData.Street = element.Street[0]
			}
		}
		for _, element := range unmarshalledAcctData.Body.Addresses {
			if element.DefaultBilling == true {
				createdEntryData.BillingAddressID = element.ID
			}
		}
		createdEntryData.SubscriptionOrigin = "web"
		createdEntryData.ShippingMethodID = 159
		if resp != nil {
			return createdEntryData, resp.Status
		}
	}
	fmt.Println("Failed getting entrydata")
	return EntryData{
		CustomerID:         0,
		ProductSizeID:      0,
		ShippingAddressID:  0,
		BillingAddressID:   0,
		PaymentMethodID:    0,
		ShippingMethodID:   0,
		WebsiteID:          0,
		Postcode:           "",
		Street:             "",
		SubscriptionOrigin: "",
	}, "nil"
}

func GetPaymentMethod(client http.Client, bearer string, useragent string) (int, string) {
	maxRetries := 3
	retries := 0
	for retries < maxRetries {
		headers := getShapeHeaders()
		req, err := http.NewRequest("GET", "https://api2.endclothing.com/gb/rest/V1/end/vault/mine", nil)
		if err != nil {
			// handle err
		}
		req.Header.Set("Authority", "api2.endclothing.com")
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Set("Authorization", "Bearer "+bearer)
		req.Header.Set("Exj5wzxnuf-A", headers.Shapeheadera)
		req.Header.Set("Exj5wzxnuf-B", headers.Shapeheaderb)
		req.Header.Set("Exj5wzxnuf-C", headers.Shapeheaderc)
		req.Header.Set("Exj5wzxnuf-D", headers.Shapeheaderd)
		req.Header.Set("Exj5wzxnuf-F", headers.Shapeheaderf)
		req.Header.Set("Exj5wzxnuf-Z", headers.Shapeheaderz)
		req.Header.Set("Origin", "https://launches.endclothing.com")
		req.Header.Set("Referer", "https://launches.endclothing.com/")
		req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-site")
		req.Header.Set("User-Agent", useragent)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			retries = retries + 1
			continue
		}
		if resp != nil {
			defer resp.Body.Close()
		} else {
			retries = retries + 1
			continue
		}
		cardData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var entryCardID int
		var unmarshalledCardData []PaymentVaultData
		unMarshallingerr := json.Unmarshal([]byte(cardData), &unmarshalledCardData)
		if unMarshallingerr != nil {
			fmt.Println(unMarshallingerr)
		}
		for _, element := range unmarshalledCardData {
			if element.IsDefault == true {
				entryCardID = element.EntityID
			}
		}
		if resp != nil {

			return entryCardID, string(resp.Status)

		}
	}
	fmt.Println("Failed getting card")
	return 0, "nil"

}

var Resp *http.Response

func EnterDraw(client http.Client, bearer string, entryData EntryData, cardData int, useragent string) string {
	maxRetries := 3
	retries := 0
	for retries < maxRetries {
		headers := getShapeHeaders()
		//for true {
		data := EntryData{
			CustomerID:         entryData.CustomerID,
			ProductSizeID:      60846,
			ShippingAddressID:  entryData.ShippingAddressID,
			BillingAddressID:   entryData.BillingAddressID,
			PaymentMethodID:    cardData,
			ShippingMethodID:   entryData.ShippingMethodID,
			WebsiteID:          entryData.WebsiteID,
			Postcode:           entryData.Postcode,
			Street:             entryData.Street,
			SubscriptionOrigin: entryData.SubscriptionOrigin,
		}
		payloadBytes, err := json.Marshal(data)

		if err != nil {
			// handle err
		}
		body := bytes.NewReader(payloadBytes)

		req, err := http.NewRequest("POST", "https://launches-api.endclothing.com/api/subscriptions", body)
		if err != nil {
			// handle err
		}
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Set("Authorization", "Bearer "+bearer)
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Origin", "https://launches.endclothing.com")
		req.Header.Set("Referer", "https://launches.endclothing.com/")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-site")
		req.Header.Set("User-Agent", useragent)
		req.Header.Set("Exj5wzxnuf-A", headers.Shapeheadera)
		req.Header.Set("Exj5wzxnuf-B", headers.Shapeheaderb)
		req.Header.Set("Exj5wzxnuf-C", headers.Shapeheaderc)
		req.Header.Set("Exj5wzxnuf-D", headers.Shapeheaderd)
		req.Header.Set("Exj5wzxnuf-F", headers.Shapeheaderf)
		req.Header.Set("Exj5wzxnuf-Z", headers.Shapeheaderz)
		req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")

		Resp, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//data1, err1 := ioutil.ReadAll(Resp.Body)
		//if err != nil {
		//	fmt.Println(err1)
		//	fmt.Println(data1)
		//	continue
		//} else {
		//	err = Resp.Body.Close()
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//	break
		//}

		//return Resp.Status
		if Resp == nil {
			retries = retries + 1
			continue
		} else {
			return Resp.Status

		}

	}

	return "nil"
}
