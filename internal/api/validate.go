package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"nursing_api/internal/common/response"
	"nursing_api/pkg/jwt"
)

func passwordValidator(fl validator.FieldLevel) bool {
	return fl.Field().Len() >= 4 && fl.Field().Len() <= 20
}
func keywordValidator(fl validator.FieldLevel) bool {
	return fl.Field().Len() >= 2
}

func getUserId(jwt *jwt.JwtClient, ctx *fiber.Ctx) (uuid.UUID, error) {
	claims, err := jwt.Parser.ExtractTokenMetadata(ctx)
	if err != nil {
		return uuid.UUID{}, response.New(response.CODE_INVALID_JWT).
			SetErrors(err).
			Error(ctx)
	}

	userId := claims.UserID
	return userId, nil
}
