package main

type MessageModel struct {
	Message string `json:"message,omitempty"`
}

type IDUri struct {
	A int `uri:"id" binding:"required"`
}

type NameUri struct {
	A string `uri:"name" binding:"required"`
}

type IdNameUri struct {
	A int    `uri:"id" binding:"required"`
	B string `uri:"name" binding:"required"`
}
