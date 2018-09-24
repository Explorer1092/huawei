package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"gopkg.in/cheggaaa/pb.v1"
)

func calc(message_2 []byte,message_1 []byte) bool {
	h := sha256.New()
	h.Write(message_2)
	if bytes.Equal(message_1,h.Sum(nil)) {
		return true
	}else{
		return false
	}
}
func main() {
	var data string
	fmt.Println("input code: ")
	// data := "08 C2 3F 05 9E 91 D7 76 96 CC C1 E7 A4 A7 DE AE C8 FC E0 90 51 50 E6 C9 DF 15 21 7C 38 3C AA 79 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 AE E7 D8 F6 DE B3 F5 6B 5A 58 A6 95 F3 B0 CD 38 43 58 E9 7E 33 69 D5 B3 D8 AC AC 47 8B 20 06 38 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00"

	inputReader := bufio.NewReader(os.Stdin)
	data, _ = inputReader.ReadString('\n')

	// oem_code := 98925750
	data = strings.Replace(data," ","",-1)
	message_1,err := hex.DecodeString(data[0:64])
	if err != nil {
		log.Fatal(err)
	}
	message_2,err := hex.DecodeString(data[64*2:64*3])
	if err != nil {
		log.Fatal(err)
	}
	bar := pb.StartNew(99999999)
	bar.ShowBar = true
	for i:=99999999; i >0 ; i--  { // 98925750

		bar.Increment()
		h := sha256.New()
		h.Write([]byte(fmt.Sprintf("%08d",i)))
		s := [][]byte{h.Sum(nil), message_2}
		ds := bytes.Join(s, []byte(""))

		if calc(ds,message_1) {

			bar.FinishPrint("The End!\nOEM code = "+fmt.Sprintf("%08d",i))
			break
		}
	}

}

