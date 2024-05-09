package maps

// 初始化地图
func InitMap() *Map {
	m := Map{"新手村", "这是一个不大的村子,看起来安静祥和,你可以在这里学习如何探索世界.", 3, 3, 10, nil, nil}
	m.Scenes = make([]*Scene, 0)
	//开始实例化地图下的场景
	s11 := Scene{&m, 1, 1, "村子", "村子不大,村子北面有一座山,西边是一边树林,东边是一片田地,南面有一条小河流过,看起来很安静,村里有小孩在嬉戏,有老人在村头的大树下下棋.", nil, nil}
	s01 := Scene{&m, 1, 1, "树林", "一片稀疏的果林,有桃树和李树", nil, nil}
	s10 := Scene{&m, 1, 1, "小河", "一条10米宽的小河,水流缓慢,水很浅,可以看到河底的鹅卵石和成群的小鱼", nil, nil}
	s12 := Scene{&m, 1, 1, "后山", "一座不太高的小山包,但是周围地势平坦,在山顶应该可以看到很远.", nil, nil}
	s21 := Scene{&m, 1, 1, "稻田", "一片绿油油的稻田,快到收获的季节了,地里有农民在除草.", nil, nil}

	s11.Paths = make([]*Path, 0, 4)
	s11.Paths = append(s11.Paths, &Path{"south", &s10}, &Path{"north", &s12}, &Path{"west", &s01}, &Path{"east", &s21})

	s01.Paths = make([]*Path, 0, 1)
	s01.Paths = append(s01.Paths, &Path{"east", &s11})

	s21.Paths = make([]*Path, 0, 1)
	s21.Paths = append(s21.Paths, &Path{"west", &s11})

	s12.Paths = make([]*Path, 0, 1)
	s12.Paths = append(s12.Paths, &Path{"south", &s11})

	s10.Paths = make([]*Path, 0, 1)
	s10.Paths = append(s10.Paths, &Path{"north", &s11})

	return &m
}
