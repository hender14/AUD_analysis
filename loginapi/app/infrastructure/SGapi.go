package infrastructure

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

func (r Sgconfig) Reqconf(jsondata []byte) (response *rest.Response, err error) {
	// requestのMethodをPostに
	r.Conf.Method = "POST"
	r.Conf.Body = jsondata
	response, err = sendgrid.API(r.Conf)
	// fmt.Println(response.StatusCode)
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	return response, err
}
