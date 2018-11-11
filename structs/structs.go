package structs

// Member struct
type Member struct {
  Name          string
  Gender        string
  Parents       *Couple
  Generation    int
  Relationships []*Couple
  // Father  *Member
  // Mother  *Member
}

// Couple struct
type Couple struct {
  Husband  *Member
  Wife     *Member
  Children []*Member
}

// TODO:
// for couple, return slice will contain only one record for father/mother in-law
// for member, return slice may contain more than one father/mother in-law

// Spouses interface to retrieve spouses of a couple/member
type Spouses interface {
  Spouses(*Member) ([]*Member, error)
}

// FathersInLaw interface to retrieve spouses of a couple/member
type FathersInLaw interface {
  FathersInLaw(*Member) ([]*Member, error)
}

// MothersInLaw interface to retrieve spouses of a couple/member
type MothersInLaw interface {
  MothersInLaw(*Member) ([]*Member, error)
}

// BrothersInLaw interface to retrieve spouses of a couple/member
type BrothersInLaw interface {
  BrothersInLaw(*Member) ([]*Member, error)
}

// SistersInLaw interface to retrieve spouses of a couple/member
type SistersInLaw interface {
  SistersInLaw(*Member) ([]*Member, error)
}

// Father get the father of the member. maybe nil in case of single parent.
type Father interface {
  Father(*Member) (*Member, error)
}

// Mother get the mother of the member. maybe nil in case of single parent.
type Mother interface {
  Mother(*Member) (*Member, error)
}
