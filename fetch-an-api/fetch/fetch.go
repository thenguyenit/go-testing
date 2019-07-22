package fetch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/thenguyenit/testing/fetch-an-api/models"
)

type AstroAPI struct {
	URL string
}

func (s *AstroAPI) fetch() ([]byte, error) {

	client := http.Client{
		// Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, s.URL, nil)
	if err != nil {
		return nil, err
	}

	response, getErr := client.Do(req)
	//Timeout
	if getErr != nil {
		return nil, getErr
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}

func (api *AstroAPI) Get() (int, error) {
	resJson, err := api.fetch()
	if err != nil {
		return 0, err
	}
	var astrosResponse models.AstrosResponse
	unMarshallErr := json.Unmarshal(resJson, &astrosResponse)
	if unMarshallErr != nil {
		return 0, unMarshallErr
	}

	return astrosResponse.Number, nil
}
