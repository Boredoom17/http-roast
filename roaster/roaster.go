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

	if r.Headers["Cache-Control"] != "" {
		s.CacheControl = 15
	}

	enc := r.Headers["Content-Encoding"]
	if strings.Contains(enc, "gzip") || strings.Contains(enc, "br") {
		s.Compression = 15
	}

	if r.Headers["Strict-Transport-Security"] != "" {
		s.SecurityHSTSS = 20
	}

	if r.Headers["X-Content-Type-Options"] != "" {
		s.SecurityXCTO = 15
	}

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

	// Response time
	bold.Println("⏱  Response Time:")
	ms := r.ResponseTime.Milliseconds()
	fmt.Printf("   %dms — ", ms)
	switch {
	case ms < 300:
		green.Println("Faster than my will to live. Respect.")
	case ms < 800:
		yellow.Println("Fine. Nobody's impressed but nobody's leaving either.")
	case ms < 1500:
		yellow.Println("My grandma loads faster and she's on dial-up in rural Nebraska.")
	case ms < 3000:
		red.Println("At this speed your bounce rate is just called 'everyone'.")
	default:
		red.Println("I started a family while waiting for this response.")
	}

	// Status code side comment
	fmt.Printf("   Status: %d", r.StatusCode)
	switch {
	case r.StatusCode == 200:
		green.Println(" — okay it's up at least. low bar, cleared.")
	case r.StatusCode >= 300 && r.StatusCode < 400:
		yellow.Println(" — redirecting. make up your mind about where you live.")
	case r.StatusCode >= 400 && r.StatusCode < 500:
		red.Println(" — client error. you broke it and you're blaming the user.")
	case r.StatusCode >= 500:
		red.Println(" — server error. it's on fire. you did this.")
	default:
		yellow.Printf(" — interesting status choice. bold.\n")
	}

	// Cache-Control
	bold.Println("\n📦  Cache-Control:")
	if s.CacheControl > 0 {
		green.Printf("   \"%s\"\n", r.Headers["Cache-Control"])
		green.Println("   Caching configured. You've done one thing right today. Don't ruin it.")
	} else {
		red.Println("   Not set. You're cooking fresh HTML for every single visitor like a deranged personal chef.")
		red.Println("   Your server is sweating. You did this to it.")
	}

	// Compression
	bold.Println("\n🗜  Compression (gzip/brotli):")
	enc := r.Headers["Content-Encoding"]
	if s.Compression > 0 {
		green.Printf("   %s detected.\n", enc)
		green.Println("   Compressed. Your bandwidth bill is slightly less embarrassing than it could be.")
	} else {
		red.Println("   Nothing. Raw. Uncompressed. In 2026.")
		red.Println("   You're shipping bytes like it's a personal vendetta against your users' data plans.")
	}

	// HSTS
	bold.Println("\n🔒  HSTS (Strict-Transport-Security):")
	if s.SecurityHSTSS > 0 {
		green.Println("   Present. Forcing HTTPS like a responsible adult who has read one security doc.")
	} else {
		red.Println("   Missing. No HSTS means browsers will happily downgrade to HTTP if nudged.")
		red.Println("   One DNS hijack away from a very awkward apology email to your users.")
	}

	// X-Content-Type-Options
	bold.Println("\n🛡  X-Content-Type-Options:")
	if s.SecurityXCTO > 0 {
		green.Println("   Set to nosniff. MIME sniffing blocked. You've clearly read at least one OWASP page.")
	} else {
		red.Println("   Missing. Browsers will sniff your content type and guess.")
		red.Println("   They will guess wrong. Congrats on your new attack surface.")
	}

	// X-Frame-Options
	bold.Println("\n🖼  X-Frame-Options:")
	if s.SecurityXFO > 0 {
		green.Printf("   Set to %s.\n", r.Headers["X-Frame-Options"])
		green.Println("   Clickjacking blocked. Your site isn't someone else's puppet. For now.")
	} else {
		red.Println("   Missing. Anyone can drop your site inside an iframe on a fake phishing page.")
		red.Println("   Very cool. Very normal. Absolutely fine.")
	}

	// Final score
	fmt.Println()
	bold.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	bold.Printf("FINAL SCORE: ")

	switch {
	case s.Total >= 85:
		green.Printf("%d/100\n", s.Total)
		green.Println("Verdict: Annoyingly good. I had a whole roast prepared and you ruined it.")
	case s.Total >= 60:
		yellow.Printf("%d/100\n", s.Total)
		yellow.Println("Verdict: The Toyota Camry of websites. Gets the job done. Inspires nothing.")
	case s.Total >= 35:
		red.Printf("%d/100\n", s.Total)
		red.Println("Verdict: Built during a hackathon, deployed on a dare, never touched again.")
	default:
		red.Printf("%d/100\n", s.Total)
		red.Println("Verdict: This is a crime scene. I'm not a cop but I'm filing a report.")
	}
	bold.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
}