package openGraph

import (
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
	"image/color"
	"path/filepath"
)

const (
	width  = 1200
	height = 628
	margin = 20.0
)

func CreateOpenGraphImage(pageTitle string) error {

	dc := gg.NewContext(width, height)
	err, done := loadBackgroundImage(dc)

	if !done {
		return err
	}

	backgroundShadow(dc)

	err, done = addlogotext(dc)
	if !done {
		return err
	}

	err, done = addurlText(dc)
	if !done {
		return err
	}

	err, done = titleText(dc, pageTitle)
	if !done {
		return err
	}

	err, done = saveImage(dc)
	if !done {
		return err
	}

	return nil

}

func titleText(dc *gg.Context, pageTitle string) (error, bool) {
	if pageTitle == "" {
		pageTitle = "Hamza U.f. Ghani"
	}

	title := pageTitle
	textShadowColor := color.Black
	textColor := color.White
	fontPath := filepath.Join("/tmp", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 40); err != nil {
		return errors.Wrap(err, "load Playfair_Display"), false
	}
	textRightMargin := 60.0
	textTopMargin := 90.0
	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.SetColor(textShadowColor)
	dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	dc.SetColor(textColor)
	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	return nil, true
}

func addurlText(dc *gg.Context) (error, bool) {
	textColor := color.White
	fontPath := filepath.Join("/tmp", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 40); err != nil {
		return errors.Wrap(err, "load Open_Sans"), false
	}
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	dc.SetColor(mutedColor)
	marginY := 30.0
	urltext := "https://hufghani.dev"
	_, textHeight := dc.MeasureString(urltext)
	x := 70.0
	y := float64(dc.Height()) - textHeight - marginY
	dc.DrawString(urltext, x, y)
	return nil, true
}

func backgroundShadow(dc *gg.Context) (float64, float64) {
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{0, 0, 0, 100})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()
	return x, y
}

func addlogotext(dc *gg.Context) (error, bool) {
	fontPath := filepath.Join("/tmp", "DrSugiyama-Regular.ttf")
	if err := dc.LoadFontFace(fontPath, 45); err != nil {
		return errors.Wrap(err, "load font"), false
	}
	dc.SetColor(color.White)
	text := "Hamza U. F. Ghani"
	marginX := 50.0
	marginY := -20.0
	textWidth, textHeight := dc.MeasureString(text)
	x := float64(dc.Width()) - textWidth - marginX
	y := float64(dc.Height()) - textHeight - marginY
	dc.DrawString(text, x, y)

	return nil, true
}

func saveImage(dc *gg.Context) (error, bool) {
	if err := dc.SavePNG("/tmp/outputFilename.png"); err != nil {
		return errors.Wrap(err, "save png"), false
	}
	return nil, true
}

func loadBackgroundImage(dc *gg.Context) (error, bool) {
	backgroundImage, err := gg.LoadImage(filepath.Join("/tmp", "background.jpeg"))
	backgroundImage = imaging.Fill(backgroundImage, dc.Width(), dc.Height(), imaging.Center, imaging.Lanczos)

	if err != nil {
		return errors.Wrap(err, " failed to load background image"), false
	}

	dc.DrawImage(backgroundImage, 0, 0)
	return nil, true

}
