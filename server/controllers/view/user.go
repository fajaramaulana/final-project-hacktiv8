package view

type ResponseWithUserId struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ResponseWithUserIdComment struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ResponseWithUserIdSocmed struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}
