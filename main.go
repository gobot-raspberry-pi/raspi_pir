package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	sensor := gpio.NewPIRMotionDriver(r, "5")
	led := gpio.NewLedDriver(r, "13")

	work := func() {
		sensor.On(gpio.MotionDetected, func(data interface{}) {
			fmt.Println(gpio.MotionDetected)
			led.On()
		})

		sensor.On(gpio.MotionStopped, func(data interface{}) {
			fmt.Println(gpio.MotionStopped)
			led.Off()
		})
	}

	robot := gobot.NewRobot("PIR-motion-bot",
		[]gobot.Connection{r},
		[]gobot.Device{sensor, led},
		work,
	)

	robot.Start()
}
