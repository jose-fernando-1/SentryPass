package controllers

type EntryController struct {
}

func NewEntryController() *EntryController {
	return &EntryController{}
}

func (ec *EntryController) CreateEntry() {
	// calls utils to encrypt data and generate a random salt
	// calls repository to create an entry in the database
	// calls logs controller to log the action
}
