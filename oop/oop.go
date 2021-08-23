package oop

// const (
// 	StrNotImplemented = "Error not implemented!"
// )

type state interface {
	Progress()
	Submit()
	Approve()
	Reject()
	Publish()
}

type dummy struct{}

func (d *dummy) Progress() {
	// fmt.Println(StrNotImplemented)
}
func (d *dummy) Submit() {
	// fmt.Println(StrNotImplemented)
}
func (d *dummy) Approve() {
	// fmt.Println(StrNotImplemented)
}
func (d *dummy) Reject() {
	// fmt.Println(StrNotImplemented)
}
func (d *dummy) Publish() {
	// fmt.Println(StrNotImplemented)
}

type (
	// Add the logic in the statuses below
	Draft struct {
		dummy
		doc *Document
	}
	Completed struct {
		dummy
		doc *Document
	}
	Submitted struct {
		dummy
		doc *Document
	}
	Approved struct {
		dummy
		doc *Document
	}
	Rejected struct {
		dummy
		doc *Document
	}
	Published struct {
		dummy
		doc *Document
	}
)

type Document struct {
	state
	stateMap map[string]state
}

func NewDocument() *Document {
	doc := &Document{}
	doc.stateMap = map[string]state{
		"draft":     &Draft{doc: doc},
		"completed": &Completed{doc: doc},
		"submitted": &Submitted{doc: doc},
		"approved":  &Approved{doc: doc},
		"rejected":  &Rejected{doc: doc},
		"published": &Published{doc: doc},
	}
	return doc
}
func (d *Document) GetStateFromMap(input string) state { return d.stateMap[input] }
func (d *Document) GetState() state                    { return d.state }
func (d *Document) SetState(s state)                   { d.state = s }

// Complete the logic in the functions below
func (d *Document) Progress() {
	d.SetState(d.GetStateFromMap("completed"))
}
func (d *Document) Submit() {
	d.SetState(d.GetStateFromMap("submitted"))
}
func (d *Document) Approve() {
	d.SetState(d.GetStateFromMap("approved"))
}
func (d *Document) Reject() {
	d.SetState(d.GetStateFromMap("rejected"))
}
func (d *Document) Publish() {
	d.SetState(d.GetStateFromMap("published"))
}
