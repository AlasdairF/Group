##Group

Group is a structure for managing groups in which a value can only be a member of one group, and adding it to any additional groups will combine those groups together. This package automatically handles the combining of groups.

Example:
     grp := group.New()
     grp.Add(1, 3, 5, 10)
     grp.Add(11, 14)
     grp.Add(20, 2)
     grp.Add(3, 11, 8)
    
The above will produce the groups:
     [0] 1, 3, 5, 8, 10, 11, 14
     [1] 2, 20
    
## Index

Create a new group:
     grp := New()
  
Add values to the group:
     grp.Add(1, 2, 3, 4, etc.)
    
Find the group of any value:
     id, exists := grp.Of(4)
    
Retrieve all groups:
     allgroups := grp.Groups() // [][]int
    
Get the total number of active (excluding empty) groups:
     numgroups := grp.Len()
    
Optimize the groups, which removes empty groups and speeds up `Groups()` & `Of()` functions. This does not change any of the group IDs.
     grp.Optimize()
  
