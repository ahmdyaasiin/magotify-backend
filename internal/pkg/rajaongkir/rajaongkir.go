package rajaongkir

import (
	"encoding/json"
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type CostDetail struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type Cost struct {
	Service     string       `json:"service"`
	Description string       `json:"description"`
	Cost        []CostDetail `json:"cost"`
}

type Result struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Costs []Cost `json:"costs"`
}

type RajaOngkirResponse struct {
	RajaOngkir struct {
		Query struct {
			Origin      string `json:"origin"`
			Destination string `json:"destination"`
			Weight      int    `json:"weight"`
			Courier     string `json:"courier"`
		} `json:"query"`
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		OriginDetails      interface{} `json:"origin_details"`
		DestinationDetails interface{} `json:"destination_details"`
		Results            []Result    `json:"results"`
	} `json:"rajaongkir"`
}

func CheckCost(postalCode int, weightOnGram float64, dest *[]model.ServicesOngkir) error {
	url := "https://api.rajaongkir.com/starter/cost"

	payload := strings.NewReader(fmt.Sprintf("origin=%d&destination=%d&weight=%.3f&courier=jne", ORIGIN, postalCode, weightOnGram))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("key", os.Getenv("API_RAJAONGKIR"))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	var rajaOngkirResponse RajaOngkirResponse
	err = json.Unmarshal(body, &rajaOngkirResponse)
	if err != nil {
		return err
	}

	for _, result := range rajaOngkirResponse.RajaOngkir.Results {
		for _, cost := range result.Costs {
			for _, costDetail := range cost.Cost {

				var etd string
				if len(costDetail.Etd) < 2 {
					etd = costDetail.Etd + " Days"
				} else {
					etd = costDetail.Etd + " Day"
				}

				service := model.ServicesOngkir{
					Name:        cost.Service,
					Description: cost.Description,
					Cost:        costDetail.Value,
					Estimation:  etd,
					Note:        costDetail.Note,
				}
				*dest = append(*dest, service)
			}
		}
	}

	return nil
}
