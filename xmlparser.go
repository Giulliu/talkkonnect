/*
 * talkkonnect headless mumble client/gateway with lcd screen and channel control
 * Copyright (C) 2018-2019, Suvir Kumar <suvir@talkkonnect.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 * Software distributed under the License is distributed on an "AS IS" basis,
 * WITHOUT WARRANTY OF ANY KIND, either express or implied. See the License
 * for the specific language governing rights and limitations under the
 * License.
 *
 * talkkonnect is the based on talkiepi and barnard by Daniel Chote and Tim Cooper
 *
 * The Initial Developer of the Original Code is
 * Suvir Kumar <suvir@talkkonnect.com>
 * Portions created by the Initial Developer are Copyright (C) Suvir Kumar. All Rights Reserved.
 *
 * Contributor(s):
 *
 * Suvir Kumar <suvir@talkkonnect.com>
 *
 * My Blog is at www.talkkonnect.com
 * The source code is hosted at github.com/talkkonnect
 *
 * xmlparser.go -> talkkonnect functionality to read from XML file and populate global variables
 */

package talkkonnect

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// lcd timer
var (
	BackLightTime    = time.NewTimer(1 * time.Millisecond)
	BackLightTimePtr = &BackLightTime
)

//account settings
var (
	Default     []bool
	Name        []string
	Server      []string
	Username    []string
	Password    []string
	Insecure    []bool
	Certificate []string
	Channel     []string
	Ident       []string
)

//software settings
var (
	OutputDevice       string
	LogFileNameAndPath string
	Logging            string
	Daemonize          bool
)

//autoprovision settings
var (
	APEnabled    bool
	TkId         string
	Url          string
	SaveFilePath string
	SaveFileName string
)

//beacon settings
var (
	BeaconEnabled         bool
	BeaconTimerSecs       int
	BeaconFileNameAndPath string
	BVolume               float32
)

//tts
var (
	TTSEnabled                           bool
	TTSVolumeLevel                       int
	TTSParticipants                      bool
	TTSChannelUp                         bool
	TTSChannelUpFileNameAndPath          string
	TTSChannelDown                       bool
	TTSChannelDownFileNameAndPath        string
	TTSMuteUnMuteSpeaker                 bool
	TTSMuteUnMuteSpeakerFileNameAndPath  string
	TTSCurrentVolumeLevel                bool
	TTSCurrentVolumeLevelFileNameAndPath string
	TTSDigitalVolumeUp                   bool
	TTSDigitalVolumeUpFileNameAndPath    string
	TTSDigitalVolumeDown                 bool
	TTSDigitalVolumeDownFileNameAndPath  string
	TTSListServerChannels                bool
	TTSListServerChannelsFileNameAndPath string
	TTSStartTransmitting                 bool
	TTSStartTransmittingFileNameAndPath  string
	TTSStopTransmitting                  bool
	TTSStopTransmittingFileNameAndPath   string
	TTSListOnlineUsers                   bool
	TTSListOnlineUsersFileNameAndPath    string
	TTSPlayChimes                        bool
	TTSPlayChimesFileNameAndPath         string
	TTSRequestGpsPosition                bool
	TTSRequestGpsPositionFileNameAndPath string
	TTSNextServer                        bool
	TTSNextServerFileNameAndPath         string
	TTSPanicSimulation                   bool
	TTSPanicSimulationFileNameAndPath    string
	TTSPrintXmlConfig                    bool
	TTSPrintXmlConfigFileNameAndPath     string
	TTSSendEmail                         bool
	TTSSendEmailFileNameAndPath          string
	TTSDisplayMenu                       bool
	TTSDisplayMenuFileNameAndPath        string
	TTSQuitTalkkonnect                   bool
	TTSQuitTalkkonnectFileNameAndPath    string
	TTSTalkkonnectLoaded                 bool
	TTSTalkkonnectLoadedFileNameAndPath  string
	TTSPingServers                       bool
	TTSPingServersFileNameAndPath        string
	TTSScan                              bool
	TTSScanFileNameAndPath               string
)

//gmail smtp settings
var (
	EmailEnabled       bool
	EmailUsername      string
	EmailPassword      string
	EmailReceiver      string
	EmailSubject       string
	EmailMessage       string
	EmailGpsDateTime   bool
	EmailGpsLatLong    bool
	EmailGoogleMapsUrl bool
)

//sound settings
var (
	EventSoundEnabled             bool
	EventSoundFilenameAndPath     string
	AlertSoundEnabled             bool
	AlertSoundFilenameAndPath     string
	AlertSoundVolume              float32
	RogerBeepSoundEnabled         bool
	RogerBeepSoundFilenameAndPath string
	RogerBeepSoundVolume          float32
	ChimesSoundEnabled            bool
	ChimesSoundFilenameAndPath    string
	ChimesSoundVolume             float32
)

//api settings
var (
	APIEnabled            bool
	APIListenPort         string
	APIDisplayMenu        bool
	APIChannelUp          bool
	APIChannelDown        bool
	APIMute               bool
	APICurrentVolumeLevel bool
	APIDigitalVolumeUp    bool
	APIDigitalVolumeDown  bool
	APIListServerChannels bool
	APIStartTransmitting  bool
	APIStopTransmitting   bool
	APIListOnlineUsers    bool
	APIPlayChimes         bool
	APIRequestGpsPosition bool
	APINextServer         bool
	APIPanicSimulation    bool
	APIPrintXmlConfig     bool
	APIEmailEnabled       bool
	APIPingServersEnabled bool
	APIScan               bool
)

//print xml config sections for easy debugging, set any section to false to prevent printing to screen
var (
	PrintAccount      bool
	PrintLogging      bool
	PrintProvisioning bool
	PrintBeacon       bool
	PrintTTS          bool
	PrintSMTP         bool
	PrintSounds       bool
	PrintTxTimeout    bool
	PrintHTTPAPI      bool
	PrintTargetboard  bool
	PrintLeds         bool
	PrintHeartbeat    bool
	PrintButtons      bool
	PrintComment      bool
	PrintLcd          bool
	PrintGps          bool
	PrintPanic        bool
)

// target board settings
var (
	TargetBoard string
)

//indicator light settings
var (
	VoiceActivityLEDPin uint
	ParticipantsLEDPin  uint
	TransmitLEDPin      uint
	OnlineLEDPin        uint
)

//heartbeat light settings
var (
	HeartBeatEnabled bool
	HeartBeatLEDPin  uint
	PeriodmSecs      int
	LEDOnmSecs       int
	LEDOffmSecs      int
)

//button settings
var (
	TxButtonPin    uint
	UpButtonPin    uint
	DownButtonPin  uint
	PanicButtonPin uint
)

//comment settings
var (
	CommentButtonPin  uint
	CommentMessageOff string
	CommentMessageOn  string
)

//lcd screen settings
var (
	LCDInterfaceType         string
	LCDI2CAddress            uint8
	LCDBackLightTimerEnabled bool
	LCDBackLightTimeoutSecs  int
	BackLightLEDPin          int
	RSPin                    int
	EPin                     int
	D4Pin                    int
	D5Pin                    int
	D6Pin                    int
	D7Pin                    int
)

//txtimeout settings
var (
	TxTimeOutEnabled bool
	TxTimeOutSecs    int
)

