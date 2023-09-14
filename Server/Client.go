package server

type ClientState int

const (
	StateHandshake ClientState = 0
	StateStatus    ClientState = 1
	StateLogin     ClientState = 2
	StatePlay      ClientState = 3
)

type Client struct {
	state    ClientState
	username string
	uuid     [16]byte
}

func NewClient(a ClientState) Client {
	return Client{state: a}
}

func (c Client) UpdateState(state ClientState) {
	c.state = state
}

func (c Client) GetState() ClientState {
	return c.state
}

func (c Client) SetUsername(name string) {
	c.username = name
}

func (c Client) SetUUID(uuid [16]byte) {
	c.uuid = uuid
}
