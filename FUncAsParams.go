package main

func dataFilter(nama string, filter func(string) string) string {
	filtered := filter(nama)
	return filtered
}

func spamFilter(nama string) string {
	if nama == "asu" {
		return "..."
	} else {
		return "Hello " + nama
	}
}

func main() {

	dataFilter("asu", spamFilter)

}
