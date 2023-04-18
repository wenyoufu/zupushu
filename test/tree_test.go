package test

import (
	"log"
	"testing"
	"treeJiazu/core"
	"treeJiazu/treebuild"
)

func TestPrint(t *testing.T) {
	db, _ := core.OpenDb(treebuild.DBUser, treebuild.DBPass, treebuild.DBIpPort, treebuild.DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温其润", Rank: 2}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root.PrintTree(0)
}

func TestPrintDetail(t *testing.T) {
	db, _ := core.OpenDb(treebuild.DBUser, treebuild.DBPass, treebuild.DBIpPort, treebuild.DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温德普", Rank: 2}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root.PrintDetailByName("温有福")
}

func TestAddChild(t *testing.T) {
	db, _ := core.OpenDb(treebuild.DBUser, treebuild.DBPass, treebuild.DBIpPort, treebuild.DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温德普", Rank: 2}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root.AddChild(db, root, &core.FamilyTree{Name: "温茂景"}, &core.FamilyTree{Name: "温怡宁2", Couple: "AAAA",Rank: 1, Sex: "女"})
	root.PrintTree(0)
}

func TestAddFatherNode(t *testing.T) {
	db, _ := core.OpenDb(treebuild.DBUser, treebuild.DBPass, treebuild.DBIpPort, treebuild.DBName)
	defer db.Close()
	addRoot := &core.FamilyTree{Name: "温其润", Rank: 1}
	oldRoot := &core.FamilyTree{Name: "温学森", Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", oldRoot.Name).Preload("Children").Find(&oldRoot).Error; err != nil {
		log.Fatal(err)
	}
	addRoot.AddChild(db,oldRoot,addRoot,oldRoot)
	addRoot.ConstructTree(db)
	addRoot.PrintTree(0)
}

func TestDelete(t *testing.T) {
	db, _ := core.OpenDb(treebuild.DBUser, treebuild.DBPass, treebuild.DBIpPort, treebuild.DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温德普", Rank: 2}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root.Delete(db, "温怡宁2")
	root.PrintTree(0)
}

func TestUpdate(t *testing.T) {
	db, _ := core.OpenDb(treebuild.DBUser, treebuild.DBPass, treebuild.DBIpPort, treebuild.DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温德普", Rank: 2}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root.UpdateCoupleByName(db, "温怡宁2", []string{"1111"})
	root.UpdateReMarkByName(db, "温怡宁2", "测试111")
	root.PrintTree(0)
}


