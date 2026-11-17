package boss

import (
	"fmt"
)

// Boss représente la structure d'un boss dans le jeu.
type Boss struct {
	Name       string
	AttackName string
	Level      int
	MaxHP      int
	CurrentHP  int
	Damage     int
}

// IsAlive vérifie si le boss est encore en vie.
func (b *Boss) IsAlive() bool {
	return b.CurrentHP > 0
}

// TakeDamage applique des dégâts au boss.
func (b *Boss) TakeDamage(amount int) {
	b.CurrentHP -= amount
	if b.CurrentHP < 0 {
		b.CurrentHP = 0
	}
}

// Attack exécute l'attaque unique du boss dans le terminal.
func (b *Boss) Attack() int {
	fmt.Printf(" [BOSS] %s lance sa capacité mortelle [%s] et inflige %d dégâts !\n", b.Name, b.AttackName, b.Damage)
	return b.Damage
}

// DisplayStatus affiche l'état du boss dans le terminal.
func (b *Boss) DisplayStatus() {
	fmt.Printf(" BOSS NIV.%d : %s | PV: %d/%d | Dégâts: %d\n", b.Level, b.Name, b.CurrentHP, b.MaxHP, b.Damage)
}

// --- Méthodes utilisées par le système de combat (interface combat.Adversaire) ---

// GetName renvoie le nom du boss.
func (b *Boss) GetName() string { return b.Name }

// GetHP renvoie les PV actuels et maximum du boss.
func (b *Boss) GetHP() (int, int) { return b.CurrentHP, b.MaxHP }

// XPReward renvoie l'XP gagnée en battant le boss (x3 par rapport à un monstre normal).
func (b *Boss) XPReward() int { return b.Damage * 3 * 3 }

// DiamondsReward renvoie les diamants gagnés en battant le boss (x3 par rapport à un monstre normal).
func (b *Boss) DiamondsReward() int { return b.Damage * 2 * 3 }

// GetAllBosses renvoie la liste fixe de tous les boss de palier.
func GetAllBosses() []Boss {
	return []Boss{
		{Name: "Chef des Gobelins", AttackName: "Coup de Grâce Perfide", Level: 10, MaxHP: 50, CurrentHP: 50, Damage: 19},
		{Name: "Reine des Araignées", AttackName: "Cocon d'Acide Venimeux", Level: 20, MaxHP: 73, CurrentHP: 73, Damage: 22},
		{Name: "Loup Blanc", AttackName: "Hurlement de la Blizzard", Level: 30, MaxHP: 98, CurrentHP: 98, Damage: 27},
		{Name: "Karaba La Sorcière", AttackName: "Malédiction des Fétiches", Level: 40, MaxHP: 123, CurrentHP: 123, Damage: 32},
		{Name: "Casper le Fantôme", AttackName: "Farce Spatiale Effroyable", Level: 50, MaxHP: 148, CurrentHP: 148, Damage: 37},
		{Name: "Dracula", AttackName: "Étreinte Sanguinaire de la Nuit", Level: 60, MaxHP: 178, CurrentHP: 178, Damage: 43},
		{Name: "Crocodile Sarcosuchus", AttackName: "Mâchoire du Crétacé", Level: 70, MaxHP: 208, CurrentHP: 208, Damage: 49},
		{Name: "Enzo le Démon", AttackName: "Cataclysme des Enfers", Level: 80, MaxHP: 238, CurrentHP: 238, Damage: 55},
		{Name: "Nathan le Faucheur", AttackName: "Jugement Dernier des Âmes", Level: 90, MaxHP: 268, CurrentHP: 268, Damage: 61},
		{Name: "Le Diable", AttackName: "Damnation Éternelle", Level: 100, MaxHP: 308, CurrentHP: 308, Damage: 68},
	}
}

// GetBossByLevel instancie le boss correspondant à un palier de niveau donné (ex: 10, 20, 30...).
func GetBossByLevel(level int) (*Boss, error) {
	for _, b := range GetAllBosses() {
		if b.Level == level {
			// Retourne une copie indépendante du boss
			bossInstance := b
			return &bossInstance, nil
		}
	}
	return nil, fmt.Errorf("aucun boss configuré pour le niveau %d", level)
}

// IsBossLevel vérifie si le niveau actuel du joueur déclenche un combat de boss.
func IsBossLevel(level int) bool {
	return level%10 == 0 && level >= 10 && level <= 100
}
