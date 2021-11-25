package eight_test

import (
	"studyGo/day12/eight"
	"testing"
)

// 单元测试 testing测试框架
func TestStore(t *testing.T) {
	monster := eight.Monster{
		Name:  "孙悟空",
		Age:   600,
		Skill: "金刚不坏",
	}
	monster.Store()
}

func TestRestore(t *testing.T) {
	var monster eight.Monster
	monster.Restore()
	if monster.Name != "悟空" {
		t.Fatalf("希望是:%v,实际是:%v", "悟空", monster.Name)
	}
	t.Logf("测试成功")
}
