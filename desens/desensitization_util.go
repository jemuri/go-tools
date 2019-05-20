package desens

// PhoneNumber desensitization phone number
// For example: 18566667777 ==>>> 185****7777
func PhoneNumber(phone string) string {
	n := []byte(phone)
	if len(n) != 11 {
		return phone
	}
	var p []byte
	for i, v := range n {
		if i >= 3 && i <= 6 {
			v = 42
		}
		p = append(p, v)
	}

	return string(p)
}
