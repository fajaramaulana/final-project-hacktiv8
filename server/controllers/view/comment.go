package view

import "time"

type ResponseCreateComment struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseGetAllComment struct {
	Id        int                        `json:"id"`
	Message   string                     `json:"message"`
	PhotoId   int                        `json:"photo_id"`
	UserId    int                        `json:"user_id"`
	UpdatedAt time.Time                  `json:"updated_at"`
	CreatedAt time.Time                  `json:"created_at"`
	User      ResponseWithUserIdComment  `json:"user"`
	Photo     ResponseWithPhotoIdComment `json:"photo"`
}