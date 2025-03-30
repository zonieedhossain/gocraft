# gocraft 🛠️

A modular CLI tool to scaffold Go backend projects with your own stack.

## 🚀 Features

- Choose your web framework (`fiber`, `echo`, `gin`, etc.)
- Select DB and ORM (`postgres`, `bun`, `gorm`)
- Add optional features: `--auth`, `--docker`, `--graphql`
- Clean, opinionated project structure

## 📦 Install

```bash
go install github.com/zonieedhossain/gocraft@latest
```

## ✨ Usage
```bash
gocraft new myapp --web=fiber --db=postgres --orm=bun --auth --docker
```

