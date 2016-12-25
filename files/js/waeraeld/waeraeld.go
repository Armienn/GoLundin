package main

import (
	"strconv"

	"github.com/Nequilich/gocto"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	character := new(Character)
	character.Genes = make([]*Gene, 0)
	genes := []*Gene{
		NewGene("Endoskelet", 10, 12, 10, 0, 0, 0, 10),
		NewGene("Exoskelet", 12, 10, 12, 0, 0, 0, 15),
		NewGene("Massivt skelet", 4, 0, 0, 0, 0, 0, 5),
		NewGene("Spinkelt skelet", -2, 2, 0, 0, 0, 0, -5),
		NewGene("Lunger", 0, 0, 0, 0, 0, 0, 1, "Trække vejret i luft", "Holde vejret", "Dør uden luft"),
		NewGene("Hud", 0, 0, 6, 0, 0, 0, 2),
	}

	destinationElement := js.Global.Get("otherplace")
	button := NewButton("Clear", func() {
		character.Genes = make([]*Gene, 0)
		js.Global.Get("outputText").Set("innerHTML", character.GetDescription())
	})
	destinationElement.Call("appendChild", button)

	for _, gene := range genes {
		addGeneButton(character, gene)
	}

	js.Global.Get("outputText").Set("innerHTML", "text")
}

type Character struct {
	Genes []*Gene
	Gene
}

func (character *Character) UpdateStats() {
	character.STR = 0
	character.DEX = 0
	character.CON = 0
	character.INT = 0
	character.WIS = 0
	character.CHA = 0
	character.WeightClass = 0
	character.Abilities = []string{}
	for _, gene := range character.Genes {
		character.STR += gene.STR
		character.DEX += gene.DEX
		character.CON += gene.CON
		character.INT += gene.INT
		character.WIS += gene.WIS
		character.CHA += gene.CHA
		character.WeightClass += gene.WeightClass
		character.Abilities = append(character.Abilities, gene.Abilities...)
	}
}

func (character *Character) GetDescription() string {
	character.UpdateStats()
	description := "Abilities:<br/>"
	for _, ability := range character.Abilities {
		description += ability + "<br/>"
	}
	description += "<br/>" + character.ToString()
	description += "<br/>Genes:<br/>"
	for _, gene := range character.Genes {
		description += gene.Name + "<br/>"
	}
	return description
}

type Gene struct {
	Name        string
	STR         int
	DEX         int
	CON         int
	INT         int
	WIS         int
	CHA         int
	WeightClass int
	Abilities   []string
}

func NewGene(name string, str int, dex int, con int, intelligence int, wis int, cha int, weightClass int, abilities ...string) *Gene {
	return &Gene{name, str, dex, con, intelligence, wis, cha, weightClass, abilities}
}

func (gene *Gene) ToString() string {
	return "STR: " + strconv.Itoa(gene.STR) + "<br/>" +
		"DEX: " + strconv.Itoa(gene.DEX) + "<br/>" +
		"CON: " + strconv.Itoa(gene.CON) + "<br/>" +
		"INT: " + strconv.Itoa(gene.INT) + "<br/>" +
		"WIS: " + strconv.Itoa(gene.WIS) + "<br/>" +
		"CHA: " + strconv.Itoa(gene.CHA) + "<br/>" +
		"WeightClass: " + strconv.Itoa(gene.WeightClass) + "<br/>"
}

func addGeneButton(character *Character, gene *Gene) {
	destinationElement := js.Global.Get("buttonplace")
	button := NewButton(gene.Name, func() {
		character.Genes = append(character.Genes, gene)
		js.Global.Get("outputText").Set("innerHTML", character.GetDescription())
	})
	destinationElement.Call("appendChild", button)
}

func NewButton(text string, onClick func()) *js.Object {
	button := gocto.CreateElement("button")
	button.Set("onclick", onClick)
	button.Set("innerHTML", text)
	button.Get("style").Set("marginBottom", "0.5rem")
	button.Get("style").Set("width", "100%")
	return button
}
