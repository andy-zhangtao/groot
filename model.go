package main

type Groot struct {
	Time string    `json:"time" bson:"time" bw:"time"`
	Bank []GrootBank `json:"bank" bson:"bank"`
}

type GrootBank struct {
	Kind    int     `json:"kind" bson:"kind"`
	Balance float32 `json:"balance" bson:"balancd"`
	Status  int     `json:"status" bson:"status"`
}

