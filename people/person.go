package people

type Person struct {
	firstName      string
	lastName       string
	age            uint8
	distanceWalked uint16
}

var Trevor Person = Person{firstName: "Trevor", lastName: "Sullivan", age: 30}

func (p *Person) Walk(distanceToWalk uint16) {
	p.distanceWalked += distanceToWalk
}
