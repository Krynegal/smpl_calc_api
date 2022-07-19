package validators

import (
	"encoding/json"
	"errors"
	"github.com/Krynegal/smpl_calc_api.git/internal/types"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	//ParameterMissed = errors.New("parameter has been missed")
	//ConvertError1   = errors.New("fail convert first operand")
	//ConvertError2   = errors.New("fail convert first operand")
	UnmarshallError = errors.New("fail unmarshall request body")
)

func GetData(r *http.Request) (types.Data, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	log.Printf("request body: %v", string(reqBody))
	data := types.Data{}
	if err = json.Unmarshal(reqBody, &data); err != nil {
		return data, UnmarshallError
	}
	return data, nil
}

// нужна проверка на то, что строка с операндом не пустая и не является строкой
