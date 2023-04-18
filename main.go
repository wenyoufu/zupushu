// package treeJiazu
package main

import "treeJiazu/treebuild"

func main() {
	treebuild.BuildWenShiJiaZu()
	treebuild.UpdateInfo()
	db, root:=treebuild.GetRootByName("温友贵")
	defer db.Close()
	root.PrintDetailByName("温宝福")
	//root.PrintTree(0)

}
