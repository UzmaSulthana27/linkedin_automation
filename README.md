# LinkedIn Automation – Technical Assignment (Proof of Concept)

## 📌 Overview
This project is a proof-of-concept browser automation system built using **Go (Golang)** and the **Rod** browser automation library.

The application demonstrates how LinkedIn-like workflows such as:
- Login
- Profile discovery
- Connection requests
- Follow-up messaging

can be automated **safely and ethically** using **mock HTML pages**.

⚠️ This project does **NOT** automate real LinkedIn accounts and is intended **only for evaluation and learning purposes**.

---

## 🛠 Tech Stack
- **Language:** Go (Golang)
- **Browser Automation:** Rod
- **Browser:** Google Chrome (DevTools Protocol)
- **State Persistence:** JSON
- **Configuration:** JSON
- **OS Tested:** Windows

---
## Environment Variables

This project uses environment variables for configuration and credentials.

Variables are stored in a `.env` file (not committed to Git) and loaded at runtime.

### Required variables
- LINKEDIN_EMAIL – mock login email
- LINKEDIN_PASSWORD – mock login password
- CHROME_WS – Chrome DevTools WebSocket URL

An example template is provided in `.env.example`.

## 🏗 Project Structure


go-learning/
│
├── cmd/
│   └── app/
│       └── main.go          # Application entry point
│
├── internal/
│   └── browser/
│       └── browser.go       # Browser automation logic
│
├── mock/
│   ├── login.html           # Mock login page
│   ├── profiles.html        # Mock profiles page
│   └── messages.html        # Mock messaging page
│
├── data/
│   └── state.json           # Persistent automation state
│
├── .env.example             # Environment variable template
├── go.mod                   # Go module file
└── README.md                # Project documentation
