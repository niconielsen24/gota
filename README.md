# 🛠️ gota

**Gota** is a personal task runner, project scaffold generator, and general-purpose dev automation tool written in Go.

Think of it like [`task`](https://taskfile.dev) or `make`, but with fewer features, more Go, and some questionable design decisions — because sometimes you just want to scratch that technical itch and build your own thing.

---

## ✨ Features

- 📦 **Project scaffolding**: `gota create TYPE APPNAME`
- ✅ **Task runner**: reads a simple YAML file and runs defined tasks (WIP or TBD depending on your implementation)
- ⚙️ **All-in-one binary**: thanks to `//go:embed`, no external template files or config mess
- 🧩 Easily modifiable and extensible
- 🧪 Made to learn Go and have fun doing it
- 🐍 Built using [Cobra](https://github.com/spf13/cobra) for CLI magic

---

---

## 📌 Usage

```bash
# List tasks (from gota.yaml in the current directory)
gota list

# Create a new CLI project
gota create cli myapp
```

---

---

### More command details:

    create TYPE APPNAME: Scaffolds a new project of the given type

    list: Lists available tasks defined in gota.yaml

    Tip: Start by copying over some templates and modifying them to suit your own setups.

---

---

### 📁 Example Template Layout

Gota uses embedded .tpl files for scaffolding, like:

> internal/templates/
> ├── main.go.tpl
> ├── mod.tpl
> └── root.go.tpl

## These are bundled into the binary and rendered with Go's standard text/template.

---

### 🚀 Why not just use task?

#### You should! task is a fantastic tool and the inspiration behind this project.

#### But sometimes, you just want to build your own — to:

    Learn Go better

    Explore CLIs, templates, and embedding

    Customize the experience

### 🧠 Goals & Philosophy

    🧰 Built for personal dev workflow automation

    🔍 Code is simple, hackable, and meant to be extended

    💡 Learn by building

    💬 Open to improvements, but no plans to productionize

### 🔧 Future Ideas

    Built-in support for build, test, lint, format

    Auto-init gota.yaml

    Watch mode / file triggers

    More project types / templates
