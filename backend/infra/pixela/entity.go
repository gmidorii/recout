package pixela

type Question string

const (
	Yes Question = "yes"
	No  Question = "no"
)

type User struct {
	Token               string   `validate:"required,gte=8,lte=32" json:"token"`
	UserName            string   `validate:"required,gte=1,lte=32" json:"username"`
	AgreeTermsOfService Question `validate:"required" json:"agreeTermsOfService"`
	NotMinor            Question `validate:"required" json:"notMinor"`
}

type PostResponse struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}

type Graph struct {
	ID             string `json:"id" validate:"required,gte=1,lte=8"`
	Name           string `json:"name" validate:"required"`
	Unit           string `json:"unit" validate:"required"`
	Type           string `json:"type" validate:"required"`
	Color          string `json:"color" validate:"required"`
	Timezone       string `json:"timezone"`
	SelfSufficient string `json:"selfSufficient"`
}
