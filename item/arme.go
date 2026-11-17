package item

import "fmt"

// Arme représente une arme équipable.
type Arme struct {
	NomArme   string
	Degats    int    // bonus d'attaque physique (AD)
	Puissance int    // bonus de puissance magique (AP)
	TypeArme  string // ex: "épée", "hache", "bâton"
	Prix      int    // prix en diamants chez le forgeron
	Niveau    int    // niveau d'amélioration (forgeron)
}

func (a Arme) Nom() string { return a.NomArme }

func (a Arme) Description() string {
	texte := a.NomArme
	if a.TypeArme != "" {
		texte += fmt.Sprintf(" (%s)", a.TypeArme)
	}
	texte += fmt.Sprintf(" — %d dégâts", a.Degats)
	if a.Puissance > 0 {
		texte += fmt.Sprintf(" +%d puissance", a.Puissance)
	}
	return texte + fmt.Sprintf(" [Niv.%d]", a.Niveau)
}

// CoutAmelioration indique le prix en diamants pour améliorer l'arme chez le forgeron.
func (a *Arme) CoutAmelioration() int {
	return 20 * a.Niveau
}

// Ameliorer augmente les dégâts de l'arme et son niveau.
func (a *Arme) Ameliorer() {
	a.Degats += 3 + a.Niveau
	a.Niveau++
}

// CatalogueArmes contient les armes achetables chez le forgeron.
var CatalogueArmes = []Arme{
	{NomArme: "Épée en fer", Degats: 8, TypeArme: "épée", Prix: 40, Niveau: 1},
	{NomArme: "Hache de guerre", Degats: 10, TypeArme: "hache", Prix: 55, Niveau: 1},
	{NomArme: "Bâton runique", Degats: 7, TypeArme: "bâton", Prix: 50, Niveau: 1},
	{NomArme: "Dague empoisonnée", Degats: 6, TypeArme: "dague", Prix: 30, Niveau: 1},
}
