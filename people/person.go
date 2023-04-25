package people

type Person struct {
	FirstName      string
	LastName       string
	Age            uint8
	DistanceWalked uint16
}

var Trevor Person = Person{FirstName: "Trevor", LastName: "Sullivan", Age: 30}

func (p *Person) Walk(distanceToWalk uint16) {
	p.DistanceWalked += distanceToWalk
}
