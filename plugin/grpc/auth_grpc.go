package appgrpc

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"fooddelivery-order-service/common"
	auth "fooddelivery-order-service/proto/authen"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

type authClient struct {
	prefix      string
	url         string
	gwSupported bool
	gwPort      int
	client      auth.AuthServiceClient
}

func NewAuthClient(prefix string) *authClient {
	return &authClient{
		prefix: prefix,
	}
}

func (uc *authClient) GetPrefix() string {
	return uc.prefix
}

func (uc *authClient) Get() interface{} {
	return uc
}

func (uc *authClient) Name() string {
	return uc.prefix
}

func (uc *authClient) InitFlags() {
	flag.StringVar(&uc.url, uc.GetPrefix()+"-url", "localhost:50051", "URL connect to grpc server")
}

func (uc *authClient) Configure() error {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial(uc.url, opts)

	if err != nil {
		return err
	}

	uc.client = auth.NewAuthServiceClient(cc)

	return nil
}

func (uc *authClient) Run() error {
	return uc.Configure()
}

func (uc *authClient) Stop() <-chan bool {
	c := make(chan bool)

	go func() {
		c <- true
	}()
	return c
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func (uc *authClient) RequiredAuth(sc goservice.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		aRes, err := uc.client.MiddlewareAuthorize(c.Request.Context(), &auth.AuthRequest{Token: token})

		if err != nil {
			panic(common.ErrNoPermission(err))
		}

		user := aRes.User

		if user.Status == 0 || user.Id <= 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		c.Set(common.CurrentUser, &common.User{
			Id:    int(user.Id),
			Email: user.Email,
			Role:  "",
		})

		c.Next()
	}
}

func (uc *authClient) ValidateToken(token string) (*common.User, error) {
	aRes, err := uc.client.MiddlewareAuthorize(context.Background(), &auth.AuthRequest{Token: token})

	if err != nil {
		return nil, common.ErrNoPermission(err)
	}

	user := aRes.User

	if user.Status == 0 || user.Id <= 0 {
		panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
	}

	return &common.User{
		Id:    int(user.Id),
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
