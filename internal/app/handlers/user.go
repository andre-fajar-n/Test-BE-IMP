package handlers

import (
	"context"
	"math"
	"test-be-IMP/gen/models"
	"test-be-IMP/gen/restapi/operations/user"
)

// GetListUserWithPagination handler contains get all user with pagination
func (h *handler) GetListUserWithPagination(ctx context.Context, params user.ListUserParams) (*user.ListUserOKBody, error) {
	var users []*models.User
	db := h.db.Model(&users)

	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return nil, err
	}

	err = db.Limit(int(*params.PerPage)).Offset(int((*params.Page - int64(1)) * *params.PerPage)).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &user.ListUserOKBody{
		Metadata: &models.Metadata{
			Page:      *params.Page,
			PerPage:   *params.PerPage,
			TotalRow:  count,
			TotalPage: int64(math.Ceil(float64(count) / float64(*params.PerPage))),
		},
		Data:    users,
		Message: "Success get list user",
	}, nil
}
