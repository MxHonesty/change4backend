package geo

type Centru struct {
	Name      string  `json:"name" bson:"name"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
	Scale     float64 `json:"scale" bson:"scale"`
}
