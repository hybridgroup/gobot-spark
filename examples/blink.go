package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-spark"
	"github.com/hybridgroup/gobot-gpio"
)

func main() {

	spark := new(gobotSpark.SparkAdaptor)
	spark.Name = "spark"
	spark.Params = make(map[string]interface{})
	spark.Params["device_id"] = "55ff6f064989495346582587"
	spark.Params["access_token"] = "043225bc38f331d9bd965a5e9bdac40ac068d5c2"

	led := gobotGPIO.NewLed(spark)
	led.Name = "led"
	led.Pin = "D7"

	work := func() {
		gobot.Every("1s", func() {
			led.Toggle()
		})
	}

	robot := gobot.Robot{
		Connections: []interface{}{ spark }, 
		Devices:     []interface{}{ led },
		Work:        work,
	}

	robot.Start()
}
