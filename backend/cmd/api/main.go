package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type quoteRequest struct {
	ID               string `json:"id"`
	FullName         string `json:"fullName"`
	PhoneNumber      string `json:"phoneNumber"`
	EmailAddress     string `json:"emailAddress"`
	ServiceType      string `json:"serviceType"`
	PickupPostcode   string `json:"pickupPostcode"`
	DeliveryPostcode string `json:"deliveryPostcode"`
	PreferredDate    string `json:"preferredDate"`
	FlexibleOnDate   bool   `json:"flexibleOnDate"`
	AdditionalNotes  string `json:"additionalNotes"`
	Status           string `json:"status"`
	CreatedAt        string `json:"createdAt"`
}

type booking struct {
	ID               string `json:"id"`
	QuoteRequestID   string `json:"quoteRequestId"`
	CustomerName     string `json:"customerName"`
	PhoneNumber      string `json:"phoneNumber"`
	ServiceType      string `json:"serviceType"`
	PickupPostcode   string `json:"pickupPostcode"`
	DeliveryPostcode string `json:"deliveryPostcode"`
	BookingDate      string `json:"bookingDate"`
	Status           string `json:"status"`
	Notes            string `json:"notes"`
	CreatedAt        string `json:"createdAt"`
}

type message struct {
	ID          string `json:"id"`
	SenderName  string `json:"senderName"`
	SenderEmail string `json:"senderEmail"`
	SenderPhone string `json:"senderPhone"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	IsRead      bool   `json:"isRead"`
	CreatedAt   string `json:"createdAt"`
}

type contentSection struct {
	ID     string            `json:"id"`
	Label  string            `json:"label"`
	Fields map[string]string `json:"fields"`
}

type store struct {
	mu       sync.RWMutex
	quotes   []quoteRequest
	bookings []booking
	messages []message
	content  []contentSection
}

func main() {
	app := &store{
		content: []contentSection{
			{ID: "hero", Label: "Hero Section", Fields: map[string]string{"badge": "UK's Most Trusted Removals", "headline": "Moving Made Simple.", "description": "Professional man and van removals across the UK."}},
			{ID: "about", Label: "About Us", Fields: map[string]string{"heading": "Moves Done Right.", "description": "Reliable, professional, always on time.", "experienceText": "2 Years Experience"}},
			{ID: "contact", Label: "Contact Info", Fields: map[string]string{"phone": "+44 7424 849252", "email": "info@dachremovals.co.uk"}},
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", app.health)
	mux.HandleFunc("POST /api/auth/login", app.login)
	mux.HandleFunc("GET /api/quotes", app.listQuotes)
	mux.HandleFunc("POST /api/quotes", app.createQuote)
	mux.HandleFunc("GET /api/bookings", app.listBookings)
	mux.HandleFunc("POST /api/bookings", app.createBooking)
	mux.HandleFunc("GET /api/messages", app.listMessages)
	mux.HandleFunc("POST /api/messages", app.createMessage)
	mux.HandleFunc("GET /api/content", app.getContent)
	mux.HandleFunc("PUT /api/content", app.updateContent)

	addr := env("ADDR", ":8080")
	log.Printf("dach removals backend listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, withCORS(mux)))
}

func (s *store) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "service": "dachremovals-backend"})
}

func (s *store) login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if body.Username != env("ADMIN_USERNAME", "admin") || body.Password != env("ADMIN_PASSWORD", "password") {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"token": "prototype-admin-token"})
}

func (s *store) listQuotes(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	writeJSON(w, http.StatusOK, s.quotes)
}

func (s *store) createQuote(w http.ResponseWriter, r *http.Request) {
	var item quoteRequest
	if !decode(w, r, &item) {
		return
	}
	item.ID = newID("quote")
	item.Status = defaultString(item.Status, "new")
	item.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	s.mu.Lock()
	s.quotes = append([]quoteRequest{item}, s.quotes...)
	s.mu.Unlock()
	writeJSON(w, http.StatusCreated, item)
}

func (s *store) listBookings(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	writeJSON(w, http.StatusOK, s.bookings)
}

func (s *store) createBooking(w http.ResponseWriter, r *http.Request) {
	var item booking
	if !decode(w, r, &item) {
		return
	}
	item.ID = newID("booking")
	item.Status = defaultString(item.Status, "new")
	item.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	s.mu.Lock()
	s.bookings = append(s.bookings, item)
	s.mu.Unlock()
	writeJSON(w, http.StatusCreated, item)
}

func (s *store) listMessages(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	writeJSON(w, http.StatusOK, s.messages)
}

func (s *store) createMessage(w http.ResponseWriter, r *http.Request) {
	var item message
	if !decode(w, r, &item) {
		return
	}
	item.ID = newID("message")
	item.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	s.mu.Lock()
	s.messages = append([]message{item}, s.messages...)
	s.mu.Unlock()
	writeJSON(w, http.StatusCreated, item)
}

func (s *store) getContent(w http.ResponseWriter, _ *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	writeJSON(w, http.StatusOK, s.content)
}

func (s *store) updateContent(w http.ResponseWriter, r *http.Request) {
	var body []contentSection
	if !decode(w, r, &body) {
		return
	}
	s.mu.Lock()
	s.content = body
	s.mu.Unlock()
	writeJSON(w, http.StatusOK, body)
}

func decode(w http.ResponseWriter, r *http.Request, target any) bool {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return false
	}
	return true
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func newID(prefix string) string {
	return prefix + "_" + strings.ReplaceAll(time.Now().UTC().Format("20060102150405.000000000"), ".", "")
}

func defaultString(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
