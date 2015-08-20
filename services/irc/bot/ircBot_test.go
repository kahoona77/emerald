package bot

import (
	"testing"

	"github.com/kahoona77/emerald/models"
)

func TestMakeOutBuffer(t *testing.T) {
	buf := makeOutBuffer(234452)

	buf2 := makeOutBuffer(560980)

	if buf[0] != 0 || buf[1] != 3 || buf[2] != 147 || buf[3] != 212 {
		t.Error("buf error")
	}

	if buf2[0] != 0 || buf2[1] != 8 || buf2[2] != 143 || buf2[3] != 84 {
		t.Error("buf error")
	}
}

func TestCleanFilename(t *testing.T) {
	filename := "☻Beckoning.The.Butcher.2013.DVDRiP.X264-TASTE.mkv☼"
	result := cleanFileName(filename)

	if result != "Beckoning.The.Butcher.2013.DVDRiP.X264-TASTE.mkv" {
		t.Error("wrong filename: " + result)
	}
}

func TestGetTempFile(t *testing.T) {
	settings := new(models.XtvSettings)
	settings.TempDir = "c:/temp"
	dcc := DccService{}
	fileEvent := DccFileEvent{"SEND", "simpsons.mkv", nil, "", 0}

	file := dcc.getTempFile(&fileEvent, settings)

	if file == nil {
		t.Error("no temp file")
	} else if file.Name() != "c:\\temp\\simpsons.mkv" {
		t.Error("wrong file: " + file.Name())
	}
}
