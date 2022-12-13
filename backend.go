package astute

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/net/html/charset"
	"gopkg.in/resty.v1"
)

type ClientResponse struct {
	Code int
	Data []byte
}

type Backend interface {
	Call(url string, ctx string, action string, requestBody string, requestData interface{}) (*ClientResponse, error)
}

func NewBackend() Backend {
	backend := &restApiBackend{}
	return backend
}

type restApiBackend struct {
	HTTPClient *resty.Client
}

func (b restApiBackend) Call(url string, ctx string, action string, requestBody string, requestData interface{}) (*ClientResponse, error) {
	var clientResponse ClientResponse
	pl, err := createTemplate(ctx, requestBody, requestData)
	if err != nil {
		return &clientResponse, err
	}
	payload := []byte(pl)

	// Prepare the request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return &clientResponse, err
	}

	// Set the content type header, as well as the oter required headers
	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", action)

	// Prepare the client request
	client := &http.Client{}

	// Dispatch the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return &clientResponse, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		responseData, err := io.ReadAll(resp.Body)
		if err != nil {
			return &clientResponse, err
		}
		clientResponse = ClientResponse{
			Code: resp.StatusCode,
			Data: responseData,
		}
		return &clientResponse, nil
	}

	clientResponse = ClientResponse{
		Code: resp.StatusCode,
		Data: nil,
	}

	return &clientResponse, nil
}

func createTemplate(ctx string, t string, data interface{}) (string, error) {
	var templateWithData bytes.Buffer
	ttp, err := template.New(ctx).Parse(t)
	if err != nil {
		return "", err
	}
	err = ttp.ExecuteTemplate(&templateWithData, ctx, data)
	if err != nil {
		return "", err
	}
	return templateWithData.String(), nil
}

func ParseResponse[T PayrollEntity](data []byte, source T) (T, error) {
	var response T
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&response)
	return response, err
}
