package actor

import (
	"fmt"
	"math/rand"
)

var (
	firstNames = []string{"Gladius", "Varro", "Doran", "Thalric", "Lucan", "Kael", "Borin", "Draven", "Garrik", "Vex", "Jay"}
	lastNames  = []string{"the Fierce", "Ironfist", "Stormborn", "Shadowblade", "Bloodfang", "the Unyielding", "Darkbane", "Grimwolf", "the Silent", "Firebrand", "The Boi"}
)

// Generate a random NPC name
func GenerateRandomName() string {
	first := firstNames[rand.Intn(len(firstNames))]
	last := lastNames[rand.Intn(len(lastNames))]
	return fmt.Sprintf("%s %s", first, last)
}
