package models

import "time"

// todo anything else zine needs?
// think about this further or research but think its fine to store file metadata on the zine object
type Zine struct {
	Id string `json:"id" firestore:"id,omitempty"`
	//TODO this file url will be dependent on a file upload service
	FileUrl   string    `json:"fileUrl" firestore:"fileUrl,omitempty"`
	CreatedOn time.Time `json:"createdOn" firestore:"createdOn,omitempty"`
	Author    string    `json:"author" firestore:"author,omitempty"`
}
