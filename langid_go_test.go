package langid_go

import (
	"testing"
)

type testVal struct {
	Text string
	Lang string
}

var testValues = []testVal{
	testVal{"Da die Anerkennung der angeborenen Würde und der gleichen und unveräußerlichen Rechte aller Mitglieder der Gemeinschaft der Menschen die Grundlage von Freiheit, Gerechtigkeit und Frie", "de"},
	testVal{"Considérant que la reconnaissance de la dignité inhérente à tous les membres de la famille humaine et de leurs droits égaux et inaliénables constitue le fondement de la liberté, de la justice et de la paix dans le monde", "fr"},
	testVal{"Considerato che il riconoscimento della dignità inerente a tutti i membri della famiglia umana e dei loro diritti, uguali ed inalienabili, costituisce il fondamento della libertà, della giustizia e della pace nel mondo; ", "it"},
	testVal{"Επειδή η αναγνώριση της αξιοπρέπειας, που είναι σύμφυτη σε όλα τα μέλη της ανθρώπινης οικογένειας, καθώς και των ίσων και αναπαλλοτρίωτων δικαιωμάτων τους αποτελεί το θεμέλιο της ελευθερίας, της δικαιοσύνης και της ειρήνης στον κόσμο.", "el"},
	testVal{"Whereas recognition of the inherent dignity and of the equal and inalienable rights of all members of the human family is the foundation of freedom, justice and peace in the world, ", "en"},
	testVal{"人類社会のすべての構成員の固有の尊厳と平等で譲ることのできない権利とを承認することは、世界における自由、正義及び平和の基礎であるので、", "ja"},
}

func TestDetectLang(t *testing.T) {
	lidfr := NewIdfr()
	defer lidfr.Destroy()
	for _, ld := range testValues {
		lang := lidfr.DetectLanguage(ld.Text)
		if lang != ld.Lang {
			t.Error("Expected: "+ld.Lang+" Got: ", lang)
		}
	}
}
