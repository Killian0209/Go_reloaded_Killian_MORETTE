package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Pour executer il faut ecrire: go run . entree.txt sortie.txt")
		return
	}

	entree := os.Args[1]
	sortie := os.Args[2]

	contenu, erreur := os.ReadFile(entree)
	if erreur != nil {
		fmt.Println("Erreur d'écriture: ", erreur)
	}

	texte := string(contenu)
	mot := strings.Fields(texte)

	// (Hex)
	for i := 1; i < len(mot); i++ {
		if mot[i] == "(hex)" {
			n, _ := strconv.ParseInt(mot[i-1], 16, 64)
			mot[i-1] = strconv.FormatInt(n, 10)
			mot = append(mot[:i], mot[i+1:]...)
			i--
		}
	}
	// (Hex)

	// (Bin)
	for i := 0; i < len(mot); i++ {
		if mot[i] == "(bin)" {
			n, _ := strconv.ParseInt(mot[i-1], 2, 64)
			mot[i-1] = strconv.FormatInt(n, 10)
			mot = append(mot[:i], mot[i+1:]...)
			i--
		}
	}
	// (Bin)

	// (up)
	for i := 0; i < len(mot); i++ {
		if mot[i] == "(up)" {
			mot[i-1] = strings.ToUpper(mot[i-1])
			mot = append(mot[:i], mot[i+1:]...)
			i--
		}
	}
	//(up)

	//(low)
	for i := 0; i < len(mot); i++ {
		if mot[i] == "(low)" {
			mot[i-1] = strings.ToLower(mot[i-1])
			mot = append(mot[:i], mot[i+1:]...)
			i--
		}
	}
	//(low)

	//(cap)
	for i := 0; i < len(mot); i++ {
		if mot[i] == "(cap)" {
			e := mot[i-1]
			mot[i-1] = (strings.ToUpper(e[:1]) + e[1:])
			mot = append(mot[:i], mot[i+1:]...)
			i--
		}
	}
	//(cap)
	
	//(up, n)
	for i := 0; i < len(mot); i++ {
		if strings.HasPrefix(mot[i], "(up,") && strings.HasSuffix(mot[i], ")") {
			entier := strings.TrimPrefix(mot[i], "(up,")
			entier = strings.TrimSuffix(entier, ")")

			n, err := strconv.Atoi(entier)
			if err != nil {
				continue
			}

			for z := 1; z <= n && i-z >= 0; z++ {
				j := i - z
				mot[j] = strings.ToUpper(mot[j])
			}

			mot = append(mot[:i], mot[i+1:]...)
			i--
		}
	}
	//(up, n)

	//(low, n)
	for i:=0; i < len(mot); i++ {
		if strings.HasPrefix(mot[i], "(low,") && strings.HasSuffix(mot[i], ")") {
			entier := strings.TrimPrefix(mot[i], "(low,")
			entier = strings.TrimSuffix(entier, ")")

			n, err := strconv.Atoi(entier)
			if err != nil {
				continue
			}

			for z := 1; z <= n && i-z >= 0; z++ {
				j := i - z
				mot[j] = strings.ToLower(mot[j])
			}

			mot = append(mot[:i], mot[i+1:]...)
			i--
		}
	}
	//(low, n)

	//(cap, n)
	for i:=0; i < len(mot); i++ {
		if strings.HasPrefix(mot[i], "(cap,") && strings.HasSuffix(mot[i], ")") {
			entier := strings.TrimPrefix(mot[i], "(cap,")
			entier = strings.TrimSuffix(entier, ")")

			n, err := strconv.Atoi(entier)
			if err != nil {
				continue
			}

			for z := 1; z <= n && i-z >= 0; z++ {
            	j := i - z
        		e := mot[j]
            	if len(e) > 0 {
                	mot[j] = strings.ToUpper(e[:1]) + e[1:]
            	}
        	}

        	mot = append(mot[:i], mot[i+1:]...)
        	i--
		}
	}
	//(cap, n)

	//ponctuation
	for i:= 0; i < len(mot); i++ {
		if string(mot[i][0]) == "!" {
			mot[i-1] += mot[i]
        	mot = append(mot[:i], mot[i+1:]...)
        	i--
		}
		if string(mot[i][0]) == "." {
			mot[i-1] += mot[i]
        	mot = append(mot[:i], mot[i+1:]...)
        	i--
		}
		if string(mot[i][0]) == "?" {
			mot[i-1] += mot[i]
        	mot = append(mot[:i], mot[i+1:]...)
        	i--
		}
		if string(mot[i][0]) == "," {
			mot[i-1] += mot[i]
        	mot = append(mot[:i], mot[i+1:]...)
        	i--
		}
		if string(mot[i][0]) == ";" {
			mot[i-1] += mot[i]
        	mot = append(mot[:i], mot[i+1:]...)
        	i--
		}
		if string(mot[i][0]) == ":" {
			mot[i-1] += mot[i]
        	mot = append(mot[:i], mot[i+1:]...)
        	i--
		}
	}
	//ponctuation

	// apostrophes
	for i := 0; i < len(mot); i++ {
		if mot[i] == "'" {
			for j := i + 1; j < len(mot); j++ {
				if mot[j] == "'" {
					mot[i+1] = "'"+mot[i+1]
					mot[j-1] = mot[j-1]+"'"
					mot = append(mot[:j], mot[j+1:]...)
					mot = append(mot[:i], mot[i+1:]...)
					break
				}
			}
		}
	}
	// apostrophes
	
	//liaison
	for i := 0; i < len(mot)-1; i++ {
		premier := mot[i][len(mot[i])-1]
		suivant := mot[i+1][0]

		if (premier == 'a' || premier == 'A') &&
			(suivant == 'a' || suivant == 'e' ||
				suivant == 'i' || suivant == 'o' ||
				suivant == 'u' || suivant == 'y' ||
				suivant == 'A' || suivant == 'E' ||
				suivant == 'I' || suivant == 'O' ||
				suivant == 'U' || suivant == 'Y') {
			mot[i] += "n"
		}
	}
	//liaison


	result := strings.Join(mot, " ")

	erreur = os.WriteFile(sortie, []byte(result), 0644)
	if erreur != nil {
		fmt.Println("Erreur d'écriture :", erreur)
		return
	}

	fmt.Println("Fichier créé")
}

