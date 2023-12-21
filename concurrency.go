package emailverifier

func VerifyEmails(emails []string) []Result {
	var results []Result
	verifier := NewVerifier()
	for _, email := range emails {
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