//gps settings
var (
	GpsEnabled          bool
	Port                string
	Baud                uint
	TxData              string
	Even                bool
	Odd                 bool
	Rs485               bool
	Rs485HighDuringSend bool
	Rs485HighAfterSend  bool
	StopBits            uint
	DataBits            uint
	CharTimeOut         uint
	MinRead             uint
	Rx                  bool
)

//panic function settings
var (
	PEnabled           bool
	PFileNameAndPath   string
	PMessage           string
	PRecursive         bool
	PVolume            float32
	PSendIdent         bool
	PSendGpsLocation   bool
	PTxLockEnabled     bool
	PTxlockTimeOutSecs uint
)

type Document struct {
	XMLName  xml.Name `xml:"document"`
	Type     string   `xml:"type,attr"`
	Accounts Accounts `xml:"accounts"`
	Global   Global   `xml:"global"`
}
type Accounts struct {
	XMLName  xml.Name  `xml:"accounts"`
	Accounts []Account `xml:"account"`
}

type Account struct {
	XMLName       xml.Name `xml:"account"`
	Name          string   `xml:"name,attr"`
	Default       bool     `xml:"default,attr"`
	ServerAndPort string   `xml:"serverandport"`
	UserName      string   `xml:"username"`
	Password      string   `xml:"password"`
	Insecure      bool     `xml:"insecure"`
	Certificate   string   `xml:"certificate"`
	Channel       string   `xml:"channel"`
	Ident         string   `xml:"ident"`
}

type Global struct {
	XMLName  xml.Name `xml:"global"`
	Software Software `xml:"software"`
	Hardware Hardware `xml:"hardware"`
}

type Software struct {
	XMLName          xml.Name         `xml:"software"`
	AutoProvisioning AutoProvisioning `xml:"autoprovisioning"`
	Beacon           Beacon           `xml:"beacon"`
	Settings         Settings         `xml:"settings"`
	Smtp             Smtp             `xml:"smtp"`
	Sounds           Sounds           `xml:"sounds"`
	TxTimeOut        TxTimeOut        `xml:"txtimeout"`
	API              API              `xml:"api"`
	PrintVariables   PrintVariables   `xml:"printvariables"`
	TTS              TTS              `xml:"tts"`
}

type Settings struct {
	XMLName            xml.Name `xml:"settings"`
	OutputDevice       string   `xml:"outputdevice"`
	LogFileNameAndPath string   `xml:"logfilenameandpath"`
	Logging            string   `xml:"logging"`
	Daemonize          bool     `xml:"daemonize"`
}

type AutoProvisioning struct {
	XMLName      xml.Name `xml:"autoprovisioning"`
	APEnabled    bool     `xml:"enabled,attr"`
	TkId         string   `xml:"tkid"`
	Url          string   `xml:"url"`
	SaveFilePath string   `xml:"savefilepath"`
	SaveFileName string   `xml:"savefilename"`
}

type Beacon struct {
	XMLName               xml.Name `xml:"beacon"`
	BeaconEnabled         bool     `xml:"enabled,attr"`
	BeaconTimerSecs       int      `xml:"beacontimersecs"`
	BeaconFileNameAndPath string   `xml:"beaconfileandpath"`
	BVolume               float32  `xml:"volume"`
}

type TTS struct {
	XMLName                              xml.Name `xml:"tts"`
	TTSEnabled                           bool     `xml:"enabled,attr"`
	TTSVolumeLevel                       int      `xml:"volumelevel"`
	TTSParticipants                      bool     `xml:"participants"`
	TTSChannelUp                         bool     `xml:"channelup"`
	TTSChannelUpFileNameAndPath          string   `xml:"channelupfilenameandpath"`
	TTSChannelDown                       bool     `xml:"channeldown"`
	TTSChannelDownFileNameAndPath        string   `xml:"channeldownfilenameandpath"`
	TTSMuteUnMuteSpeaker                 bool     `xml:"muteunmutespeaker"`
	TTSMuteUnMuteSpeakerFileNameAndPath  string   `xml:"muteunmutespeakerfilenameandpath"`
	TTSCurrentVolumeLevel                bool     `xml:"currentvolumelevel"`
	TTSCurrentVolumeLevelFileNameAndPath string   `xml:"currentvolumelevelfilenameandpath"`
	TTSDigitalVolumeUp                   bool     `xml:"digitalvolumeup"`
	TTSDigitalVolumeUpFileNameAndPath    string   `xml:"digitalvolumeupfilenameandpath"`
	TTSDigitalVolumeDown                 bool     `xml:"digitalvolumedown"`
	TTSDigitalVolumeDownFileNameAndPath  string   `xml:"digitalvolumedownfilenameandpath"`
	TTSListServerChannels                bool     `xml:"listserverchannels"`
	TTSListServerChannelsFileNameAndPath string   `xml:"listserverchannelsfilenameandpath"`
	TTSStartTransmitting                 bool     `xml:"starttransmitting"`
	TTSStartTransmittingFileNameAndPath  string   `xml:"starttransmittingfilenameandpath"`
	TTSStopTransmitting                  bool     `xml:"stoptransmitting"`
	TTSStopTransmittingFileNameAndPath   string   `xml:"stoptransmittingfilenameandpath"`
	TTSListOnlineUsers                   bool     `xml:"listonlineusers"`
	TTSListOnlineUsersFileNameAndPath    string   `xml:"listonlineusersfilenameandpath"`
	TTSPlayChimes                        bool     `xml:"playchimes"`
	TTSPlayChimesFileNameAndPath         string   `xml:"playchimesfilenameandpath"`
	TTSRequestGpsPosition                bool     `xml:"requestgpsposition"`
	TTSRequestGpsPositionFileNameAndPath string   `xml:"requestgpspositionfilenameandpath"`
	TTSNextServer                        bool     `xml:"nextserver"`
	TTSNextServerFileNameAndPath         string   `xml:"nextserverfilenameandpath"`
	TTSPanicSimulation                   bool     `xml:"panicsimulation"`
	TTSPanicSimulationFileNameAndPath    string   `xml:"panicsimulationfilenameandpath"`
	TTSPrintXmlConfig                    bool     `xml:"printxmlconfig"`
	TTSPrintXmlConfigFileNameAndPath     string   `xml:"printxmlconfigfilenameandpath"`
	TTSSendEmail                         bool     `xml:"sendemail"`
	TTSSendEmailFileNameAndPath          string   `xml:"sendemailfilenameandpath"`
	TTSDisplayMenu                       bool     `xml:"displaymenu"`
	TTSDisplayMenuFileNameAndPath        string   `xml:"displaymenufilenameandpath"`
	TTSQuitTalkkonnect                   bool     `xml:"quittalkkonnect"`
	TTSQuitTalkkonnectFileNameAndPath    string   `xml:"quittalkkonnectfilenameandpath"`
	TTSTalkkonnectLoaded                 bool     `xml:"talkkonnectloaded"`
	TTSTalkkonnectLoadedFileNameAndPath  string   `xml:"talkkonnectloadedfilenameandpath"`
	TTSPingServers                       bool     `xml:"pingservers"`
	TTSPingServersFileNameAndPath        string   `xml:"pingserversfilenameandpath"`
}

