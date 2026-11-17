package marchand

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"projetred/couleur"
	"projetred/item"
	"projetred/personnage"
)

func lireLigne(reader *bufio.Reader, question string) string {
	fmt.Print(question)
	reponse, _ := reader.ReadString('\n')
	return strings.TrimSpace(reponse)
}

func afficherCatalogue(hero *personnage.Personnage) {
	couleur.Effacer()
	fmt.Println(couleur.Titre("\n╔══════════════ 🛒 MARCHAND ═══════════════╗"))
	fmt.Println(couleur.Doré(fmt.Sprintf(" 💎 Diamants : %d", hero.Diamants)))
	fmt.Println(couleur.Titre("╠═══════════════════════════════════════════╣"))
	fmt.Println(couleur.Doré(" Potions disponibles :"))
	for i, p := range item.CataloguePotions {
		fmt.Println(couleur.Doré(fmt.Sprintf("   [%d] %s", i+1, p.Description())))
	}
	fmt.Println(couleur.Doré(" [0] Retour au menu"))
	fmt.Println(couleur.Titre("╚═════════════════════════════════════════╝"))
}

// Ouvrir lance la boutique du marchand : achat de potions contre des diamants.
func Ouvrir(reader *bufio.Reader, hero *personnage.Personnage) {
	for {
		afficherCatalogue(hero)
		choix := lireLigne(reader, "Ton choix : ")
		if choix == "0" {
			return
		}

		idx, err := strconv.Atoi(choix)
		if err != nil || idx < 1 || idx > len(item.CataloguePotions) {
			fmt.Println(couleur.TexteRouge("⚠️  Choix invalide."))
			continue
		}

		potion := item.CataloguePotions[idx-1]
		if !hero.DepenserDiamants(potion.Prix) {
			fmt.Println(couleur.TexteRouge("❌ Pas assez de diamants."))
			continue
		}

		hero.Inventaire.AjouterPotion(potion, 1)
		fmt.Println(couleur.TexteVert(fmt.Sprintf("✅ %s achetée !", potion.NomPotion)))
	}
}
