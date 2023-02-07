// @Author: Ciusyan 2023/2/6
package impl

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Go-To-Byte/DouSheng/user_center/apps/token"
)

// 获取Token
func (s *tokenServiceImpl) get(ctx context.Context, ak string) (*token.Token, error) {

	// 查询条件
	filter := bson.M{}
	filter["_id"] = ak

	dk := token.NewDefaultToken()
	// 查询并且反序列化
	if err := s.col.FindOne(ctx, filter).Decode(dk); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("不存在Token")
		}
		return nil, fmt.Errorf("获取Token[%s]失败：%s", ak, err.Error())
	}

	return dk, nil
}

// 刷新Token
func (s *tokenServiceImpl) update(ctx context.Context, tk *token.Token) error {

	if _, err := s.col.UpdateByID(ctx, tk.AccessToken, bson.M{"$set": tk}); err != nil {
		return fmt.Errorf("刷新Token[%s]失败：%s", tk.AccessToken, err.Error())
	}
	return nil
}