type Smtp struct {
	XMLName            xml.Name `xml:"smtp"`
	EmailEnabled       bool     `xml:"enabled,attr"`
	EmailUsername      string   `xml:"username"`
	EmailPassword      string   `xml:"password"`
	EmailReceiver      string   `xml:"receiver"`
	EmailSubject       string   `xml:"subject"`
	EmailMessage       string   `xml:"message"`
	EmailGpsDateTime   bool     `xml:"gpsdatetime"`
	EmailGpsLatLong    bool     `xml:"gpslatlong"`
	EmailGoogleMapsUrl bool     `xml:"googlemapsurl"`
}

type Sounds struct {
	XMLName   xml.Name  `xml:"sounds"`
	Event     Event     `xml:"event"`
	Alert     Alert     `xml:"alert"`
	RogerBeep RogerBeep `xml:"rogerbeep"`
	Chimes    Chimes    `xml:"chimes"`
}

type API struct {
	XMLName               xml.Name `xml:"api"`
	APIEnabled            bool     `xml:"enabled,attr"`
	APIListenPort         string   `xml:"apilistenport"`
	APIDisplayMenu        bool     `xml:"displaymenu"`
	APIChannelUp          bool     `xml:"channelup"`
	APIChannelDown        bool     `xml:"channeldown"`
	APIMute               bool     `xml:"mute"`
	APICurrentVolumeLevel bool     `xml:"currentvolumelevel"`
	APIDigitalVolumeUp    bool     `xml:"digitalvolumeup"`
	APIDigitalVolumeDown  bool     `xml:"digitalvolumedown"`
	APIListServerChannels bool     `xml:"listserverchannels"`
	APIStartTransmitting  bool     `xml:"starttransmitting"`
	APIStopTransmitting   bool     `xml:"stoptransmitting"`
	APIListOnlineUsers    bool     `xml:"listonlineusers"`
	APIPlayChimes         bool     `xml:"playchimes"`
	APIRequestGpsPosition bool     `xml:"requestgpsposition"`
	APINextServer         bool     `xml:"nextserver"`
	APIPanicSimulation    bool     `xml:"panicsimulation"`
	APIPrintXmlConfig     bool     `xml:"printxmlconfig"`
	APIEmailEnabled       bool     `xml:"sendemail"`
	APIPingServersEnabled bool     `xml:"pingservers"`
}

type PrintVariables struct {
	XMLName           xml.Name `xml:"printvariables"`
	PrintAccount      bool     `xml:"printaccount"`
	PrintLogging      bool     `xml:"printlogging"`
	PrintProvisioning bool     `xml:"printprovisioning"`
	PrintBeacon       bool     `xml:"printbeacon"`
	PrintTTS          bool     `xml:"printts"`
	PrintSMTP         bool     `xml:"printsmtp"`
	PrintSounds       bool     `xml:"printsounds"`
	PrintTxTimeout    bool     `xml:"printtxtimeout"`
	PrintHTTPAPI      bool     `xml:"printhttpapi"`
	PrintTargetboard  bool     `xml:"printtargetboard"`
	PrintLeds         bool     `xml:"printleds"`
	PrintHeartbeat    bool     `xml:"printheartbeat"`
	PrintButtons      bool     `xml:"printbuttons"`
	PrintComment      bool     `xml:"printcomment"`
	PrintLcd          bool     `xml:"printlcd"`
	PrintGps          bool     `xml:"printgps"`
	PrintPanic        bool     `xml:"printpanic"`
}

type Event struct {
	XMLName          xml.Name `xml:"event"`
	EEnabled         bool     `xml:"enabled,attr"`
	EFileNameAndPath string   `xml:"filenameandpath"`
}

type Alert struct {
	XMLName          xml.Name `xml:"alert"`
	AEnabled         bool     `xml:"enabled,attr"`
	AFileNameAndPath string   `xml:"filenameandpath"`
	AVolume          float32  `xml:"volume"`
}

type RogerBeep struct {
	XMLName          xml.Name `xml:"rogerbeep"`
	REnabled         bool     `xml:"enabled,attr"`
	RFileNameAndPath string   `xml:"filenameandpath"`
	RBeepVolume      float32  `xml:"volume"`
}

type TxTimeOut struct {
	XMLName          xml.Name `xml:"txtimeout"`
	TxTimeOutEnabled bool     `xml:"enabled,attr"`
	TxTimeOutSecs    int      `xml:"txtimeoutsecs"`
}

type Chimes struct {
	XMLName          xml.Name `xml:"chimes"`
	CEnabled         bool     `xml:"enabled,attr"`
	CFileNameAndPath string   `xml:"filenameandpath"`
	CVolume          float32  `xml:"volume"`
}

type Hardware struct {
	XMLName       xml.Name      `xml:"hardware"`
	TargetBoard   string        `xml:"targetboard,attr"`
	Lights        Lights        `xml:"lights"`
	HeartBeat     HeartBeat     `xml:"heartbeat"`
	Buttons       Buttons       `xml:"buttons"`
	Comment       Comment       `xml:"comment"`
	Screen        Screen        `xml:"screen"`
	GPS           GPS           `xml:"gps"`
	PanicFunction PanicFunction `xml:"panicfunction"`
}

type Lights struct {
	XMLName             xml.Name `xml:"lights"`
	VoiceActivityLedPin uint     `xml:"voiceactivityledpin"`
	ParticipantsLedPin  uint     `xml:"participantsledpin"`
	TransmitLedPin      uint     `xml:"transmitledpin"`
	OnlineLedPin        uint     `xml:"onlineledpin"`
}

type HeartBeat struct {
	XMLName          xml.Name `xml:"heartbeat"`
	HeartBeatEnabled bool     `xml:"enabled,attr"`
	HeartBeatLEDPin  uint     `xml:"heartbeatledpin"`
	PeriodmSecs      int      `xml:"periodmsecs"`
	LEDOnmSecs       int      `xml:"ledonmsecs"`
	LEDOffmSecs      int      `xml:"ledoffmsecs"`
}

type Buttons struct {
	XMLName        xml.Name `xml:"buttons"`
	TxButtonPin    uint     `xml:"txbuttonpin"`
	UpButtonPin    uint     `xml:"upbuttonpin"`
	DownButtonPin  uint     `xml:"downbuttonpin"`
	PanicButtonPin uint     `xml:"panicbuttonpin"`
}

type Comment struct {
	XMLName           xml.Name `xml:"comment"`
	CommentButtonPin  uint     `xml:"commentbuttonpin"`
	CommentMessageOff string   `xml:"commentmessageoff"`
	CommentMessageOn  string   `xml:"commentmessageon"`
}

type Screen struct {
	XMLName                  xml.Name `xml:"screen"`
	LCDInterfaceType         string   `xml:"lcdinterfacetype"`
	LCDI2CAddress            uint8    `xml:"lcdi2caddress"`
	LCDBackLightTimerEnabled bool     `xml:"lcdbacklighttimerenabled"`
	LCDBackLightTimeoutSecs  int      `xml:"lcdbacklighttimeoutsecs"`
	BackLightLEDPin          int      `xml:"backlightpin"`
	RsPin                    int      `xml:"rspin"`
	EsPin                    int      `xml:"epin"`
	D4Pin                    int      `xml:"d4pin"`
	D5Pin                    int      `xml:"d5pin"`
	D6Pin                    int      `xml:"d6pin"`
	D7Pin                    int      `xml:"d7pin"`
}

