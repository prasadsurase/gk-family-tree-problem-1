package search

import (
  "github.com/prasadsurase/geektrust-family-tree-problem/structs"
)

// Relatives find relatives for a member based on the provided relation
func Relatives(root *structs.Couple, member *structs.Member, relation string) ([]*structs.Member, error) {
  switch relation {
    case "Paternal Uncles": // DONE
    // Father’s brothers, Father’s brother-in-laws
    uncles := []*structs.Member{}
    if member.Parents != nil {
      // get father's brothers
      for _, child := range member.Parents.Husband.Parents.Children {
        if child.Gender == "M" && member.Parents.Husband != child {
          uncles = append(uncles, child)
        }
      }
      // get father's sisters' spouses. ie. father's brothers-in-laws
      sisters, _ := Relatives(member.Parents, member.Parents.Husband, "Sisters")
      for _, sis := range sisters {
        if len(sis.Relationships) > 0 {
          for _, pair := range sis.Relationships {
            uncles = append(uncles, pair.Husband)
          }
        }
      }
    }
    return uncles, nil

    case "Maternal Uncles": // DONE
    // Mother’s brother, Mother’s brother-in-laws
    uncles := []*structs.Member{}
    if member.Parents != nil {
      // get mother's brothers
      for _, child := range member.Parents.Wife.Parents.Children {
        if child.Gender == "M" {
          uncles = append(uncles, child)
        }
      }
      // get mother's sisters' spouses. ie. father's brothers-in-laws
      sisters, _ := Relatives(member.Parents, member.Parents.Wife, "Sisters")
      for _, sis := range sisters {
        if len(sis.Relationships) > 0 {
          for _, pair := range sis.Relationships {
            uncles = append(uncles, pair.Husband)
          }
        }
      }
    }
    return uncles, nil

    case "Paternal Aunts": // DONE
    // Father’s sister, Father’s sister-in-laws
    aunties := []*structs.Member{}
    if member.Parents != nil {
      for _, child := range member.Parents.Husband.Parents.Children {
        if child.Gender == "F" {
          aunties = append(aunties, child)
        }
      }
      // get father's brothers' spouses. ie. father's sisters-in-laws
      brothers, _ := Relatives(member.Parents, member.Parents.Husband, "Brothers")
      for _, bro := range brothers {
        if len(bro.Relationships) > 0 {
          for _, pair := range bro.Relationships {
            aunties = append(aunties, pair.Wife)
          }
        }
      }
    }
    return aunties, nil

    case "Maternal Aunts": // DONE
    // Mother’s sister, Mother’s sister-in-law
    aunties := []*structs.Member{}
    if member.Parents != nil {
      for _, child := range member.Parents.Wife.Parents.Children {
        if child.Gender == "F" && child != member.Parents.Wife {
          aunties = append(aunties, child)
        }
      }
      // get mother's brothers' spouses. ie. mother's sisters-in-laws
      brothers, _ := Relatives(member.Parents, member.Parents.Wife, "Brothers")
      for _, bro := range brothers {
        if len(bro.Relationships) > 0 {
          for _, pair := range bro.Relationships {
            aunties = append(aunties, pair.Wife)
          }
        }
      }
    }
    return aunties, nil

    case "Sisters-In-Law": // DONE
    // Spouse’s sisters, Wives of siblings
    sisters := []*structs.Member{}
    if member.Parents != nil {
      // find wives of siblings(brothers)
      for _, child := range member.Parents.Children {
        if child.Gender == "M" && child != member {
          if len(child.Relationships) > 0 {
            for _, pair := range child.Relationships {
              sisters = append(sisters, pair.Wife)
            }
          }
        }
      }
    }
    if len(member.Relationships) > 0 {
      for _, pair := range member.Relationships {
        if pair.Husband == member {
          sisters, _ := Relatives(pair, pair.Wife, "Sisters")
          for _, sis := range sisters {
            if len(sis.Relationships) > 0 {
              for _, pair := range sis.Relationships {
                sisters = append(sisters, pair.Wife)
              }
            }
          }
        }
      }
    }
    return sisters, nil
    case "Brothers-In-Law": // DONE
    // Spouse’s brothers, Husbands of siblings
    brothers := []*structs.Member{}
    if member.Parents != nil {
      // find husbands of siblings(sisters)
      for _, child := range member.Parents.Children {
        if child.Gender == "F" && child != member {
          if len(child.Relationships) > 0 {
            for _, pair := range child.Relationships {
              brothers = append(brothers, pair.Husband)
            }
          }
        }
      }
    }
    if len(member.Relationships) > 0 {
      for _, pair := range member.Relationships {
        if pair.Wife == member {
          brothers, _ := Relatives(pair, pair.Husband, "Brothers")
          for _, bro := range brothers {
            if len(bro.Relationships) > 0 {
              for _, pair := range bro.Relationships {
                brothers = append(brothers, pair.Wife)
              }
            }
          }
        }
      }
    }
    return brothers, nil

    case "Spouses": // DONE
    spouses := []*structs.Member{}
    if len(member.Relationships) > 0 {
      for _, pair := range member.Relationships {
        if member == pair.Husband {
          spouses = append(spouses, pair.Wife)
        } else {
          spouses = append(spouses, pair.Husband)
        }
      }
    }
    return spouses, nil
  case "Cousins":
    // Children of parent’s siblings

    case "Father": // DONE
    return []*structs.Member{member.Parents.Husband}, nil

    case "Mother": // DONE
    return []*structs.Member{member.Parents.Wife}, nil

    case "Sisters": // DONE
    sisters := []*structs.Member{}
    for _, child := range member.Parents.Children {
      if child.Gender == "F" && child != member {
        sisters = append(sisters, child)
      }
    }
    return sisters, nil

    case "Brothers": // DONE
    brothers := []*structs.Member{}
    for _, child := range member.Parents.Children {
      if child.Gender == "M" && child != member {
        brothers = append(brothers, child)
      }
    }
    return brothers, nil

    case "Siblings": // DONE
    children := []*structs.Member{}
    for _, child := range member.Parents.Children {
      if child != member {
        children = append(children, child)
      }
    }
    return children, nil

    case "Daughters": // DONE
    daughters := []*structs.Member{}
    if len(member.Relationships) > 0 {
      for _, child := range member.Relationships[0].Children {
        if child.Gender == "F" {
          daughters = append(daughters, child)
        }
      }
    }
    //TODO: define a method for the Couple and Member struct
    return daughters, nil

    case "Sons": // DONE
    sons := []*structs.Member{}
    if len(member.Relationships) > 0 {
      for _, child := range member.Relationships[0].Children {
        if child.Gender == "M" {
          sons = append(sons, child)
        }
      }
    }
    //TODO: define a method for the Couple and Member struct
    return sons, nil

    case "Children": // DONE
    children := []*structs.Member{}
    if len(member.Relationships) > 0 {
      for _, pair := range member.Relationships {
        children = append(children, pair.Children...)
      }
    }
    //TODO: define a method for the Couple and Member struct
    return children, nil

    case "Grand Daughters": // DONE
    grandDaughters := []*structs.Member{}
    if len(member.Relationships) > 0 {
      for _, pair := range member.Relationships {
        for _, child := range pair.Children {
          if child.Gender == "F" {
            grandDaughters = append(grandDaughters, child)
          }
        }
      }
    }
    return grandDaughters, nil

    // Daughters of daughters and sons
    case "Grand Sons": // DONE
    grandSons := []*structs.Member{}
    if len(member.Relationships) > 0 {
      for _, pair := range member.Relationships {
        for _, child := range pair.Children {
          if child.Gender == "M" {
            grandSons = append(grandSons, child)
          }
        }
      }
    }
    return grandSons, nil
    // Sons of daughters and sons

    case "Fathers In Law": // DONE
    // TODO: add method to retreive spouses of a member
    fathersInLaw := []*structs.Member{}
    if len(member.Relationships) > 0 {
      for _, pair := range member.Relationships {
        if pair.Husband == member {
          // TODO: check if spouse's parents are present.
          fathersInLaw = append(fathersInLaw, pair.Wife.Parents.Husband)
        } else {
          // TODO: check if spouse's parents are present.
          fathersInLaw = append(fathersInLaw, pair.Husband.Parents.Husband)
        }
      }
    }
    return fathersInLaw, nil

    case "Mothers In Law": // DONE
    mothersInLaw := []*structs.Member{}
    if len(member.Relationships) > 0 {
      for _, pair := range member.Relationships {
        if pair.Husband == member {
          // TODO: check if spouse's parents are present.
          mothersInLaw = append(mothersInLaw, pair.Wife.Parents.Wife)
        } else {
          // TODO: check if spouse's parents are present.
          mothersInLaw = append(mothersInLaw, pair.Husband.Parents.Wife)
        }
      }
    }
    return mothersInLaw, nil
  }
  return []*structs.Member{}, nil
}
