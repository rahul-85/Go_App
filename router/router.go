package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"Go_App/controller"
	"Go_App/model"
)

func getResponseBytes(url string) []byte {
	myClient := http.Client{
		Timeout: time.Second * 6, // Timeout after 6 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func StoreCovidData() {
	url := "https://data.covid19india.org/v4/min/data.min.json"
	body := getResponseBytes(url)
	var data_covid map[string]interface{}
	json.Unmarshal([]byte(body), &data_covid)

	client := controller.ConnectMongoDB()
	controller.RemoveAllData(client)
	for key := range data_covid { 
		var data model.CovidData
		data.Last_updated = data_covid[key].(map[string]interface{})["meta"].(map[string]interface{})["last_updated"].(string)
		data.Regional_id = key
		if (key == "TT") {
			data.Regional_id = "IND"
		}
		data.CovidCases = data_covid[key].(map[string]interface{})["total"].(map[string]interface{})["confirmed"].(float64)
		fmt.Println(data)
		controller.InsertOneData(client,data)
	}
}

func RetrieveCovidData(loc string) []byte{
	client := controller.ConnectMongoDB()
	var data []model.CovidData
	data = append(data, controller.RetrieveData(client, loc))
	data = append(data, controller.RetrieveData(client, "IND"))
    jsonData, err := json.Marshal(data)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(jsonData))
	return jsonData
}

func RetrieveDataFromLatLong(lat string, long string) []byte{
	url := "http://api.positionstack.com/v1/reverse?access_key=55f9c71e0249cba0217c3d4e00406807&query=" + lat + "," + long
    body := getResponseBytes(url)
	var data_loc map[string]interface{}
	json.Unmarshal([]byte(body), &data_loc)
	regionID := data_loc["data"].([]interface{})[0].(map[string]interface{})["region_code"].(string)
	fmt.Println("Region id: ", regionID)
	return RetrieveCovidData(regionID)
}