type GPS struct {
	XMLName             xml.Name `xml:"gps"`
	GpsEnabled          bool     `xml:"enabled,attr"`
	Port                string   `xml:"port"`
	Baud                uint     `xml:"baud"`
	TxData              string   `xml:"txdata"`
	Even                bool     `xml:"even"`
	Odd                 bool     `xml:"odd"`
	Rs485               bool     `xml:"rs485"`
	Rs485highduringsend bool     `xml:"rs485highduringsend"`
	Rs485highaftersend  bool     `xml:"rs485highaftersend"`
	StopBits            uint     `xml:"stopbits"`
	DataBits            uint     `xml:"databits"`
	CharTimeOut         uint     `xml:"chartimeout"`
	MinRead             uint     `xml:"minread"`
	Rx                  bool     `xml:"rx"`
}

type PanicFunction struct {
	XMLName            xml.Name `xml:"panicfunction"`
	PEnabled           bool     `xml:"enabled,attr"`
	PMessage           string   `xml:"panicmessage"`
	PRecursive         string   `xml:"recursivesendmessage"`
	PFileNameAndPath   string   `xml:"filenameandpath"`
	PVolume            float32  `xml:"volume"`
	PSendIdent         bool     `xml:"sendident"`
	PSendGpsLocation   bool     `xml:"sendgpslocation"`
	PTxLockEnabled     bool     `xml:"txlockenabled"`
	PTxlockTimeOutSecs uint     `xml:"txlocktimeoutsecs"`
}

