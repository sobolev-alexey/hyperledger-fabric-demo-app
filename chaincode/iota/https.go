package iota

import (
     "fmt"
     "net/http"
     "strconv"
     "io/ioutil"
     "encoding/json"
)

type ErrRequest struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error"`
	Exception    string `json:"exception"`
}

func (er *ErrRequest) Error() string {
	var msg string
	if er.ErrorMessage != "" {
		msg += fmt.Sprintf("error message: %s;", er.ErrorMessage)
	}
	if er.Exception != "" {
		msg += fmt.Sprintf("exception message: %s;", er.Exception)
	}
	msg += fmt.Sprintf("http status code: %d;", er.Code)
	return msg
}

// See https://golang.org/pkg/net/http/
func MakeRequest(url string, number int, ch chan<- string) {
  res, err := http.Get(url + strconv.Itoa(number))
  if err != nil {
      close(ch)
      return
  }
  defer res.Body.Close()
  defer close(ch) //don't forget to close the channel as well

  body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
  }

  ch <- string(body)
}

func MakeRequest1(url string, out interface{}) error {
  resp, err := http.Get(url)
  if err != nil {
      return err
  }
  defer resp.Body.Close()
  bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errResp := &ErrRequest{Code: resp.StatusCode}
		json.Unmarshal(bs, errResp)
		return errResp
	}

	return json.Unmarshal(bs, out)
}

func MakeRequest2(url string, number int, ch chan<- string) {
  var data []byte

  data = make([]byte, 128)
  res, err := http.Get(url + strconv.Itoa(number))
  if err != nil {
      close(ch)
      return
  }
  defer res.Body.Close()
  defer close(ch) //don't forget to close the channel as well

  for n, err := res.Body.Read(data); err == nil; n, err = res.Body.Read(data) {
      ch <- string(data[:n])
  }
}

