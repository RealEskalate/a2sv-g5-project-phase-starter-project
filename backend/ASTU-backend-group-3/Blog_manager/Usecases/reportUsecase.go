package Usecases

import (
	"ASTU-backend-group-3/Blog_manager/Domain"
	"errors"
)

func (uc *blogUsecase) ReportBlog(blogID string , report Domain.Report) error{

	if uc.blogRepo.IsUserReported(report.UserId , report.BlogId){
		return errors.New("user reported")
	}
	return uc.blogRepo.AddReport(blogID, report)
	
}