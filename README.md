![Nephtegen Logo](media/neph-logo.png)

[![Go](https://img.shields.io/badge/Go-1.21%2B-blue?style=flat-square)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Terminal-lightgrey?style=flat-square)](#)
[![GitHub issues](https://img.shields.io/github/issues/bxavaby/mdice?style=flat-square)](https://github.com/bxavaby/mdice/issues)
[![GitHub forks](https://img.shields.io/github/forks/bxavaby/mdice?style=flat-square)](https://github.com/bxavaby/mdice/network)
[![GitHub stars](https://img.shields.io/github/stars/bxavaby/mdice?style=flat-square)](https://github.com/bxavaby/mdice/stargazers)

**Nephtegen** is a terminal-based creative assistant for conjuring names, phrases, haikus, acronyms, mottos, and more. Built in Go and powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Fabric](https://github.com/danielmiessler/fabric), it blends procedural logic with poetic invention. **Nephtegen** is ideal for naming systems, speculative fiction, or linguistic experimentation.

---

## ‣ Demo
![Nephtegen Demo](assets/neph.gif)  
*⊹ a walkthrough ⊹*

---

## ‣ Screenshots

---

## ‣ Showcase (exemplary products)

### One-word
- **Cytos** ⇝ *can be defined as a network or system of cells, reflecting the intersection of biological (neolithic, in terms of ancient, primal connections) and technological (cyberpunk) 
- **Tyomorph** ⇝ *can be defined as a powerful, isolated entity or device, possibly a weapon, that embodies the raw energy of the universe ("om") while warning of its tyrannical or overwhelming power ("ty")*

### Two-word
- **** ⇝ **
- **** ⇝ **

### Two-word (hyphenated)
- **** ⇝ **
- **** ⇝ **

### Three-word
- **** ⇝ **
- **** ⇝ **

### Three-word (hyphenated)
- **** ⇝ **
- **** ⇝ **

### Abstract Phrase
- ** ⇝ **
- ** ⇝ **

### Acronym
- **** - **
- **** - **

### Alliterative Phrase
- **** ⇝ **
- **** ⇝ **

### Haiku
- **** ⇝ **
- **** ⇝ **

### Motto
- *.* ⇝ **
- *.* ⇝ **

### Portmanteau
- **** ⇝ **
- **** ⇝ **

### Symbolic Term
- **** ⇝ **
- **** ⇝ **


 *Crafted by Nephtegen*

---

## ‣ Features

- **Interactive CLI**: Responsive TUI flow built with Bubble Tea & Huh.
- **Pattern-Based Prompting**: Uses markdown templates per format type, stored in `patterns/`.
- **Alphabet Injection**: Injects rare letterforms to reduce LLM repetition.
- **Multiple Naming Forms**:
  - One-word, Two-word, Three-word
  - Haikus, Acronyms, Mottos
  - Portmanteaus, Symbolic Terms
  - Abstract or Alliterative Phrases
- **Logging**: Names and haikus are archived automatically.
- **Modular Expansion**: Easy to add new output types or templates.

---

## ‣ Getting Started

### Prerequisites
- **Go 1.21+**
- **Fabric** installed and accessible via CLI
- Ensure `fabric` command works in your shell (`zsh` or compatible)

### Installation
```bash
git clone https://github.com/bxavaby/nephtegen.git
cd nephtegen
go build -o neph
```

You can optionally move the binary for global use:
```bash
sudo mv neph /usr/local/bin/
```

Then run it:
```bash
neph
```

---

## ‣ Project Structure

```plaintext
.
├── assets/             # GIFs and screenshots
├── patterns/           # markdown-based prompt templates
│   ├── haiku/
│   ├── motto/
│   ├── one-word/
│   ├── two-word/
│   ├── ...
├── neph.go             # entry point
├── plex.go             # prompt flow logic and UI routing
├── gen.go              # generation logic for each format
├── tagen.go            # styling, spinners, progress, logo
├── setts.go            # fabric session manager
├── letters.json        # letterforms for injection
├── LICENSE
└── README.md
```

---

## ‣ Patterns

Nephtegen does **not** load templates directly from the `patterns/` folder at runtime.

Instead, you must install them following [Fabric’s official pattern system](https://github.com/danielmiessler/fabric#custom-patterns). Each pattern must be placed inside the correct subdirectory of your local Fabric installation:

```bash
cp patterns/haiku/haikugen.md ~/.fabric/patterns/haiku/
```

---

## ‣ Technical Notes

By default, Nephtegen uses **Fabric’s `--session` flag** in its CLI prompts to group requests by output type (e.g., `haiku`, `motto`, `port`). This helps Fabric maintain contextual memory for iterative improvements — but depending on your **LLM backend**, this may result in:

- Slower response times
- Context length limitations
- Higher token usage or API throttling (especially when using OpenAI APIs)

If you experience performance issues or API rate errors, **you can safely remove the `--session` flags** in `gen.go`. Each generation function uses a helper like:
```go
callFabricX("haiku", "haikugen", prompt)
```

To disable session handling, simply swap it for:
```go
callFabric(prompt)
```

This invokes Fabric without persisting session memory. Output creativity may vary slightly, but performance will improve for API-based setups.

LLMs:
All examples below were generated using llama-3.3-70b-veratile (with Fabric, via Groq). While Nephtegen is model-agnostic, quality may vary across models.

---

## ‣ License
This project is licensed under the MIT License. See the LICENSE file for details.

