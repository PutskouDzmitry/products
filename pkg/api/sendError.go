package api

import (
	"fmt"
)

func SendErrorFromAPI(err error, param string) string {
	if err.Error() == "or you set incorrect parameter or we cannot find data with this parameter" {
		return fmt.Sprint("or you set incorrect parameter or we cannot find data with this parameter: ", param)
	} else {
		return fmt.Sprintf("got an error when tried to get parameters: ", err.Error())
	}
}