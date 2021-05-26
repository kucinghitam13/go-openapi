package model

type (
	Person struct {
		ID      int64  `json:"id" example:"2"`
		Name    string `json:"name" example:"dika"`
		Age     int    `json:"age" example:"24"`
		Address string `json:"address" example:"Depok"`
	}
)

// for dto
type (
	GetPersonsResponse struct {
		Total   int      `json:"total" example:"2"`
		Persons []Person `json:"persons"`
	}

	AddPersonRequest struct {
		PersonMutationRequest
	}
	AddPersonResponse struct {
		PersonMutationResponse
	}

	EditPersonRequest struct {
		PersonMutationRequest
	}
	EditPersonResponse struct {
		PersonMutationResponse
	}

	DeletePersonResponse struct {
		PersonMutationResponse
	}

	PersonMutationRequest struct {
		Name    string `json:"name" example:"dika"`
		Age     int    `json:"age" example:"24"`
		Address string `json:"address" example:"Depok"`
	}
	PersonMutationResponse struct {
		Operation string `json:"operation"`
		Success   bool   `json:"success"`
		Person    Person `json:"person"`
	}
)
