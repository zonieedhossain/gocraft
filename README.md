# gocraft 🛠️

A modular CLI tool to scaffold Go backend projects with your preferred stack in seconds.

---

## 🚀 Features

- Choose your web framework: `fiber`, `echo`, `gin`
- Select database: `postgres`, `mysql`, `sqlite`
- Choose ORM: `bun`, `gorm`, `sqlc`
- Add optional features: `--auth`, `--docker`
- Define module path via:
- `--github=username`
- `--gitlab=username`
- `--bitbucket=username`
- `--module-path=your/custom/modulepath`
- Generates `.env`, `go.mod`, `Dockerfile`, project structure and boilerplate

---

## 📦 Install

```bash
go install github.com/zonieedhossain/gocraft@latest
```

## ✨ Usage
```bash
gocraft new myapp \
--web=fiber \
--db=postgres \
--orm=bun \
--auth \
--docker \
--github=zonieedhossain
```

```bash
--gitlab=yourname
--bitbucket=yourname
--module-path=your.custom.vcs/path/to/project
```

