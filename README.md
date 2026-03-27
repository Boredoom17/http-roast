<div align="center">

# 🔥 http-roast

### A CLI tool that roasts your website's HTTP quality

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![CLI](https://img.shields.io/badge/CLI-Tool-black?style=for-the-badge&logo=gnometerminal&logoColor=white)]()
[![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)]()
[![Release](https://img.shields.io/github/v/release/Boredoom17/http-roast?style=for-the-badge)](https://github.com/Boredoom17/http-roast/releases)

Point it at any URL. Get a brutal, honest breakdown of its HTTP performance and security headers.

</div>

---

## 📖 Overview

`http-roast` is a Go CLI tool that analyzes any public website and roasts it based on real engineering metrics — response time, caching, compression, and security headers.

No browser. No UI. Just run it in your terminal and watch websites get destroyed.

### ✨ What it checks

| Check | Max Points | What it means |
|---|---|---|
| ⏱ Response Time | 20 | How fast the server responds |
| 📦 Cache-Control | 15 | Whether caching headers are set |
| 🗜 Compression | 15 | gzip or brotli enabled |
| 🔒 HSTS | 20 | Forces HTTPS on all connections |
| 🛡 X-Content-Type-Options | 15 | Blocks MIME sniffing attacks |
| 🖼 X-Frame-Options | 15 | Prevents clickjacking via iframes |

**Total: 100 points.** Most sites don't even hit 60.

---

## 🚀 Usage

**Build from source:**
```bash
git clone https://github.com/Boredoom17/http-roast
cd http-roast
go build -o http-roast main.go
```

**Run it:**
```bash
./http-roast https://yoursite.com
```

**JSON output (for APIs or piping):**
```bash
./http-roast --json https://yoursite.com
```

---

## 💀 Example Output
```
🔥 HTTP ROAST REPORT 🔥
Target: https://craigslist.org

⏱  Response Time:
   7372ms — I started a family while waiting for this response.
   Status: 200 — okay it's up at least. low bar, cleared.

📦  Cache-Control:
   Caching configured. You've done one thing right today. Don't ruin it.

🗜  Compression (gzip/brotli):
   Nothing. Raw. Uncompressed. In 2026.
   You're shipping bytes like it's a personal vendetta against your users' data plans.

🔒  HSTS:
   Missing. One DNS hijack away from a very awkward apology email to your users.

🛡  X-Content-Type-Options:
   Missing. Browsers will sniff your content type and guess.
   They will guess wrong. Congrats on your new attack surface.

🖼  X-Frame-Options:
   Missing. Anyone can drop your site inside an iframe on a fake phishing page.
   Very cool. Very normal. Absolutely fine.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
FINAL SCORE: 35/100
Verdict: Built during a hackathon, deployed on a dare, never touched again.
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

---

## 🛠️ Tech Stack

- **Language**: Go
- **HTTP**: `net/http` standard library
- **CLI**: [Cobra](https://cobra.dev/)
- **Colors**: [fatih/color](https://github.com/fatih/color)
- **Deploy**: Single compiled binary — no runtime, no dependencies

## 📁 Project Structure
```
http-roast/
├── main.go
├── cmd/
│   └── root.go          ← CLI entry point (cobra)
├── analyzer/
│   └── analyzer.go      ← HTTP fetch + header analysis
├── roaster/
│   └── roaster.go       ← Scoring + roast output
├── go.mod
└── README.md
```

---

<div align="center">

Built by **[Boredoom17](https://github.com/Boredoom17)** — because someone had to say it lol

</div>