func readxmlconfig(file string) error {
	var counter int = 0
	xmlFile, err := os.Open(file)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot open configuration file talkkonnect.xml", err))
	}

	log.Println("info: Successfully Opened file talkkonnect.xml")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var document Document

	err = xml.Unmarshal(byteValue, &document)
	if err != nil {
		errors.New(fmt.Sprintf("File talkkonnect.xml formatting error Please fix! ", err))
	}
	log.Println("Document               : " + document.Type)

	for i := 0; i < len(document.Accounts.Accounts); i++ {
		if document.Accounts.Accounts[i].Default == true {
			Name = append(Name, document.Accounts.Accounts[i].Name)
			Server = append(Server, document.Accounts.Accounts[i].ServerAndPort)
			Username = append(Username, document.Accounts.Accounts[i].UserName)
			Password = append(Password, document.Accounts.Accounts[i].Password)
			Insecure = append(Insecure, document.Accounts.Accounts[i].Insecure)
			Certificate = append(Certificate, document.Accounts.Accounts[i].Certificate)
			Channel = append(Channel, document.Accounts.Accounts[i].Channel)
			Ident = append(Ident, document.Accounts.Accounts[i].Ident)
			counter++
		}
	}

	if counter == 0 {
		log.Fatal("No Default Accounts Found! Please Add at least 1 Default Account in XML File")
	}

	OutputDevice = document.Global.Software.Settings.OutputDevice
	LogFileNameAndPath = document.Global.Software.Settings.LogFileNameAndPath
	Logging = document.Global.Software.Settings.Logging
	Daemonize = document.Global.Software.Settings.Daemonize

	APEnabled = document.Global.Software.AutoProvisioning.APEnabled
	TkId = document.Global.Software.AutoProvisioning.TkId
	Url = document.Global.Software.AutoProvisioning.Url
	SaveFilePath = document.Global.Software.AutoProvisioning.SaveFilePath
	SaveFileName = document.Global.Software.AutoProvisioning.SaveFileName

	BeaconEnabled = document.Global.Software.Beacon.BeaconEnabled
	BeaconTimerSecs = document.Global.Software.Beacon.BeaconTimerSecs
	BeaconFileNameAndPath = document.Global.Software.Beacon.BeaconFileNameAndPath
	BVolume = document.Global.Software.Beacon.BVolume

	TTSEnabled = document.Global.Software.TTS.TTSEnabled
	TTSVolumeLevel = document.Global.Software.TTS.TTSVolumeLevel
	TTSParticipants = document.Global.Software.TTS.TTSParticipants
	TTSChannelUp = document.Global.Software.TTS.TTSChannelUp
	TTSChannelUpFileNameAndPath = document.Global.Software.TTS.TTSChannelUpFileNameAndPath
	TTSChannelDown = document.Global.Software.TTS.TTSChannelDown
	TTSChannelDownFileNameAndPath = document.Global.Software.TTS.TTSChannelDownFileNameAndPath
	TTSMuteUnMuteSpeaker = document.Global.Software.TTS.TTSMuteUnMuteSpeaker
	TTSMuteUnMuteSpeakerFileNameAndPath = document.Global.Software.TTS.TTSMuteUnMuteSpeakerFileNameAndPath
	TTSCurrentVolumeLevel = document.Global.Software.TTS.TTSCurrentVolumeLevel
	TTSCurrentVolumeLevelFileNameAndPath = document.Global.Software.TTS.TTSCurrentVolumeLevelFileNameAndPath
	TTSDigitalVolumeUp = document.Global.Software.TTS.TTSDigitalVolumeUp
	TTSDigitalVolumeUpFileNameAndPath = document.Global.Software.TTS.TTSDigitalVolumeUpFileNameAndPath
	TTSDigitalVolumeDown = document.Global.Software.TTS.TTSDigitalVolumeDown
	TTSDigitalVolumeDownFileNameAndPath = document.Global.Software.TTS.TTSDigitalVolumeDownFileNameAndPath
	TTSListServerChannels = document.Global.Software.TTS.TTSListServerChannels
	TTSListServerChannelsFileNameAndPath = document.Global.Software.TTS.TTSListServerChannelsFileNameAndPath
	TTSStartTransmitting = document.Global.Software.TTS.TTSStartTransmitting
	TTSStartTransmittingFileNameAndPath = document.Global.Software.TTS.TTSStartTransmittingFileNameAndPath
	TTSStopTransmitting = document.Global.Software.TTS.TTSStopTransmitting
	TTSStopTransmittingFileNameAndPath = document.Global.Software.TTS.TTSStopTransmittingFileNameAndPath
	TTSListOnlineUsers = document.Global.Software.TTS.TTSListOnlineUsers
	TTSListOnlineUsersFileNameAndPath = document.Global.Software.TTS.TTSListOnlineUsersFileNameAndPath
	TTSPlayChimes = document.Global.Software.TTS.TTSPlayChimes
	TTSPlayChimesFileNameAndPath = document.Global.Software.TTS.TTSPlayChimesFileNameAndPath
	TTSRequestGpsPosition = document.Global.Software.TTS.TTSRequestGpsPosition
	TTSRequestGpsPositionFileNameAndPath = document.Global.Software.TTS.TTSRequestGpsPositionFileNameAndPath
	TTSNextServer = document.Global.Software.TTS.TTSNextServer
	TTSNextServerFileNameAndPath = document.Global.Software.TTS.TTSNextServerFileNameAndPath
	TTSPanicSimulation = document.Global.Software.TTS.TTSPanicSimulation
	TTSPanicSimulationFileNameAndPath = document.Global.Software.TTS.TTSPanicSimulationFileNameAndPath
	TTSPrintXmlConfig = document.Global.Software.TTS.TTSPrintXmlConfig
	TTSPrintXmlConfigFileNameAndPath = document.Global.Software.TTS.TTSPrintXmlConfigFileNameAndPath
	TTSSendEmail = document.Global.Software.TTS.TTSSendEmail
	TTSSendEmailFileNameAndPath = document.Global.Software.TTS.TTSSendEmailFileNameAndPath
	TTSDisplayMenu = document.Global.Software.TTS.TTSDisplayMenu
	TTSDisplayMenuFileNameAndPath = document.Global.Software.TTS.TTSDisplayMenuFileNameAndPath
	TTSQuitTalkkonnect = document.Global.Software.TTS.TTSQuitTalkkonnect
	TTSQuitTalkkonnectFileNameAndPath = document.Global.Software.TTS.TTSQuitTalkkonnectFileNameAndPath
	TTSTalkkonnectLoaded = document.Global.Software.TTS.TTSTalkkonnectLoaded
	TTSTalkkonnectLoadedFileNameAndPath = document.Global.Software.TTS.TTSTalkkonnectLoadedFileNameAndPath
	TTSPingServers = document.Global.Software.TTS.TTSPingServers
	TTSPingServersFileNameAndPath = document.Global.Software.TTS.TTSPingServersFileNameAndPath

	EmailEnabled = document.Global.Software.Smtp.EmailEnabled
	EmailUsername = document.Global.Software.Smtp.EmailUsername
	EmailPassword = document.Global.Software.Smtp.EmailPassword
	EmailReceiver = document.Global.Software.Smtp.EmailReceiver
	EmailSubject = document.Global.Software.Smtp.EmailSubject
	EmailMessage = document.Global.Software.Smtp.EmailMessage
	EmailGpsDateTime = document.Global.Software.Smtp.EmailGpsDateTime
	EmailGpsLatLong = document.Global.Software.Smtp.EmailGpsLatLong
	EmailGoogleMapsUrl = document.Global.Software.Smtp.EmailGoogleMapsUrl

	EventSoundEnabled = document.Global.Software.Sounds.Event.EEnabled
	EventSoundFilenameAndPath = document.Global.Software.Sounds.Event.EFileNameAndPath

	AlertSoundEnabled = document.Global.Software.Sounds.Alert.AEnabled
	AlertSoundFilenameAndPath = document.Global.Software.Sounds.Alert.AFileNameAndPath
	AlertSoundVolume = document.Global.Software.Sounds.Alert.AVolume

	RogerBeepSoundEnabled = document.Global.Software.Sounds.RogerBeep.REnabled
	RogerBeepSoundFilenameAndPath = document.Global.Software.Sounds.RogerBeep.RFileNameAndPath
	RogerBeepSoundVolume = document.Global.Software.Sounds.RogerBeep.RBeepVolume

	ChimesSoundEnabled = document.Global.Software.Sounds.Chimes.CEnabled
	ChimesSoundFilenameAndPath = document.Global.Software.Sounds.Chimes.CFileNameAndPath
	ChimesSoundVolume = document.Global.Software.Sounds.Chimes.CVolume

	TxTimeOutEnabled = document.Global.Software.TxTimeOut.TxTimeOutEnabled
	TxTimeOutSecs = document.Global.Software.TxTimeOut.TxTimeOutSecs

	APIEnabled = document.Global.Software.API.APIEnabled
	APIListenPort = document.Global.Software.API.APIListenPort
	APIDisplayMenu = document.Global.Software.API.APIDisplayMenu
	APIChannelUp = document.Global.Software.API.APIChannelUp
	APIChannelDown = document.Global.Software.API.APIChannelDown
	APIMute = document.Global.Software.API.APIMute
	APICurrentVolumeLevel = document.Global.Software.API.APICurrentVolumeLevel
	APIDigitalVolumeUp = document.Global.Software.API.APIDigitalVolumeUp
	APIDigitalVolumeDown = document.Global.Software.API.APIDigitalVolumeDown
	APIListServerChannels = document.Global.Software.API.APIListServerChannels
	APIStartTransmitting = document.Global.Software.API.APIStartTransmitting
	APIStopTransmitting = document.Global.Software.API.APIStopTransmitting
	APIListOnlineUsers = document.Global.Software.API.APIListOnlineUsers
	APIPlayChimes = document.Global.Software.API.APIPlayChimes
	APIRequestGpsPosition = document.Global.Software.API.APIRequestGpsPosition
	APINextServer = document.Global.Software.API.APINextServer
	APIPanicSimulation = document.Global.Software.API.APIPanicSimulation
	APIPrintXmlConfig = document.Global.Software.API.APIPrintXmlConfig
	APIEmailEnabled = document.Global.Software.API.APIEmailEnabled
	APIPingServersEnabled = document.Global.Software.API.APIPingServersEnabled

	PrintAccount = document.Global.Software.PrintVariables.PrintAccount
	PrintLogging = document.Global.Software.PrintVariables.PrintLogging
	PrintProvisioning = document.Global.Software.PrintVariables.PrintProvisioning
	PrintBeacon = document.Global.Software.PrintVariables.PrintBeacon
	PrintTTS = document.Global.Software.PrintVariables.PrintTTS
	PrintSMTP = document.Global.Software.PrintVariables.PrintSMTP
	PrintSounds = document.Global.Software.PrintVariables.PrintSounds
	PrintTxTimeout = document.Global.Software.PrintVariables.PrintTxTimeout
	PrintHTTPAPI = document.Global.Software.PrintVariables.PrintHTTPAPI
	PrintTargetboard = document.Global.Software.PrintVariables.PrintTargetboard
	PrintLeds = document.Global.Software.PrintVariables.PrintLeds
	PrintHeartbeat = document.Global.Software.PrintVariables.PrintHeartbeat
	PrintButtons = document.Global.Software.PrintVariables.PrintButtons
	PrintComment = document.Global.Software.PrintVariables.PrintComment
	PrintLcd = document.Global.Software.PrintVariables.PrintLcd
	PrintGps = document.Global.Software.PrintVariables.PrintGps
	PrintPanic = document.Global.Software.PrintVariables.PrintPanic

	TargetBoard = document.Global.Hardware.TargetBoard

	VoiceActivityLEDPin = document.Global.Hardware.Lights.VoiceActivityLedPin
	ParticipantsLEDPin = document.Global.Hardware.Lights.ParticipantsLedPin
	TransmitLEDPin = document.Global.Hardware.Lights.TransmitLedPin
	OnlineLEDPin = document.Global.Hardware.Lights.OnlineLedPin

	HeartBeatEnabled = document.Global.Hardware.HeartBeat.HeartBeatEnabled
	HeartBeatLEDPin = document.Global.Hardware.HeartBeat.HeartBeatLEDPin
	PeriodmSecs = document.Global.Hardware.HeartBeat.PeriodmSecs
	LEDOnmSecs = document.Global.Hardware.HeartBeat.LEDOnmSecs
	LEDOffmSecs = document.Global.Hardware.HeartBeat.LEDOffmSecs

	TxButtonPin = document.Global.Hardware.Buttons.TxButtonPin
	UpButtonPin = document.Global.Hardware.Buttons.UpButtonPin
	DownButtonPin = document.Global.Hardware.Buttons.DownButtonPin
	PanicButtonPin = document.Global.Hardware.Buttons.PanicButtonPin

	CommentButtonPin = document.Global.Hardware.Comment.CommentButtonPin
	CommentMessageOff = document.Global.Hardware.Comment.CommentMessageOff
	CommentMessageOn = document.Global.Hardware.Comment.CommentMessageOn

	LCDInterfaceType = document.Global.Hardware.Screen.LCDInterfaceType

	LCDI2CAddress = document.Global.Hardware.Screen.LCDI2CAddress

	LCDBackLightTimerEnabled = document.Global.Hardware.Screen.LCDBackLightTimerEnabled
	LCDBackLightTimeoutSecs = document.Global.Hardware.Screen.LCDBackLightTimeoutSecs
	BackLightLEDPin = document.Global.Hardware.Screen.BackLightLEDPin
	RSPin = document.Global.Hardware.Screen.RsPin
	EPin = document.Global.Hardware.Screen.EsPin
	D4Pin = document.Global.Hardware.Screen.D4Pin
	D5Pin = document.Global.Hardware.Screen.D5Pin
	D6Pin = document.Global.Hardware.Screen.D6Pin
	D7Pin = document.Global.Hardware.Screen.D7Pin

	GpsEnabled = document.Global.Hardware.GPS.GpsEnabled
	Port = document.Global.Hardware.GPS.Port
	Baud = document.Global.Hardware.GPS.Baud
	TxData = document.Global.Hardware.GPS.TxData
	Even = document.Global.Hardware.GPS.Even
	Odd = document.Global.Hardware.GPS.Odd
	Rs485 = document.Global.Hardware.GPS.Rs485
	Rs485HighDuringSend = document.Global.Hardware.GPS.Rs485highduringsend
	Rs485HighAfterSend = document.Global.Hardware.GPS.Rs485highaftersend
	StopBits = document.Global.Hardware.GPS.StopBits
	DataBits = document.Global.Hardware.GPS.DataBits
	CharTimeOut = document.Global.Hardware.GPS.CharTimeOut
	MinRead = document.Global.Hardware.GPS.MinRead
	Rx = document.Global.Hardware.GPS.Rx

	PEnabled = document.Global.Hardware.PanicFunction.PEnabled
	PFileNameAndPath = document.Global.Hardware.PanicFunction.PFileNameAndPath
	PMessage = document.Global.Hardware.PanicFunction.PMessage
	PVolume = document.Global.Hardware.PanicFunction.PVolume
	PSendIdent = document.Global.Hardware.PanicFunction.PSendIdent
	PSendGpsLocation = document.Global.Hardware.PanicFunction.PSendGpsLocation
	PTxLockEnabled = document.Global.Hardware.PanicFunction.PTxLockEnabled
	PTxlockTimeOutSecs = document.Global.Hardware.PanicFunction.PTxlockTimeOutSecs

	log.Println("Successfully loaded configuration file into memory")
	return nil
}

