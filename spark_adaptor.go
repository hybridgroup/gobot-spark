package gobotSpark

import (
	"encoding/json"
	"fmt"
	"github.com/hybridgroup/gobot"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SparkAdaptor struct {
	gobot.Adaptor
}

func (me *SparkAdaptor) Connect() {
}

func (me *SparkAdaptor) Disconnect() {
}

func (me *SparkAdaptor) AnalogRead(pin string) float64 {
	params := url.Values{
		"params":       {pin},
		"access_token": {me.Params["access_token"].(string)},
	}
	url := fmt.Sprintf("%v/analogread", me.deviceUrl())
	resp := me.postToSpark(url, params)
	if resp != nil {
		return resp["return_value"].(float64)
	}
	return 0
}

func (me *SparkAdaptor) AnalogWrite(pin string, level uint8) {
	params := url.Values{
		"params":       {fmt.Sprintf("%v,%v", pin, level)},
		"access_token": {me.Params["access_token"].(string)},
	}
	url := fmt.Sprintf("%v/analogwrite", me.deviceUrl())
	me.postToSpark(url, params)
}

func (me *SparkAdaptor) DigitalWrite(pin string, level string) {
	params := url.Values{
		"params":       {fmt.Sprintf("%v,%v", pin, me.pinLevel(level))},
		"access_token": {me.Params["access_token"].(string)},
	}
	url := fmt.Sprintf("%v/digitalwrite", me.deviceUrl())
	me.postToSpark(url, params)
}

func (me *SparkAdaptor) DigitalRead(pin string) float64 {
	params := url.Values{
		"params":       {pin},
		"access_token": {me.Params["access_token"].(string)},
	}
	url := fmt.Sprintf("%v/digitalread", me.deviceUrl())
	resp := me.postToSpark(url, params)
	if resp != nil {
		return resp["return_value"].(float64)
	}
	return 0
}

func (me *SparkAdaptor) deviceUrl() string {
	return fmt.Sprintf("https://api.spark.io/v1/devices/%v", me.Params["device_id"])
}

func (me *SparkAdaptor) pinLevel(level string) string {
	if level == "1" {
		return "HIGH"
	} else {
		return "LOW"
	}
}

func (me *SparkAdaptor) postToSpark(url string, params url.Values) map[string]interface{} {
	resp, err := http.PostForm(url, params)
	if err != nil {
		fmt.Println(me.Name, "Error writing to spark device", err)
		return nil
	}
	m := make(map[string]interface{})
	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &m)
	if resp.Status != "200 OK" {
		fmt.Println(me.Name, "Error: ", m["error"])
		return nil
	}
	return m
}
