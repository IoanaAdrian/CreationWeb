package utilities

import (
	"fmt"
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
func PathVisualisation(){
	cmd := exec.Command("/home/adrian/ADI/2WayEncryption/EncDecrCpp","-e","pukla")
	out, err := cmd.CombinedOutput()
	HandleErr(err)
	fmt.Println(string(out))
}