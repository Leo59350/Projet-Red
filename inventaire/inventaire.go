package inventaire

import (
	"fmt"

	"projetred/affichage"
	"projetred/item"
)

// LimiteInventaire est le nombre maximum d'emplacements dans l'inventaire..
const LimiteInventaire = 10

// StockPotion représente une potion possédée et sa quantité.
type StockPotion struct {
	Potion   item.Potion
	Quantite int
}

// Inventaire regroupe tout ce que possède un personnage :
type Inventaire struct {
	Potions       []StockPotion
	ArmeEquipee   *item.Arme
	ArmesStockees []item.Arme
}

// Nouveau crée un inventaire vide correctement initialisé.
func Nouveau() Inventaire {
	return Inventaire{
		Potions:       []StockPotion{},
		ArmesStockees: []item.Arme{},
	}
}

// NombreEmplacements renvoie le nombre d'emplacements actuellement occupés
// (une pile de potions = 1 emplacement, une arme en réserve = 1 emplacement).
func (inv *Inventaire) NombreEmplacements() int {
	return len(inv.Potions) + len(inv.ArmesStockees)
}

// EmplacementsLibres renvoie le nombre d'emplacements encore disponibles
// avant d'atteindre la limite de l'inventaire (le "countdown").
func (inv *Inventaire) EmplacementsLibres() int {
	return LimiteInventaire - inv.NombreEmplacements()
}

// AjouterPotion ajoute une (ou plusieurs) potion(s) dans l'inventaire.
// Chaque pile est limitée à item.LimiteStackPotion exemplaires : les piles
// existantes sont d'abord complétées (ça ne prend pas de nouvel emplacement),
// puis de nouvelles piles sont créées pour le surplus tant qu'il reste de la
// place dans l'inventaire (LimiteInventaire).
// Renvoie la quantité réellement ajoutée (peut être 0 si tout est plein).
func (inv *Inventaire) AjouterPotion(p item.Potion, quantite int) int {
	ajoutTotal := 0

	// On remplit d'abord les piles déjà existantes de ce type de potion.
	for i := range inv.Potions {
		if inv.Potions[i].Potion.NomPotion != p.NomPotion {
			continue
		}
		placeRestante := item.LimiteStackPotion - inv.Potions[i].Quantite
		if placeRestante <= 0 {
			continue
		}
		ajout := quantite
		if ajout > placeRestante {
			ajout = placeRestante
		}
		inv.Potions[i].Quantite += ajout
		quantite -= ajout
		ajoutTotal += ajout
		if quantite == 0 {
			return ajoutTotal
		}
	}

	// Ce qu'il reste forme de nouvelles piles (5 maximum chacune),
	// tant qu'il reste un emplacement libre dans l'inventaire.
	for quantite > 0 && inv.EmplacementsLibres() > 0 {
		taille := quantite
		if taille > item.LimiteStackPotion {
			taille = item.LimiteStackPotion
		}
		inv.Potions = append(inv.Potions, StockPotion{Potion: p, Quantite: taille})
		quantite -= taille
		ajoutTotal += taille
	}

	return ajoutTotal
}

// RetirerPotion consomme un exemplaire de la potion nommée.
// Renvoie la potion consommée et true si elle était disponible.
func (inv *Inventaire) RetirerPotion(nom string) (item.Potion, bool) {
	for i := range inv.Potions {
		if inv.Potions[i].Potion.NomPotion == nom && inv.Potions[i].Quantite > 0 {
			p := inv.Potions[i].Potion
			inv.Potions[i].Quantite--
			if inv.Potions[i].Quantite == 0 {
				inv.Potions = append(inv.Potions[:i], inv.Potions[i+1:]...)
			}
			return p, true
		}
	}
	return item.Potion{}, false
}

// AjouterArme range une nouvelle arme dans le stock (non équipée), si
// l'inventaire a encore de la place. Renvoie false si l'inventaire est plein.
func (inv *Inventaire) AjouterArme(a item.Arme) bool {
	if inv.EmplacementsLibres() <= 0 {
		return false
	}
	inv.ArmesStockees = append(inv.ArmesStockees, a)
	return true
}

// EquiperArme équipe directement une arme ; l'ancienne (s'il y en a une)
// retourne dans la réserve. Le nombre d'emplacements occupés ne change pas
// (on retire l'arme équipée, on y remet l'ancienne).
func (inv *Inventaire) EquiperArme(a item.Arme) {
	if inv.ArmeEquipee != nil {
		inv.ArmesStockees = append(inv.ArmesStockees, *inv.ArmeEquipee)
	}
	arme := a
	inv.ArmeEquipee = &arme
}

// EquiperArmeParNom équipe une arme présente dans la réserve, par son nom.
func (inv *Inventaire) EquiperArmeParNom(nom string) bool {
	for i, a := range inv.ArmesStockees {
		if a.NomArme == nom {
			inv.ArmesStockees = append(inv.ArmesStockees[:i], inv.ArmesStockees[i+1:]...)
			inv.EquiperArme(a)
			return true
		}
	}
	return false
}

// Afficher liste le contenu de l'inventaire (arme équipée, réserve, potions),
// avec le compteur d'emplacements et le même habillage doré que le reste
// de l'interface.
func (inv Inventaire) Afficher() {
	affichage.Section("🎒 INVENTAIRE")
	fmt.Println(affichage.Doré(fmt.Sprintf(" Emplacements : %d/%d", inv.NombreEmplacements(), LimiteInventaire)))

	if inv.ArmeEquipee != nil {
		fmt.Println(affichage.Doré(fmt.Sprintf(" ⚔️  Arme équipée : %s", inv.ArmeEquipee.Description())))
	} else {
		fmt.Println(affichage.Doré(" ⚔️  Arme équipée : (aucune)"))
	}

	if len(inv.ArmesStockees) > 0 {
		fmt.Println(affichage.Doré(" 🛡️  Armes en réserve :"))
		for i, a := range inv.ArmesStockees {
			fmt.Println(affichage.Doré(fmt.Sprintf("   [%d] %s", i+1, a.Description())))
		}
	}

	fmt.Println(affichage.Doré(" 🧪 Potions :"))
	if len(inv.Potions) == 0 {
		fmt.Println(affichage.Doré("   (aucune)"))
	}
	for i, s := range inv.Potions {
		fmt.Println(affichage.Doré(fmt.Sprintf("   [%d] %s x%d/%d", i+1, s.Potion.NomPotion, s.Quantite, item.LimiteStackPotion)))
	}

	affichage.Fin()
}