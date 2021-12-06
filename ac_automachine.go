package lad

const rootRaw = "/"

type acNode struct {
	raw      string
	children map[string]*acNode
	isEnd    bool
	length   int
	fail     *acNode
}

func NewAcNode(raw string) *acNode {
	return &acNode{raw: raw}
}

func (ac *acNode) Add(pattern string) error {
	return nil
}

func (ac *acNode) Load(path string) error {
	return nil
}

func (ac *acNode) Build() {

}

func (ac *acNode) Match(text string) {

}
