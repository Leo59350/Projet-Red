package classe

type Classe struct {
	Nom        string
	AD         int
	AP         int
	PvMax      int
	ArmeDeBase []any
}

func NewClasse(nom string, att int, puiss int, pvmax int, arme []any) *Classe {
	return &Classe{
		Nom:        nom,
		AD:         att,
		AP:         puiss,
		PvMax:      pvmax,
		ArmeDeBase: arme,
	}
}

var ListeClasse []Classe

func CreerListeClasse() {
	ListeClasse = []Classe{
		*NewClasse("Le Chevalier", 12, 0, 59, []any{"Épée", "arme", 1, 1, 18, 0}),
		*NewClasse("Le Mage", 4, 12, 35, []any{"Bâton de Bois", "arme", 1, 1, 1, 21}),
		*NewClasse("L'Archer", 9, 8, 40, []any{"Arc", "arme", 1, 1, 18, 7}),
		*NewClasse("Le Voleur", 13, 4, 45, []any{"Dague Pointue", "arme", 1, 1, 12, 0}),
		*NewClasse("Le Géant", 8, 0, 72, []any{"Le Tronc d'Arbre", "arme", 1, 1, 12, 0}),
	}
}
