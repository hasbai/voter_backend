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
