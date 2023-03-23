package utils

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func HomeMenu() string {
	fmt.Println(" ______     ______     ______     ______   ______     ______     ______    \n/\\  ___\\   /\\  ___\\   /\\  __ \\   /\\__  _\\ /\\  ___\\   /\\  == \\   /\\  __ \\   \n\\ \\  __\\   \\ \\___  \\  \\ \\ \\/\\ \\  \\/_/\\ \\/ \\ \\  __\\   \\ \\  __<   \\ \\  __ \\  \n \\ \\_____\\  \\/\\_____\\  \\ \\_____\\    \\ \\_\\  \\ \\_____\\  \\ \\_\\ \\_\\  \\ \\_\\ \\_\\ \n  \\/_____/   \\/_____/   \\/_____/     \\/_/   \\/_____/   \\/_/ /_/   \\/_/\\/_/ \n                                                                           ")
	prompt := promptui.Select{
		Label: "Select Site",
		Items: []string{"1. END. Clothing", "2. JD Sports", "3. FootPatrol", "4. Size?", "5. Adidas Confirmed",
			"6. Adidas Confirmed", "7. Nike"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Selection failed %v\n", err)

	}
	return result
}

func EndMenu() string {
	fmt.Println(" ______     ______     ______     ______   ______     ______     ______    \n/\\  ___\\   /\\  ___\\   /\\  __ \\   /\\__  _\\ /\\  ___\\   /\\  == \\   /\\  __ \\   \n\\ \\  __\\   \\ \\___  \\  \\ \\ \\/\\ \\  \\/_/\\ \\/ \\ \\  __\\   \\ \\  __<   \\ \\  __ \\  \n \\ \\_____\\  \\/\\_____\\  \\ \\_____\\    \\ \\_\\  \\ \\_____\\  \\ \\_\\ \\_\\  \\ \\_\\ \\_\\ \n  \\/_____/   \\/_____/   \\/_____/     \\/_/   \\/_____/   \\/_/ /_/   \\/_/\\/_/ \n                                                                           ")
	prompt := promptui.Select{
		Label: "Select Mode",
		Items: []string{"1. Entry", "2. Account Gen", "3. Address Changer"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Selection failed %v\n", err)

	}
	return result
}
