package item

import "fmt"

// TypeEffet représente l'effet produit par une potion.
type TypeEffet string

const (
	EffetSoin      TypeEffet = "Soin"
	EffetForce     TypeEffet = "Force"
	EffetPuissance TypeEffet = "Puissance"
)

// Cible représente un personnage pouvant recevoir l'effet d'une potion.
type Cible interface {
	Soigner(montant int)
	AugmenterForce(montant int)
	AugmenterPuissance(montant int)
}

// Potion représente un consommable (soin, force, puissance...).
type Potion struct {
	NomPotion string
	Effet     TypeEffet
	Valeur    int
	Prix      int // prix en diamants chez le marchand
}

func (p Potion) Nom() string { return p.NomPotion }

func (p Potion) Description() string {
	return fmt.Sprintf("%s (+%d %s) — %d💎", p.NomPotion, p.Valeur, p.Effet, p.Prix)
}

// Consommer applique l'effet de la potion à la cible.
func (p Potion) Consommer(cible Cible) bool {
	if cible == nil {
		return false
	}

	switch p.Effet {
	case EffetSoin:
		cible.Soigner(p.Valeur)
	case EffetForce:
		cible.AugmenterForce(p.Valeur)
	case EffetPuissance:
		cible.AugmenterPuissance(p.Valeur)
	default:
		return false
	}
	return true
}

// LimiteStackPotion est le nombre maximum de potions identiques par pile.
// Au-delà, une nouvelle pile (un nouvel emplacement) est créée.
const LimiteStackPotion = 5

// Potions prédéfinies et catalogue vendu par le marchand.
var (
	PotionSoin      = Potion{NomPotion: "Potion de soin", Effet: EffetSoin, Valeur: 15, Prix: 10}
	PotionForce     = Potion{NomPotion: "Potion de force", Effet: EffetForce, Valeur: 5, Prix: 15}
	PotionPuissance = Potion{NomPotion: "Potion de puissance", Effet: EffetPuissance, Valeur: 5, Prix: 15}
)

// CataloguePotions contient les potions achetables chez le marchand.
var CataloguePotions = []Potion{PotionSoin, PotionForce, PotionPuissance}