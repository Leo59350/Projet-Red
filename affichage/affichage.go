// Package affichage regroupe TOUT ce qui concerne l'affichage dans le
// terminal : couleurs, barres de vie, titres de section, effacement de
// l'écran et lecture des saisies utilisateur. Tous les autres packages
// (combat, boutique, personnage, main...) passent par ici pour afficher
// quoi que ce soit, afin que le rendu reste identique partout dans le jeu.
package affichage

import (
	"bufio"
	"fmt"
	"strings"
)

// ---------- Couleurs ----------

const (
	Reset  = "\033[0m"
	Or     = "\033[38;5;220m" // doré : couleur de base de toute l'interface
	OrGras = "\033[1;38;5;220m"
	Rouge  = "\033[38;5;196m" // dégâts, erreurs, danger, camp ennemi
	Vert   = "\033[38;5;82m"  // soin, succès, victoire
	Orange = "\033[38;5;214m"
	Cyan   = "\033[38;5;51m" // expérience
	Bleu   = "\033[38;5;39m" // camp du héros
)

func Doré(texte string) string       { return Or + texte + Reset }
func Titre(texte string) string      { return OrGras + texte + Reset }
func TexteRouge(texte string) string { return Rouge + texte + Reset }
func TexteVert(texte string) string  { return Vert + texte + Reset }

// ---------- Écran ----------

// Effacer nettoie l'écran du terminal. Appelé au début de chaque nouvel
// écran (menu, combat, boutique...) pour garder un affichage propre.
func Effacer() {
	fmt.Print("\033[H\033[2J")
}

// largeurSection est la longueur des lignes "═══" utilisées par Section et Fin.
// Comme ces lignes ne contiennent jamais d'emoji, leur largeur est toujours
// exacte à l'écran : c'est ce qui évite les cadres qui débordent.
const largeurSection = 48

// Section affiche l'en-tête d'un écran : une ligne, le titre, une ligne.
// Le contenu qui suit s'affiche librement (sans bordure sur les côtés),
// donc les emojis ou accents ne peuvent jamais décaler un cadre.
func Section(titre string) {
	ligne := strings.Repeat("═", largeurSection)
	fmt.Println(Titre("\n" + ligne))
	fmt.Println(Titre("  " + titre))
	fmt.Println(Titre(ligne))
}

// Fin affiche la ligne de fermeture d'un écran, symétrique à Section.
func Fin() {
	fmt.Println(Titre(strings.Repeat("═", largeurSection)))
}

// ---------- Saisie utilisateur ----------

// LireLigne affiche une question et renvoie la réponse tapée par le joueur,
// sans espaces superflus autour.
func LireLigne(reader *bufio.Reader, question string) string {
	fmt.Print(question)
	reponse, _ := reader.ReadString('\n')
	return strings.TrimSpace(reponse)
}

// ---------- Barres de progression ----------

// Barre construit une barre de progression générique, ex: [██████░░░░].
func Barre(actuel, max, largeur int, couleurBarre string) string {
	if max <= 0 {
		max = 1
	}
	if actuel < 0 {
		actuel = 0
	}
	if actuel > max {
		actuel = max
	}

	rempli := (actuel * largeur) / max
	vide := largeur - rempli
	contenu := strings.Repeat("█", rempli) + strings.Repeat("░", vide)
	return couleurBarre + "[" + contenu + "]" + Reset
}

// BarreDeVie : verte/orange/rouge selon le pourcentage de vie restant.
// Utilisée pour la fiche personnage (hors combat).
func BarreDeVie(actuel, max, largeur int) string {
	coul := Vert
	if max > 0 {
		ratio := float64(actuel) / float64(max)
		switch {
		case ratio <= 0.2:
			coul = Rouge
		case ratio <= 0.5:
			coul = Orange
		}
	}
	return Barre(actuel, max, largeur, coul)
}

// BarreHeros et BarreEnnemi utilisent des couleurs FIXES (et différentes)
// pour qu'on distingue immédiatement, en combat, qui est qui : bleu pour
// toi, rouge pour l'adversaire — peu importe leurs PV restants.
func BarreHeros(actuel, max, largeur int) string {
	return Barre(actuel, max, largeur, Bleu)
}

func BarreEnnemi(actuel, max, largeur int) string {
	return Barre(actuel, max, largeur, Rouge)
}

// BarreXP : barre d'expérience (cyan), utilisée sur la fiche personnage.
func BarreXP(actuel, max, largeur int) string {
	return Barre(actuel, max, largeur, Cyan)
}
