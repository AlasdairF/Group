package group

type groupStruct struct {
 groups [][]int
 groupmap map[int]int
 num int
 messy bool
}

func New() *groupStruct {
	g := new(groupStruct)
	g.groupmap = make(map[int]int)
	g.groups = make([][]int, 0, 10)
	return g
}

func (g *groupStruct) Add(a ...int) {
	
	if len(a) == 0 {
		return
	}
	
	// Check how many groups already exist
	var numgroups, grp, lastgroup, i, id int
	var ok bool
	thisgroup := make([]int, len(a))
	for i, id = range a {
		if grp, ok = g.groupmap[id]; ok {
			thisgroup[i] = grp
			numgroups++
			lastgroup = i
		} else {
			thisgroup[i] = -1
		}
	}
	
	if numgroups == 0 {
		// Make a new group for all and add all into the new group
		numgroups = len(g.groups)
		g.groups = append(g.groups, a)
		for _, id = range a {
			g.groupmap[id] = numgroups
		}
		g.num++
		return
	} else if numgroups == 1 {
		// Add all to lastgroup
		grp = thisgroup[lastgroup]
		for i, id = range a {
			if i != lastgroup {
				g.groups[grp] = append(g.groups[grp], id)
				g.groupmap[id] = grp
			}
		}
		return
	}
	
	// There are multiple groups which need to be joined together
	var id2 int
	newgroup := make([]int, 0, len(a))
	for i, id = range a {
		grp = thisgroup[i]
		if grp == -1 {
			newgroup = append(newgroup, id)
		} else {
			if len(g.groups[grp]) > 0 {
				for _, id2 = range g.groups[grp] {
					newgroup = append(newgroup, id2)
				}
				g.groups[grp] = make([]int, 0)
				g.num--
			}
		}
	}
	grp = thisgroup[lastgroup]
	for _, id = range newgroup {
		g.groupmap[id] = grp
	}
	g.groups[grp] = newgroup
	g.num++
	g.messy = true
	return
}

func (g *groupStruct) Groups() [][]int {
	if !g.messy {
		return g.groups
	}
	num := len(g.groups)
	groups := make([][]int, num)
	var on int
	for i:=0; i<num; i++ {
		if len(g.groups[i]) > 0 {
			groups[on] = g.groups[i]
			on++
		}
	}
	return groups[0:on]
}

func (g *groupStruct) Len() int {
	return g.num
}

func (g *groupStruct) Of(a int) (int, bool) {
	grp, ok := g.groupmap[a]
	if !ok {
		return -1, false
	}
	if !g.messy {
		return grp, ok
	}
	newgrp := grp
	for i:=0; i<grp; i++ {
		if len(g.groups[i]) == 0 {
			newgrp--
		}
	}
	return newgrp, true
}

func (g *groupStruct) Optimize() {
	if !g.messy {
		return
	}
	l := len(g.groups)
	var on, id int
	for i:=0; i<l; i++ {
		if len(g.groups[i]) > 0 {
			g.groups[on] = g.groups[i]
			for _, id = range g.groups[on] {
				g.groupmap[id] = on
			}
			on++
		}
	}
	g.groups = g.groups[0:on]
	g.num = on
	g.messy = false
	return
}

