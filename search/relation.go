package search

import (
  "fmt"

  "github.com/prasadsurase/geektrust-family-tree-problem/structs"
)

// Relation among 2 members
func Relation(root *structs.Couple, nameOne, nameTwo string) (string, error) {
  member, memberGen, _ := Member(root, nameOne, 1)
  relative, relativeGen, _ := Member(root, nameTwo, 1)
  fmt.Println("Member:", member.Name, ", Gen:", memberGen)
  fmt.Println("Relative:", relative.Name, " Gen:", relativeGen)

  // if belong to the same generation
  if memberGen == relativeGen {
    if (member.Parents != nil && relative.Parents != nil) && member.Parents == relative.Parents {
      return "Siblings", nil
    } else if len(member.Relationships) > 0 && len(relative.Relationships) > 0 && member.Relationships[0] == relative.Relationships[0] {
      return "Spouse", nil
    } else if (member.Parents != nil && relative.Parents == nil) || (member.Parents == nil && relative.Parents != nil) || (member.Parents == nil && relative.Parents == nil) {
      if relative.Gender == "M" {
        return "Brother-in-law", nil
      }
      return "Sister-in-law", nil
    } else {
      return "Cousins", nil
    }
  } else if memberGen-relativeGen == 1 || memberGen-relativeGen == -1 {

  } else if memberGen-relativeGen > 1 || memberGen-relativeGen < -1 {

  }
  return "Something", nil
}
