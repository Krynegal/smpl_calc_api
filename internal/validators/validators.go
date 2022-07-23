package validators

import (
	"encoding/json"
	"errors"
	"github.com/Krynegal/smpl_calc_api.git"
	"io/ioutil"
	"log"
	"net/http"
)

func GetData(r *http.Request) (smpl_calc_api.Data, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	log.Printf("request body: %v", string(reqBody))
	data := smpl_calc_api.Data{}
	if err = json.Unmarshal(reqBody, &data); err != nil {
		return data, errors.New("fail unmarshall request body")
	}
	return data, nil
}
