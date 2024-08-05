package models

type (
	WordchainGetLastModel struct {
		w string
	}
	WordchainListModel struct {
		wc  []string
		len int
	}
)

func NewWordchainGetLastModel(w string) *WordchainGetLastModel {
	return &WordchainGetLastModel{
		w: w,
	}
}

func (m *WordchainGetLastModel) GetWord() string {
	return m.w
}

func NewWordchainListModel(wc []string, len int) *WordchainListModel {
	return &WordchainListModel{
		wc:  wc,
		len: len,
	}
}

func (m *WordchainListModel) GetWordchain() []string {
	return m.wc
}
func (m *WordchainListModel) GetLen() int {
	return m.len
}
