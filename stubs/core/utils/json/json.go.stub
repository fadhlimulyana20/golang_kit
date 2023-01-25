package json

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func Decode(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		logrus.Error("Cannot read data")
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		logrus.Error("Error when unmarshal data")
	}

	return err
}
