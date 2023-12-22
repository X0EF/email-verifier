package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	emailVerifier "github.com/AfterShip/email-verifier"
)

func GetEmailVerification(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	verifier := emailVerifier.NewVerifier()
	ret, err := verifier.Verify(ps.ByName("email"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if !ret.Syntax.Valid {
		_, _ = fmt.Fprint(w, "email address syntax is invalid")
		return
	}

	bytes, err := json.Marshal(ret)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	_, _ = fmt.Fprint(w, string(bytes))
}

func GetEmailsVerification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type RequestBody struct {
		Emails []string `json:"emails"`
	}
	var requestBody RequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	startTime := time.Now()
	results := emailVerifier.VerifyEmails(&requestBody.Emails)
	endTime := time.Now()
	fmt.Printf("without routine took %s to execute \n", endTime.Sub(startTime))
	startTime = time.Now()
	results2 := emailVerifier.VerifyEmails(&requestBody.Emails)
	endTime = time.Now()
	fmt.Printf("with routine took %s to execute \n", endTime.Sub(startTime))
	responseMap := map[string]interface{}{
		"data": results,
	}
	errx := json.NewEncoder(w).Encode(responseMap)
	if errx != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	println(len(results2))
}

func main() {
	router := httprouter.New()

	router.POST("/v1/bulk/verifications", GetEmailsVerification)
	router.GET("/v1/:email/verification", GetEmailVerification)

	log.Fatal(http.ListenAndServe(":8080", router))
}
