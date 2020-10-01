package main

import (
	"fmt"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	g "gobot.io/x/gobot/platforms/dexter/gopigo3"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func robotMainLoop(piProcessor *raspi.Adaptor, gopigo3 *g.Driver, lidarSensor *i2c.LIDARLiteDriver,
	lcd *i2c.GroveLcdDriver) {
	lidarReading, err := lidarSensor.Distance()
	if err != nil {
		fmt.Println("Error reading lidar sensor %+v", err)
	}
	message := fmt.Sprintf("Lidar Reading: %d", lidarReading)
	if message != nil {
		_ = lcd.Write(message)
	} else {
		fmt.Println("error with message")
	}
	time.Sleep(time.Second * 3)

}

func main() {
	raspberryPi := raspi.NewAdaptor()
	gopigo3 := g.NewDriver(raspberryPi)
	lidarSensor := i2c.NewLIDARLiteDriver(raspberryPi)
	lcd := i2c.NewGroveLcdDriver(raspberryPi)
	workerThread := func() {
		robotMainLoop(raspberryPi, gopigo3, lidarSensor, lcd)
	}
	robot := gobot.NewRobot("Gopigo Pi4 Bot",
		[]gobot.Connection{raspberryPi},
		[]gobot.Device{gopigo3, lidarSensor},
		workerThread)

	robot.Start()

}
