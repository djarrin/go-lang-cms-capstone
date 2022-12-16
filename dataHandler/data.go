package dataHandler

var Data = map[int]Customer{
	1: Customer{Name: "David", Role: "Admin", Email: "david@gmail.com", Phone: "4443335678", Contacted: true},
	2: Customer{Name: "Sarah", Role: "Subscriber", Email: "sarah@gmail.com", Phone: "4445555678", Contacted: false},
	3: Customer{Name: "Keith", Role: "Admin", Email: "keith@gmail.com", Phone: "4446665678"},
	4: Customer{Name: "Linda", Role: "Admin", Email: "linda@gmail.com", Phone: "4447775678", Contacted: true},
}

func GetNextAvailaleID() int {
	ID := 0
	counter := 1
	for ID <= 0 {
		_, isPresent := Data[counter]
		if isPresent {
			counter += 1
		} else {
			ID = counter
		}
	}
	return ID
}