func printxmlconfig() {

	if PrintAccount {
		log.Println("info: ---------- Account Information -------- ")
		log.Println("info: Default     " + fmt.Sprintf("%t", Default))
		log.Println("info: Server      " + Server[0])
		log.Println("info: Username    " + Username[0])
		log.Println("info: Password    " + Password[0])
		log.Println("info: Insecure    " + fmt.Sprintf("%t", Insecure[0]))
		log.Println("info: Certificate " + Certificate[0])
		log.Println("info: Channel     " + Channel[0])
		log.Println("info: Ident       " + Ident[0])
	} else {
		log.Println("info: ---------- Account Information -------- SKIPPED ")
	}

	if PrintLogging {
		log.Println("info: -------- Logging & Daemonizing -------- ")
		log.Println("info: Output Device " + OutputDevice)
		log.Println("info: Log File      " + LogFileNameAndPath)
		log.Println("info: Logging       " + Logging)
		log.Println("info: Daemonize     " + fmt.Sprintf("%t", Daemonize))
	} else {
		log.Println("info: --------   Logging & Daemonizing -------- SKIPPED ")
	}

	if PrintProvisioning {
		log.Println("info: --------   AutoProvisioning   --------- ")
		log.Println("info: AutoProvisioning Enabled    " + fmt.Sprintf("%t", APEnabled))
		log.Println("info: Talkkonned ID (tkid)        " + TkId)
		log.Println("info: AutoProvisioning Server URL " + Url)
		log.Println("info: Config Local Path           " + SaveFilePath)
		log.Println("info: Config Local FileName       " + SaveFileName)
	} else {
		log.Println("info: --------   AutoProvisioning   --------- SKIPPED ")
	}

	if PrintBeacon {
		log.Println("info: --------  Beacon   --------- ")
		log.Println("info: Beacon Enabled         " + fmt.Sprintf("%t", BeaconEnabled))
		log.Println("info: Beacon Time (Secs)     " + fmt.Sprintf("%v", BeaconTimerSecs))
		log.Println("info: Beacon FileName & Path " + BeaconFileNameAndPath)
		log.Println("info: Beacon Playback Volume " + fmt.Sprintf("%v", BVolume))
	} else {
		log.Println("info: --------   Beacon   --------- SKIPPED ")
	}

	if PrintTTS {
		log.Println("info: -------- TTS  -------- ")
		log.Println("info: TTS Global Enabled     ", fmt.Sprintf("%t", TTSEnabled))
		log.Println("info: TTS Volume Level (%)   ", fmt.Sprintf("%d", TTSVolumeLevel))
		log.Println("info: TTS Participants       ", fmt.Sprintf("%t", TTSParticipants))
		log.Println("info: TTS ChannelUp          ", fmt.Sprintf("%t", TTSChannelUp))
		log.Println("info: TTS ChannelUpFileNameAndPath ", TTSChannelUpFileNameAndPath)
		log.Println("info: TTS ChannelDown        ", fmt.Sprintf("%t", TTSChannelDown))
		log.Println("info: TTS ChannelDownFileNameAndPath  ", TTSChannelDownFileNameAndPath)
		log.Println("info: TTS MuteUnMuteSpeaker  ", fmt.Sprintf("%t", TTSMuteUnMuteSpeaker))
		log.Println("info: TTS MuteUnMuteSpeakerFileNameAndPath ", TTSMuteUnMuteSpeakerFileNameAndPath)
		log.Println("info: TTS CurrentVolumeLevel ", fmt.Sprintf("%t", TTSCurrentVolumeLevel))
		log.Println("info: TTS CurrentVolumeLevelFileNameAndPath ", TTSCurrentVolumeLevelFileNameAndPath)
		log.Println("info: TTS DigitalVolumeUp    ", fmt.Sprintf("%t", TTSDigitalVolumeUp))
		log.Println("info: TTS DigitalVolumeUpFileNameAndPath ", TTSDigitalVolumeUpFileNameAndPath)
		log.Println("info: TTS DigitalVolumeDown  ", fmt.Sprintf("%t", TTSDigitalVolumeDown))
		log.Println("info: TTS DigitalVolumeDownFileNameAndPath ", TTSDigitalVolumeDownFileNameAndPath)
		log.Println("info: TTS ListServerChannels ", fmt.Sprintf("%t", TTSListServerChannels))
		log.Println("info: TTS ListServerChannelsFileNameAndPath  ", TTSListServerChannelsFileNameAndPath)
		log.Println("info: TTS StartTransmitting  ", fmt.Sprintf("%t", TTSStartTransmitting))
		log.Println("info: TTS StartTransmittingFileNameAndPath ", TTSStartTransmittingFileNameAndPath)
		log.Println("info: TTS StopTransmitting   ", fmt.Sprintf("%t", TTSStopTransmitting))
		log.Println("info: TTS StopTransmittingFileNameAndPath ", TTSStopTransmittingFileNameAndPath)
		log.Println("info: TTS ListOnlineUsers    ", fmt.Sprintf("%t", TTSListOnlineUsers))
		log.Println("info: TTS ListOnlineUsersFileNameAndPath ", TTSListOnlineUsersFileNameAndPath)
		log.Println("info: TTS PlayChimes         ", fmt.Sprintf("%t", TTSPlayChimes))
		log.Println("info: TTS PlayChimesFileNameAndPath ", TTSPlayChimesFileNameAndPath)
		log.Println("info: TTS RequestGpsPosition ", fmt.Sprintf("%t", TTSRequestGpsPosition))
		log.Println("info: TTS RequestGpsPositionFileNameAndPath ", TTSRequestGpsPositionFileNameAndPath)
		log.Println("info: TTS NextServer         ", fmt.Sprintf("%t", TTSNextServer))
		log.Println("info: TTS NextServerFileNameAndPath         ", TTSNextServerFileNameAndPath)
		log.Println("info: TTS PanicSimulation    ", fmt.Sprintf("%t", TTSPanicSimulation))
		log.Println("info: TTS PanicSimulationFileNameAndPath ", TTSPanicSimulationFileNameAndPath)
		log.Println("info: TTS PrintXmlConfig     ", fmt.Sprintf("%t", TTSPrintXmlConfig))
		log.Println("info: TTS PrintXmlConfigFileNameAndPath ", TTSPrintXmlConfigFileNameAndPath)
		log.Println("info: TTS SendEmail          ", fmt.Sprintf("%t", TTSSendEmail))
		log.Println("info: TTS SendEmailFileNameAndPath ", TTSSendEmailFileNameAndPath)
		log.Println("info: TTS DisplayMenu        ", fmt.Sprintf("%t", TTSDisplayMenu))
		log.Println("info: TTS DisplayMenuFileNameAndPath ", TTSDisplayMenuFileNameAndPath)
		log.Println("info: TTS QuitTalkkonnect    ", fmt.Sprintf("%t", TTSQuitTalkkonnect))
		log.Println("info: TTS QuitTalkkonnectFileNameAndPath ", TTSQuitTalkkonnectFileNameAndPath)
		log.Println("info: TTS TalkkonnectLoaded  ", fmt.Sprintf("%t", TTSTalkkonnectLoaded))
		log.Println("info: TTS TalkkonnectLoadedFileNameAndPath ", TTSTalkkonnectLoadedFileNameAndPath)
		log.Println("info: TTS TalkkonnectLoaded  " + fmt.Sprintf("%t", TTSTalkkonnectLoaded))
		log.Println("info: TTS PingServersFileNameAndPath ", TTSPingServersFileNameAndPath)
		log.Println("info: TTS PingServers " + fmt.Sprintf("%t", TTSPingServers))
	} else {
		log.Println("info: --------   TTS  -------- SKIPPED ")
	}

	if PrintSMTP {
		log.Println("info: --------  Gmail SMTP Settings  -------- ")
		log.Println("info: Email Enabled   " + fmt.Sprintf("%t", EmailEnabled))
		log.Println("info: Username        " + EmailUsername)
		log.Println("info: Password        " + EmailPassword)
		log.Println("info: Receiver        " + EmailReceiver)
		log.Println("info: Subject         " + EmailSubject)
		log.Println("info: Message         " + EmailMessage)
		log.Println("info: GPS Date/Time   " + fmt.Sprintf("%t", EmailGpsDateTime))
		log.Println("info: GPS Lat/Long    " + fmt.Sprintf("%t", EmailGpsLatLong))
		log.Println("info: Google Maps Url " + fmt.Sprintf("%t", EmailGoogleMapsUrl))
	} else {
		log.Println("info: --------   Gmail SMTP Settings  -------- SKIPPED ")
	}

	if PrintSounds {
		log.Println("info: ------------- Sounds  ------------------ ")
		log.Println("info: Event Sound Enabled  " + fmt.Sprintf("%t", EventSoundEnabled))
		log.Println("info: Event Sound Filename " + EventSoundFilenameAndPath)
		log.Println("info: Alert Sound Enabled  " + fmt.Sprintf("%t", AlertSoundEnabled))
		log.Println("info: Alert Sound Filename " + AlertSoundFilenameAndPath)
		log.Println("info: Alert Sound Volume   " + fmt.Sprintf("%v", AlertSoundVolume))
		log.Println("info: Roger Beep Enabled " + fmt.Sprintf("%t", RogerBeepSoundEnabled))
		log.Println("info: Roger Beep File    " + RogerBeepSoundFilenameAndPath)
		log.Println("info: Roger Beep Volume  " + fmt.Sprintf("%v", RogerBeepSoundVolume))
		log.Println("info: Chimes Enabled     " + fmt.Sprintf("%t", ChimesSoundEnabled))
		log.Println("info: Chimes File        " + ChimesSoundFilenameAndPath)
		log.Println("info: Chimes Volume      " + fmt.Sprintf("%v", ChimesSoundVolume))
	} else {
		log.Println("info: ------------ Sounds  ------------------ SKIPPED ")
	}

	if PrintTxTimeout {
		log.Println("info: ------------ TX Timeout ------------------ ")
		log.Println("info: Tx Timeout Enabled  " + fmt.Sprintf("%t", TxTimeOutEnabled))
		log.Println("info: Tx Timeout Secs     " + fmt.Sprintf("%v", TxTimeOutSecs))
	} else {
		log.Println("info: ------------ TX Timeout ------------------ SKIPPED ")
	}

	if PrintHTTPAPI {
		log.Println("info: ------------ HTTP API  ----------------- ")
		log.Println("info: API Enabled        " + fmt.Sprintf("%t", APIEnabled))
		log.Println("info: API Listen Port    " + APIListenPort)
		log.Println("info: DisplayMenu        " + fmt.Sprintf("%t", APIDisplayMenu))
		log.Println("info: ChannelUp          " + fmt.Sprintf("%t", APIChannelUp))
		log.Println("info: ChannelDown        " + fmt.Sprintf("%t", APIChannelDown))
		log.Println("info: Mute               " + fmt.Sprintf("%t", APIMute))
		log.Println("info: CurentVolumeLevel  " + fmt.Sprintf("%t", APICurrentVolumeLevel))
		log.Println("info: DigitalVolumeUp    " + fmt.Sprintf("%t", APIDigitalVolumeUp))
		log.Println("info: DigitalVolumeDown  " + fmt.Sprintf("%t", APIDigitalVolumeDown))
		log.Println("info: ListServerChannels " + fmt.Sprintf("%t", APIListServerChannels))
		log.Println("info: StartTransmitting  " + fmt.Sprintf("%t", APIStartTransmitting))
		log.Println("info: StopTransmitting   " + fmt.Sprintf("%t", APIStopTransmitting))
		log.Println("info: GPS Position       " + fmt.Sprintf("%t", APIRequestGpsPosition))
		log.Println("info: PlayChimes         " + fmt.Sprintf("%t", APIPlayChimes))
		log.Println("info: RequestGpsPosition " + fmt.Sprintf("%t", APIRequestGpsPosition))
		log.Println("info: NextServer         " + fmt.Sprintf("%t", APINextServer))
		log.Println("info: PanicSimulation    " + fmt.Sprintf("%t", APIPanicSimulation))
		log.Println("info: PrintXmlConfig     " + fmt.Sprintf("%t", APIPrintXmlConfig))
		log.Println("info: EmailEnabled       " + fmt.Sprintf("%t", APIEmailEnabled))
		log.Println("info: PingServersEnabled " + fmt.Sprintf("%t", APIPingServersEnabled))
	} else {
		log.Println("info: ------------ HTTP API  ----------------- SKIPPED ")
	}

	if PrintTargetboard {
		log.Println("info: ------------ Target Board --------------- ")
		log.Println("info: Target Board " + fmt.Sprintf("%v", TargetBoard))
	} else {
		log.Println("info: ------------ Target Board --------------- SKIPPED ")
	}

	if PrintLeds {
		log.Println("info: ------------ LEDS  ---------------------- ")
		log.Println("info: Voice Activity Led Pin " + fmt.Sprintf("%v", VoiceActivityLEDPin))
		log.Println("info: Participants Led Pin   " + fmt.Sprintf("%v", ParticipantsLEDPin))
		log.Println("info: Transmit Led Pin       " + fmt.Sprintf("%v", TransmitLEDPin))
		log.Println("info: Online Led Pin         " + fmt.Sprintf("%v", OnlineLEDPin))
	} else {
		log.Println("info: ------------ LEDS  ---------------------- SKIPPED ")
	}

	if PrintHeartbeat {
		log.Println("info: ---------- HEARTBEAT -------------------- ")
		log.Println("info: HeartBeat Enabled " + fmt.Sprintf("%v", HeartBeatEnabled))
		log.Println("info: HeartBeat LED Pin " + fmt.Sprintf("%v", HeartBeatLEDPin))
		log.Println("info: Period  mSecs     " + fmt.Sprintf("%v", PeriodmSecs))
		log.Println("info: Led On  mSecs     " + fmt.Sprintf("%v", LEDOnmSecs))
		log.Println("info: Led Off mSecs     " + fmt.Sprintf("%v", LEDOffmSecs))
	}

	if PrintButtons {
		log.Println("info: ------------ Buttons  ------------------- ")
		log.Println("info: Tx Button Pin           " + fmt.Sprintf("%v", TxButtonPin))
		log.Println("info: Channel Up Button Pin   " + fmt.Sprintf("%v", UpButtonPin))
		log.Println("info: Channel Down Button Pin " + fmt.Sprintf("%v", DownButtonPin))
		log.Println("info: Panic Button Pin        " + fmt.Sprintf("%v", PanicButtonPin))
	} else {
		log.Println("info: ------------ Buttons  ------------------- SKIPPED ")
	}

	if PrintComment {
		log.Println("info: ------------ Comment  ------------------- ")
		log.Println("info: Comment Button Pin         " + fmt.Sprintf("%v", CommentButtonPin))
		log.Println("info: Comment Message State 1    " + fmt.Sprintf("%v", CommentMessageOff))
		log.Println("info: Comment Message State 2    " + fmt.Sprintf("%v", CommentMessageOn))
	} else {
		log.Println("info: ------------ Comment  ------------------- SKIPPED ")
	}

	if PrintLcd {
		log.Println("info: ------------ LCD  ----------------------- ")
		log.Println("info: Lcd Interface Type       " + fmt.Sprintf("%v", LCDInterfaceType))
		log.Println("info: Lcd I2C Address          " + fmt.Sprintf("%x", LCDI2CAddress))
		log.Println("info: Back Light Timer Enabled " + fmt.Sprintf("%t", LCDBackLightTimerEnabled))
		log.Println("info: Back Light Timer Timeout " + fmt.Sprintf("%v", LCDBackLightTimeoutSecs))
		log.Println("info: Back Light Pin " + fmt.Sprintf("%v", BackLightLEDPin))
		log.Println("info: RS Pin " + fmt.Sprintf("%v", RSPin))
		log.Println("info: E  Pin " + fmt.Sprintf("%v", EPin))
		log.Println("info: D4 Pin " + fmt.Sprintf("%v", D4Pin))
		log.Println("info: D5 Pin " + fmt.Sprintf("%v", D5Pin))
		log.Println("info: D6 Pin " + fmt.Sprintf("%v", D6Pin))
		log.Println("info: D7 Pin " + fmt.Sprintf("%v", D7Pin))
	} else {
		log.Println("info: ------------ LCD  ----------------------- SKIPPED ")
	}

	if PrintGps {
		log.Println("info: ------------ GPS  ------------------------ ")
		log.Println("info: GPS Enabled            " + fmt.Sprintf("%t", GpsEnabled))
		log.Println("info: Port                   " + fmt.Sprintf("%s", Port))
		log.Println("info: Baud                   " + fmt.Sprintf("%v", Baud))
		log.Println("info: TxData                 " + fmt.Sprintf("%s", TxData))
		log.Println("info: Even                   " + fmt.Sprintf("%v", Even))
		log.Println("info: Odd                    " + fmt.Sprintf("%v", Odd))
		log.Println("info: RS485                  " + fmt.Sprintf("%v", Rs485))
		log.Println("info: RS485 High During Send " + fmt.Sprintf("%v", Rs485HighDuringSend))
		log.Println("info: RS485 High After Send  " + fmt.Sprintf("%v", Rs485HighAfterSend))
		log.Println("info: Stop Bits              " + fmt.Sprintf("%v", StopBits))
		log.Println("info: Data Bits              " + fmt.Sprintf("%v", DataBits))
		log.Println("info: Char Time Out          " + fmt.Sprintf("%v", CharTimeOut))
		log.Println("info: Min Read               " + fmt.Sprintf("%v", MinRead))
		log.Println("info: Rx                     " + fmt.Sprintf("%t", Rx))
	} else {
		log.Println("info: ------------ GPS  ------------------------ SKIPPED ")
	}

	if PrintPanic {
		log.Println("info: ------------ PANIC Function -------------- ")
		log.Println("info: Panic Function Enable          " + fmt.Sprintf("%t", PEnabled))
		log.Println("info: Panic Sound Filename and Path  " + fmt.Sprintf("%s", PFileNameAndPath))
		log.Println("info: Panic Message                  " + fmt.Sprintf("%s", PMessage))
		log.Println("info: Panic Message Send Recursively " + fmt.Sprintf("%t", PRecursive))
		log.Println("info: Panic Volume                   " + fmt.Sprintf("%v", PVolume))
		log.Println("info: Panic Send Ident               " + fmt.Sprintf("%t", PSendIdent))
		log.Println("info: Panic Send GPS Location        " + fmt.Sprintf("%t", PSendGpsLocation))
		log.Println("info: Panic TX Lock Enabled          " + fmt.Sprintf("%t", PTxLockEnabled))
		log.Println("info: Panic TX Lock Timeout Secs     " + fmt.Sprintf("%v", PTxlockTimeOutSecs))
	} else {
		log.Println("info: ------------ PANIC Function -------------- SKIPPED ")
	}

}
