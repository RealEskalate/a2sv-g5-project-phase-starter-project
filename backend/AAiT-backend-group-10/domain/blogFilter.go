package domain

type BlogFilter struct {
	Author          string 
	Tags            []string   
	SortBy			string     
	Page            int        
	PageSize        int       
}