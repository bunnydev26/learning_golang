package main

import "fmt"

type car struct {
	gas_pedal uint16 //min 0 max 65535
	break_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}


func main() {
	a_car := car{
					gas_pedal: 22341,
					break_pedal: 38341,
					steering_wheel: 18341,
					top_speed_kmh: 255.0,
			}
	fmt.Println(a_car.gas_pedal, a_car.top_speed_kmh)
}