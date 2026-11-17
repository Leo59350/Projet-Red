package combat

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"projetred/affichage"
	"projetred/item"
	"projetred/personnage"
)

// Adversaire est tout ce que le héros peut combattre : un monstre
// (ennemi.Monster) ou un boss (boss.Boss).
type Adversaire interface {
	IsAlive() bool
	TakeDamage(amount int)
	Attack() int
	GetName() string
	GetHP() (actuels, max int)
	XPReward() int
	DiamondsReward() int
}

// Combattre lance un combat au tour par tour entre le héros et un ennemi.
// Renvoie true si le héros gagne, false s'il perd ou s'il fuit.
func Combattre(reader *bufio.Reader, hero *personnage.Personnage, e Adversaire) bool {
	for hero.EstVivant() && e.IsAlive() {
		afficherCombat(hero, e)

		joueurAJoue := true

		switch affichage.LireLigne(reader, affichage.Doré("Ton choix : ")) {

		case "1":
			degats := hero.DegatsAttaque() + rand.Intn(5) - 2
			if degats < 1 {
				degats = 1
			}
			e.TakeDamage(degats)
			afficherAction(fmt.Sprintf("⚔️  Tu infliges %d dégâts à %s !", degats, e.GetName()), affichage.Vert)

		case "2":
			joueurAJoue = utiliserPotion(reader, hero)

		case "3":
			if rand.Intn(2) == 0 {
				afficherAction("🏃 Tu prends la fuite avec succès !", affichage.Or)
				return false
			}
			afficherAction("❌ La fuite échoue !", affichage.Rouge)

		default:
			afficherAction("⚠️  Choix invalide.", affichage.Rouge)
			joueurAJoue = false
		}

		if !e.IsAlive() {
			break
		}
		if !joueurAJoue {
			continue // choix invalide ou potion annulée : on ne perd pas son tour
		}

		// Tour de l'ennemi.
		time.Sleep(500 * time.Millisecond)
		fmt.Println()
		degatsEnnemi := e.Attack() // l'ennemi affiche lui-même son attaque
		hero.Pv -= degatsEnnemi
		if hero.Pv < 0 {
			hero.Pv = 0
		}
		time.Sleep(1200 * time.Millisecond)
	}

	return terminerCombat(hero, e)
}

// afficherCombat montre clairement les deux camps : nom, barre de vie
// (bleue pour le héros, rouge pour l'ennemi) et PV, puis le menu d'actions.
func afficherCombat(hero *personnage.Personnage, e Adversaire) {
	affichage.Effacer()
	affichage.Section("⚔️  COMBAT")

	fmt.Println(affichage.Doré(fmt.Sprintf("🛡️  %s (toi)", hero.Nom)))
	fmt.Println(fmt.Sprintf("   %s  %d/%d PV", affichage.BarreHeros(hero.Pv, hero.PvMax, 24), hero.Pv, hero.PvMax))

	pvE, pvEMax := e.GetHP()
	fmt.Println()
	fmt.Println(affichage.TexteRouge(fmt.Sprintf("👹 %s", e.GetName())))
	fmt.Println(fmt.Sprintf("   %s  %d/%d PV", affichage.BarreEnnemi(pvE, pvEMax, 24), pvE, pvEMax))

	affichage.Fin()
	fmt.Println(affichage.Doré("\n [1] ⚔️  Attaquer    [2] 🧪 Potion    [3] 🏃 Fuir"))
}

// afficherAction affiche un message d'action coloré puis marque une courte
// pause, pour que chaque tour reste lisible au lieu de défiler trop vite.
func afficherAction(message, couleur string) {
	fmt.Println("\n" + couleur + message + affichage.Reset)
	time.Sleep(400 * time.Millisecond)
}

// utiliserPotion affiche les potions disponibles et applique celle choisie.
// Renvoie false si aucune potion n'a été utilisée (pour ne pas faire perdre
// un tour au joueur inutilement).
func utiliserPotion(reader *bufio.Reader, hero *personnage.Personnage) bool {
	if len(hero.Inventaire.Potions) == 0 {
		afficherAction("❌ Tu n'as aucune potion.", affichage.Rouge)
		return false
	}

	fmt.Println(affichage.Doré("\nTes potions :"))
	for i, s := range hero.Inventaire.Potions {
		fmt.Println(affichage.Doré(fmt.Sprintf("   [%d] %s x%d/%d", i+1, s.Potion.NomPotion, s.Quantite, item.LimiteStackPotion)))
	}

	choix := affichage.LireLigne(reader, "Laquelle utiliser ? (0 pour annuler) : ")
	if choix == "0" {
		return false
	}

	idx, err := strconv.Atoi(choix)
	if err != nil || idx < 1 || idx > len(hero.Inventaire.Potions) {
		afficherAction("⚠️  Choix invalide.", affichage.Rouge)
		return false
	}

	nom := hero.Inventaire.Potions[idx-1].Potion.NomPotion
	hero.UtiliserPotion(nom)
	afficherAction(fmt.Sprintf("🧪 Tu utilises %s.", nom), affichage.Vert)
	return true
}

// terminerCombat affiche l'issue du combat (victoire ou défaite) et
// applique les récompenses ou les conséquences de la défaite.
func terminerCombat(hero *personnage.Personnage, e Adversaire) bool {
	affichage.Effacer()

	if !e.IsAlive() {
		affichage.Section("🎉 VICTOIRE")
		fmt.Println(affichage.TexteVert(fmt.Sprintf("Tu as vaincu %s !", e.GetName())))
		hero.GagnerXP(e.XPReward())
		hero.GagnerDiamants(e.DiamondsReward())
		fmt.Println(affichage.Doré(fmt.Sprintf("+%d XP    +%d 💎 Diamants", e.XPReward(), e.DiamondsReward())))
		affichage.Fin()
		return true
	}

	affichage.Section("💀 DÉFAITE")
	fmt.Println(affichage.TexteRouge(fmt.Sprintf("%s a été vaincu par %s...", hero.Nom, e.GetName())))
	hero.EstMort()
	fmt.Println(affichage.Doré(fmt.Sprintf("Tu te réveilles avec %d/%d PV.", hero.Pv, hero.PvMax)))
	affichage.Fin()
	return false
}