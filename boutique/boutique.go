// Package boutique regroupe le marchand (achat de potions) et le forgeron
// (amélioration/achat d'armes) : deux écrans qui fonctionnent pareil
// (afficher un catalogue, dépenser des diamants), regroupés ici pour ne
// pas dupliquer la même logique dans deux fichiers séparés.
package boutique

import (
	"bufio"
	"fmt"
	"strconv"

	"projetred/affichage"
	"projetred/item"
	"projetred/personnage"
)

// ---------- Marchand : achat de potions ----------

// OuvrirMarchand permet d'acheter des potions contre des diamants.
// OuvrirMarchand affiche la boutique et permet d'acheter des potions
// avec des diamants. Elle tourne jusqu'à ce que le joueur tape 0.
// reader : sert à lire ce que le joueur tape
// hero   : le personnage (pointeur, donc ses diamants et son inventaire sont modifiés directement)
func OuvrirMarchand(reader *bufio.Reader, hero *personnage.Personnage) {

	// Boucle infinie : le marchand reste ouvert tant qu'on ne fait pas "return".
	for {

		// ----- 1. AFFICHAGE -----
		affichage.Effacer()                 // efface l'écran
		affichage.Section("🛒 MARCHAND")    // affiche le titre encadré
		fmt.Println(affichage.Doré(fmt.Sprintf("💎 Diamants : %d\n", hero.Diamants))) // montre l'argent du joueur

		// On parcourt le catalogue de potions (défini dans item/potion.go).
		// i = position dans la liste (commence à 0), p = la potion.
		// On affiche i+1 pour que le joueur voie [1], [2], [3].
		for i, p := range item.CataloguePotions {
			fmt.Println(affichage.Doré(fmt.Sprintf(" [%d] %s", i+1, p.Description())))
		}
		fmt.Println(affichage.Doré(" [0] Retour au menu"))
		affichage.Fin() // ferme le cadre

		// ----- 2. LECTURE DU CHOIX -----
		// Le joueur tape quelque chose, on le récupère sous forme de texte.
		choix := affichage.LireLigne(reader, "Ton choix : ")

		// Si c'est "0", on quitte la fonction : retour au menu principal.
		if choix == "0" {
			return
		}

		// ----- 3. VÉRIFICATION DU NUMÉRO -----
		// Atoi convertit le texte en nombre. Si ce n'est pas un nombre (ex : "abc"),
		// err contient une erreur.
		idx, err := strconv.Atoi(choix)

		// Choix refusé si : ce n'est pas un nombre, OU il est < 1, OU il dépasse la liste.
		if err != nil || idx < 1 || idx > len(item.CataloguePotions) {
			fmt.Println(affichage.TexteRouge("⚠️  Choix invalide."))
			continue // recommence la boucle depuis le début
		}

		// ----- 4. CHOIX DE LA POTION -----
		// Le joueur tape 1, mais une liste commence à 0 : on retire 1.
		potion := item.CataloguePotions[idx-1]

		// ----- 5. VÉRIFICATION DE L'ARGENT -----
		if hero.Diamants < potion.Prix {
			fmt.Println(affichage.TexteRouge("❌ Pas assez de diamants."))
			continue
		}

		// ----- 6. VÉRIFICATION DE LA PLACE -----
		// AjouterPotion ajoute la potion et renvoie combien ont été ajoutées.
		// Si elle renvoie 0, l'inventaire est plein.
		if hero.Inventaire.AjouterPotion(potion, 1) == 0 {
			fmt.Println(affichage.TexteRouge("❌ Inventaire plein ! (10/10 emplacements)"))
			continue
		}

		// ----- 7. PAIEMENT -----
		// Les diamants sont retirés en dernier, quand tout a réussi :
		// le joueur ne perd donc rien si l'achat échoue avant.
		hero.DepenserDiamants(potion.Prix)
		fmt.Println(affichage.TexteVert(fmt.Sprintf("✅ %s achetée !", potion.NomPotion)))
	}
}

// ---------- Forgeron : amélioration et achat d'armes ----------

// OuvrirForgeron permet d'améliorer l'arme équipée ou d'en acheter une nouvelle.
func OuvrirForgeron(reader *bufio.Reader, hero *personnage.Personnage) {
	for {
		affichage.Effacer()
		affichage.Section("🔨 FORGERON")
		fmt.Println(affichage.Doré(fmt.Sprintf("💎 Diamants : %d\n", hero.Diamants)))

		if hero.Inventaire.ArmeEquipee != nil {
			cout := hero.Inventaire.ArmeEquipee.CoutAmelioration()
			fmt.Println(affichage.Doré(fmt.Sprintf(" Arme équipée : %s", hero.Inventaire.ArmeEquipee.Description())))
			fmt.Println(affichage.Doré(fmt.Sprintf(" [1] Améliorer cette arme (%d💎)", cout)))
		} else {
			fmt.Println(affichage.Doré(" Arme équipée : (aucune)"))
		}
		fmt.Println(affichage.Doré(" [2] Acheter une nouvelle arme"))
		fmt.Println(affichage.Doré(" [0] Retour au menu"))
		affichage.Fin()

		switch affichage.LireLigne(reader, "Ton choix : ") {
		case "1":
			ameliorerArme(hero)
		case "2":
			acheterArme(reader, hero)
		case "0":
			return
		default:
			fmt.Println(affichage.TexteRouge("⚠️  Choix invalide."))
		}
	}
}

func ameliorerArme(hero *personnage.Personnage) {
	if hero.Inventaire.ArmeEquipee == nil {
		fmt.Println(affichage.TexteRouge("❌ Aucune arme équipée."))
		return
	}
	cout := hero.Inventaire.ArmeEquipee.CoutAmelioration()
	if !hero.DepenserDiamants(cout) {
		fmt.Println(affichage.TexteRouge("❌ Pas assez de diamants."))
		return
	}
	hero.Inventaire.ArmeEquipee.Ameliorer()
	hero.SyncArmeEquipee()
	fmt.Println(affichage.TexteVert(fmt.Sprintf("✅ Arme améliorée ! %s", hero.Inventaire.ArmeEquipee.Description())))
}

func acheterArme(reader *bufio.Reader, hero *personnage.Personnage) {
	fmt.Println(affichage.Titre("\n--- Armes disponibles ---"))
	for i, a := range item.CatalogueArmes {
		fmt.Println(affichage.Doré(fmt.Sprintf(" [%d] %s — %d💎", i+1, a.Description(), a.Prix)))
	}
	fmt.Println(affichage.Doré(" [0] Retour"))

	choix := affichage.LireLigne(reader, "Ton choix : ")
	if choix == "0" {
		return
	}

	idx, err := strconv.Atoi(choix)
	if err != nil || idx < 1 || idx > len(item.CatalogueArmes) {
		fmt.Println(affichage.TexteRouge("⚠️  Choix invalide."))
		return
	}

	arme := item.CatalogueArmes[idx-1]
	if hero.Diamants < arme.Prix {
		fmt.Println(affichage.TexteRouge("❌ Pas assez de diamants."))
		return
	}
	if !hero.Inventaire.AjouterArme(arme) {
		fmt.Println(affichage.TexteRouge("❌ Inventaire plein ! (10/10 emplacements)"))
		return
	}

	hero.DepenserDiamants(arme.Prix)
	fmt.Println(affichage.TexteVert(fmt.Sprintf("✅ %s achetée ! Équipe-la depuis ton profil.", arme.NomArme)))
}