package util

import (
	"fmt"

	"github.com/ghodss/yaml"
)

type Kbell struct {
	Name     string `json:"name"` // Affects YAML field names too.
	Number   string `json:"number"`
	Location string `json:"location"`
}

func AAA() {
	// Marshal a Kbell struct to YAML.
	k := Kbell{"Kbell", "042-863-5388", "Daejeon"}
	j, err := yaml.Marshal(k)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j))
}
