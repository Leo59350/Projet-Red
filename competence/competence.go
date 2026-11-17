package competence

// Competence représente une capacité utilisable en combat.
type Competence struct {
	Nom       string
	Puissance int
	CoutMana  int
}

// Utiliser applique la compétence (logique à compléter plus tard).
func (c Competence) Utiliser() int {
	return c.Puissance
}
