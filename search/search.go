package search

import (
  "errors"
  "fmt"
  "sort"

  "github.com/prasadsurase/geektrust-family-tree-problem/structs"
)

type keyValue struct {
  Key   string
  Value int
}

// Member search in the provided tree.
func Member(currentPair *structs.Couple, searchName string, gen int) (*structs.Member, int, error) {
  switch searchName {
  case currentPair.Husband.Name:
    if currentPair.Husband.Parents == nil {
      return currentPair.Husband, gen - 1, nil
    }
    return currentPair.Husband, gen, nil

  case currentPair.Wife.Name:
    if currentPair.Wife.Parents == nil {
      return currentPair.Wife, gen - 1, nil
    }
    return currentPair.Wife, gen, nil

  default:
    for _, child := range currentPair.Children {
      if child.Name == searchName {
        return child, gen, nil
      } else {
        for _, couple := range child.Relationships {
          member, level, _ := Member(couple, searchName, gen+1)
          if member != nil {
            return member, level, nil
          }
        }
      }
    }
  }
  return nil, 0, errors.New("Member with provided name '" + searchName + "' not found")
}

// MothersWithMaxKids if gender is 'F' then retrieve mothers with max daughters, 'M' for max sons, 'N' for max children
func MothersWithMaxKids(currentPair *structs.Couple, gender string) (map[string]int, error) {
  stats := make(map[string]int)
  data, _ := countKids(currentPair, gender, &stats)
  var ss []keyValue
  for k, v := range *data {
    ss = append(ss, keyValue{k, v})
  }

  sort.Slice(ss, func(i, j int) bool {
    return ss[i].Value > ss[j].Value
  })
  fmt.Println("Data:", ss)
  return stats, nil
}

func countKids(pair *structs.Couple, gender string, stats *map[string]int) (*map[string]int, error) {
  switch gender {
  case "F":
    total := 0
    for _, member := range pair.Children {
      if member.Gender == "F" {
        total++
      }
    }
    if total > 0 {
      (*stats)[pair.Wife.Name] = total
    }
  case "M":
    total := 0
    for _, member := range pair.Children {
      if member.Gender == "M" {
        total++
      }
    }
    if total > 0 {
      (*stats)[pair.Wife.Name] = total
    }
  case "N":
    if len(pair.Children) > 0 {
      (*stats)[pair.Wife.Name] = len(pair.Children)
    }
  }

  for _, child := range pair.Children {
    for _, couple := range child.Relationships {
      countKids(couple, gender, stats)
    }
  }
  return stats, nil
}
