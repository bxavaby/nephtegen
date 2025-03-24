package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
)

func RunSelectionFlow() {
	styles := []string{
		"Abstract",
		"Anthropomorphic",
		"Kinetic",
		"Oozographic",
		"Arcane",
		"Bathysophic",
		"Chthonic",
		"Cryptomnesic",
		"Eidolonetic",
		"Mystical",
		"Grotesque",
		"Macabre",
		"Malevolent",
		"Martyrial",
		"Obscurivenous",
		"Ossified",
		"Parasitic",
		"Stygian",
		"Thalassomantic",
		"Underworld",
		"Cyberpunk",
		"Dieselpunk",
		"Experimentalist",
		"Technocratic",
		"Cosmic",
		"Empyrean",
		"Igniluminous",
		"Surreal",
		"Industrial",
		"Hyperborean",
		"Neolithic",
		"Tectofloral",
		"Bohemian",
		"Elegant",
		"Gothic",
		"Literary",
		"Philosophical",
		"Dystopian",
		"Delirium",
		"Phaneromania",
		"Weaponry",
	}
	selectedStyles := SelectMultiple("Summon up to 4 essences\n", styles, 4)

	formatOptions := []string{
		"One-word",
		"Two-word",
		"Two-word (Hyphenated)",
		"Three-word (Hyphenated)",
		"Abstract Phrase",
		"Acronym",
		"Alliterative",
		"Haiku",
		"Motto",
		"Portmanteau",
		"Symbolic Term",
	}
	wordFormat := SelectOne("Select the structure of thy appellation\n", formatOptions)

	if wordFormat == "One-word" {
		prefixStyles := []string{
			"Arabic-based",
			"Coptic-based",
			"Heian-Japanese",
			"Japanese Folklore",
			"Hellenic",
			"Indus-Valley",
			"Latin-based",
			"Old Slavic",
			"Russian Folklore",
			"Phoenician",
			"Teutonic",
			"Ugaritic",
			"Aetherborne",
			"Bathyclastic",
			"Cybernetic",
			"Druidic",
			"Ecliptic",
			"Gravimetric",
			"Machiniform",
			"Noctivagant",
			"Oblivionic",
			"Orogenic",
			"Petrichoral",
			"Pyroclastic",
			"Quantum",
			"Runebound",
			"Steelflow",
			"Surrealist",
			"Sylvanic",
			"Ornitology Jargon",
			"Microbiology Jargon",
			"Schizofrenic",
			"Aquafungal",
			"Vortexian",
			"Lunarbotic",
			"Insectopunk",
			"Hologothic",
			"Plasma",
			"Glitchfolk",
			"Bloodpunk",
			"Hallucinogenic",
			"Paranoia-driven",
			"Horror-core",
			"Elasmobranchii",
			"Amnesia",
			"Laserbeam",
			"Symbiosis",
			"Blizzard",
			"Dissolving",
			"Orgasmic",
			"Fever Dream",
			"Overkill",
			"Ghost Duet",
			"Omnidirectional",
			"Inertia",
			"Decay",
			"Gluttonous",
			"Mentalist",
			"Kenjutsu",
			"Molecular",
			"Agony",
			"Cataclysm",
			"Sonar",
			"Neurosis",
			"Alien Gods",
			"Man-eating Lizard Dragon",
			"Ending",
			"Prostration",
			"Ominous",
			"Being or Entity",
			"Sensation",
			"Ability or Technique",
			"Terrain Manipulation",
			"Hazard Creation",
			"Mnemonic Device",
			"Strategic Sacrifice",
			"Weaponry",
			"Spell",
			"Cryptomnesia Recursive Trigger",
			"Antimatter Dimensional Collapse",
			"Schizotypal Probability Infiltration",
			"Apeirophobic Sensory Recursion",
			"Synesthetic Pain Transduction",
			"Capgras Dimensional Synthesis",
			"Zero-Point Energy Erosion",
			"Cotard Delusion",
			"Ekphrastic Neurological Breach",
			"Prosopagnosia Recursion",
			"Quantum Decoherence",
			"Quantum Chromodynamic Dissolution",
			"Dermatographic Spacial Synthesis",
			"Catastrophe Threshold Fragmentation",
			"Reduplicative Sensory Transduction",
			"Tachyonic Shockwave",
			"Synkinesis Recursion",
			"Quantum Critical Point",
			"Jamais Vu Synthesis",
			"Reduplicative Paramnesia",
			"Quantum Vacuum Metastability",
			"Olfactory Hallucination",
			"Micropsia Breach",
			"Quantum Tunneling Negation",
			"Non-Orientable Topology Drift",
			"Quantum Vacuum Polarization",
			"Quantum Spin Liquid",
			"Quantum Phase Transition",
			"Semantic Satiation",
			"Quantum Scarring Erosion",
			"Apophenia Infiltration",
			"Noetic Super-Sealer",
			"Stochastic Phenomenological Shear",
			"Memetic Infiltration",
			"Dimensional Drift",
			"Neurodivergent Recursion",
			"Pataphysical Collapse",
			"Epistemic Occultation",
			"Biomechanical Infiltration",
			"Chronopathic Resonance",
			"Neuroplastic Corrosion",
			"Extrasensory Adaptation",
			"Cryptospatial Inability",
			"Metasemantic Warp",
			"Morphic Transduction",
			"Linguistic Quantum Field",
			"Emergent Topology",
			"Temporal Apophenia",
			"Semiotic Entanglement",
			"Akashic Cipher",
			"Hyperdimensional Logic",
			"Phenomenological Warp",
		}

		suffixStyles := []string{
			"Arabic-based",
			"Coptic-based",
			"Heian-Japanese",
			"Japanese Folklore",
			"Hellenic",
			"Indus-Valley",
			"Latin-based",
			"Old Slavic",
			"Russian Folklore",
			"Phoenician",
			"Teutonic",
			"Ugaritic",
			"Anime-style",
			"Aetherborne",
			"Bathyclastic",
			"Cybernetic",
			"Druidic",
			"Ecliptic",
			"Gravimetric",
			"Machiniform",
			"Noctivagant",
			"Oblivionic",
			"Orogenic",
			"Petrichoral",
			"Pyroclastic",
			"Quantum",
			"Runebound",
			"Steelflow",
			"Surrealist",
			"Sylvanic",
			"Ornitology Jargon",
			"Microbiology Jargon",
			"Schizofrenic",
			"Aquafungal",
			"Vortexian",
			"Lunarbotic",
			"Insectopunk",
			"Hologothic",
			"Plasma",
			"Glitchfolk",
			"Bloodpunk",
			"Hallucinogenic",
			"Paranoia-driven",
			"Horror-core",
			"Elasmobranchii",
			"Amnesia",
			"Laserbeam",
			"Symbiosis",
			"Blizzard",
			"Dissolving",
			"Orgasmic",
			"Fever Dream",
			"Overkill",
			"Ghost Duet",
			"Omnidirectional",
			"Inertia",
			"Decay",
			"Gluttonous",
			"Mentalist",
			"Kenjutsu",
			"Molecular",
			"Agony",
			"Cataclysm",
			"Sonar",
			"Neurosis",
			"Alien Gods",
			"Man-eating Lizard Dragon",
			"Ending",
			"Prostration",
			"Ominous",
			"Being or Entity",
			"Sensation",
			"Ability or Technique",
			"Terrain Manipulation",
			"Hazard Creation",
			"Mnemonic Device",
			"Strategic Sacrifice",
			"Weaponry",
			"Spell",
			"Cryptomnesia Recursive Trigger",
			"Antimatter Dimensional Collapse",
			"Schizotypal Probability Infiltration",
			"Apeirophobic Sensory Recursion",
			"Synesthetic Pain Transduction",
			"Capgras Dimensional Synthesis",
			"Zero-Point Energy Erosion",
			"Cotard Delusion",
			"Ekphrastic Neurological Breach",
			"Prosopagnosia Recursion",
			"Quantum Decoherence",
			"Quantum Chromodynamic Dissolution",
			"Dermatographic Spacial Synthesis",
			"Catastrophe Threshold Fragmentation",
			"Reduplicative Sensory Transduction",
			"Tachyonic Shockwave",
			"Synkinesis Recursion",
			"Quantum Critical Point",
			"Jamais Vu Synthesis",
			"Reduplicative Paramnesia",
			"Quantum Vacuum Metastability",
			"Olfactory Hallucination",
			"Micropsia Breach",
			"Quantum Tunneling Negation",
			"Non-Orientable Topology Drift",
			"Quantum Vacuum Polarization",
			"Quantum Spin Liquid",
			"Quantum Phase Transition",
			"Semantic Satiation",
			"Quantum Scarring Erosion",
			"Apophenia Infiltration",
			"Noetic Super-Sealer",
			"Stochastic Phenomenological Shear",
			"Memetic Infiltration",
			"Dimensional Drift",
			"Neurodivergent Recursion",
			"Pataphysical Collapse",
			"Epistemic Occultation",
			"Biomechanical Infiltration",
			"Chronopathic Resonance",
			"Neuroplastic Corrosion",
			"Extrasensory Adaptation",
			"Cryptospatial Inability",
			"Metasemantic Warp",
			"Morphic Transduction",
			"Linguistic Quantum Field",
			"Emergent Topology",
			"Temporal Apophenia",
			"Semiotic Entanglement",
			"Akashic Cipher",
			"Hyperdimensional Logic",
			"Phenomenological Warp",
		}

		prefixStyle := SelectOne("Summon an origin for the prefix\n", prefixStyles)
		suffixStyle := SelectOne("Summon an origin for the suffix\n", suffixStyles)

		lengthMin, lengthMax := SelectLengthLimits()

		GenOneWord(selectedStyles, wordFormat, lengthMin, lengthMax, prefixStyle, suffixStyle)

	} else if wordFormat == "Two-word" || wordFormat == "Two-word (Hyphenated)" {
		fullWordStyles := []string{
			"Arcane",
			"Cryptic",
			"Eldritch",
			"Occultic",
			"Phantasmal",
			"Runic",
			"Abyssal",
			"Baleful",
			"Exanimate",
			"Somber",
			"Thanatotic",
			"Voidborne",
			"Cybernetic",
			"Hyperreal",
			"Quantum",
			"Technomantic",
			"Posthuman",
			"Aetherial",
			"Dreamlike",
			"Liminal",
			"Numinous",
			"Subliminal",
			"Transcendent",
			"Archaic",
			"Feral",
			"Majestic",
			"Weirdling",
			"Zenithal",
			"Chimeric",
			"Eerie",
			"Spectral",
			"Uncanny",
		}

		firstWordStyle := SelectOne("Invoke an essence for the primal word\n", fullWordStyles)
		secondWordStyle := SelectOne("Invoke an essence for the subsequent word\n", fullWordStyles)

		lengthMin, lengthMax := SelectLengthLimits()

		GenTwoWord(selectedStyles, wordFormat, lengthMin, lengthMax, firstWordStyle, secondWordStyle)

	} else if wordFormat == "Three-word (Hyphenated)" {
		fullWordStyles := []string{
			"Arcane",
			"Cryptic",
			"Eldritch",
			"Occultic",
			"Phantasmal",
			"Runic",
			"Abyssal",
			"Baleful",
			"Exanimate",
			"Somber",
			"Thanatotic",
			"Voidborne",
			"Cybernetic",
			"Hyperreal",
			"Quantum",
			"Technomantic",
			"Posthuman",
			"Aetherial",
			"Dreamlike",
			"Liminal",
			"Numinous",
			"Subliminal",
			"Transcendent",
			"Archaic",
			"Feral",
			"Majestic",
			"Weirdling",
			"Zenithal",
			"Chimeric",
			"Eerie",
			"Spectral",
			"Uncanny",
		}

		firstWordStyle := SelectOne("Invoke an essence for the primal word\n", fullWordStyles)
		secondWordStyle := SelectOne("Invoke an essence for the subsequent word\n", fullWordStyles)
		thirdWordStyle := SelectOne("Invoke an essence for the ultimate word\n", fullWordStyles)

		lengthMin, lengthMax := SelectLengthLimits()

		GenThreeWord(selectedStyles, wordFormat, lengthMin, lengthMax, firstWordStyle, secondWordStyle, thirdWordStyle)

	} else if wordFormat == "Abstract Phrase" {
		phraseStyles := []string{
			"Abyssal Elegies",
			"Baudelarian",
			"Celestial Visions",
			"Cryptic Whispers",
			"Dark Reverberations",
			"Dreamlike Passages",
			"Ethereal Hymns",
			"Futuristic Paradoxes",
			"Liminal Insights",
			"Mythic Allegories",
			"Numinous Reveries",
			"Ontological Riddles",
			"Phantasmal Sonatas",
			"Rimbaudian Vignettes",
			"Subliminal Currents",
			"Zen Cataclysms",
		}

		selectedPhraseStyle := SelectOne("Invoke an essence for the abstract phrase\n", phraseStyles)

		lengthMin, lengthMax := SelectLengthLimits()

		GenAbstractPhrase(selectedStyles, wordFormat, lengthMin, lengthMax, selectedPhraseStyle)

	} else if wordFormat == "Acronym" {
		acronymStyles := []string{
			"Arcane-Scientific",
			"Cyber-Mystical",
			"Ethereal",
			"Liminal-Cosmic",
			"Neo-Mythic",
			"Philosophical-Tech",
			"Posthuman",
			"Quantum",
			"Surreal-Futuristic",
			"Temporal",
		}

		selectedAcronymStyle := SelectOne("Invoke an essence for the acronym\n", acronymStyles)

		lengthMin, lengthMax := SelectLengthLimits()

		GenAcronym(selectedStyles, wordFormat, lengthMin, lengthMax, selectedAcronymStyle)

	} else if wordFormat == "Alliterative" {
		alliterativeStyles := []string{
			"Abyssal Musings",
			"Celestial Visions",
			"Cryptic Whispers",
			"Dark Epiphanies",
			"Dreamlike Fragments",
			"Ethereal Echoes",
			"Futuristic Paradoxes",
			"Liminal Insights",
			"Mythic Allegories",
			"Numinous Reveries",
			"Ontological Riddles",
			"Phantasmal Threads",
			"Subliminal Currents",
			"Zen Contradictions",
		}

		selectedAlliterativeStyle := SelectOne("Invoke an essence for the alliteration\n", alliterativeStyles)

		lengthMin, lengthMax := SelectLengthLimits()

		GenAlliterative(selectedStyles, wordFormat, lengthMin, lengthMax, selectedAlliterativeStyle)

	} else if wordFormat == "Haiku" {
		haikuThemes := []string{
			"Apollyonic Reckoning",
			"Byronic Defiance",
			"Celestial Yearnings",
			"Chthonic Truths",
			"Dantesque Redemption",
			"Decadent Elegance",
			"Digital Oblivion",
			"Kabbalistic Insight",
			"Luciferian Resolve",
			"Miltonian Grandeur",
			"Nietzschean Will",
			"Orphic Mysteries",
			"Poe-esque Lament",
			"Posthuman Ascension",
			"Revolutionary Ardor",
			"Thanatological Reflections",
			"Venusian Decay",
			"Voidborne Yearnings",
		}

		selectedHaikuTheme := SelectOne("Invoke an essence for the haiku\n", haikuThemes)

		GenHaiku(selectedStyles, wordFormat, selectedHaikuTheme)

	} else if wordFormat == "Motto" {
		mottoThemes := []string{
			"Apollyonic Reckoning",
			"Byronic Defiance",
			"Celestial Yearnings",
			"Chthonic Truths",
			"Dantesque Redemption",
			"Decadent Elegance",
			"Digital Oblivion",
			"Kabbalistic Insight",
			"Luciferian Resolve",
			"Miltonian Grandeur",
			"Nietzschean Will",
			"Orphic Mysteries",
			"Poe-esque Lament",
			"Posthuman Ascension",
			"Revolutionary Ardor",
			"Thanatological Reflections",
			"Venusian Decay",
			"Voidborne Yearnings",
		}

		selectedMottoTheme := SelectOne("Invoke an essence for the motto\n", mottoThemes)

		GenMotto(selectedStyles, wordFormat, selectedMottoTheme)

	} else if wordFormat == "Portmanteau" {
		portmanteauStyles := []string{
			"Apollyonic Reckoning",
			"Byronic Defiance",
			"Celestial Yearnings",
			"Chthonic Truths",
			"Dantesque Redemption",
			"Decadent Elegance",
			"Digital Oblivion",
			"Kabbalistic Insight",
			"Luciferian Resolve",
			"Miltonian Grandeur",
			"Nietzschean Will",
			"Orphic Mysteries",
			"Poe-esque Lament",
			"Posthuman Ascension",
			"Revolutionary Ardor",
			"Thanatological Reflections",
			"Venusian Decay",
			"Voidborne Yearnings",
		}

		selectedPortmanteauStyle := SelectOne("Invoke an essence for the portmanteau\n", portmanteauStyles)

		lengthMin, lengthMax := SelectLengthLimits()

		GenPortmanteau(selectedStyles, wordFormat, lengthMin, lengthMax, selectedPortmanteauStyle)

	} else if wordFormat == "Symbolic Term" {
		symbolicThemes := []string{
			"Arcane",
			"Cryptic",
			"Eldritch",
			"Occultic",
			"Phantasmal",
			"Runic\n",

			"Abyssal",
			"Baleful",
			"Exanimate",
			"Somber",
			"Thanatotic",
			"Voidborne\n",

			"Cybernetic",
			"Hyperreal",
			"Quantum",
			"Technomantic",
			"Posthuman\n",

			"Aetherial",
			"Dreamlike",
			"Liminal",
			"Numinous",
			"Subliminal",
			"Transcendent\n",

			"Archaic",
			"Feral",
			"Majestic",
			"Weirdling",
			"Zenithal\n",

			"Chimeric",
			"Eerie",
			"Spectral",
			"Uncanny",
		}

		selectedSymbolicTheme := SelectOne("Invoke an essence for the symbolic name\n", symbolicThemes)

		lengthMin, lengthMax := SelectLengthLimits()

		GenSymbolic(selectedStyles, wordFormat, lengthMin, lengthMax, selectedSymbolicTheme)

	} else {
		fmt.Println("Unsupported word format selected.")
	}
}

