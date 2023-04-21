package treebuild

import (
	"github.com/jinzhu/gorm"
	"log"
	"treeJiazu/core"
)

const (
	DBUser   = "root"
	DBPass   = "Abc@123456"
	DBName   = "db_tree"
	DBIpPort = "127.0.0.1:3306"
)

func GetRootByName(name string) (*gorm.DB, *core.FamilyTree) {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	root := &core.FamilyTree{Name: name, Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err == nil {
		log.Printf("已存在 %v", root.Name)
		root.ConstructTree(db)
		return db, root
	}
	return db, nil
}

// BuildL1To5Tree 一到五辈的祖宗树建立
func BuildL1To5Tree() *core.FamilyTree {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温友贵", Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err == nil {
		log.Printf("已存在，无需创建 %v", root.Name)
		root.ConstructTree(db)
		return root
	}
	db.Create(root)
	root.ConstructTree(db)
	root.UpdateCoupleByName(db, "温友贵", []string{"孙氏", "裴氏"})

	root1 := &core.FamilyTree{Name: "温大智", Rank: 1}
	root.AddChild(db, root, root, root1)
	root.UpdateCoupleByName(db, "温大智", []string{"杨氏", "李氏"})

	root2 := &core.FamilyTree{Name: "温自富", Rank: 1}
	root.AddChild(db, root, root1, root2)
	root.UpdateCoupleByName(db, "温自富", []string{"薛氏"})

	root3 := &core.FamilyTree{Name: "温成", Rank: 1}
	root.AddChild(db, root, root2, root3)
	root.UpdateCoupleByName(db, "温成", []string{"葛氏"})

	root4_1 := &core.FamilyTree{Name: "温贵德", Rank: 1}
	root.AddChild(db, root, root3, root4_1)
	root.UpdateCoupleByName(db, "温贵德", []string{"裴氏"})

	root4_2 := &core.FamilyTree{Name: "温贵生", Rank: 2}
	root.AddChild(db, root, root3, root4_2)
	root.UpdateCoupleByName(db, "温贵生", []string{"杨氏"})
	root4_3 := &core.FamilyTree{Name: "温贵全", Rank: 3}
	root.AddChild(db, root, root3, root4_3)
	root.UpdateCoupleByName(db, "温贵全", []string{"胡氏"})
	root4_4 := &core.FamilyTree{Name: "温贵天", Rank: 4}
	root.AddChild(db, root, root3, root4_4)
	root.UpdateCoupleByName(db, "温贵天", []string{"张氏", "王氏"})
	root4_5 := &core.FamilyTree{Name: "温贵勤", Rank: 5}
	root.AddChild(db, root, root3, root4_5)
	root.UpdateCoupleByName(db, "温贵勤", []string{"张氏"})
	root4_6 := &core.FamilyTree{Name: "温贵检", Rank: 6}
	root.AddChild(db, root, root3, root4_6)
	root.UpdateCoupleByName(db, "温贵检", []string{"孙氏"})
	root4_7 := &core.FamilyTree{Name: "温贵发", Rank: 7}
	root.AddChild(db, root, root3, root4_7)
	root.UpdateCoupleByName(db, "温贵发", []string{"王氏"})

	return root
}

