package item

// Item est l'interface commune à tous les objets du jeu (armes, potions...).
type Item interface {
	Nom() string
	Description() string
}
