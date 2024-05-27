package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Reminder struct {
	ID      string    `json:"id"`
	Creator string    `json:"creator"`
	Content string    `json:"content"`
	DueDate time.Time `json:"due_date"`
}

type ReminderStore struct {
	Reminders map[string]Reminder
}

func NewReminderStore() *ReminderStore {
	return &ReminderStore{
		Reminders: make(map[string]Reminder),
	}
}

func (rs *ReminderStore) AddReminder(reminder Reminder) {
	rs.Reminders[reminder.ID] = reminder
}

func (rs *ReminderStore) GetRemindersByCreator(creatorID string) []Reminder {
	var reminders []Reminder
	for _, reminder := range rs.Reminders {
		if reminder.Creator == creatorID {
			reminders = append(reminders, reminder)
		}
	}
	return reminders
}

func (rs *ReminderStore) DeleteReminder(reminderID string) {
	delete(rs.Reminders, reminderID)
}

func (rs *ReminderStore) UpdateReminder(reminder Reminder) {
	rs.Reminders[reminder.ID] = reminder
}

func addReminderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter add")
	var reminder Reminder
	err := json.NewDecoder(r.Body).Decode(&reminder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reminder.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	//fmt.Println(reminder.ID)
	//fmt.Println(reminder.Creator)
	//fmt.Println(reminder.Content)
	//fmt.Println(reminder.DueDate)
	reminderStore.AddReminder(reminder)
	w.WriteHeader(http.StatusCreated)
}

func getRemindersByCreatorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter get")

	params := mux.Vars(r)
	creatorID := params["creator"]
	if creatorID == "" {
		http.Error(w, "Creator ID not provided", http.StatusBadRequest)
		return
	}

	reminders := reminderStore.GetRemindersByCreator(creatorID)
	json.NewEncoder(w).Encode(reminders)
}

func deleteReminderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter delete")
	reminderID := mux.Vars(r)["id"]
	if reminderID == "" {
		http.Error(w, "Reminder ID not provided", http.StatusBadRequest)
		return
	}

	reminderStore.DeleteReminder(reminderID)
	w.WriteHeader(http.StatusNoContent)
}

func updateReminderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter update")

	var updatedReminder Reminder
	err := json.NewDecoder(r.Body).Decode(&updatedReminder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println(updatedReminder.ID)
	//fmt.Println(updatedReminder.Creator)
	//fmt.Println(updatedReminder.Content)
	//fmt.Println(updatedReminder.DueDate)
	if updatedReminder.ID == "" {
		http.Error(w, "Reminder ID not provided", http.StatusBadRequest)
		return
	}
	reminderStore.UpdateReminder(updatedReminder)
	w.WriteHeader(http.StatusNoContent)
}

var reminderStore = NewReminderStore()

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/reminders", addReminderHandler).Methods("POST")
	router.HandleFunc("/reminders/{creator}", getRemindersByCreatorHandler).Methods("GET")
	router.HandleFunc("/reminders/{id}", deleteReminderHandler).Methods("DELETE")
	router.HandleFunc("/reminders", updateReminderHandler).Methods("PUT")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}