func BuildL5To6Tree() *core.FamilyTree {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温友贵", Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Println("不存在 BuildL5To9Tree", root.Name)
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root1 := root.FindByName("温贵德")
	root2 := root.FindByName("温贵生")
	root3 := root.FindByName("温贵全")
	root4 := root.FindByName("温贵天")
	root5 := root.FindByName("温贵勤")
	root6 := root.FindByName("温贵检")
	root7 := root.FindByName("温贵发")
	// 第六层
	//找到第六层则无需创建
	root1_1 := &core.FamilyTree{Name: "温克祥", Rank: 1}
	if err := db.Where("name = ?", "温克祥").Preload("Children").Find(&root1_1).Error; err == nil {
		log.Println("已存在第六层，无需创建 ")
		return root
	}

	root.AddChild(db, root, root1, root1_1)
	root.UpdateCoupleByName(db, "温克祥", []string{"钱氏"})
	root1_2 := &core.FamilyTree{Name: "温克焕", Rank: 2}
	root.AddChild(db, root, root1, root1_2)
	root.UpdateCoupleByName(db, "温克焕", []string{"蒋氏"})
	root2_1 := &core.FamilyTree{Name: "温克仓", Rank: 1}
	root.AddChild(db, root, root2, root2_1)
	root.UpdateCoupleByName(db, "温克仓", []string{"李氏"})
	root2_2 := &core.FamilyTree{Name: "温克有", Rank: 2}
	root.AddChild(db, root, root2, root2_2)
	root.UpdateCoupleByName(db, "温克有", []string{"刘氏"})
	root2_3 := &core.FamilyTree{Name: "温克积", Rank: 3}
	root.AddChild(db, root, root2, root2_3)
	root.UpdateCoupleByName(db, "温克积", []string{"易氏", "马氏"})

	root3_1 := &core.FamilyTree{Name: "温克惠", Rank: 1}
	root.AddChild(db, root, root3, root3_1)
	root.UpdateCoupleByName(db, "温克惠", []string{"马氏"})

	root4_1 := &core.FamilyTree{Name: "温克恭", ReMark: "承继", Rank: 1}
	root.AddChild(db, root, root4, root4_1)
	root.UpdateCoupleByName(db, "温克恭", []string{"任氏"})
	root5_1 := &core.FamilyTree{Name: "温克忠", Rank: 1}
	root.AddChild(db, root, root5, root5_1)
	root.UpdateCoupleByName(db, "温克忠", []string{"周氏", "王氏", "常氏"})
	root5_2 := &core.FamilyTree{Name: "温克典", Rank: 2}
	root.AddChild(db, root, root5, root5_2)
	root.UpdateCoupleByName(db, "温克典", []string{"李氏", "王氏"})
	root6_1 := &core.FamilyTree{Name: "温克良", Rank: 1}
	root.AddChild(db, root, root6, root6_1)
	root.UpdateCoupleByName(db, "温克良", []string{"邢氏"})
	root6_2 := &core.FamilyTree{Name: "温克恭（过继）", ReMark: "过继", Rank: 2}
	root.AddChild(db, root, root6, root6_2)
	root.UpdateCoupleByName(db, "温克恭（过继）", []string{"马氏"})
	root7_1 := &core.FamilyTree{Name: "温克让", Rank: 1}
	root.AddChild(db, root, root7, root7_1)
	root.UpdateCoupleByName(db, "温克让", []string{"邢氏"})
	return root
}
func BuildL6To7Tree() *core.FamilyTree {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温友贵", Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Println("不存在 BuildL7To8Tree ", root.Name)
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root1 := root.FindByName("温克祥")
	root2 := root.FindByName("温克焕")
	root3 := root.FindByName("温克仓")
	root4 := root.FindByName("温克积")
	root5 := root.FindByName("温克惠")
	root6 := root.FindByName("温克恭")
	root7 := root.FindByName("温克忠")
	root8 := root.FindByName("温克典")
	root9 := root.FindByName("温克良")
	root10 := root.FindByName("温克让")

	// 第七层
	//找到第七层则无需创建
	root1_1 := &core.FamilyTree{Name: "温国栋", ReMark: "承继", Rank: 1}
	if err := db.Where("name = ?", "温国栋").Preload("Children").Find(&root1_1).Error; err == nil {
		log.Println("已存在第七层，无需创建 ")
		return root
	}

	root.AddChild(db, root, root1, root1_1)
	root.UpdateCoupleByName(db, "温国栋", []string{"韩氏", "侯氏"})
	root2_1 := &core.FamilyTree{Name: "温国延", Rank: 1}
	root.AddChild(db, root, root2, root2_1)
	root.UpdateCoupleByName(db, "温国延", []string{"李氏"})
	root2_2 := &core.FamilyTree{Name: "温国用", Rank: 2}
	root.AddChild(db, root, root2, root2_2)
	root.UpdateCoupleByName(db, "温国用", []string{"谷氏", "李氏", "韩氏"})
	root2_3 := &core.FamilyTree{Name: "温国栋(过继)", ReMark: "过继", Rank: 3}
	root.AddChild(db, root, root2, root2_3)
	root.UpdateCoupleByName(db, "温国栋(过继)", []string{"韩氏", "侯氏"})

	root3_1 := &core.FamilyTree{Name: "温国玺", Rank: 1}
	root.AddChild(db, root, root3, root3_1)
	root.UpdateCoupleByName(db, "温国玺", []string{"郭氏"})

	root4_1 := &core.FamilyTree{Name: "温国富", Rank: 1}
	root.AddChild(db, root, root4, root4_1)
	root.UpdateCoupleByName(db, "温国富", []string{"张氏", "李氏"})

	root5_1 := &core.FamilyTree{Name: "温国尚", Rank: 1}
	root.AddChild(db, root, root5, root5_1)
	root.UpdateCoupleByName(db, "温国尚", []string{"杨氏", "习氏"})

	root6_1 := &core.FamilyTree{Name: "温国玺（过继）", Rank: 1}
	root.AddChild(db, root, root6, root6_1)
	root.UpdateCoupleByName(db, "温国玺（过继）", []string{"郭氏"})
	root6_2 := &core.FamilyTree{Name: "温国印", Rank: 2}
	root.AddChild(db, root, root6, root6_2)
	root.UpdateCoupleByName(db, "温国印", []string{"王氏"})
	root7_1 := &core.FamilyTree{Name: "温国辅", ReMark: "承继", Rank: 1}
	root.AddChild(db, root, root7, root7_1)
	root.UpdateCoupleByName(db, "温国辅", []string{"郭氏", "薛式"})

	root8_1 := &core.FamilyTree{Name: "温国辅(过继)", ReMark: "过继", Rank: 1}
	root.AddChild(db, root, root8, root8_1)
	root.UpdateCoupleByName(db, "温国辅(过继)", []string{"郭氏", "薛式"})
	root8_2 := &core.FamilyTree{Name: "温国相", Rank: 2}
	root.AddChild(db, root, root8, root8_2)
	root.UpdateCoupleByName(db, "温国相", []string{"李氏"})

	root9_1 := &core.FamilyTree{Name: "温国桢", Rank: 1}
	root.AddChild(db, root, root9, root9_1)
	root.UpdateCoupleByName(db, "温国桢", []string{"雷式"})
	root10_1 := &core.FamilyTree{Name: "温国泰", Rank: 1}
	root.AddChild(db, root, root10, root10_1)
	root.UpdateCoupleByName(db, "温国泰", []string{"刘氏", "李氏"})
	return root
}