func SelectMultiple(label string, options []string, maxSelect int) []string {
	var selected []string

	prompt := huh.NewMultiSelect[string]().
		Title(label).
		Options(huh.NewOptions(options...)...).
		Value(&selected).
		Limit(maxSelect)

	err := prompt.Run()
	if err != nil {
		fmt.Println("Error during multi-selection:", err)
		return nil
	}

	if len(selected) == 0 {
		fmt.Println("No options selected. Please select at least one.")
		return SelectMultiple(label, options, maxSelect)
	}

	return selected
}

func SelectOne(label string, options []string) string {
	var selected string

	prompt := huh.NewSelect[string]().
		Title(label).
		Options(huh.NewOptions(options...)...).
		Value(&selected)

	err := prompt.Run()
	if err != nil {
		fmt.Println("Error during single-selection:", err)
		return ""
	}

	return selected
}

func SelectLengthLimits() (int, int) {
	fmt.Println(AccentStyle.Render("Impose bounds upon its length (default 3-12): "))
	var min, max int

	fmt.Print(SecondaryStyle.Render("Lower Boundary: "))
	fmt.Scanln(&min)

	fmt.Print(SecondaryStyle.Render("Upper Boundary: "))
	fmt.Scanln(&max)

	if min == 0 {
		min = 3
	}
	if max == 0 {
		max = 12
	}

	Wiper()
	return min, max
}
