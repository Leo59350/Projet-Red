package ennemi

import (
	"fmt"
	"math/rand"
	"time"
)

type Category string

const (
	CategoryFauneNaturels Category = "Faune Sauvage & Monstres Naturels"
	CategoryPeauxVertes   Category = "Peaux-Vertes & Humanoïdes"
	CategoryMortsVivants  Category = "Morts-Vivants & Spectres"
	CategoryDemonBoss     Category = "Créatures Démoniaques & Légendaires"
)

// LevelRange définit la plage de niveaux d'apparition.
type LevelRange struct {
	MinLevel int
	MaxLevel int
}

// Monster représente les statistiques d'un monstre en mémoire.
type Monster struct {
	Name       string
	AttackName string
	Category   Category
	Levels     LevelRange
	MaxHP      int
	CurrentHP  int
	Damage     int
}

// IsAlive vérifie si le monstre est encore en vie.
func (m *Monster) IsAlive() bool {
	return m.CurrentHP > 0
}

// TakeDamage applique des dégâts au monstre.
func (m *Monster) TakeDamage(amount int) {
	m.CurrentHP -= amount
	if m.CurrentHP < 0 {
		m.CurrentHP = 0
	}
}

// Attack exécute l'attaque du monstre et renvoie les dégâts infligés.
func (m *Monster) Attack() int {
	fmt.Printf("⚔️  %s lance [%s] et inflige %d dégâts !\n", m.Name, m.AttackName, m.Damage)
	return m.Damage
}

// DisplayStatus affiche une ligne d'état pour le terminal.
func (m *Monster) DisplayStatus() {
	fmt.Printf("[%s] %s | PV: %d/%d | Dégâts: %d\n", m.Category, m.Name, m.CurrentHP, m.MaxHP, m.Damage)
}

// --- Méthodes utilisées par le système de combat (interface combat.Adversaire) ---

// GetName renvoie le nom du monstre.
func (m *Monster) GetName() string { return m.Name }

// GetHP renvoie les PV actuels et maximum du monstre.
func (m *Monster) GetHP() (int, int) { return m.CurrentHP, m.MaxHP }

// XPReward renvoie l'XP gagnée en battant le monstre.
func (m *Monster) XPReward() int { return m.Damage * 3 }

// DiamondsReward renvoie les diamants gagnés en battant le monstre.
func (m *Monster) DiamondsReward() int { return m.Damage * 2 }

