package personnage

import (
	"fmt"
	"strings"

	"projetred/affichage"
	"projetred/classe"
	"projetred/inventaire"
	"projetred/item"
	boss "projetred/monstre"
)

type Personnage struct {
	Nom         string
	Classe      classe.Classe
	Niveau      int
	PvMax       int
	Pv          int
	AD          int
	AP          int
	Inventaire  inventaire.Inventaire
	ArmeEquipee []any

	// Champs nécessaires au reste du jeu (combat, boutique, montée de niveau).
	XP       int
	XPMax    int
	Diamants int
}

func NewPersonnage(nom string, classe classe.Classe, niveau int, pvmax int, pv int, att int, puiss int, inv inventaire.Inventaire, arme []any) *Personnage {
	return &Personnage{
		Nom:         nom,
		Classe:      classe,
		Niveau:      niveau,
		PvMax:       pvmax,
		Pv:          pv,
		AD:          att,
		AP:          puiss,
		Inventaire:  inv,
		ArmeEquipee: arme,

		XP:       0,
		XPMax:    50,
		Diamants: 30,
	}
}

// Creer construit un nouveau héros de niveau 1 à partir d'une classe :
// stats de la classe, arme de base équipée et 2 potions de soin.
func Creer(nom string, c classe.Classe) *Personnage {
	inv := inventaire.Nouveau()

	// Copie de l'arme de base pour ne pas modifier la définition de la classe.
	armeBase := append([]any(nil), c.ArmeDeBase...)
	if a, ok := armeDepuisSlice(armeBase); ok {
		inv.EquiperArme(a)
	}
	inv.AjouterPotion(item.PotionSoin, 2)

	return NewPersonnage(nom, c, 1, c.PvMax, c.PvMax, c.AD, c.AP, inv, armeBase)
}

// armeDepuisSlice convertit une arme au format []any
// {nom, "arme", niveau, quantité, bonus AD, bonus AP} en item.Arme.
func armeDepuisSlice(d []any) (item.Arme, bool) {
	if len(d) < 6 {
		return item.Arme{}, false
	}
	nom, ok0 := d[0].(string)
	niveau, ok2 := d[2].(int)
	ad, ok4 := d[4].(int)
	ap, ok5 := d[5].(int)
	if !ok0 || !ok2 || !ok4 || !ok5 {
		return item.Arme{}, false
	}
	return item.Arme{NomArme: nom, Degats: ad, Puissance: ap, Niveau: niveau}, true
}

// sliceDepuisArme fait l'inverse : item.Arme -> format []any.
func sliceDepuisArme(a item.Arme) []any {
	return []any{a.NomArme, "arme", a.Niveau, 1, a.Degats, a.Puissance}
}

// SyncArmeEquipee remet ArmeEquipee (format []any) à jour avec l'arme réellement
// équipée dans l'inventaire. À appeler après un changement d'arme ou une amélioration.
func (p *Personnage) SyncArmeEquipee() {
	if p.Inventaire.ArmeEquipee == nil {
		p.ArmeEquipee = nil
		return
	}
	p.ArmeEquipee = sliceDepuisArme(*p.Inventaire.ArmeEquipee)
}

// AfficherStats affiche la fiche de personnage en doré, avec barres de vie/XP.
func (p *Personnage) AfficherStats() {
	affichage.Section(fmt.Sprintf("⚜️  FICHE DE %s", strings.ToUpper(p.Nom)))
	fmt.Println(affichage.Doré(fmt.Sprintf(" Classe    : %s", p.Classe.Nom)))
	fmt.Println(affichage.Doré(fmt.Sprintf(" Niveau    : %d", p.Niveau)))
	fmt.Println(affichage.Doré(fmt.Sprintf(" PV        : %s %d/%d", affichage.BarreDeVie(p.Pv, p.PvMax, 20), p.Pv, p.PvMax)))
	fmt.Println(affichage.Doré(fmt.Sprintf(" XP        : %s %d/%d", affichage.BarreXP(p.XP, p.XPMax, 20), p.XP, p.XPMax)))
	fmt.Println(affichage.Doré(fmt.Sprintf(" AD (attaque)   : %d", p.AD)))
	fmt.Println(affichage.Doré(fmt.Sprintf(" AP (puissance) : %d", p.AP)))
	fmt.Println(affichage.Doré(fmt.Sprintf(" 💎 Diamants : %d", p.Diamants)))
	affichage.Fin()
}

// BossDisponible : vrai tous les 10 niveaux (10, 20, 30... jusqu'à 100).
func (p *Personnage) BossDisponible() bool {
	return boss.IsBossLevel(p.Niveau)
}

// EstVivant indique si le personnage a encore des points de vie.
func (p *Personnage) EstVivant() bool {
	return p.Pv > 0
}

// EstMort vérifie si le personnage est à 0 point de vie. Si c'est le cas,
// il "meurt" puis ressuscite immédiatement avec 50% de ses PV max.
// Renvoie true si le personnage était mort (et vient d'être ressuscité).
func (p *Personnage) EstMort() bool {
	if p.Pv > 0 {
		return false
	}
	p.Pv = p.PvMax / 2
	return true
}

// --- Implémentation de item.Cible : effets appliqués par les potions ---

func (p *Personnage) Soigner(montant int) {
	p.Pv += montant
	if p.Pv > p.PvMax {
		p.Pv = p.PvMax
	}
}

// AugmenterForce augmente l'attaque physique (AD).
func (p *Personnage) AugmenterForce(montant int) {
	p.AD += montant
}

// AugmenterPuissance augmente la puissance magique (AP).
func (p *Personnage) AugmenterPuissance(montant int) {
	p.AP += montant
}

// UtiliserPotion consomme une potion de l'inventaire (par son nom) et applique son effet.
func (p *Personnage) UtiliserPotion(nom string) bool {
	potion, ok := p.Inventaire.RetirerPotion(nom)
	if !ok {
		return false
	}
	return potion.Consommer(p)
}

// DegatsAttaque calcule les dégâts infligés lors d'une attaque
// (AD + AP du héros + bonus de l'arme équipée).
func (p *Personnage) DegatsAttaque() int {
	base := p.AD + p.AP
	if p.Inventaire.ArmeEquipee != nil {
		base += p.Inventaire.ArmeEquipee.Degats + p.Inventaire.ArmeEquipee.Puissance
	}
	return base
}

// GagnerXP ajoute de l'expérience et gère la montée de niveau (stats en hausse, soin complet).
func (p *Personnage) GagnerXP(montant int) {
	p.XP += montant
	for p.XP >= p.XPMax {
		p.XP -= p.XPMax
		p.Niveau++
		p.XPMax += 25
		p.PvMax += 15
		p.AD += 3
		p.AP += 2
		p.Pv = p.PvMax
		fmt.Println(affichage.Titre(fmt.Sprintf("\n✨ %s passe au niveau %d !", p.Nom, p.Niveau)))
	}
}

// GagnerDiamants ajoute des diamants au personnage.
func (p *Personnage) GagnerDiamants(montant int) {
	p.Diamants += montant
}

// DepenserDiamants tente de dépenser des diamants ; renvoie false si le solde est insuffisant.
func (p *Personnage) DepenserDiamants(montant int) bool {
	if p.Diamants < montant {
		return false
	}
	p.Diamants -= montant
	return true
}
