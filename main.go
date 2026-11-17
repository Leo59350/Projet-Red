package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"projetred/affichage"
	"projetred/boutique"
	"projetred/classe"
	"projetred/combat"
	"projetred/ennemi"
	boss "projetred/monstre"
	"projetred/personnage"
)

// ---------- Intro ----------

// afficherIntro joue une animation de "décryptage" du titre : les lettres
// de FANTAISIE sont d'abord des symboles aléatoires puis se révèlent une
// par une. Pas de barre de téléchargement : juste un effet qui donne
// l'impression que le titre s'affiche tout seul.
func afficherIntro() {
	rand.Seed(time.Now().UnixNano())

	titre := "FANTAISIE"
	symboles := []rune("!@#$%&*+=/<>?01")

	for etape := 0; etape <= len(titre); etape++ {
		var construit strings.Builder
		for i, r := range titre {
			if i < etape {
				construit.WriteRune(r)
			} else {
				construit.WriteRune(symboles[rand.Intn(len(symboles))])
			}
		}

		affichage.Effacer()
		fmt.Println(affichage.Titre("\n\n\n\n              ⚔️   " + construit.String() + "   🐉"))
		fmt.Println(affichage.Doré("\n              Le Royaume Oublié"))
		time.Sleep(90 * time.Millisecond)
	}

	time.Sleep(500 * time.Millisecond)
}

// afficherArtistesCaches révèle l'easter egg du jeu avant de commencer l'aventure.
func afficherArtistesCaches(reader *bufio.Reader) {
	affichage.Effacer()
	affichage.Section("🎭 ARTISTES CACHÉS DU JEU")
	fmt.Println(affichage.Doré("Deux légendes se cachent dans cette aventure :\n"))
	fmt.Println(affichage.Doré("  🥇 Artiste caché n°1 : ABBA"))
	fmt.Println(affichage.Doré("  🥈 Artiste caché n°2 : Steven Spielberg"))
	affichage.Fin()
	affichage.LireLigne(reader, affichage.Doré("\nAppuie sur Entrée pour commencer l'aventure..."))
}

// ---------- Création du personnage ----------

func choisirClasse(reader *bufio.Reader) classe.Classe {
	classe.CreerListeClasse()
	classes := classe.ListeClasse

	fmt.Println(affichage.Doré("\nChoisis ta classe :"))
	for i, c := range classes {
		fmt.Println(affichage.Doré(fmt.Sprintf("  [%d] %-14s PV: %-3d AD: %-3d AP: %d", i+1, c.Nom, c.PvMax, c.AD, c.AP)))
	}

	for {
		choix := affichage.LireLigne(reader, "\nTon choix (1-"+strconv.Itoa(len(classes))+") : ")
		num, err := strconv.Atoi(choix)
		if err == nil && num >= 1 && num <= len(classes) {
			return classes[num-1]
		}
		fmt.Println(affichage.TexteRouge("⚠️  Choix invalide, réessaie."))
	}
}

// ---------- Menu principal ----------

func afficherMenu(hero *personnage.Personnage) {
	affichage.Effacer()
	affichage.Section("⚔️  LE ROYAUME OUBLIÉ — MENU PRINCIPAL  🐉")
	fmt.Println(affichage.Doré(fmt.Sprintf("Diamants : %d\n", hero.Diamants)))

	fmt.Println(affichage.Doré(" [1] 📜 Profil"))
	fmt.Println(affichage.Doré(" [2] 🛒 Marchand"))
	fmt.Println(affichage.Doré(" [3] 🔨 Forgeron"))
	fmt.Println(affichage.Doré(" [4] ⚔️  Combat"))
	fmt.Println(affichage.Doré(" [5] 🚪 Quitter"))
	if hero.BossDisponible() {
		fmt.Println(affichage.TexteRouge(" [6] 👹 Combattre le boss"))
	}
	affichage.Fin()
}

