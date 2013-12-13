package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-gpio"
	"github.com/hybridgroup/gobot-spark"
)

func main() {

	spark := new(gobotSpark.SparkAdaptor)
	spark.Name = "spark"
	spark.Params = map[string]interface{}{
		"device_id":    "",
		"access_token": "",
	}

	button := gobotGPIO.NewButton(spark)
	button.Name = "button"
	button.Pin = "D5"
	button.Interval = "2s"

	led := gobotGPIO.NewLed(spark)
	led.Name = "led"
	led.Pin = "D7"

	work := func() {
		led.Off()
		go func() {
			for {
				fmt.Println("update", gobot.On(button.Events["update"]))
			}
		}()
		go func() {
			for {
				fmt.Println("push", gobot.On(button.Events["push"]))
				led.On()
			}
		}()
		go func() {
			for {
				fmt.Println("release", gobot.On(button.Events["release"]))
				led.Off()
			}
		}()
	}

	robot := gobot.Robot{
		Connections: []interface{}{spark},
		Devices:     []interface{}{button, led},
		Work:        work,
	}

	robot.Start()
}
