package tutorials

import "fmt"

type Vehicle interface {
	GetVehicleDescription() string
	GetVehicleParts() string
}

type Bicicleta struct {
	ModelName string
	Year      string
	Parts     []string
}

type Carro struct {
	ModelName string
	Year      string
	Parts     []string
}

// Bicycle implements Vehicle interface
func (b Bicicleta) GetVehicleDescription() string {
	return "Type: Bicycle --- Name: " + b.ModelName + " Year: " + b.Year
}

// This is a dummy interface example; there is no added utility by making additional function of the Vehicle interface
// Just showing that all defined methods on an imterface must be implemented
func (b Bicicleta) GetVehicleParts() (parts string) {
	parts = "Bicycle Parts\t"
	for _, part := range b.Parts {
		parts = parts + part + " "
	}
	return
}

// Car implements Vehicle interface
func (c Carro) GetVehicleDescription() string {
	return "Type: Car --- Name: " + c.ModelName + " Year: " + c.Year
}

func (c Carro) GetVehicleParts() (parts string) {
	parts = "Car Parts\t"
	for _, part := range c.Parts {
		parts = parts + part + " "
	}
	return
}

func InterfaceWithMultipleMethodsDemo() {
	mi_bicicleta := Bicicleta{ModelName: "Giant Anthem 25", Year: "2025", Parts: []string{"Saddle", "Derailer", "Shifter", "Crank", "Fork"}}
	mi_carro := Carro{ModelName: "Hyundai X24", Year: "2025", Parts: []string{"Engine", "Carburator", "Bonnet", "Windshield"}}

	var vehicle Vehicle
	vehicle = mi_bicicleta
	fmt.Println(vehicle.GetVehicleDescription())

	vehicle = mi_carro
	fmt.Println(vehicle.GetVehicleParts())

}
