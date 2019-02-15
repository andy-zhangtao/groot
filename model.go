package main

//type Groot struct {
//	Time string      `json:"time" bson:"time" bw:"time"`
//	Bank []GrootBank `json:"bank" bson:"bank"`
//}
//
type GrootBank struct {
	Kind    int     `json:"kind" bson:"kind"`
	Balance float32 `json:"balance" bson:"balancd"`
	Status  int     `json:"status" bson:"status"`
}

//字段首字母大写后面必须是小写
type Groot struct {
	Time    string  `json:"time" pq:"time" bw:"time"`
	B_icbc  float32 `json:"b_icbc" pq:"b_icbc"`
	B_abc   float32 `json:"b_abc" pq:"b_abc"`
	B_bocom float32 `json:"b_bocom" pq:"b_bocom"`
	B_cmb   float32 `json:"b_cmb" pq:"b_cmb"`
	B_citic float32 `json:"b_citic" pq:"b_citic"`
	B_ccb   float32 `json:"b_ccb" pq:"b_ccb"`
	B_bj    float32 `json:"b_bj" pq:"b_bj"`
	B_ali   float32 `json:"b_ali" pq:"b_ali"`
	B_oth   float32 `json:"b_oth" pq:"b_oth"`
}