func BuildL7To8Tree() *core.FamilyTree {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温友贵", Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Println("不存在 BuildL6To7Tree ", root.Name)
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root1 := root.FindByName("温国栋")
	root2 := root.FindByName("温国延")
	root3 := root.FindByName("温国用")
	root4 := root.FindByName("温国玺")
	root5 := root.FindByName("温国富")
	root6 := root.FindByName("温国尚")
	root7 := root.FindByName("温国印")
	root8 := root.FindByName("温国辅")
	root9 := root.FindByName("温国相")
	root10 := root.FindByName("温国桢")
	root11 := root.FindByName("温国泰")

	// 第八层
	//找到第八层则无需创建
	root1_1 := &core.FamilyTree{Name: "温萃玉", Rank: 1}
	if err := db.Where("name = ?", "温萃玉").Preload("Children").Find(&root1_1).Error; err == nil {
		log.Println("已存在第八层BuildL7To8Tree，无需创建 ")
		return root
	}

	root.AddChild(db, root, root1, root1_1)
	root.UpdateCoupleByName(db, "温萃玉", []string{"王氏"})
	root1_2 := &core.FamilyTree{Name: "温大玉", Rank: 2}
	root.AddChild(db, root, root1, root1_2)
	root.UpdateCoupleByName(db, "温大玉", []string{"孙氏"})
	root1_3 := &core.FamilyTree{Name: "温如玉", Rank: 3}
	root.AddChild(db, root, root1, root1_3)
	root.UpdateCoupleByName(db, "温如玉", []string{"张氏"})
	root1_4 := &core.FamilyTree{Name: "温式玉", Rank: 4}
	root.AddChild(db, root, root1, root1_4)
	root.UpdateCoupleByName(db, "温式玉", []string{"姜氏"})

	root2_1 := &core.FamilyTree{Name: "温润玉", Rank: 1}
	root.AddChild(db, root, root2, root2_1)
	root.UpdateCoupleByName(db, "温润玉", []string{"冯氏", "雷氏"})
	root2_2 := &core.FamilyTree{Name: "温若玉", Rank: 2}
	root.AddChild(db, root, root2, root2_2)
	root.UpdateCoupleByName(db, "温若玉", []string{"裴氏"})

	root3_1 := &core.FamilyTree{Name: "温成玉", Rank: 1}
	root.AddChild(db, root, root3, root3_1)
	root.UpdateCoupleByName(db, "温成玉", []string{"王氏"})

	root4_1 := &core.FamilyTree{Name: "温宝玉", ReMark: "承继", Rank: 1}
	root.AddChild(db, root, root4, root4_1)
	//root.UpdateCoupleByName(db, "温宝玉", []string{"张氏", "李氏"})

	root5_1 := &core.FamilyTree{Name: "温鸣玉", Rank: 1}
	root.AddChild(db, root, root5, root5_1)
	root.UpdateCoupleByName(db, "温鸣玉", []string{"汪氏"})
	root5_2 := &core.FamilyTree{Name: "温泽玉", Rank: 2}
	root.AddChild(db, root, root5, root5_2)
	root.UpdateCoupleByName(db, "温泽玉", []string{"薛氏"})

	root6_1 := &core.FamilyTree{Name: "温润太", Rank: 1}
	root.AddChild(db, root, root6, root6_1)
	root.UpdateCoupleByName(db, "温润太", []string{"高氏"})
	root7_1 := &core.FamilyTree{Name: "温*玉", ReMark: "三兄弟，记录缺失", Rank: 1}
	root.AddChild(db, root, root7, root7_1)
	//root.UpdateCoupleByName(db, "温*玉", []string{""})
	root7_2 := &core.FamilyTree{Name: "温宝玉（过继）", ReMark: "过继", Rank: 4}
	root.AddChild(db, root, root7, root7_2)
	//root.UpdateCoupleByName(db, "温宝玉（过继）", []string{""})

	root8_1 := &core.FamilyTree{Name: "温佩玉", Rank: 1}
	root.AddChild(db, root, root8, root8_1)
	root.UpdateCoupleByName(db, "温佩玉", []string{"张式"})
	root8_2 := &core.FamilyTree{Name: "温璞玉", Rank: 2}
	root.AddChild(db, root, root8, root8_2)
	root.UpdateCoupleByName(db, "温璞玉", []string{"张氏"})
	root8_3 := &core.FamilyTree{Name: "温兰玉", Rank: 3}
	root.AddChild(db, root, root8, root8_3)
	root.UpdateCoupleByName(db, "温兰玉", []string{"李氏"})

	root9_1 := &core.FamilyTree{Name: "温金玉", Rank: 1}
	root.AddChild(db, root, root9, root9_1)
	root.UpdateCoupleByName(db, "温金玉", []string{"邢式"})
	root9_2 := &core.FamilyTree{Name: "温满玉", Rank: 2}
	root.AddChild(db, root, root9, root9_2)
	root.UpdateCoupleByName(db, "温满玉", []string{"朱式"})
	root9_3 := &core.FamilyTree{Name: "温堂玉", Rank: 3}
	root.AddChild(db, root, root9, root9_3)
	//root.UpdateCoupleByName(db, "温堂玉", []string{})

	root10_1 := &core.FamilyTree{Name: "温振玉", Rank: 1}
	root.AddChild(db, root, root10, root10_1)
	root.UpdateCoupleByName(db, "温振玉", []string{"张氏"})
	root10_2 := &core.FamilyTree{Name: "温怀玉", Rank: 2}
	root.AddChild(db, root, root10, root10_2)
	root.UpdateCoupleByName(db, "温怀玉", []string{"唐氏"})
	root10_3 := &core.FamilyTree{Name: "温瑞玉", Rank: 3}
	root.AddChild(db, root, root10, root10_3)
	root.UpdateCoupleByName(db, "温瑞玉", []string{"杨氏"})

	root11_1 := &core.FamilyTree{Name: "温明玉", Rank: 1}
	root.AddChild(db, root, root11, root11_1)
	root.UpdateCoupleByName(db, "温明玉", []string{"李氏"})
	return root
}