func menuProfil(reader *bufio.Reader, hero *personnage.Personnage) {
	affichage.Effacer()
	hero.AfficherStats()
	hero.Inventaire.Afficher()

	if len(hero.Inventaire.ArmesStockees) > 0 {
		choix := affichage.LireLigne(reader, affichage.Doré("\nÉquiper une arme de la réserve ? (numéro, ou Entrée pour passer) : "))
		if choix != "" {
			idx, err := strconv.Atoi(choix)
			if err == nil && idx >= 1 && idx <= len(hero.Inventaire.ArmesStockees) {
				nom := hero.Inventaire.ArmesStockees[idx-1].NomArme
				hero.Inventaire.EquiperArmeParNom(nom)
				hero.SyncArmeEquipee()
				fmt.Println(affichage.TexteVert(fmt.Sprintf("✅ %s équipée !", nom)))
			} else {
				fmt.Println(affichage.TexteRouge("⚠️  Choix invalide."))
			}
		}
	}

	affichage.LireLigne(reader, affichage.Doré("\nAppuie sur Entrée pour revenir au menu..."))
}

func ecranCombat(reader *bufio.Reader, hero *personnage.Personnage, e combat.Adversaire) {
	combat.Combattre(reader, hero, e)
	affichage.LireLigne(reader, affichage.Doré("\nAppuie sur Entrée pour revenir au menu..."))
}

func menuPrincipal(reader *bufio.Reader, hero *personnage.Personnage) {
	for {
		afficherMenu(hero)

		switch affichage.LireLigne(reader, "Ton choix : ") {
		case "1":
			menuProfil(reader, hero)
		case "2":
			boutique.OuvrirMarchand(reader, hero)
		case "3":
			boutique.OuvrirForgeron(reader, hero)
		case "4":
			monstre, err := ennemi.SpawnRandomMonster(hero.Niveau)
			if err != nil {
				fmt.Println(affichage.TexteRouge("⚠️  " + err.Error()))
				affichage.LireLigne(reader, affichage.Doré("\nAppuie sur Entrée pour revenir au menu..."))
				continue
			}
			ecranCombat(reader, hero, monstre)
		case "5":
			fmt.Println(affichage.Titre("À bientôt, aventurier ! 👋"))
			return
		case "6":
			if !hero.BossDisponible() {
				fmt.Println(affichage.TexteRouge("⚠️  Choix invalide."))
				continue
			}
			b, err := boss.GetBossByLevel(hero.Niveau)
			if err != nil {
				fmt.Println(affichage.TexteRouge("⚠️  " + err.Error()))
				affichage.LireLigne(reader, affichage.Doré("\nAppuie sur Entrée pour revenir au menu..."))
				continue
			}
			ecranCombat(reader, hero, b)
		default:
			fmt.Println(affichage.TexteRouge("⚠️  Choix invalide."))
		}
	}
}

// ---------- Programme principal ----------

func main() {
	reader := bufio.NewReader(os.Stdin)

	afficherIntro()
	afficherArtistesCaches(reader)

	affichage.Effacer()
	affichage.Section("⚔️  LE ROYAUME OUBLIÉ  🐉")

	nom := ""
	for nom == "" {
		nom = affichage.LireLigne(reader, affichage.Doré("\nEntre ton nom : "))
		if nom == "" {
			fmt.Println(affichage.TexteRouge("⚠️  Le nom ne peut pas être vide."))
		}
	}

	laClasse := choisirClasse(reader)

	affichage.Fin()
	fmt.Println(affichage.Doré(fmt.Sprintf("\nBienvenue, %s ! Tu joues : %s", nom, laClasse.Nom)))
	fmt.Println(affichage.Doré("Ton aventure commence maintenant !"))
	time.Sleep(800 * time.Millisecond)

	hero := personnage.Creer(nom, laClasse)
	menuPrincipal(reader, hero)
}
