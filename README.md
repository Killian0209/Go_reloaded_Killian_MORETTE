# PROJET GO-RELOADED
# README — Projet de traitement de texte en Go

## Présentation

Ce projet est un programme écrit en Go (Golang) qui permet de lire un fichier texte, de modifier son contenu automatiquement et d'enregistrer le résultat dans un autre fichier.

Le programme reconnaît plusieurs instructions qui permettent de modifier les mots, les nombres et la ponctuation.

## Prérequis

Pour utiliser ce programme, il faut :
- Installer Go sur son ordinateur.
- Avoir un fichier `main.go` contenant le code du programme.
- Préparer un fichier texte à modifier.

Pour vérifier que Go est installé, exécute cette commande dans le terminal :

```bash
go version
```

## Installation et exécution

### 1. Préparer les fichiers

Le projet peut être organisé comme ceci :

```text
gp-reloaded/
├── main.go
├── entree.txt
├── sortie.txt
└── README.md
```

- `main.go` : le code du programme.
- `entree.txt` : le fichier contenant le texte original.
- `sortie.txt` : le fichier qui recevra le texte modifié.
- `README.md` : le fichier qui explique le projet.

### 2. Lancer le programme

Dans le terminal, place-toi dans le dossier du projet et exécute :

```bash
go run . entree.txt sortie.txt
```

La commande indique au programme quel fichier lire et dans quel fichier enregistrer le résultat.

Si le projet ne contient pas encore de fichier `go.mod`, crée-le avec :

```bash
go mod init mon-projet
```

### 3. Récupérer le résultat

Si le traitement se déroule correctement, le programme affiche :

```text
Fichier créé
```

Le texte transformé est ensuite disponible dans `sortie.txt`.

## Fonctionnement du programme

### 1. Les bibliothèques utilisées

Au début du code, plusieurs bibliothèques sont importées :

- `fmt` : affiche des messages dans le terminal.
- `os` : lit les fichiers, écrit les résultats et récupère les arguments de la commande.
- `strconv` : convertit les nombres d'une base à une autre et transforme les chaînes en nombres.
- `strings` : permet de découper, rechercher et modifier du texte.

### 2. Vérification des arguments

Le programme vérifie que deux noms de fichiers ont été fournis : le fichier d'entrée et le fichier de sortie.

Si les arguments sont incorrects, il affiche une instruction d'utilisation et arrête le programme.

### 3. Lecture du fichier

Le programme utilise `os.ReadFile()` pour lire le contenu du fichier d'entrée.

Le contenu est ensuite converti en texte avec `string()`, puis découpé en mots grâce à `strings.Fields()`.

Les mots sont stockés dans une liste appelée `mot`, que le programme peut parcourir et modifier.

### 4. Conversion hexadécimale : `(hex)`

Cette instruction convertit le nombre situé juste avant `(hex)` de la base 16 vers la base 10.

Exemple :

```text
1A (hex)
```

Résultat :

```text
26
```

### 5. Conversion binaire : `(bin)`

Cette instruction convertit le nombre situé juste avant `(bin)` de la base 2 vers la base 10.

Exemple :

```text
1010 (bin)
```

Résultat :

```text
10
```

### 6. Mettre un mot en majuscules : `(up)`

Cette instruction transforme le mot précédent en majuscules.

Exemple :

```text
bonjour (up)
```

Résultat :

```text
BONJOUR
```

### 7. Mettre un mot en minuscules : `(low)`

Cette instruction transforme le mot précédent en minuscules.

Exemple :

```text
BONJOUR (low)
```

Résultat :

```text
bonjour
```

### 8. Mettre une majuscule au début d'un mot : `(cap)`

Cette instruction met la première lettre du mot précédent en majuscule.

Exemple :

```text
bonjour (cap)
```

Résultat :

```text
Bonjour
```

### 9. Mettre plusieurs mots en majuscules : `(up, n)`

Cette instruction met en majuscules les `n` mots qui se trouvent avant le marqueur.

Exemple :

```text
bonjour tout le monde (up, 3)
```

Résultat :

```text
bonjour TOUT LE MONDE
```

Les trois mots précédant `(up, 3)` sont transformés.

### 10. Mettre plusieurs mots en minuscules : `(low, n)`

Cette instruction met en minuscules les `n` mots qui précèdent le marqueur.

Exemple :

```text
BONJOUR TOUT LE MONDE (low, 3)
```

Résultat :

```text
BONJOUR tout le monde
```

### 11. Mettre une majuscule à plusieurs mots : `(cap, n)`

Cette instruction met en majuscule la première lettre des `n` mots précédents.

Exemple :

```text
bonjour tout le monde (cap, 3)
```

Résultat :

```text
bonjour Tout Le Monde
```

### 12. Gestion de la ponctuation

Le programme parcourt les mots et repère les éléments commençant par certains signes de ponctuation.

Il traite notamment :
- `!` : point d'exclamation.
- `.` : point.
- `?` : point d'interrogation.
- `,` : virgule.
- `;` : point-virgule.
- `:` : deux-points.

Lorsqu'un élément commence par l'un de ces signes, il est ajouté à la fin du mot précédent. Le marqueur de ponctuation est ensuite retiré de la liste.

Cela permet, par exemple, de transformer :

```text
Bonjour !
```

en :

```text
Bonjour!
```

### 13. Gestion des apostrophes

Le programme recherche les apostrophes simples utilisées comme marqueurs de début et de fin d'un groupe de mots.

Il ajoute une apostrophe au début du premier mot du groupe et à la fin du dernier mot, puis supprime les marqueurs séparés.

Exemple :

```text
' salut tout le monde '
```

Le résultat visé est :

```text
'salut tout le monde'
```

Cette partie du code dépend de la manière dont les apostrophes sont séparées dans le texte d'entrée.

### 14. Ajout de la lettre `n`

Le programme vérifie les mots deux par deux.

Si un mot se termine par `a` ou `A` et que le mot suivant commence par une voyelle, il ajoute la lettre `n` à la fin du premier mot.

Les voyelles prises en compte sont `a`, `e`, `i`, `o`, `u`, `y` et leurs versions majuscules.

Exemple :

```text
a amiga
```

Résultat :

```text
an amiga
```

### 15. Création du fichier de sortie

Après toutes les transformations, le programme rassemble les mots avec `strings.Join()` en utilisant un espace comme séparateur.

Le texte final est ensuite écrit dans le fichier de sortie avec `os.WriteFile()`.

Si l'écriture échoue, un message d'erreur est affiché. Sinon, le programme confirme que le fichier a été créé.

### Exemple complet

Contenu de `entree.txt` :

```text
hello (up)
bonjour (cap)
1010 (bin)
1A (hex)
a apple
```

Contenu attendu de `sortie.txt` :

```text
HELLO
Bonjour
10
26
an apple
```

Le résultat exact dépend de la disposition du texte d'entrée, car le programme remplace les espaces et les retours à la ligne par un espace entre chaque mot.

### Conclusion

Ce programme permet d'automatiser plusieurs opérations courantes sur un texte grâce à des marqueurs simples.

Il utilise les bibliothèques standard de Go pour lire un fichier, parcourir les mots, appliquer des transformations et enregistrer le résultat.

### Remarque : certaines parties du code supposent que les marqueurs sont utilisés correctement. Par exemple, `(up)` ne doit pas être le premier élément du texte et les nombres doivent être valides pour leur base. Une version plus robuste devrait vérifier ces situations avant de modifier les mots.