func BuildL8To9Tree() *core.FamilyTree {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温友贵", Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err != nil {
		log.Println("不存在 BuildL8To9Tree ", root.Name)
		log.Fatal(err)
	}
	root.ConstructTree(db)
	root1 := root.FindByName("温萃玉")
	root2 := root.FindByName("温大玉")
	root3 := root.FindByName("温如玉")
	root4 := root.FindByName("温润玉")
	root5 := root.FindByName("温若玉")
	root6 := root.FindByName("温成玉")
	root7 := root.FindByName("温鸣玉")
	root8 := root.FindByName("温泽玉")
	root9 := root.FindByName("温润太")
	root10 := root.FindByName("温兰玉")
	root11 := root.FindByName("温满玉")
	root12 := root.FindByName("温怀玉")
	root13 := root.FindByName("温瑞玉")
	// 第9层
	//找到第9层则无需创建
	root1_1 := &core.FamilyTree{Name: "温其瑞", Rank: 1}
	if err := db.Where("name = ?", "温其瑞").Preload("Children").Find(&root1_1).Error; err == nil {
		log.Println("已存在第9层BuildL8To9Tree，无需创建 ")
		return root
	}

	root.AddChild(db, root, root1, root1_1)
	root.UpdateCoupleByName(db, "温其瑞", []string{"石氏"})

	root2_1 := &core.FamilyTree{Name: "温其润", Rank: 1}
	root.AddChild(db, root, root2, root2_1)
	//root.UpdateCoupleByName(db, "温其润", []string{"冯氏", "雷氏"})

	root3_1 := &core.FamilyTree{Name: "温其光", Rank: 1}
	root.AddChild(db, root, root3, root3_1)
	root.UpdateCoupleByName(db, "温其光", []string{"李氏"})
	root3_2 := &core.FamilyTree{Name: "温其明", Rank: 2}
	root.AddChild(db, root, root3, root3_2)
	root.UpdateCoupleByName(db, "温其明", []string{"张氏"})
	root3_3 := &core.FamilyTree{Name: "温其禄", Rank: 3}
	root.AddChild(db, root, root3, root3_3)
	//root.UpdateCoupleByName(db, "温其圣", []string{"李氏"})
	root3_4 := &core.FamilyTree{Name: "温其圣", Rank: 4}
	root.AddChild(db, root, root3, root3_4)
	root.UpdateCoupleByName(db, "温其圣", []string{"李氏", "王氏"})

	root4_1 := &core.FamilyTree{Name: "温其昌", Rank: 1}
	root.AddChild(db, root, root4, root4_1)
	root.UpdateCoupleByName(db, "温其昌", []string{"孙氏"})
	root4_2 := &core.FamilyTree{Name: "温其永", Rank: 2}
	root.AddChild(db, root, root4, root4_2)
	root.UpdateCoupleByName(db, "温其昌", []string{"李氏"})
	root4_3 := &core.FamilyTree{Name: "温其春", Rank: 3}
	root.AddChild(db, root, root4, root4_3)
	//root.UpdateCoupleByName(db, "温其春", []string{"张氏", "李氏"})

	root5_1 := &core.FamilyTree{Name: "温其允", Rank: 1}
	root.AddChild(db, root, root5, root5_1)
	root.UpdateCoupleByName(db, "温其允", []string{"刘氏"})
	root5_2 := &core.FamilyTree{Name: "温其顺", Rank: 2}
	root.AddChild(db, root, root5, root5_2)
	root.UpdateCoupleByName(db, "温其顺", []string{"杨氏"})

	root6_1 := &core.FamilyTree{Name: "温其凤", Rank: 1}
	root.AddChild(db, root, root6, root6_1)
	//root.UpdateCoupleByName(db, "温其凤", []string{"高氏"})
	root6_2 := &core.FamilyTree{Name: "温其才", Rank: 1}
	root.AddChild(db, root, root6, root6_2)
	root.UpdateCoupleByName(db, "温其才", []string{"张氏"})

	root7_1 := &core.FamilyTree{Name: "温其连", Rank: 1}
	root.AddChild(db, root, root7, root7_1)
	root.UpdateCoupleByName(db, "温其连", []string{"高氏", "翟氏"})
	root7_2 := &core.FamilyTree{Name: "温其璋", Rank: 2}
	root.AddChild(db, root, root7, root7_2)
	root.UpdateCoupleByName(db, "温其璋", []string{"冯氏"})

	root8_1 := &core.FamilyTree{Name: "温其俊", Rank: 1}
	root.AddChild(db, root, root8, root8_1)
	root.UpdateCoupleByName(db, "温其俊", []string{"崔氏"})
	root8_2 := &core.FamilyTree{Name: "温其印", Rank: 2}
	root.AddChild(db, root, root8, root8_2)
	//root.UpdateCoupleByName(db, "温其印", []string{"张氏"})

	root9_1 := &core.FamilyTree{Name: "温其美", Rank: 1}
	root.AddChild(db, root, root9, root9_1)
	root.UpdateCoupleByName(db, "温其美", []string{"王式"})

	root10_1 := &core.FamilyTree{Name: "温其荣", Rank: 1}
	root.AddChild(db, root, root10, root10_1)
	//root.UpdateCoupleByName(db, "温振玉", []string{"张氏"})
	root10_2 := &core.FamilyTree{Name: "温其华", Rank: 2}
	root.AddChild(db, root, root10, root10_2)
	//root.UpdateCoupleByName(db, "温其华", []string{"唐氏"})

	root11_1 := &core.FamilyTree{Name: "温其贵", Rank: 1}
	root.AddChild(db, root, root11, root11_1)
	//root.UpdateCoupleByName(db, "温其贵", []string{"李氏"})

	root12_1 := &core.FamilyTree{Name: "温福全", Rank: 1}
	root.AddChild(db, root, root12, root12_1)
	//root.UpdateCoupleByName(db, "温其贵", []string{"李氏"})

	root13_1 := &core.FamilyTree{Name: "温其惠", Rank: 1}
	root.AddChild(db, root, root13, root13_1)
	root.UpdateCoupleByName(db, "温其惠", []string{"裴氏"})
	return root
}

