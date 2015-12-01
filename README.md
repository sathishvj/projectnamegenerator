# Project Name Generator (in Go)

 * Generates names in the form: Adverb Verb Adjective Noun
 * You can choose to exclude one or more of those parts
 * You can choose to hyphenate it
 * You can choose it to be returned in lower case

### Installation
```
	go get github.com/sathishvj/projectnamegenerator
```

### Examples:

```
	import png "github.com/sathishvj/projectnamegenerator"

	for i := 0; i < 10; i++ {
		fmt.Println(png.NewName())
	}
```

 * Justly Chart Salty Rate
 * Intensely Tie Clumsy Farmer
 * Knowledgeably Enlist Faint Fireman
 * Freely Connect Quick Decision
 * Frenetically Trot Quick Surprise
 * Scarily Let Salty Doctor
 * Doubtfully Fear Mysterious Park
 * Busily Manipulate Old Vessel
 * Extremely Finalize Delicious Teeth
 * Arrogantly List Thundering Dolls

```
	import png "github.com/sathishvj/projectnamegenerator"

	png1 := png.NewGenerator()
	png1.Lowercase = true
	png1.Hyphenate = true

	for i := 0; i < 9; i++ {
		fmt.Println(png1.NewName())
	}
```


 * shyly-transfer-sparkling-yak
 * sternly-switch-screeching-look
 * seriously-wet-quaint-can
 * tenderly-alight-round-finger
 * upbeat-increase-tart-juice
 * monthly-nest-orange-goose
 * unfortunately-tickle-teeny-tiny-zinc
 * unexpectedly-melt-curved-bit
 * greedily-launch-damaged-direction

```
	import png "github.com/sathishvj/projectnamegenerator"

	png1 := png.NewGenerator()
	png1.Lowercase = true
	png1.Hyphenate = true
	png1.HasAdverb = false
	png1.HasVerb = false


	for i := 0; i < 10; i++ {
		fmt.Println(png1.NewName())
	}
```

 * early-basket
 * worried-effect
 * adorable-breakfast
 * plain-mitten
 * fancy-carpenter
 * dusty-channel
 * kind-stitch
 * yummy-minute
 * massive-pizzas
 * brave-title

#### Warning
The words in the lists are not filtered for happy, positive ones.

