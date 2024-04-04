package jsons

import (
	"encoding/json"
	"fmt"
	"log"
)

type ram struct {
	Model    string `json: model`
	Quantity uint   `json: quantity`
}

type storage struct {
	Disk1 string `json: disk1`
	Disk2 string `json: disk2`
}

type pc struct {
	Cpu         string  `json: cpu`
	Motherboard string  `json: motherboard`
	Ram         ram     `json: ram`
	Storage     storage `json: storage`
}

func UseUnmarshal() {
	pcJSON := `{
		"cpu": "amd ryzen 7",
		"motherboard": "msi b550 vdh pro wifi",
		"ram": {
			"model": "asgard 16gb ddr4 2600mhz",
			"quantity": 2
		},
		"storage": {
			"disk1": "ssd nvme kingston 1tb",
			"disk2": "ssd sata kingston 240gb"
		}
	}`

	var meuPc pc
	fmt.Println(meuPc)
	if err := json.Unmarshal([]byte(pcJSON), &meuPc); err != nil {
		log.Fatal(err)
	}
	fmt.Println(meuPc)
}
