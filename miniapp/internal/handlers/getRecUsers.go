// //package handlers
// //
// //import (
// //	"iskra/miniapp/internal/tools/response"/
// 	"iskra/shared/models"
// 	"iskra/shared/storage/postgres"
// 	"net/http"

// 	"github.com/go-chi/render"
// )

// func GetRecUsersHandler(s *postgres.Storage) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// заглушка (ввывод всех пользователей)
// 		users, err := s.UserRepo.GetAll()
// 		if err != nil {
// 			render.JSON(w, r, response.Error("Server error"))
// 			return
// 		}

// 		resp := make([]models.UserResponse, len(users))
// 		for i, user := range users {
// 			resp[i] = models.UserDBToUser(user)
// 		}

//			render.JSON(w, r, struct {
//				response.Response
//				models.ManyUserResponse
//			}{
//				response.Ok(),
//				models.ManyUserResponse{Users: resp},
//			})
//		}
//	}
package handlers