func BuildQiRunTree(father *core.FamilyTree) {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root00 := &core.FamilyTree{Name: "温学增", Rank: 1}
	root := &core.FamilyTree{Name: "温学森", Rank: 2}
	root02 := &core.FamilyTree{Name: "温学富", Rank: 3}
	if father.FindByName("温学森") != nil {
		log.Printf("BuildQiRunTree已经创建，无需再创建！")
		return
	}
	//db.Create(root)
	father.AddChild(db, father, father, root00)
	father.AddChild(db, father, father, root)
	father.AddChild(db, father, father, root02)
	// 读取数据
	//root.ConstructTree(db)
	father.UpdateCoupleByName(db, "温学森", []string{"孙氏", "孙氏"})
	father.UpdateCoupleByName(db, "温学增", []string{"裴氏"})
	//father.UpdateCoupleByName(db, "温学富", []string{"孙氏"})

	root1 := &core.FamilyTree{Name: "温德厚", Rank: 1}
	root2 := &core.FamilyTree{Name: "温德普", Rank: 2}
	father.AddChild(db, father, root, root1)
	father.AddChild(db, father, root, root2)

	root1_1 := &core.FamilyTree{Name: "温茂春", Rank: 1}
	father.AddChild(db, father, root1, root1_1)
	root2_1 := &core.FamilyTree{Name: "温茂顺", Rank: 1}
	father.AddChild(db, father, root2, root2_1)
	root2_2 := &core.FamilyTree{Name: "温茂清", Rank: 2}
	father.AddChild(db, father, root2, root2_2)
	root2_3 := &core.FamilyTree{Name: "温茂金", Rank: 3}
	father.AddChild(db, father, root2, root2_3)
	root2_4 := &core.FamilyTree{Name: "温茂景", Rank: 4}
	father.AddChild(db, father, root2, root2_4)

	root1_1_1 := &core.FamilyTree{Name: "温树福", Rank: 1}
	father.AddChild(db, father, root1_1, root1_1_1)

	root1_1_2 := &core.FamilyTree{Name: "温树芳", Sex: "女", Rank: 2}
	father.AddChild(db, father, root1_1, root1_1_2)
	root2_1_1 := &core.FamilyTree{Name: "温平福", Rank: 1}
	father.AddChild(db, father, root2_1, root2_1_1)
	root2_2_2 := &core.FamilyTree{Name: "温宝福", Rank: 1}
	father.AddChild(db, father, root2_2, root2_2_2)
	root2_3_1 := &core.FamilyTree{Name: "温有福", Rank: 1}
	father.AddChild(db, father, root2_3, root2_3_1)
	root2_4_1 := &core.FamilyTree{Name: "温怡宁", Sex: "女", Rank: 1}
	father.AddChild(db, father, root2_4, root2_4_1)

	root1_1_1_1 := &core.FamilyTree{Name: "温睿萱", Sex: "女", Rank: 1}
	father.AddChild(db, father, root1_1_1, root1_1_1_1)
	root2_1_1_1 := &core.FamilyTree{Name: "温可心", Sex: "女", Rank: 1}
	father.AddChild(db, father, root2_1_1, root2_1_1_1)

	father.UpdateCoupleByName(db, "温德厚", []string{"张氏"})
	father.UpdateCoupleByName(db, "温德普", []string{"李氏"})
	father.UpdateCoupleByName(db, "温茂春", []string{"王恩凤"})
	father.UpdateCoupleByName(db, "温茂顺", []string{"马香英"})
	father.UpdateCoupleByName(db, "温茂清", []string{"裴印凤"})
	father.UpdateCoupleByName(db, "温茂金", []string{"马秀云"})
	father.UpdateCoupleByName(db, "温茂景", []string{"刘桂玲"})

	father.UpdateCoupleByName(db, "温平福", []string{"王晓梅"})
	father.UpdateCoupleByName(db, "温树福", []string{"孙显文"})
	father.UpdateCoupleByName(db, "温有福", []string{"王芳"})
}

