package modname

import "testing"

func TestGetModName(t *testing.T) {
	modName := GetModuleName()
	t.Logf("ModName : %s", modName)

	if modName == "" {
		t.Errorf("Get Modname Failed")
	}
}
