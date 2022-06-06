package wxwork

import "context"

var (
	accessTokenKey = struct{}{}
)

func ContextWithAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, accessTokenKey, token)
}

func AccessTokenFromContext(ctx context.Context) string {
	tok := ctx.Value(accessTokenKey)
	if tok != nil {
		return tok.(string)
	}
	return ""
}
