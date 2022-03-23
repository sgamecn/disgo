package discord

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/disgoorg/disgo/json"
)

type IconType string

//goland:noinspection GoUnusedConst
const (
	IconTypeJPEG    IconType = "image/jpeg"
	IconTypePNG     IconType = "image/png"
	IconTypeWEBP    IconType = "image/webp"
	IconTypeGIF     IconType = "image/gif"
	IconTypeUnknown          = IconTypeJPEG
)

func (t IconType) GetMIME() string {
	return string(t)
}

func (t IconType) GetHeader() string {
	return "data:" + string(t) + ";base64"
}

var _ json.Marshaler = (*Icon)(nil)
var _ fmt.Stringer = (*Icon)(nil)

//goland:noinspection GoUnusedExportedFunction
func NewIcon(iconType IconType, reader io.Reader) (*Icon, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return NewIconRaw(iconType, data), nil
}

//goland:noinspection GoUnusedExportedFunction
func NewIconRaw(iconType IconType, src []byte) *Icon {
	var data []byte
	base64.StdEncoding.Encode(data, src)
	return &Icon{Type: iconType, Data: data}
}

type Icon struct {
	Type IconType
	Data []byte
}

func (i Icon) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i Icon) String() string {
	if len(i.Data) == 0 {
		return ""
	}
	return i.Type.GetHeader() + "," + string(i.Data)
}