func BuildQiShengTree(father *core.FamilyTree) {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温学义", Rank: 1}
	root00 := &core.FamilyTree{Name: "温学伯", Rank: 2}
	root02 := &core.FamilyTree{Name: "温学武", Rank: 3}
	if father.FindByName("温学义") != nil {
		log.Printf("已经创建 BuildQiShengTree ，无需再创建！")
		return
	}
	//db.Create(root)
	father.AddChild(db, father, father, root00)
	father.AddChild(db, father, father, root)
	father.AddChild(db, father, father, root02)
	// 读取数据
	father.UpdateCoupleByName(db, "温学义", []string{"无氏"})
	father.UpdateCoupleByName(db, "温学伯", []string{"张氏"})
	//father.UpdateCoupleByName(db, "温学武", []string{"薄氏"})

}
func BuildQihuiTree(father *core.FamilyTree) {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温学朗", Rank: 1}
	root00 := &core.FamilyTree{Name: "温学谦", Rank: 2}
	root02 := &core.FamilyTree{Name: "温学讲", Rank: 3}
	if father.FindByName("温学朗") != nil {
		log.Printf("已经创建 BuildQihuiTree ，无需再创建！")
		return
	}
	//db.Create(root)
	father.AddChild(db, father, father, root00)
	father.AddChild(db, father, father, root)
	father.AddChild(db, father, father, root02)
	// 读取数据
	father.UpdateCoupleByName(db, "温学朗", []string{"汪氏"})
	father.UpdateCoupleByName(db, "温学谦", []string{"王氏"})
	father.UpdateCoupleByName(db, "温学讲", []string{"薄氏"})

	root1 := &core.FamilyTree{Name: "温德修", Rank: 1}
	root2 := &core.FamilyTree{Name: "温德森", Rank: 2}
	root3 := &core.FamilyTree{Name: "温德新", Rank: 3}
	root4 := &core.FamilyTree{Name: "温德安", Rank: 4}
	root5 := &core.FamilyTree{Name: "温德荣", Rank: 5}
	root6 := &core.FamilyTree{Name: "温德祥", Rank: 6}
	father.AddChild(db, father, root, root1)
	father.AddChild(db, father, root, root2)
	father.AddChild(db, father, root, root3)
	father.AddChild(db, father, root, root4)
	father.AddChild(db, father, root, root5)
	father.AddChild(db, father, root, root6)
	father.UpdateCoupleByName(db, "温德修", []string{"王氏"})
	father.UpdateCoupleByName(db, "温德森", []string{"张氏"})
	father.UpdateCoupleByName(db, "温德新", []string{"薄氏"})
	father.UpdateCoupleByName(db, "温德安", []string{"裴氏"})
	father.UpdateCoupleByName(db, "温德荣", []string{"裴氏"})
	father.UpdateCoupleByName(db, "温德祥", []string{"汪氏"})

	root1_1 := &core.FamilyTree{Name: "温茂福", Rank: 1}
	father.AddChild(db, father, root1, root1_1)
	root1_2 := &core.FamilyTree{Name: "温茂禄", Rank: 2}
	father.AddChild(db, father, root1, root1_2)
	root1_3 := &core.FamilyTree{Name: "温茂贵", Rank: 3}
	father.AddChild(db, father, root1, root1_3)

	root2_1 := &core.FamilyTree{Name: "温茂圣", Rank: 1}
	father.AddChild(db, father, root2, root2_1)
	root2_2 := &core.FamilyTree{Name: "温茂永", Rank: 2}
	father.AddChild(db, father, root2, root2_2)
	root2_3 := &core.FamilyTree{Name: "温茂来", Rank: 3}
	father.AddChild(db, father, root2, root2_3)

	root3_1 := &core.FamilyTree{Name: "温茂立", Rank: 1}
	father.AddChild(db, father, root3, root3_1)
	root3_2 := &core.FamilyTree{Name: "温茂志", Rank: 2}
	father.AddChild(db, father, root3, root3_2)

	root4_1 := &core.FamilyTree{Name: "温茂海", Rank: 1}
	father.AddChild(db, father, root4, root4_1)

	root5_1 := &core.FamilyTree{Name: "温茂军", Rank: 1}
	father.AddChild(db, father, root5, root5_1)
	root5_2 := &core.FamilyTree{Name: "温茂岭", Rank: 2}
	father.AddChild(db, father, root5, root5_2)

	root6_1 := &core.FamilyTree{Name: "温茂凯", Rank: 1}
	father.AddChild(db, father, root6, root6_1)

	root1_1_1 := &core.FamilyTree{Name: "温晓旭", Rank: 1}
	father.AddChild(db, father, root1_1, root1_1_1)
	root1_2_1 := &core.FamilyTree{Name: "温洁", Sex: "女", Rank: 1}
	father.AddChild(db, father, root1_2, root1_2_1)
	root1_3_1 := &core.FamilyTree{Name: "温晓龙", Rank: 1}
	father.AddChild(db, father, root1_3, root1_3_1)
	root2_1_1 := &core.FamilyTree{Name: "温彬", Rank: 1}
	father.AddChild(db, father, root2_1, root2_1_1)
	root2_2_1 := &core.FamilyTree{Name: "温福龙", Rank: 1}
	father.AddChild(db, father, root2_2, root2_2_1)
	root2_3_1 := &core.FamilyTree{Name: "温浩然", Rank: 1}
	father.AddChild(db, father, root2_3, root2_3_1)
	root3_1_1 := &core.FamilyTree{Name: "温涛", Rank: 1}
	father.AddChild(db, father, root3_1, root3_1_1)
	root3_2_1 := &core.FamilyTree{Name: "温宝龙", Rank: 1}
	father.AddChild(db, father, root3_2, root3_2_1)
	root4_1_1 := &core.FamilyTree{Name: "温强", Rank: 1}
	father.AddChild(db, father, root4_1, root4_1_1)
	root5_1_1 := &core.FamilyTree{Name: "温超", Rank: 1}
	father.AddChild(db, father, root5_1, root5_1_1)
	root5_2_1 := &core.FamilyTree{Name: "温帅", Rank: 1}
	father.AddChild(db, father, root5_2, root5_2_1)
	root6_1_1 := &core.FamilyTree{Name: "温宇航", Rank: 1}
	father.AddChild(db, father, root6_1, root6_1_1)

	father.UpdateCoupleByName(db, "温茂福", []string{"裴淑梅"})
	father.UpdateCoupleByName(db, "温茂禄", []string{"任京花"})
	father.UpdateCoupleByName(db, "温茂贵", []string{"裴晓霞"})
	father.UpdateCoupleByName(db, "温茂圣", []string{"王淑花"})
	father.UpdateCoupleByName(db, "温茂永", []string{"裴海兰"})
	father.UpdateCoupleByName(db, "温茂来", []string{"李贺之"})
	father.UpdateCoupleByName(db, "温茂立", []string{"裴桂玲"})
	father.UpdateCoupleByName(db, "温茂志", []string{"裴晓萍"})
	father.UpdateCoupleByName(db, "温茂海", []string{"淑艳"})
	father.UpdateCoupleByName(db, "温茂军", []string{"李素焕"})
	father.UpdateCoupleByName(db, "温茂岭", []string{"刘景花"})
	father.UpdateCoupleByName(db, "温茂凯", []string{"王翠艳"})
}

