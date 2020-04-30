package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Post(url string, param interface{}) (interface{}, bool) {
	return post(url, param, "result")
}

func PostDingTalk(url string, param interface{}) (interface{}, bool) {
	return post(url, param, "")
}

//  headers map[string]string
func post(url string, param interface{}, resultKey string) (interface{}, bool) {
	paramBytes, err := json.Marshal(param)
	if err != nil {
		Logger.Errorf("marshal param error, url = %s, param = %v, err = %s", url, param, err)
		return nil, false
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(paramBytes))
	if err != nil {
		Logger.Errorf("new request error, url = %s, param = %v, err = %s", url, param, err)
		return nil, false
	}

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		Logger.Errorf("post request error, url = %s, param = %v, err = %s", url, param, err)
		return nil, false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Logger.Errorf("read response error, url = %s, param = %v, err = %s", url, param, err)
			return nil, false
		}

		result := make(map[string]interface{})
		decoder := json.NewDecoder(bytes.NewBuffer(body))
		decoder.UseNumber()
		err = decoder.Decode(&result)
		if err != nil {
			Logger.Errorf("decode response error, url = %s, param = %v, err = %s", url, param, string(body))
			return nil, false
		}

		if result["error"] != nil || (result["code"] != nil && result["code"].(json.Number).String() != "0") || (result["errcode"] != nil && result["errcode"].(json.Number).String() != "0") {
			Logger.Errorf("read response error, url = %s, param = %v, err = %s", url, param, string(body))
			return nil, false
		}

		if resultKey != "" {
			return result[resultKey], true
		} else {
			return result, true
		}
	} else {
		Logger.Errorf("response failed, url = %s, param = %v, http_code = %d", url, param, resp.StatusCode)
		return nil, false
	}
}
