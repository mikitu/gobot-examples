package main

import (
        "time"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/raspi"
        "log"
        "flag"
)

func main() {
        pin1 := flag.String("pin1", "11", "pin 1")
        pin2 := flag.String("pin2", "12", "pin 2")
        flag.Parse()

        r := raspi.NewAdaptor()
        ledred := gpio.NewLedDriver(r, *pin1)
        ledgreen := gpio.NewLedDriver(r, *pin2)
        work := func() {
                ledgreen.Off()
                ledred.On()
                dur := 1 * time.Second
                gobot.Every(dur, func() {
                        err := ledred.Toggle()
                        if err != nil {
                                log.Println(err)
                        }
                })
                dur = 2 * time.Second
                gobot.Every(dur, func() {
                        err := ledgreen.Toggle()
                        if err != nil {
                                log.Println(err)
                        }
                })
        }

        robot := gobot.NewRobot("blinkBot",
                []gobot.Connection{r},
                []gobot.Device{ledred,ledgreen},
                work,
        )
        ledred.Start()
        ledgreen.Start()

        robot.Start()
}
