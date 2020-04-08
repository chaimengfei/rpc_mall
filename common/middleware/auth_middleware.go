package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"github.com/go-redis/redis"
	"time"
)

var ErrAuthorization = errors.New("Please Login 请先登录")
const (
	AuthorizationExpire = 604800 * time.Second
)
type Authoriztion struct {
	redisCache *redis.Client
}
func NewAuthorization( r *redis.Client) *Authoriztion {
	return &Authoriztion{redisCache:r}
}

func (a *Authoriztion) Auth(c *gin.Context)  {
	token := c.GetHeader("Authorization")
	if strings.TrimSpace(token) == ""{
		c.JSON(500,"auth failed")
		c.Abort()
	}
	src:=a.redisCache.Get(token)
	userId,_:=strconv.Atoi(src.Val())
	c.Set("userId",userId)
	c.Next()
}