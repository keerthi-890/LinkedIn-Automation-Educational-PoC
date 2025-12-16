package core

import (
	"database/sql"
	"log"
	"time"
	"math/rand"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"

	"linkedin-automation-poc/stealth"
	"linkedin-automation-poc/storage"
)

type Bot struct {
	Browser *rod.Browser
	Page    *rod.Page
	DB      *sql.DB
	Running bool
	Stealth *stealth.Humanizer
}

func NewBot(db *sql.DB) *Bot {
	return &Bot{
		DB:      db,
		Stealth: stealth.NewHumanizer(),
	}
}

func (b *Bot) Start() error {
	if b.Running {
		return nil
	}

	// Absolute path for PoC mock site
	path := "C:\\Users\\mkeer\\Desktop\\LinkedIn\\mock-site\\network.html"
	url := "file:///" + path

	// Launch browser (visible for demo)
	u := launcher.New().
		Headless(false).
		MustLaunch()

	b.Browser = rod.New().ControlURL(u).MustConnect()
	b.Running = true

	log.Println("Browser launched")

	// Open mock site
	var err error
	b.Page, err = b.Browser.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		return err
	}

	b.Page.MustWaitLoad()
	log.Println("Navigated to mock site")

	// Run automation loop
	go b.runAutomationLoop()

	return nil
}

func (b *Bot) runAutomationLoop() {
	log.Println("Starting automation loop...")

	for b.Running {

		buttons, err := b.Page.Elements(".btn-connect")
		if err != nil {
			log.Println("Error finding buttons:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if len(buttons) == 0 {
			log.Println("No buttons found, waiting...")
			time.Sleep(5 * time.Second)
			continue
		}

		for _, btn := range buttons {
			if !b.Running {
				return
			}

			txt, _ := btn.Text()
			if txt == "Pending" {
				continue
			}

			// Move mouse (Rod-supported API)
			b.moveMouseToElement(btn)

			// Think time
			b.Stealth.SimulateThinkTime()

			// Click
			btn.MustClick()

			// Log activity
			storage.LogActivity(b.DB, "CONNECT", "Sent connection request")
			log.Println("Sent connection request")

			// Reading delay
			b.Stealth.SimulateReadingDelay()
		}

		time.Sleep(2 * time.Second)
	}
}

func (b *Bot) moveMouseToElement(el *rod.Element) {
	// Hover over the element
	el.MustHover()

	// Small random delay to simulate human behavior
	time.Sleep(100*time.Millisecond + time.Duration(rand.Intn(100))*time.Millisecond)
}


func (b *Bot) Stop() {
	b.Running = false
	if b.Browser != nil {
		b.Browser.MustClose()
	}
	log.Println("Bot stopped")
}
