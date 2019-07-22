package fetch_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenguyenit/testing/fetch-an-api/fetch"
	"github.com/thenguyenit/testing/fetch-an-api/mocks"
)

func TestFetchAstro(t *testing.T) {
	astroAPI := &fetch.AstroAPI{
		URL: "http://api.open-notify.org/astros.json",
	}

	number, err := astroAPI.Get()

	assert.NoError(t, err)
	assert.True(t, number >= 0)
}

func TestGetAstroWrongURL(t *testing.T) {
	testTables := []struct {
		URL              string
		ExpectedResponse int
	}{
		{
			"httpss://abc.com/a.json",
			0,
		},
		{
			"",
			0,
		},
	}

	for _, rowTest := range testTables {
		astroAPI := fetch.AstroAPI{
			URL: rowTest.URL,
		}

		res, err := astroAPI.Get()
		assert.Equal(t, rowTest.ExpectedResponse, res)
		assert.Error(t, err)
	}
}

func TestFetchAstroTimeOut(t *testing.T) {
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(time.Second * 5)

	// 	message := map[string]interface{}{
	// 		"id":    "12",
	// 		"scope": "test-scope",
	// 	}

	// 	b, _ := json.Marshal(message)
	// 	io.WriteString(w, string(b))
	// }
	// req := httptest.NewRequest(http.MethodGet, "http://api.open-notify.org/astros.json", nil)
	// w := httptest.NewRecorder()
	// handler(w, req)

	// resp := w.Result()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	mockAstroAPI := new(mocks.RESTAPI)
	mockAstroAPI.On("fetch").Return(nil, errors.New("TimeOut"))

	number, err := mockAstroAPI.Get()
	assert.Error(t, err)
	assert.True(t, number == 0)

}