// GetAllMonsters renvoie la liste modèle de tous les monstres du jeu.
func GetAllMonsters() []Monster {
	return []Monster{
		// 1. Faune Sauvage & Monstres Naturels (Niveau 1 à 100)
		{Name: "Insecte", AttackName: "Incision Chitinieuse", Category: CategoryFauneNaturels, Levels: LevelRange{MinLevel: 1, MaxLevel: 100}, MaxHP: 25, CurrentHP: 25, Damage: 6},
		{Name: "Gobelin", AttackName: "Estoc Traître", Category: CategoryFauneNaturels, Levels: LevelRange{MinLevel: 1, MaxLevel: 100}, MaxHP: 40, CurrentHP: 40, Damage: 15},
		{Name: "Loup", AttackName: "Morsure de la Meute", Category: CategoryFauneNaturels, Levels: LevelRange{MinLevel: 1, MaxLevel: 100}, MaxHP: 35, CurrentHP: 35, Damage: 10},
		{Name: "Araignée", AttackName: "Toile Toilée & Cocon", Category: CategoryFauneNaturels, Levels: LevelRange{MinLevel: 1, MaxLevel: 100}, MaxHP: 30, CurrentHP: 30, Damage: 12},
		{Name: "Serpent", AttackName: "Venin Neurotoxique", Category: CategoryFauneNaturels, Levels: LevelRange{MinLevel: 1, MaxLevel: 100}, MaxHP: 28, CurrentHP: 28, Damage: 13},
		{Name: "Sanglier", AttackName: "Charge Encroûtée", Category: CategoryFauneNaturels, Levels: LevelRange{MinLevel: 1, MaxLevel: 100}, MaxHP: 50, CurrentHP: 50, Damage: 14},
		{Name: "Crabe", AttackName: "Pince Broyeuse", Category: CategoryFauneNaturels, Levels: LevelRange{MinLevel: 1, MaxLevel: 100}, MaxHP: 54, CurrentHP: 54, Damage: 8},

		// 2. Peaux-Vertes & Humanoïdes (Niveau 25 à 100)
		{Name: "Crocodile", AttackName: "Roulade Mortelle", Category: CategoryPeauxVertes, Levels: LevelRange{MinLevel: 25, MaxLevel: 100}, MaxHP: 45, CurrentHP: 45, Damage: 16},
		{Name: "Sorcière", AttackName: "Malédiction d'Athanor", Category: CategoryPeauxVertes, Levels: LevelRange{MinLevel: 25, MaxLevel: 100}, MaxHP: 39, CurrentHP: 39, Damage: 23},
		{Name: "Troll", AttackName: "Lancer de Rocher Brutal", Category: CategoryPeauxVertes, Levels: LevelRange{MinLevel: 25, MaxLevel: 100}, MaxHP: 57, CurrentHP: 57, Damage: 17},
		{Name: "Cyclope", AttackName: "Écrasement Titanesque", Category: CategoryPeauxVertes, Levels: LevelRange{MinLevel: 25, MaxLevel: 100}, MaxHP: 65, CurrentHP: 65, Damage: 19},

		// 3. Morts-Vivants & Spectres (Niveau 50 à 100)
		{Name: "Squelette", AttackName: "Entaille d'Acier Rouillé", Category: CategoryMortsVivants, Levels: LevelRange{MinLevel: 50, MaxLevel: 100}, MaxHP: 62, CurrentHP: 62, Damage: 16},
		{Name: "Fantôme", AttackName: "Lamentation Douloureuse", Category: CategoryMortsVivants, Levels: LevelRange{MinLevel: 50, MaxLevel: 100}, MaxHP: 68, CurrentHP: 68, Damage: 18},
		{Name: "Vampire", AttackName: "Morsure Sanguinaire", Category: CategoryMortsVivants, Levels: LevelRange{MinLevel: 50, MaxLevel: 100}, MaxHP: 85, CurrentHP: 85, Damage: 30},

		// 4. Créatures Démoniaques & Légendaires (Niveau 75 à 100)
		{Name: "Faucheur", AttackName: "Moisson des Âmes", Category: CategoryDemonBoss, Levels: LevelRange{MinLevel: 75, MaxLevel: 100}, MaxHP: 100, CurrentHP: 100, Damage: 32},
		{Name: "Démon", AttackName: "Flammes Infernales", Category: CategoryDemonBoss, Levels: LevelRange{MinLevel: 75, MaxLevel: 100}, MaxHP: 150, CurrentHP: 150, Damage: 41},
		{Name: "Dragon", AttackName: "Souffle du Cataclysme", Category: CategoryDemonBoss, Levels: LevelRange{MinLevel: 75, MaxLevel: 100}, MaxHP: 250, CurrentHP: 250, Damage: 62},
	}
}

// GetMonstersForLevel filtre les monstres disponibles pour le niveau du joueur.
func GetMonstersForLevel(playerLevel int) []Monster {
	var available []Monster
	for _, m := range GetAllMonsters() {
		if playerLevel >= m.Levels.MinLevel && playerLevel <= m.Levels.MaxLevel {
			available = append(available, m)
		}
	}
	return available
}

// SpawnRandomMonster renvoie une nouvelle instance de monstre aléatoire selon le niveau.
func SpawnRandomMonster(playerLevel int) (*Monster, error) {
	available := GetMonstersForLevel(playerLevel)
	if len(available) == 0 {
		return nil, fmt.Errorf("aucun monstre disponible au niveau %d", playerLevel)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	selected := available[rng.Intn(len(available))]

	return &selected, nil
}
