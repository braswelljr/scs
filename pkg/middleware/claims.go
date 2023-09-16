package middleware

import (
	"context"
)

// GetVerifiedClaims - is a function that handles the retrieval of verified claims.
//
//	@param ctx - context.Context
//	@param role - string
//	@return *SignedParams
//	@return error
func GetVerifiedClaims(ctx context.Context, role string) (*SignedParams, error) {
	// get the claims from the context
	claims := ctx.Value(ContextKey).(*SignedParams)
	if claims == nil {
		return &SignedParams{}, ErrorInvalidToken
	}

	// check if the user has the required role
	if role != "" && !claims.HasRole(role) {
		return &SignedParams{}, ErrorUnauthorizedUser
	}

	return claims, nil
}
