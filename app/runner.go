package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Run() {
	printBanner()

	for {
		profile, err := readLine("⌨️ Introducir el nombre del Perfil: ")
		if err != nil {
			fmt.Println("❌ Error leyendo la entrada. Intentá de nuevo.")
			continue
		}
		if profile == "" {
			profile = "default"
		}

		fmt.Println("✔️ Perfil seleccionado:", profile)
		execute(profile)

		time.Sleep(1 * time.Second) //Para darle tiempo a la Go runtime de logs

		opt, err := readLine("⌨️ ¿Desea continuar? (Y/N): ")
		if err != nil || strings.EqualFold(opt, "n") {
			fmt.Println("🔌 Shutdown, Bye!")
			break
		}
	}
}

func printBanner() {
	fmt.Println("────────────────────────────────────────────────────────")
	fmt.Println("🔐 App: GoPGPFileCipher v1.0.0")
	fmt.Println("────────────────────────────────────────────────────────")
}

func readLine(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(input), err
}