func BuildQiChangTree(father *core.FamilyTree) {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温学文", Rank: 1}
	root00 := &core.FamilyTree{Name: "温学诗", Rank: 2}
	root02 := &core.FamilyTree{Name: "温学礼", Rank: 3}
	if father.FindByName("温学文") != nil {
		log.Printf("已经创建 BuildQiChangTree ，无需再创建！")
		return
	}
	//db.Create(root)
	father.AddChild(db, father, father, root00)
	father.AddChild(db, father, father, root)
	father.AddChild(db, father, father, root02)
	// 读取数据
	//father.UpdateCoupleByName(db, "温学文", []string{"汪氏"})
	father.UpdateCoupleByName(db, "温学诗", []string{"王氏"})
	father.UpdateCoupleByName(db, "温学礼", []string{"张氏", "梁氏"})

	root1 := &core.FamilyTree{Name: "温德辉", Rank: 1}
	root2 := &core.FamilyTree{Name: "温德合", Rank: 1}
	father.AddChild(db, father, root00, root1)
	father.AddChild(db, father, root02, root2)

	//father.UpdateCoupleByName(db, "温德辉", []string{"王氏"})
	father.UpdateCoupleByName(db, "温德合", []string{"葛氏"})

	root1_1 := &core.FamilyTree{Name: "温茂合", Rank: 1}
	father.AddChild(db, father, root1, root1_1)
	root1_2 := &core.FamilyTree{Name: "温茂林", Rank: 2}
	father.AddChild(db, father, root1, root1_2)

	root2_1 := &core.FamilyTree{Name: "温茂发", Rank: 1}
	father.AddChild(db, father, root2, root2_1)
	root2_2 := &core.FamilyTree{Name: "温茂圣", Rank: 2}
	father.AddChild(db, father, root2, root2_2)

	root1_1_1 := &core.FamilyTree{Name: "温锦章", Rank: 1}
	father.AddChild(db, father, root1_1, root1_1_1)

	root1_1_1_1 := &core.FamilyTree{Name: "温路军", Rank: 1}
	father.AddChild(db, father, root1_1_1, root1_1_1_1)
	root1_1_1_2 := &core.FamilyTree{Name: "温海军", Rank: 2}
	father.AddChild(db, father, root1_1_1, root1_1_1_2)

	father.UpdateCoupleByName(db, "温茂合", []string{"姚氏"})
	father.UpdateCoupleByName(db, "温茂林", []string{"陈氏"})
}

func BuildQiYunTree(father *core.FamilyTree) {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root00 := &core.FamilyTree{Name: "温学朱", Rank: 1}
	root01 := &core.FamilyTree{Name: "温学名", Rank: 2}
	if father.FindByName("温学朱") != nil {
		log.Printf("已经创建 BuildQiYunTree ，无需再创建！")
		return
	}
	father.AddChild(db, father, father, root00)
	father.AddChild(db, father, father, root01)
	// 读取数据
	//father.UpdateCoupleByName(db, "温学文", []string{"汪氏"})
	father.UpdateCoupleByName(db, "温学朱", []string{"张氏"})
	father.UpdateCoupleByName(db, "温学名", []string{"裴氏", "汪氏"})

	root1 := &core.FamilyTree{Name: "温德洪", Rank: 1}
	root2 := &core.FamilyTree{Name: "温德清", Rank: 1}
	root3 := &core.FamilyTree{Name: "温德广", Rank: 2}
	father.AddChild(db, father, root00, root1)
	father.AddChild(db, father, root01, root2)
	father.AddChild(db, father, root01, root3)
	//father.UpdateCoupleByName(db, "温德辉", []string{"王氏"})
	father.UpdateCoupleByName(db, "温德洪", []string{"李氏"})
	father.UpdateCoupleByName(db, "温德清", []string{"孙氏"})
	father.UpdateCoupleByName(db, "温德广", []string{"刘氏"})

	root1_1 := &core.FamilyTree{Name: "温茂生", Rank: 1}
	father.AddChild(db, father, root1, root1_1)
	root1_2 := &core.FamilyTree{Name: "温茂旺", Rank: 2}
	father.AddChild(db, father, root1, root1_2)

	root2_1 := &core.FamilyTree{Name: "温茂维", Rank: 1}
	father.AddChild(db, father, root2, root2_1)
	root2_2 := &core.FamilyTree{Name: "温茂才", Rank: 2}
	father.AddChild(db, father, root2, root2_2)
	root3_1 := &core.FamilyTree{Name: "温茂春-", Rank: 1}
	father.AddChild(db, father, root3, root3_1)

	root1_2_1 := &core.FamilyTree{Name: "温波", Rank: 1}
	father.AddChild(db, father, root1_2, root1_2_1)
	root2_2_1 := &core.FamilyTree{Name: "温景武", Rank: 1}
	father.AddChild(db, father, root2_2, root2_2_1)
	root2_2_2 := &core.FamilyTree{Name: "温景军", Rank: 2}
	father.AddChild(db, father, root2_2, root2_2_2)

	father.UpdateCoupleByName(db, "温茂生", []string{"崔氏"})
	father.UpdateCoupleByName(db, "温茂旺", []string{"庭晓芳"})
	father.UpdateCoupleByName(db, "温茂维", []string{"俞氏"})
	father.UpdateCoupleByName(db, "温茂才", []string{"薄素荣"})
	father.UpdateCoupleByName(db, "温景军", []string{"薛梅"})

}

