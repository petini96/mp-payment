package mp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestGetPreferencia(t *testing.T) {
	requestFields := NewDefaultRequestFields("http://localhost:5000")

	requestImplementer := NewMPRequestImplementer()

	varzipCode, streetName, streetNumber := "", "", ""
	requestImplementer.NewMPAddress(address.Cep.String, address.Street.String, address.HouseNumber.String)
	requestImplementer.NewMPPayer(user.Name, user.Email, "67", client.CellNumber.String, user.Cpf.String, time.Now())
	requestImplementer.NewMPItems(tour.ID, tour.Name, tour.Description, "base64", "travelers", 1, "BRL", tourSchedule.Price)
	requestImplementer.NewMetadata("Vinicius", "544")

	requestFields.CompleteMPRequest(requestImplementer)

	endpoints := NewEndpoints()
	resultRequest, _ := MakePostRequest(requestFields, endpoints.Preferences, "application/json")
	var dat map[string]interface{}
	if err := json.Unmarshal(resultRequest, &dat); err != nil {
		panic(err)
	}
	log.Fatal(dat["sandbox_init_point"])
	//log.Fatal(string(resultRequest))
}

type ResponsePost struct{}

func MakeGetRequest(endpoint string) string {
	apiUrl := "https://api.mercadopago.com"
	resource := "/checkout/preferences"
	data := url.Values{}
	data.Set("name", "foo")
	data.Set("surname", "bar")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Authorization", "Bearer "+TokenSandbox)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	return resp.Status
}

func MakePostRequest(object interface{}, url string, headerType string) ([]byte, error) {
	postBody, _ := json.Marshal(object)
	responseBody := bytes.NewBuffer(postBody)

	str := responseBody.Bytes()

	var jsonStr = str
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+TokenSandbox)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
