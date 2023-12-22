package emailverifier

import (
	"sync"
)

type ErrorEmail struct {
	Email   string `json:"email"` // passed email address
	Message string `json:"message"`
}

type ResultAndError struct {
	Results []Result     `json:"results"`
	Errors  []ErrorEmail `json:"errors"`
}

func VerifyEmails(emails *[]string) ResultAndError {
	var wg sync.WaitGroup
	resultChan := make(chan Result, len(*emails))
	var results []Result
	var errors []ErrorEmail
	verifier := NewVerifier()
	for _, email := range *emails {
		wg.Add(1)
		go func(email string) {
			defer wg.Done()
			result, err := verifier.Verify(email)
			if err != nil {
				error := ErrorEmail{Email: email, Message: err.Error()}
				errors = append(errors, error)
			}
			resultChan <- *result
		}(email)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for result := range resultChan {
		results = append(results, result)
	}
	return ResultAndError{Results: results, Errors: errors}
}
