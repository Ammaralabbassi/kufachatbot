package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", chatHandler)
	http.ListenAndServe(":8080", nil)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "text/html")

	// Set background color
	fmt.Fprintf(w, "<body style='background-color:#529c2b;'>")
	// Set title
	fmt.Fprintf(w, " <title>Chatbot</title>")
	// set header
	fmt.Fprintf(w, " <h1>Kufa Chatbot</h1>")

	if r.Method == "POST" {
		// Get user's input
		msg := r.FormValue("message")

		// Print user's message
		fmt.Fprintf(w, "<div style='background-color:#e6ffe6; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;'><strong>You:</strong> %s</div>", msg)

		// Check for keywords in user's message and send reply
		if strings.Contains(strings.ToLower(msg), "hello") {
			fmt.Fprintf(w, "<div style='background-color:#ffffff; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;'><strong>Chatbot:</strong> Hi there!</div>")
		} else if strings.Contains(strings.ToLower(msg), "time") {
			fmt.Fprintf(w, "<div style='background-color:#ffffff; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;'><strong>Chatbot:</strong> The current time is %s</div>", time.Now().Format("15:04:05"))
		} else if strings.Contains(strings.ToLower(msg), "weather") {
			fmt.Fprintf(w, "<div style='background-color:#ffffff; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;'><strong>Chatbot:</strong> The current weather is cloudy and cold </div>")
		} else if strings.Contains(strings.ToLower(msg), "traffic") {
			fmt.Fprintf(w, "<div style='background-color:#ffffff; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;'><strong>Chatbot:</strong> The traffic is good this morning </div>")
		} else if strings.Contains(strings.ToLower(msg), "to do list") {
			fmt.Fprintf(w, "<div style='background-color:#ffffff; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;'><strong>Chatbot:</strong> You have meeting at 10 Am. and another at 1 Pm. </div>")
		} else {
			fmt.Fprintf(w, "<div style='background-color:#ffffff; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;'><strong>Chatbot:</strong> I'm sorry, I didn't understand that.</div>")
		}
	}

	// Display chatbot input field and send button
	fmt.Fprintf(w, "<form method='POST'><input type='text' name='message' style='width: 80%%; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;' placeholder='Type your message here...'><br><button type='submit' style='background-color:#008CBA; color:#ffffff; padding:10px; border-radius:5px; margin-top:10px; margin-bottom:10px;'>Send</button></form>")

	// End HTML page
	fmt.Fprintf(w, "</body>")
}
