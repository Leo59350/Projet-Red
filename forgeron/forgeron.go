package forgeron

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

func afficherMenu(hero *personnage.Personnage) {
	couleur.Effacer()
	fmt.Println(couleur.Titre("\n╔══════════════ 🔨 FORGERON ════════════════╗"))
	fmt.Println(couleur.Doré(fmt.Sprintf(" 💎 Diamants : %d", hero.Diamants)))
	fmt.Println(couleur.Titre("╠════════════════════════════════════════════╣"))

	if hero.Inventaire.ArmeEquipee != nil {
		cout := hero.Inventaire.ArmeEquipee.CoutAmelioration()
		fmt.Println(couleur.Doré(fmt.Sprintf(" Arme équipée : %s", hero.Inventaire.ArmeEquipee.Description())))
		fmt.Println(couleur.Doré(fmt.Sprintf(" [1] Améliorer cette arme (%d💎)", cout)))
	} else {
		fmt.Println(couleur.Doré(" Arme équipée : (aucune)"))
	}
	fmt.Println(couleur.Doré(" [2] Acheter une nouvelle arme"))
	fmt.Println(couleur.Doré(" [0] Retour au menu"))
	fmt.Println(couleur.Titre("╚════════════════════════════════════════════╝"))
}

func afficherCatalogueArmes() {
	fmt.Println(couleur.Titre("\n--- Armes disponibles ---"))
	for i, a := range item.CatalogueArmes {
		fmt.Println(couleur.Doré(fmt.Sprintf("   [%d] %s — %d💎", i+1, a.Description(), a.Prix)))
	}
	fmt.Println(couleur.Doré("   [0] Retour"))
}

// Ouvrir lance l'atelier du forgeron : amélioration de l'arme équipée
// ou achat d'une arme neuve, contre des diamants.
func Ouvrir(reader *bufio.Reader, hero *personnage.Personnage) {
	for {
		afficherMenu(hero)

		switch lireLigne(reader, "Ton choix : ") {

		case "1":
			if hero.Inventaire.ArmeEquipee == nil {
				fmt.Println(couleur.TexteRouge("❌ Aucune arme équipée."))
				continue
			}
			cout := hero.Inventaire.ArmeEquipee.CoutAmelioration()
			if !hero.DepenserDiamants(cout) {
				fmt.Println(couleur.TexteRouge("❌ Pas assez de diamants."))
				continue
			}
			hero.Inventaire.ArmeEquipee.Ameliorer()
			fmt.Println(couleur.TexteVert(fmt.Sprintf("✅ Arme améliorée ! %s", hero.Inventaire.ArmeEquipee.Description())))

		case "2":
			afficherCatalogueArmes()
			choix := lireLigne(reader, "Ton choix : ")
			if choix == "0" {
				continue
			}
			idx, err := strconv.Atoi(choix)
			if err != nil || idx < 1 || idx > len(item.CatalogueArmes) {
				fmt.Println(couleur.TexteRouge("⚠️  Choix invalide."))
				continue
			}
			arme := item.CatalogueArmes[idx-1]
			if !hero.DepenserDiamants(arme.Prix) {
				fmt.Println(couleur.TexteRouge("❌ Pas assez de diamants."))
				continue
			}
			hero.Inventaire.AjouterArme(arme)
			fmt.Println(couleur.TexteVert(fmt.Sprintf("✅ %s achetée ! Équipe-la depuis ton profil.", arme.NomArme)))

		case "0":
			return

		default:
			fmt.Println(couleur.TexteRouge("⚠️  Choix invalide."))
		}
	}
}
