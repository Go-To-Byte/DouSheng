// @Author: Ciusyan 2023/3/4
package impl

const (
	// 查询朋友的SQL
	friendSql = `
			SELECT
				u2.user_id,
				u2.follower_id,
				u2.follower_flag
			FROM
				( user_follower u1, user_follower u2 )
			WHERE
				u1.user_id = u2.follower_id
				AND u1.follower_id = u2.user_id
				AND u1.follower_flag = ?
				AND u2.follower_flag = ?
				AND u1.user_id = ?;
`
)
