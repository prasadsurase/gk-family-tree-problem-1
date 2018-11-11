package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"

  "github.com/prasadsurase/geektrust-family-tree-problem/create"
  "github.com/prasadsurase/geektrust-family-tree-problem/display"
  "github.com/prasadsurase/geektrust-family-tree-problem/search"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  kingQueen := create.Family()
  var input int
  for {
    fmt.Println("======================================================")
    fmt.Println("Please choose an option")
    fmt.Println("1) Add child to a couple")
    fmt.Println("2) Display relations among 2 members")
    fmt.Println("3) Display mothers with max daughters")
    fmt.Println("4) Display mothers with max sons")
    fmt.Println("5) Display relatives for a member")
    fmt.Println("6) Display family tree")
    fmt.Println("7) Exit")
    fmt.Scan(&input)
    switch input {
    case 1:
      // add a child to a couple
    case 2:
      // display relation amoung 2 members
      var memberOne, memberTwo string
      fmt.Println("Enter name of member:")
      memberOne, _ = reader.ReadString([]byte("\n")[0])
      memberOne = strings.TrimSuffix(memberOne, "\n")
      fmt.Println("Enter name of relative:")
      memberTwo, _ = reader.ReadString([]byte("\n")[0])
      memberTwo = strings.TrimSuffix(memberTwo, "\n")
      fmt.Println("Member One:", memberOne)
      fmt.Println("Member Two:", memberTwo)
      rel, _ := search.Relation(kingQueen, memberOne, memberTwo)
      fmt.Println("Relation between ", memberOne, " & ", memberTwo, ":", rel)
    case 3:
      // display mother(s) with most daughters
      search.MothersWithMaxKids(kingQueen, "F")
    case 4:
      // display mother(s) with most sons
      search.MothersWithMaxKids(kingQueen, "M")
    case 5:
      // display relatives for a member
      var name, relation string
      fmt.Println("Enter name of member:")
      name, _ = reader.ReadString([]byte("\n")[0])
      name = strings.TrimSuffix(name, "\n")
      fmt.Println("Enter relation:")
      relation, _ = reader.ReadString([]byte("\n")[0])
      relation = strings.TrimSuffix(relation, "\n")
      fmt.Println("Member:", name)
      fmt.Println("Relation:", relation)
      member, _, _ := search.Member(kingQueen, name, 0)
      relatives, _ := search.Relatives(kingQueen, member, relation)
      fmt.Println("Relatives:")
      for _, r := range relatives {
        fmt.Println(r.Name)
      }
    case 6:
      fmt.Println("Family tree:")
      display.FamilyTree(kingQueen, 0)
    case 7:
      return
    default:
      fmt.Println("Wrong choice. Try Again.")
    }
  }
}
