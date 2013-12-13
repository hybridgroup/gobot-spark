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

func (me *SparkAdaptor) DigitalWrite(pin string, level string) {
	params := url.Values{
		"params":       {fmt.Sprintf("%v,%v", pin, me.pinLevel(level))},
		"access_token": {me.Params["access_token"].(string)},
	}
	url := fmt.Sprintf("%v/digitalwrite", me.deviceUrl())
	_, err := http.PostForm(url, params)
	if err != nil {
		fmt.Println("Error writing to spark device", me.Name, err)
	}
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
