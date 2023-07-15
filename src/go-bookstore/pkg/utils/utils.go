package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, x); err != nil {
		return err
	}
	return nil
}

func ParseBookId(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	return strconv.ParseInt(bookId, 0, 0)

}
