README 

Titre du projet :  

Le nom du projet se nomme : Le royaume oublié 

Contexte : 

le contexte de pourquoi il a était créer est que c'est un projet qu'on doit réaliser pour les fin des Ymerssions on devait faire un Jeu CLI sur le thème qu'on voulait 

Fonctionnalités : 

Un jeu de rôle en console écrit en Go.

- **5 classes** : Chevalier, Mage, Archer, Voleur, Géant. Chacune a ses stats et son arme de départ.
- **Combat au tour par tour** : attaquer, utiliser une potion ou fuir.
- **17 monstres** qui apparaissent selon ton niveau, dans 4 catégories : faune sauvage, peaux-vertes, morts-vivants et créatures démoniaques.
- **10 boss**, un tous les 10 niveaux (de 10 à 100).
- **Progression** : chaque victoire donne de l'XP et des diamants 💎, et chaque niveau augmente tes stats.
- **Marchand** : achète des potions de soin, de force ou de puissance.
- **Forgeron** : améliore ton arme ou achètes-en une nouvelle.
- **Inventaire** : 10 emplacements, avec des piles de 5 potions maximum.

Prérequis : 

- [Go](https://go.dev/dl/) 1.21 ou plus
- Un terminal avec couleurs et emojis (UTF-8)

guide de lancements : 

Lancement : 

il faut installe [Go](https://go.dev/dl/) 1.21 ou plus.
2. dans le dossier du projet (celui qui contient `main.go`), lance sur le terminal :

       go run main.go

3. Navigue dans les menus en tapant un numéro puis Entrée. Tape `5` dans le menu principal pour quitter.

ong du code source : 

Le numéro de série du volume est 2A92-EDDE
C:.
├───affichage
├───boutique
├───classe
├───combat
├───competence
├───couleur
├───ennemi
├───forgeron
├───inventaire
├───item
├───marchand
├───monstre
├───personnage
main.go

