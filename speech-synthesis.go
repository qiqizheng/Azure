package main

import (
  "bufio"
  "fmt"
   "os"
   "strings"
   "time"
   
   "github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
   "github.com/Microsoft/cognitive-services-speech-sdk-go/common"
   "github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
)


func main(){

   speechkey := os.Getenv("SPEECH_KEY")  //定义密钥
   speechRegion := os.Getenv("SPEECH_REGION") //定义可用区
   
   audioConfig, err := audio.NewAudioConfigFromDefaultSpeakerOutput()
   if err != nil{
     fmt.Println("Got an error:" , err)
     return
   }
   
   defer

}
