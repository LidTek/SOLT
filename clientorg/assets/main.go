package assets

import (
	"embed"
	"path"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed **/**/*.ttf
var FS embed.FS

var DisplayFont = "Barlow"

type FontGroup struct {
	Fonts      map[text.Weight]*text.GoTextFace
	LineHeight float64
}

func (fg FontGroup) GetFace(weight text.Weight) *text.GoTextFace {
	if fg.Fonts[weight] != nil {
		return fg.Fonts[weight]
	}
	return fg.Fonts[text.WeightNormal]
}

var FontGroups = make(map[string]FontGroup)

func init() {
	dir, err := FS.ReadDir("fonts")
	if err != nil {
		panic("no fonts dir")
	}
	for _, d := range dir {
		if d.IsDir() {
			group := FontGroup{
				Fonts: make(map[text.Weight]*text.GoTextFace),
			}
			dir, err := FS.ReadDir("fonts/" + d.Name())
			if err != nil {
				continue
			}
			for _, f := range dir {
				if !f.IsDir() {
					if !strings.HasSuffix(f.Name(), ".ttf") {
						continue
					}
					fh, err := FS.Open(path.Join("fonts", d.Name(), f.Name()))
					if err != nil {
						continue
					}
					s, err := text.NewGoTextFaceSource(fh)
					if err != nil {
						continue
					}
					if strings.Contains(f.Name(), "Bold") {
						group.Fonts[text.WeightBold] = &text.GoTextFace{
							Source: s,
							Size:   32,
						}
					} else {
						group.Fonts[text.WeightNormal] = &text.GoTextFace{
							Source: s,
							Size:   32,
						}
					}
				}
			}
			if len(group.Fonts) > 0 {
				FontGroups[d.Name()] = group
			}
		}
	}
}
