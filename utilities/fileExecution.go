package utilities

import (
	"os/exec"
)

func Encrypt(s string) string {
	cmd := exec.Command("EncDecrCpp", "-e", s)
	out, err := cmd.CombinedOutput()
	HandleErr(err)
	return string(out)
}
func Decrypt(s string, num string) string {
	cmd := exec.Command("EncDecrCpp", "-d", s,num)
	out, err := cmd.CombinedOutput()
	HandleErr(err)
	return string(out)
}
