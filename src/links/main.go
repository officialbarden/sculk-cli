package links

import (
	"errors"
	"fmt"

	"github.com/savioxavier/termlink"
)

var github string = "https://github.com/officialbarden/sculk-cli"
var website string = "https://officialbarden.github.io/sculk"
var libraries string = "https://officialbarden.github.io/sculk/libraries"
var discord string = "https://discord.gg/JWZkAgsyry"

func Main(args []string) error {
	providers := args

	// If providers is empty, provide all links:
	if len(providers) == 0 {
		fmt.Printf("\n\nAll Links Affiliated with Sculk \n\n")
		fmt.Println(termlink.Link("🔗 Github", github, true))
		fmt.Println(termlink.Link("🌐 Website", website, true))
		fmt.Println(termlink.Link("🗃  Libraries", libraries, true))
		fmt.Println(termlink.Link("💬 Discord", discord, true))
		fmt.Printf("\n\n")
		return nil
	}
	
	// iterate through provided arguments and provide links
	for i := range len(providers) {
		providerName := providers[i]
		switch providerName {
			case "github":
				fmt.Println(termlink.Link("🔗 Github", github, true))
			case "website":
				fmt.Println(termlink.Link("🌐 Website", website, true))
			case "libraries":
				fmt.Println(termlink.Link("🗃  Libraries", libraries, true))
			case "discord":
				fmt.Println(termlink.Link("💬 Discord", discord, true))
			default:
				error := "No Provider '" + providerName + "' found"
				return errors.New(error)
		}
	}

	return nil
}


