package main

import (
	"os"
	"testing"
)

func TestForMultipleAddressWithFlag(t *testing.T) {

	os.Args = append(os.Args, "adjust.com")
	os.Args = append(os.Args, "google.com")
	os.Args = append(os.Args, "facebook.com")
	os.Args = append(os.Args, "yahoo.com")
	os.Args = append(os.Args, "yandex.com")

	main()
}

func TestArgsWithoutFlagThenPrintHash(t *testing.T) {
	os.Args = append(os.Args, "-parallel=4")
	os.Args = append(os.Args, "adjust.com")
	os.Args = append(os.Args, "google.com")
	os.Args = append(os.Args, "facebook.com")
	os.Args = append(os.Args, "yahoo.com")
	os.Args = append(os.Args, "yandex.com")

	main()
}
