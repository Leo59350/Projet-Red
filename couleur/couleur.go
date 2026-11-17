// Package couleur centralise l'habillage visuel du jeu dans le terminal :
// tous les affichages (intro, menu, fiche personnage, inventaire, combat,
// marchand, forgeron...) utilisent ces mêmes fonctions pour rester cohérents.
package couleur

import (
	"fmt"
	"strings"
)

const (
	Reset  = "\033[0m"
	Or     = "\033[38;5;220m" // doré, couleur "par défaut" de l'interface
	OrGras = "\033[1;38;5;220m"
	Rouge  = "\033[38;5;196m" // dégâts subis, erreurs, danger
	Vert   = "\033[38;5;82m"  // soin, succès, victoire
	Orange = "\033[38;5;214m" // avertissement (PV intermédiaires)
	Cyan   = "\033[38;5;51m"  // expérience / progression
)

// Doré colore un texte en doré. C'est la couleur "par défaut" de tous
// les affichages de menu/fiche/inventaire du jeu.
func Doré(texte string) string {
	return Or + texte + Reset
}

// Titre colore un texte en doré gras (titres, bannières, séparateurs).
func Titre(texte string) string {
	return OrGras + texte + Reset
}

// TexteRouge colore un texte en rouge (dégâts, erreurs, défaite).
func TexteRouge(texte string) string {
	return Rouge + texte + Reset
}

// TexteVert colore un texte en vert (soin, victoire, succès).
func TexteVert(texte string) string {
	return Vert + texte + Reset
}

// Effacer nettoie l'écran du terminal (utilisé entre les différents écrans du jeu).
func Effacer() {
	fmt.Print("\033[H\033[2J")
}

// Barre construit une barre de progression colorée, ex: [██████░░░░].
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

// BarreDeVie construit une barre de vie : verte si santé haute,
// orange si intermédiaire, rouge si critique.
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

// BarreXP construit une barre d'expérience (cyan).
func BarreXP(actuel, max, largeur int) string {
	return Barre(actuel, max, largeur, Cyan)
}
