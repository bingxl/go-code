package datastruct

import "testing"

func TestSqlList(t *testing.T) {

	t.Run("创建测试", func(t *testing.T) {
		sq := NewSqlList(10)
		hasErr := false
		for i := 0; i < 10; i++ {
			err := sq.Insert(i, i+3)
			if err != nil {
				hasErr = true
				t.Error("Insert 出现失败" + err.Error())
			}
		}

		if !hasErr {
			t.Log("顺序表创建和insert成功", sq)
		}

		err := sq.Insert(1, 9)
		if err.Error() != "顺序表已满" {
			t.Error("顺序表已满时插入返回错误不正确")
		} else {
			t.Log("顺序表已满时插入正确返回错误信息")
		}

	})
}
