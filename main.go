package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	sensor := gpio.NewPIRMotionDriver(r, "11")
	led := gpio.NewLedDriver(r, "7")
	work := func() {
		sensor.On(gpio.MotionDetected, func(data interface{}) {
			fmt.Printf("Motion Detected at: %v\n", time.Now().Format(time.StampMilli))
			led.On()
		})

		sensor.On(gpio.MotionStopped, func(data interface{}) {
			fmt.Printf("Motion Stopped at: %v\n", time.Now().Format(time.StampMilli))
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
