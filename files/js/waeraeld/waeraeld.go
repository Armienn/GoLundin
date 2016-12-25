package main

import "github.com/gopherjs/gopherjs/js"
import "fmt"

func main() {
	character := new(Character)
	character.Genes = make([]Gene, 0)
	genes := []Gene{
		{"Endoskelet", 10, 12, 10, 10, 10, 10, 10},
		{"Exoskelet", 12, 10, 12, 10, 10, 10, 15},
	}
	document := GetDocument()

	destinationElement := js.Global.Get("otherplace")
	button := NewButton(document, "Clear", func() {
		character.Genes = make([]Gene, 0)
		js.Global.Get("outputText").Set("innerHTML", character.GetDescription())
	})
	destinationElement.Call("appendChild", button)

	for _, gene := range genes {
		addGeneButton(document, character, gene)
	}

	js.Global.Get("outputText").Set("innerHTML", "text")
}

type Character struct {
	Genes       []Gene
	STR         int
	DEX         int
	CON         int
	INT         int
	WIS         int
	CHA         int
	WeightClass int
}

func (character *Character) UpdateStats() {
	character.STR = 0
	character.DEX = 0
	character.CON = 0
	character.INT = 0
	character.WIS = 0
	character.CHA = 0
	character.WeightClass = 0
	for _, gene := range character.Genes {
		character.STR += gene.STR
		character.DEX += gene.DEX
		character.CON += gene.CON
		character.INT += gene.INT
		character.WIS += gene.WIS
		character.CHA += gene.CHA
		character.WeightClass += gene.WeightClass
	}
}

func (character *Character) GetDescription() string {
	character.UpdateStats()
	return fmt.Sprintf(`
	STR: %d <br/>
	DEX: %d <br/>
	CON: %d <br/>
	INT: %d <br/>
	WIS: %d <br/>
	CHA: %d <br/>
	WeightClass: %d
	`,
		character.STR,
		character.DEX,
		character.CON,
		character.INT,
		character.WIS,
		character.CHA,
		character.WeightClass)
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
}

func addGeneButton(document *Document, character *Character, gene Gene) {
	destinationElement := js.Global.Get("buttonplace")
	button := NewButton(document, gene.Name, func() {
		character.Genes = append(character.Genes, gene)
		js.Global.Get("outputText").Set("innerHTML", character.GetDescription())
	})
	destinationElement.Call("appendChild", button)
}

func NewButton(document *Document, text string, onClick func()) *js.Object {
	button := document.CreateElement("button")
	button.Set("onclick", onClick)
	button.Set("innerHTML", text)
	button.Get("style").Set("marginBottom", "0.5rem")
	return button
}

type Document struct {
	*js.Object
}

func GetDocument() *Document {
	return &Document{js.Global.Get("document")}
}

func (document *Document) CreateElement(tagName string) *js.Object {
	return document.Call("createElement", tagName)
}
