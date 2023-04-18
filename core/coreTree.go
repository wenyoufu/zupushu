package core

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

type FamilyTree struct {
	ID     int
	Name   string
	Sex    string
	Rank   uint // 排行
	Couple string
	Born   string // 出生日期
	ReMark string // 备注
	//Children []*FamilyTree
	ParentID *int
	Children []*FamilyTree `gorm:"foreignkey:ParentID"`
}

func OpenDb(usr string, pwd string, ipPort string, dbName string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", usr+":"+pwd+"@tcp("+
		ipPort+")/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(false)
	retDb := db.AutoMigrate(&FamilyTree{})
	if retDb.Error != nil {
		log.Fatal(retDb.Error)
	}
	return db, err
}

func (t *FamilyTree) ConstructTree(db *gorm.DB) {

	if err := db.Preload("Children").Find(&t).Error; err != nil {
		log.Fatal(err)
	}
	for _, child := range t.Children {
		child.ConstructTree(db)
	}
}

func (t *FamilyTree) FindByName(name string) *FamilyTree {
	if t.Name == name {
		return t
	}
	for _, child := range t.Children {
		if found := child.FindByName(name); found != nil {
			return found
		}
	}
	return nil
}

func (t *FamilyTree) FindByID(id int) *FamilyTree {
	if t.ID == id {
		return t
	}
	for _, child := range t.Children {
		if found := child.FindByID(id); found != nil {
			return found
		}
	}
	return nil
}

//func (t *FamilyTree) Delete() {
//	if t == nil {
//		return
//	}
//	for _, child := range t.Children {
//		child.Delete()
//	}
//	t.Name = ""
//	t.Children = nil
//}

func (t *FamilyTree) Delete(db *gorm.DB, delName string) {
	if t == nil {
		return
	}
	if f := t.FindByName(delName); f != nil {
		db.Model(&f).Delete(f)
	}
	//log.Printf("Delete %v not found ", delName)
	t.ConstructTree(db)
}

func (t *FamilyTree) AddChild(db *gorm.DB, root *FamilyTree, father *FamilyTree, child *FamilyTree) {
	//t.SubAddChild(db, root, father, child)
	f := root.FindByName(father.Name)
	// 异常判断处理
	// 添加父亲标识
	addFather := false
	if f == nil {
		f = father
		addFather = true
	}
	if child.ParentID == nil {
		f.Children = []*FamilyTree{child}
		child.ParentID = &(f.ID)
		if addFather {
			db.Create(f)
		}
		db.Model(&t).Update(t)
		db.Model(&child).Update(child)
		t.ConstructTree(db)
	}

}

func (t *FamilyTree) UpdateName(db *gorm.DB, name string) {
	if t == nil {
		return
	}
	if f := t.FindByName(name); f != nil {
		f.Name = name
		db.Model(&f).Update(f)
	}
}

func (t *FamilyTree) UpdateReMarkByName(db *gorm.DB, name string, remark string) {
	if t == nil {
		return
	}
	if f := t.FindByName(name); f != nil {
		f.ReMark = remark
		db.Model(&f).Update(f)
	}
}

func (t *FamilyTree) UpdateBornByName(db *gorm.DB, name string, born string) {
	if t == nil {
		return
	}
	if f := t.FindByName(name); f != nil {
		f.Born = born
		db.Model(&f).Update(f)
	}
}

func (t *FamilyTree) UpdateCoupleByName(db *gorm.DB, name string, couples []string) {
	if t == nil {
		return
	}
	if f := t.FindByName(name); f != nil {
		//tmp, _ := json.Marshal(couples)
		//f.Couple = string(tmp)
		tmp := strings.Join(couples, " ")
		f.Couple = tmp
		db.Model(&f).Update(f)
	}
}

func (t *FamilyTree) PrintTree(level int) {

	//var copule []string
	//_ = json.Unmarshal([]byte(t.Couple), &copule)
	if t.Couple == "" {
		fmt.Printf("%s第%v代:%s\n", strings.Repeat("-", level*4), level+1, t.Name)
	} else {
		fmt.Printf("%s第%v代:%s 配偶:%v\n", strings.Repeat("-", level*4), level+1, t.Name, t.Couple)
	}
	for _, child := range t.Children {
		child.PrintTree(level + 1)
	}
}

func (t *FamilyTree) PrintDetailByName(name string) {
	if t == nil {
		return
	}
	if f := t.FindByName(name); f != nil {
		//var copule []string
		//_ = json.Unmarshal([]byte(f.Couple), &copule)
		if f.Sex == "" {
			f.Sex = "男"
		}
		childNum := 0
		if f.Children != nil {
			childNum = len(f.Children)
		}
		fmt.Printf("姓名:%v\n性别:%v\n排行:%v\n出生日期：%v\n配偶:%v\n"+
			"子嗣数量:%v\n备注:%v\n",
			f.Name, f.Sex, f.Rank, f.Born, f.Couple, childNum, f.ReMark)

	}
}
