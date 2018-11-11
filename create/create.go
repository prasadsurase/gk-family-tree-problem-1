package create

import "github.com/prasadsurase/geektrust-family-tree-problem/structs"

// Family creates the family tree
func Family() *structs.Couple {
  king := structs.Member{Name: "Shan", Gender: "M"}
  queen := structs.Member{Name: "Anga", Gender: "F"}
  kingQueen := structs.Couple{Husband: &king, Wife: &queen}
  king.Relationships = []*structs.Couple{&kingQueen}
  queen.Relationships = []*structs.Couple{&kingQueen}

  ish := structs.Member{Name: "Ish", Gender: "M", Parents: &kingQueen}
  chit := structs.Member{Name: "Chit", Gender: "M", Parents: &kingQueen}
  vich := structs.Member{Name: "Vich", Gender: "M", Parents: &kingQueen}
  satya := structs.Member{Name: "Satya", Gender: "F", Parents: &kingQueen}

  ambi := structs.Member{Name: "Ambi", Gender: "F"}
  chitAmbi := structs.Couple{Husband: &chit, Wife: &ambi}
  chit.Relationships = []*structs.Couple{&chitAmbi}
  ambi.Relationships = []*structs.Couple{&chitAmbi}

  lika := structs.Member{Name: "Lika", Gender: "F"}
  vichLika := structs.Couple{Husband: &vich, Wife: &lika}
  vich.Relationships = []*structs.Couple{&vichLika}
  lika.Relationships = []*structs.Couple{&vichLika}

  vyan := structs.Member{Name: "Vyan", Gender: "M"}
  vyanSatya := structs.Couple{Husband: &vyan, Wife: &satya}
  vyan.Relationships = []*structs.Couple{&vyanSatya}
  satya.Relationships = []*structs.Couple{&vyanSatya}

  kingQueen.Children = []*structs.Member{&ish, &chit, &vich, &satya}

  drita := structs.Member{Name: "Drita", Gender: "M", Parents: &chitAmbi}
  jaya := structs.Member{Name: "Jaya", Gender: "F"}
  dritaJaya := structs.Couple{Husband: &drita, Wife: &jaya}
  drita.Relationships = []*structs.Couple{&dritaJaya}
  jaya.Relationships = []*structs.Couple{&dritaJaya}
  vrita := structs.Member{Name: "Vrita", Gender: "M", Parents: &chitAmbi}
  chitAmbi.Children = []*structs.Member{&drita, &vrita}

  jata := structs.Member{Name: "Jata", Gender: "M", Parents: &dritaJaya}
  driya := structs.Member{Name: "Driya", Gender: "F", Parents: &dritaJaya}
  manu := structs.Member{Name: "Manu", Gender: "M"}
  manuDriya := structs.Couple{Husband: &manu, Wife: &driya}
  manu.Relationships = []*structs.Couple{&manuDriya}
  driya.Relationships = []*structs.Couple{&manuDriya}
  dritaJaya.Children = []*structs.Member{&jata, &driya}

  vila := structs.Member{Name: "Vila", Gender: "M", Parents: &vichLika}
  jnki := structs.Member{Name: "Jnki", Gender: "F"}
  vilaJnki := structs.Couple{Husband: &vila, Wife: &jnki}
  vila.Relationships = []*structs.Couple{&vilaJnki}
  jnki.Relationships = []*structs.Couple{&vilaJnki}
  chika := structs.Member{Name: "Chika", Gender: "F", Parents: &vichLika}
  kpila := structs.Member{Name: "Kpila", Gender: "M"}
  kpilaChika := structs.Couple{Husband: &kpila, Wife: &chika}
  kpila.Relationships = []*structs.Couple{&kpilaChika}
  chika.Relationships = []*structs.Couple{&kpilaChika}
  vichLika.Children = []*structs.Member{&vila, &chika}

  lavnya := structs.Member{Name: "Lavnya", Gender: "F", Parents: &vilaJnki}
  gru := structs.Member{Name: "Gru", Gender: "M"}
  gruLavnya := structs.Couple{Husband: &gru, Wife: &lavnya}
  gru.Relationships = []*structs.Couple{&gruLavnya}
  lavnya.Relationships = []*structs.Couple{&gruLavnya}
  vilaJnki.Children = []*structs.Member{&lavnya}

  satvy := structs.Member{Name: "Satvy", Gender: "F", Parents: &vyanSatya}
  asva := structs.Member{Name: "Asva", Gender: "M"}
  asvaSatvy := structs.Couple{Husband: &asva, Wife: &satvy}
  asva.Relationships = []*structs.Couple{&asvaSatvy}
  satvy.Relationships = []*structs.Couple{&asvaSatvy}
  // vyanSatya.Children = []*structs.Member{&}

  savya := structs.Member{Name: "Savya", Gender: "M", Parents: &vyanSatya}
  krpi := structs.Member{Name: "Krpi", Gender: "F"}
  savyaKrpi := structs.Couple{Husband: &savya, Wife: &krpi}
  savya.Relationships = []*structs.Couple{&savyaKrpi}
  krpi.Relationships = []*structs.Couple{&savyaKrpi}

  saayan := structs.Member{Name: "Saayan", Gender: "M", Parents: &vyanSatya}
  mina := structs.Member{Name: "Mina", Gender: "F"}
  saayanMina := structs.Couple{Husband: &saayan, Wife: &mina}
  saayan.Relationships = []*structs.Couple{&saayanMina}
  mina.Relationships = []*structs.Couple{&saayanMina}

  vyanSatya.Children = []*structs.Member{&satvy, &savya, &saayan}

  kriya := structs.Member{Name: "Kriya", Gender: "M", Parents: &savyaKrpi}
  savyaKrpi.Children = []*structs.Member{&kriya}

  misa := structs.Member{Name: "Misa", Gender: "M", Parents: &saayanMina}
  saayanMina.Children = []*structs.Member{&misa}

  return &kingQueen
}
