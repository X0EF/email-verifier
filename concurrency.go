package emailverifier

import (
	"sync"
)

func VerifyEmails(emails *[]string) []Result {
	var results []Result
	verifier := NewVerifier()
	for _, email := range *emails {
		ret, err := verifier.Verify(email)
		if err != nil {
			continue
		}
		// if !ret.Syntax.Valid {
		// 	_, _ = fmt.Fprint(w, "email address syntax is invalid")
		// 	continue
		// }
		results = append(results, *ret)
	}
	return results
}

func VerifyEmailsV2(emails *[]string) []Result {
	var wg sync.WaitGroup
	resultChan := make(chan Result, len(*emails))
	var results []Result
	verifier := NewVerifier()
	for _, email := range *emails {
		wg.Add(1)
		go func(email string) {
			defer wg.Done()

			result, err := verifier.Verify(email)
			if err != nil {
			}

			// Send the result to the channel
			resultChan <- *result
		}(email)
		// if !ret.Syntax.Valid {
		// 	_, _ = fmt.Fprint(w, "email address syntax is invalid")
		// 	continue
		// }
		// results = append(results, *ret)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	for result := range resultChan {
		results = append(results, result)
	}
	return results
}
