package dto

type CaServer struct {
	Name string
}

type CaServerDto struct {
	ID   int
	Name string
}

// type ReqCreateCaServer struct {
// 	Name string `uri:"name" binding:"required"`
// }

type ReqCreateCaServer struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type ReqUserCaServer struct {
	Name string `form:"name" json:"name" binding:"required"`
}
