package gobotSpark

import (
	"fmt"
	"github.com/hybridgroup/gobot"
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

func (me *SparkAdaptor) AnalogRead(pin string, level string) {
	params := url.Values{
		"params":       {pin},
		"access_token": {me.Params["access_token"].(string)},
	}
	resp := me.postToSpark(fmt.Sprintf("%v/analogread", me.deviceUrl()))
	fmt.Println(resp)
}

func (me *SparkAdaptor) AnalogWrite(pin string, level string) {
	params := url.Values{
		"params":       {fmt.Sprintf("%v,%v", pin, level)},
		"access_token": {me.Params["access_token"].(string)},
	}
	me.postToSpark(fmt.Sprintf("%v/analogwrite", me.deviceUrl()))
}

func (me *SparkAdaptor) DigitalWrite(pin string, level string) {
	params := url.Values{
		"params":       {fmt.Sprintf("%v,%v", pin, me.pinLevel(level))},
		"access_token": {me.Params["access_token"].(string)},
	}
	me.postToSpark(fmt.Sprintf("%v/digitalwrite", me.deviceUrl()))
}

func (me *SparkAdaptor) DigitalRead(pin string, level string) {
	params := url.Values{
		"params":       {pin},
		"access_token": {me.Params["access_token"].(string)},
	}
	resp := me.postToSpark(fmt.Sprintf("%v/digitalread", me.deviceUrl()))
	fmt.Println(resp)
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

func (me *SparkAdaptor) postToSpark(url string) *Response {
	resp, err := http.PostForm(url, params)
	if err != nil {
		fmt.Println("Error writing to spark device", me.Name, err)
	}
	return resp
}
