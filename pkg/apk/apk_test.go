package apk

import (
	"log"
	"testing"
)

func TestApk_PackageName(t *testing.T) {


	pkg, _ := OpenFile("DxLampLauncher_0.0.9_9.apk")

	defer pkg.Close()

	pkgName := pkg.PackageName() // returns the pakcage name

	log.Println(pkgName)


}