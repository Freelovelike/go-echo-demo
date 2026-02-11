package auth_service

import (
	"errors"
	"fmt"
	"time"

	"go-echo-demo/internal/dto"
	"go-echo-demo/internal/model"
	"go-echo-demo/internal/vo"
	"go-echo-demo/pkg"
	"go-echo-demo/pkg/db"

	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

func UserLoginService(c *echo.Context, dto dto.LoginAndRegisterDto) (vo.LoginAndRegisterVo, error) {
	var user model.User

	// 1. 只根据用户名查找，不带密码条件
	result := db.DB.Where("username = ?", dto.Username).First(&user)

	// 2. 检查数据库查询是否出错（比如用户不存在）
	if result.Error != nil {
		return vo.LoginAndRegisterVo{}, errors.New("用户名或密码错误") // 建议模糊提示，更安全
	}

	// 对比密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		// 报错说明密码对不上
		return vo.LoginAndRegisterVo{}, errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := pkg.GenerateToken(user.ID)
	if err != nil {
		return vo.LoginAndRegisterVo{}, errors.New("生成token失败")
	}

	// 将token存入redis
	redisKey := fmt.Sprintf("auth:token:%d", user.ID)
	err = db.Redis.Set(c.Request().Context(), redisKey, token, 72*time.Hour).Err()
	if err != nil {
		// 如果 Redis 挂了，这里可以选择记录日志并继续，或者直接报错
		fmt.Printf("Redis 写入失败: %v\n", err)
		return vo.LoginAndRegisterVo{}, errors.New("服务器内部错误")
	}
	return vo.LoginAndRegisterVo{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}, nil
}

func UserRegisterService(dto dto.LoginAndRegisterDto) (*model.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := model.User{
		Username: dto.Username,
		Password: string(password),
	}
	fmt.Println("user", user, "dto", dto)
	result := db.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
