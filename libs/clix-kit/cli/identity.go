package cli

type Identity struct {
	name        string
	description string
}

func NewIdentity(name, description string) *Identity {
	return &Identity{name: name, description: description}
}

func (i *Identity) GetName() string {
	return i.name
}

func (i *Identity) GetDescription() string {
	return i.description
}