func BuildQiCaiTree(father *core.FamilyTree) {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root00 := &core.FamilyTree{Name: "温学孟", Rank: 1}
	root01 := &core.FamilyTree{Name: "温学彦", Rank: 2}
	root02 := &core.FamilyTree{Name: "温学司", Rank: 2}
	if father.FindByName("温学孟") != nil {
		log.Printf("已经创建 BuildQiCaiTree ，无需再创建！")
		return
	}
	father.AddChild(db, father, father, root00)
	father.AddChild(db, father, father, root01)
	father.AddChild(db, father, father, root02)
	// 读取数据
	//father.UpdateCoupleByName(db, "温学文", []string{"汪氏"})
	father.UpdateCoupleByName(db, "温学孟", []string{"赵氏"})

	root1 := &core.FamilyTree{Name: "温德耀", Rank: 1}
	father.AddChild(db, father, root01, root1)

}

func BuildQiRuiTree(root *core.FamilyTree) {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	if root.FindByName("温学仙") != nil {
		log.Printf("BuildQiRuiTree 已经创建，无需再创建！")
		return
	}

	root1 := &core.FamilyTree{Name: "温学仙", Rank: 1}
	root2 := &core.FamilyTree{Name: "温学广", Rank: 2}
	root3 := &core.FamilyTree{Name: "温学鲁", Rank: 3}
	root.AddChild(db, root, root, root1)
	root.AddChild(db, root, root, root2)
	root.AddChild(db, root, root, root3)

}

func BuildXueRongJunGongTree() {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	father := &core.FamilyTree{Name: "温友贵", Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", father.Name).Preload("Children").Find(&father).Error; err != nil {
		log.Println("不存在 BuildXueRongJunGongTree", father.Name)
		log.Fatal(err)
	}
	father.ConstructTree(db)
	if father.FindByName("温学荣") == nil {
		father1 := father.FindByName("温其连")
		root1 := &core.FamilyTree{Name: "温学荣", Rank: 1}
		father.AddChild(db, father, father1, root1)
		father.UpdateCoupleByName(db, "温学荣", []string{"夏氏"})
	}
	if father.FindByName("温学君") == nil {
		father1 := father.FindByName("温其璋")
		root1 := &core.FamilyTree{Name: "温学君", Rank: 1}
		father.AddChild(db, father, father1, root1)
		father.UpdateCoupleByName(db, "温学君", []string{"赵氏"})
	}
	if father.FindByName("温学公") == nil {
		father1 := father.FindByName("温其俊")
		root1 := &core.FamilyTree{Name: "温学公", Rank: 1}
		father.AddChild(db, father, father1, root1)
		father.UpdateCoupleByName(db, "温学公", []string{"陈氏"})
	}
}
func BuildXueFuTree() *core.FamilyTree {
	db, _ := core.OpenDb(DBUser, DBPass, DBIpPort, DBName)
	defer db.Close()
	root := &core.FamilyTree{Name: "温其瑞", Rank: 1}
	// 读取数据
	if err := db.Where("name = ?", root.Name).Preload("Children").Find(&root).Error; err == nil {
		//log.Fatal(err)
		log.Println("数据库已存在，无需创建")
	} else {
		// 创建流程
	}
	return root
}

func BuildWenShiJiaZu() {
	_ = BuildL1To5Tree()
	_ = BuildL5To6Tree()
	_ = BuildL6To7Tree()
	_ = BuildL7To8Tree()
	rootE := BuildL8To9Tree()
	rootFather := rootE.FindByName("温其润")
	BuildQiRunTree(rootFather)
	rootFather1 := rootE.FindByName("温其瑞")
	BuildQiRuiTree(rootFather1)
	rootFather2 := rootE.FindByName("温其惠")
	BuildQihuiTree(rootFather2)
	rootFather3 := rootE.FindByName("温其昌")
	BuildQiChangTree(rootFather3)
	rootFather5 := rootE.FindByName("温其圣")
	BuildQiShengTree(rootFather5)
	rootFather6 := rootE.FindByName("温其允")
	BuildQiYunTree(rootFather6)
	rootFather7 := rootE.FindByName("温其才")
	BuildQiCaiTree(rootFather7)
	BuildXueRongJunGongTree()
}
func UpdateInfo() {
	db, root := GetRootByName("温友贵")
	defer db.Close()
	if root != nil {
		root.UpdateReMarkByName(db, "温自富", "以上埋葬于南苗坐兹爽坨于西北")
		root.UpdateReMarkByName(db, "温其禄", "少亡")
		root.UpdateReMarkByName(db, "温其禄", "少亡")
		root.UpdateReMarkByName(db, "温其光", "出嫁")
		root.UpdateReMarkByName(db, "温其明", "出嫁")
		root.UpdateReMarkByName(db, "温学武", "抗日牺牲")
		root.UpdateReMarkByName(db, "温德合", "全家现住吉林省白山市")
		root.UpdateReMarkByName(db, "温茂生", "生子一人,离婚带走")
		root.UpdateReMarkByName(db, "温德耀", "现住赵各庄")
		root.UpdateReMarkByName(db, "温其美", "移居关外，下落不明")
		root.UpdateReMarkByName(db, "温有福", "统招硕士-北京交通大学")
		root.AddChild(db, root, root.FindByName("温有福"),
			&core.FamilyTree{Name: "温暖",Rank: 1,ReMark: "温暖为其小名"})
		root.UpdateReMarkByName(db, "温宝福", "已移居江苏南京;统招本科-南京邮电大学")
	}
}
