package display

import (
  "fmt"

  "github.com/prasadsurase/geektrust-family-tree-problem/structs"
)

// FamilyTree displays the family tree from the provided root
func FamilyTree(pair *structs.Couple, level int) {
  nodeName := pair.Husband.Name + "<->" + pair.Wife.Name
  for i := 0; i < level; i++ {
    nodeName = "    " + nodeName
  }
  fmt.Println(nodeName)

  for _, child := range pair.Children {
    childName := child.Name
    for i := 0; i < level; i++ {
      childName = "    " + childName
    }
    fmt.Println(childName + "(" + child.Gender + ")")
    for _, couple := range child.Relationships {
      FamilyTree(couple, level+1)
    }
  }
}
