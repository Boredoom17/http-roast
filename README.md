# 🔥 http-roast

A CLI tool that analyzes any website's HTTP performance and security headers — then roasts it.

## Install

Download the binary from [Releases](../../releases) or build from source:

go build -o http-roast main.go

## Usage

./http-roast https://example.com

# JSON output
./http-roast --json https://example.com

## What it checks

| Check | Max Points |
|---|---|
| Response Time | 20 |
| Cache-Control | 15 |
| Compression (gzip/br) | 15 |
| HSTS | 20 |
| X-Content-Type-Options | 15 |
| X-Frame-Options | 15 |

## Example output

🔥 HTTP ROAST REPORT 🔥
Target: https://craigslist.org

⏱  Response Time:
   7372ms — I've seen continental drift faster than this response time.

FINAL SCORE: 35/100
Verdict: This site is held together with duct tape and optimism.

## Stack
- Go + net/http
- Cobra CLI
- fatih/color