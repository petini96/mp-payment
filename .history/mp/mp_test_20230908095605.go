package mp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestGetPreferencia(t *testing.T) {
	requestFields := NewDefaultRequestFields("http://localhost:8080")

	requestImplementer := NewMPRequestImplementer()

	varzipCode := "79005030"
	streetName := "domingos tenuta, 534"
	streetNumber := "54"
	requestImplementer.NewMPAddress(varzipCode, streetName, streetNumber)

	payerID := "24242"
	name := "Edmara"
	email := "vspetini96@gmail.com"
	areaCode := "55"
	cellNumber := "67984087417"
	identificationNumber := "04942179106"
	createdAt := time.Now()
	requestImplementer.NewMPPayer(payerID, name, email, areaCode, cellNumber, identificationNumber, createdAt)

	itemID := 0
	title := "Pintada atômica"
	description := "vixe"
	pictureUrl := "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSQsLSFnF027f6lE7u3DXIti4JWoLxyARUNHxEVIUSFlA&s"
	categoryID := "1"
	quantity := 1
	currencyID := "BRL"
	unitPrice := 100.0
	requestImplementer.NewMPItems(itemID, title, description, pictureUrl, categoryID, int(quantity), currencyID, unitPrice)

	responsible := "vinicius"
	register := "222"
	requestImplementer.NewMetadata(responsible, register)

	requestFields.CompleteMPRequest(requestImplementer)

	endpoints := NewEndpoints()
	resultRequest, _ := MakePostRequest(requestFields, endpoints.Preferences, "application/json")
	var dat map[string]interface{}
	if err := json.Unmarshal(resultRequest, &dat); err != nil {
		panic(err)
	}
	log.Fatal(dat["init_point"])
	//log.Fatal(dat["sandbox_init_point"])
	//log.Fatal(string(resultRequest))
}

type ResponsePost struct{}

// func MakeGetRequest(endpoint string) string {
// 	apiUrl := "https://api.mercadopago.com"
// 	resource := "/checkout/preferences"
// 	data := url.Values{}
// 	data.Set("name", "foo")
// 	data.Set("surname", "bar")

// 	u, _ := url.ParseRequestURI(apiUrl)
// 	u.Path = resource
// 	urlStr := u.String()

// 	client := &http.Client{}
// 	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
// 	r.Header.Add("Authorization", "Bearer "+TokenSandbox)
// 	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

// 	resp, _ := client.Do(r)
// 	return resp.Status
// }

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
