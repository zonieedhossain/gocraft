# gocraft 🛠️

> CLI tool to scaffold production-ready Go microservices with Clean Architecture, configurable ORM, auth bootstrapping and Docker — out of the box.

Built by a Senior Backend Engineer who got tired of setting up the same structure across every project.

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

---

## 🧱 Generated Structure
Clean Architecture out of the box — handlers, usecases, repositories and infrastructure separated by concern.

---

## 👤 Author
**Md. Zonieed Hossain** — Senior Backend Engineer (Go)
[LinkedIn](https://linkedin.com/in/zonieedhossain) · [GitHub](https://github.com/zonieedhossain)
