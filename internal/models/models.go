package models

type ProblemDetails struct {
    Type     string `json:"type,omitempty"`     
    Title    string `json:"title,omitempty"`    
    Status   int    `json:"status,omitempty"`   
    Detail   string `json:"detail,omitempty"`   
}