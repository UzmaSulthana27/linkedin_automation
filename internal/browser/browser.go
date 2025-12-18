package browser

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
)

/* =====================
   STATE STRUCTURE
===================== */

type State struct {
	SentConnections []string `json:"sentConnections"`
	SentMessages    []string `json:"sentMessages"`
}

func loadState() State {
	file, err := os.ReadFile("data/state.json")
	if err != nil {
		return State{}
	}
	var state State
	_ = json.Unmarshal(file, &state)
	return state
}

func saveState(state State) {
	data, _ := json.MarshalIndent(state, "", "  ")
	_ = os.WriteFile("data/state.json", data, 0644)
}

/* =====================
   MAIN AUTOMATION FLOW
===================== */

func StartBrowser() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Load environment variables
	wsURL := os.Getenv("CHROME_WS")
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if wsURL == "" {
		log.Fatal("CHROME_WS not set in .env")
	}

	// Connect to Chrome
	browser := rod.New().
		ControlURL(wsURL).
		MustConnect()

	log.Println("Connected to Chrome")

	// Load state
	state := loadState()

	// Use ONE tab for full demo
	page := browser.MustPage()
	page.MustWindowMaximize()

	/* =====================
	   1️⃣ LOGIN AUTOMATION
	===================== */

	log.Println("Starting LOGIN automation")

	page.MustNavigate("file:///C:/Users/gamer/go-learning/mock/login.html")
	page.MustWaitLoad()

	page.MustElement("#email").MustInput(email)
	time.Sleep(1 * time.Second)

	page.MustElement("#password").MustInput(password)
	time.Sleep(1 * time.Second)

	page.MustElement("#loginBtn").MustClick()
	time.Sleep(3 * time.Second)

	log.Println("Login completed successfully")

	/* =====================
	   2️⃣ PROFILE AUTOMATION
	===================== */

	log.Println("Starting PROFILE automation")

	page.MustNavigate("file:///C:/Users/gamer/go-learning/mock/profiles.html")
	page.MustWaitLoad()

	profiles := page.MustElements(".profile")
	log.Println("Profiles found:", len(profiles))

	for i, profile := range profiles {

		name := profile.MustElement(".name").MustText()

		alreadyConnected := false
		for _, sent := range state.SentConnections {
			if sent == name {
				alreadyConnected = true
			}
		}

		if alreadyConnected {
			log.Println("Already connected to:", name)
			continue
		}

		log.Println("Sending connection to:", name)

		time.Sleep(time.Duration(2+i) * time.Second) // human delay
		profile.MustElement(".connect").MustClick()

		state.SentConnections = append(state.SentConnections, name)
		saveState(state)

		time.Sleep(2 * time.Second)
	}

	log.Println("Profile automation completed")

	/* =====================
	   3️⃣ MESSAGING AUTOMATION
	===================== */

	log.Println("Starting MESSAGING automation")

	page.MustNavigate("file:///C:/Users/gamer/go-learning/mock/messages.html")
	page.MustWaitLoad()

	chats := page.MustElements(".chat")

	for _, chat := range chats {

		name := chat.MustElement(".name").MustText()

		alreadyMessaged := false
		for _, sent := range state.SentMessages {
			if sent == name {
				alreadyMessaged = true
			}
		}

		if alreadyMessaged {
			log.Println("Message already sent to:", name)
			continue
		}

		message := "Hi " + name + ", thanks for connecting! Looking forward to staying in touch."

		log.Println("Sending message to:", name)

		chat.MustElement(".messageBox").MustInput(message)
		time.Sleep(2 * time.Second)

		chat.MustElement(".send").MustClick()
		time.Sleep(2 * time.Second)

		state.SentMessages = append(state.SentMessages, name)
		saveState(state)
	}

	log.Println("Messaging automation completed")
	log.Println("All automation steps finished successfully")

	// Keep browser open for demo
	select {}
}
