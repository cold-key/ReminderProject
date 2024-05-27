package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddReminder(t *testing.T) {
	reminderJSON := []byte(`{"creator":"user_1", "content":"Test reminder", "due_date":"2024-05-30T12:00:00Z"}`)
	req, err := http.NewRequest("POST", "/reminders", bytes.NewBuffer(reminderJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("/reminders", addReminderHandler)
	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("返回错误状态码: 获取 %v 期望 %v",
			status, http.StatusCreated)
	}
}

func TestDeleteReminder(t *testing.T) {

	//reminderStore := NewReminderStore()
	reminder := Reminder{
		ID:      "1",
		Creator: "user1",
		Content: "Test reminder",
		DueDate: time.Now(),
	}
	reminderStore.AddReminder(reminder)

	req, err := http.NewRequest("DELETE", "/reminders/1", nil)
	rr := httptest.NewRecorder()
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	deleteReminderHandler(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("返回错误状态码: 获取 %v 期望 %v",
			status, http.StatusNoContent)
	}

}
func TestGetRemindersByCreator(t *testing.T) {

	//reminderStore := NewReminderStore()
	reminder := Reminder{
		ID:      "1",
		Creator: "user1",
		Content: "Test reminder",
		DueDate: time.Now(),
	}
	reminderStore.AddReminder(reminder)

	req, err := http.NewRequest("Get", "/reminders/user1", nil)
	rr := httptest.NewRecorder()
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"creator": "user1",
	}
	req = mux.SetURLVars(req, vars)
	getRemindersByCreatorHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("返回错误状态码: 获取 %v 期望 %v",
			status, http.StatusOK)
	}

}
func TestUpdateReminder(t *testing.T) {
	reminder := Reminder{
		ID:      "1",
		Creator: "user1",
		Content: "Test reminder",
		DueDate: time.Now(),
	}
	reminderStore.AddReminder(reminder)
	reminderJSON := []byte(`{"id":"1", "creator":"user_1", "content":"New Test reminder", "due_date":"2024-05-30T12:00:00Z"}`)
	req, err := http.NewRequest("PUT", "/reminders", bytes.NewBuffer(reminderJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("/reminders", addReminderHandler)
	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("返回错误状态码: 获取 %v 期望 %v",
			status, http.StatusCreated)
	}
}
