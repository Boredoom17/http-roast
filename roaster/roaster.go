package roaster

import (
	"fmt"
	"http-roast/analyzer"
	"strings"

	"github.com/fatih/color"
)

type Score struct {
	ResponseTime  int
	CacheControl  int
	Compression   int
	SecurityHSTSS int
	SecurityXCTO  int
	SecurityXFO   int
	Total         int
}

func ScoreResult(r *analyzer.Result) Score {
	s := Score{}

	// Response time scoring (lower is better)
	ms := r.ResponseTime.Milliseconds()
	if ms < 300 {
		s.ResponseTime = 20
	} else if ms < 800 {
		s.ResponseTime = 15
	} else if ms < 1500 {
		s.ResponseTime = 8
	} else if ms < 3000 {
		s.ResponseTime = 3
	} else {
		s.ResponseTime = 0
	}

	// Cache-Control
	if r.Headers["Cache-Control"] != "" {
		s.CacheControl = 15
	}

	// Compression (gzip/br)
	enc := r.Headers["Content-Encoding"]
	if strings.Contains(enc, "gzip") || strings.Contains(enc, "br") {
		s.Compression = 15
	}

	// HSTS
	if r.Headers["Strict-Transport-Security"] != "" {
		s.SecurityHSTSS = 20
	}

	// X-Content-Type-Options
	if r.Headers["X-Content-Type-Options"] != "" {
		s.SecurityXCTO = 15
	}

	// X-Frame-Options
	if r.Headers["X-Frame-Options"] != "" {
		s.SecurityXFO = 15
	}

	s.Total = s.ResponseTime + s.CacheControl + s.Compression +
		s.SecurityHSTSS + s.SecurityXCTO + s.SecurityXFO

	return s
}

func Roast(r *analyzer.Result, s Score) {
	bold := color.New(color.Bold)
	red := color.New(color.FgRed, color.Bold)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgGreen)
	cyan := color.New(color.FgCyan, color.Bold)

	cyan.Println("\n🔥 HTTP ROAST REPORT 🔥")
	fmt.Printf("Target: %s\n\n", r.URL)

	// Response time roast
	bold.Println("⏱  Response Time:")
	ms := r.ResponseTime.Milliseconds()
	fmt.Printf("   %dms — ", ms)
	switch {
	case ms < 300:
		green.Println("Actually fast. Did you just flex on me?")
	case ms < 800:
		yellow.Println("Acceptable. Not impressive. Acceptable.")
	case ms < 1500:
		yellow.Println("Loading... still loading... oh there it is. Users are already gone.")
	case ms < 3000:
		red.Println("This site loads slower than a government PDF.")
	default:
		red.Println("I've seen continental drift faster than this response time.")
	}

	// Cache-Control roast
	bold.Println("\n📦  Cache-Control:")
	if s.CacheControl > 0 {
		green.Println("   Set. Good. You've heard of caching. Gold star.")
	} else {
		red.Println("   Missing. Serving fresh HTML to every request like a artisanal bread shop. Cute. Expensive.")
	}

	// Compression roast
	bold.Println("\n🗜  Compression:")
	if s.Compression > 0 {
		green.Println("   gzip/brotli enabled. You compressed your ego too? Impressive.")
	} else {
		red.Println("   No compression. Sending raw HTML like it's 2003. Dial-up users are crying.")
	}

	// HSTS roast
	bold.Println("\n🔒  HSTS (Strict-Transport-Security):")
	if s.SecurityHSTSS > 0 {
		green.Println("   Present. You remembered HTTPS exists. Have a cookie.")
	} else {
		red.Println("   Missing. One MITM attack away from a very bad day.")
	}

	// X-Content-Type-Options roast
	bold.Println("\n🛡  X-Content-Type-Options:")
	if s.SecurityXCTO > 0 {
		green.Println("   Set. MIME sniffing blocked. You're not completely reckless.")
	} else {
		red.Println("   Missing. Browsers can MIME-sniff your content. Hope nothing sensitive is in there.")
	}

	// X-Frame-Options roast
	bold.Println("\n🖼  X-Frame-Options:")
	if s.SecurityXFO > 0 {
		green.Println("   Set. Clickjacking blocked. Your users thank you (they don't know to, but they should).")
	} else {
		red.Println("   Missing. Your site can be iframed by anyone. Clickjacking party, you're invited.")
	}

	// Final score
	fmt.Println()
	bold.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	bold.Printf("FINAL SCORE: ")

	switch {
	case s.Total >= 85:
		green.Printf("%d/100\n", s.Total)
		green.Println("Verdict: Surprisingly not terrible. Ship it.")
	case s.Total >= 60:
		yellow.Printf("%d/100\n", s.Total)
		yellow.Println("Verdict: Mid. Like your site. It exists, technically.")
	case s.Total >= 35:
		red.Printf("%d/100\n", s.Total)
		red.Println("Verdict: This site is held together with duct tape and optimism.")
	default:
		red.Printf("%d/100\n", s.Total)
		red.Println("Verdict: Who deployed this? Name them. For the investigation.")
	}
	bold.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}