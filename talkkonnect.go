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
 * talkkonnect.go -> function in talkkonnect for printing banners to screen
 */

package talkkonnect

import (
	"fmt"
	"github.com/talkkonnect/gumble/gumble"
	"github.com/talkkonnect/volume-go"
	"log"
	"net"
	"time"
)

func talkkonnectBanner() {
	log.Println("info: ┌────────────────────────────────────────────────────────────────┐")
	log.Println("info: │  _        _ _    _                               _             │")
	log.Println("info: │ | |_ __ _| | | _| | _____  _ __  _ __   ___  ___| |_           │")
	log.Println("info: │ | __/ _` | | |/ / |/ / _ \\| '_ \\| '_ \\ / _ \\/ __|  __|         │")
	log.Println("info: │ | || (_| | |   <|   < (_) | | | | | | |  __/ (__| |_           │")
	log.Println("info: │  \\__\\__,_|_|_|\\_\\_|\\_\\___/|_| |_|_| |_|\\___|\\_ _|\\__|          │")
	log.Println("info: ├────────────────────────────────────────────────────────────────┤")
	log.Println("info: │A Flexible Headless Mumble Transceiver/Gateway for RPi/PC/VM    │")
	log.Println("info: ├────────────────────────────────────────────────────────────────┤")
	log.Println("info: │Created By : Suvir Kumar  <suvir@talkkonnect.com>               │")
	log.Println("info: ├────────────────────────────────────────────────────────────────┤")
	log.Println("info: │Version 1.41 Released February 10 2019                          │")
	log.Println("info: │Additional Modifications Released under MPL 2.0 License         │")
	log.Println("info: ├────────────────────────────────────────────────────────────────┤")
	log.Println("info: │visit us at www.talkkonnect.com and github.com/talkkonnect      │")
	log.Println("info: └────────────────────────────────────────────────────────────────┘")
	log.Println("info: Press the <Del> key for Menu Options or <Ctrl-c> to Quit talkkonnect")
}

func talkkonnectAcknowledgements() {
	log.Println("info: ┌───────────────────────────────────────────────────────────────────────────────────────────┐")
	log.Println("info: │Acknowledgements & Inspriation from the talkkonnect team of developers and maintainers     │")
	log.Println("info: ├───────────────────────────────────────────────────────────────────────────────────────────┤")
	log.Println("info: │talkkonnect is based on the works of many people and many open source projects             │")
	log.Println("info: │                                                                                           │")
	log.Println("info: │Thanks to :-                                                                               │")
	log.Println("info: │                                                                                           │")
	log.Println("info: │Organizations :-                                                                           │")
	log.Println("info: │The Mumble Development team, Raspberry Pi Foundation, Developers and Maintainers of Debian │")
	log.Println("info: │The Creators and Maintainers of Golang and all the libraries available on github.com       │")
	log.Println("info: │                                                                                           │")
	log.Println("info: │Individuals :-                                                                             │")
	log.Println("info: │Daniel Chote Creator of talkiepi and Tim Cooper Creator of Barnard and gumble library      │")
	log.Println("info: │Tayeb Meftah and other people who wish to remain anonymous for their feedback and testing  │")
	log.Println("info: ├───────────────────────────────────────────────────────────────────────────────────────────┤")
	log.Println("info: │visit us at www.talkkonnect.com and github.com/talkkonnect <suvir@talkkonnect.com>         │")
	log.Println("info: └───────────────────────────────────────────────────────────────────────────────────────────┘")
}

func (b *Talkkonnect) talkkonnectMenu() {
	log.Println("info: ┌────────────────────────────────────────────────────────────────┐")
	log.Println("info: │                 _                                              │")
	log.Println("info: │ _ __ ___   __ _(_)_ __    _ __ ___   ___ _ __  _   _           │")
	log.Println("info: │| '_ ` _ \\ / _` | | '_ \\  | '_ ` _ \\ / _ \\ '_ \\| | | |          │")
	log.Println("info: │| | | | | | (_| | | | | | | | | | | |  __/ | | | |_| |          │")
	log.Println("info: │|_| |_| |_|\\__,_|_|_| |_| |_| |_| |_|\\___|_| |_|\\__,_|          │")
	log.Println("info: ├─────────────────────────────┬──────────────────────────────────┤")
	log.Println("info: │ <Del> to Display this Menu  | Ctrl-C to Quit talkkonnect       │")
	log.Println("info: ├─────────────────────────────┼──────────────────────────────────┤")
	log.Println("info: │ <F1>  Channel Up (+)        │ <F2>  Channel Down (-)           │")
	log.Println("info: │ <F3>  Mute/Unmute Speaker   │ <F4>  Current Volume Level       │")
	log.Println("info: │ <F5>  Digital Volume Up (+) │ <F6>  Digital Volume Down (-)    │")
	log.Println("info: │ <F7>  List Server Channels  │ <F8>  Start Transmitting         │")
	log.Println("info: │ <F9>  Stop Transmitting     │ <F10> List Online Users          │")
	log.Println("info: │ <F11> Playback/Stop Chimes  │ <F12> For GPS Position           │")
	log.Println("info: ├─────────────────────────────┼──────────────────────────────────┤")
	log.Println("info: │<Ctrl-E> Send Email          │<Ctrl-L> Clear Screen             │")
	log.Println("info: │<Ctrl-M> Ping Servers        │<Ctrl-N> Connect Next Server      │")
	log.Println("info: │<Ctrl-P> Panic Simulation    │<Ctrl-S> Scan Channels            │ ")
	log.Println("info: │<Ctrl-X> Dump XML Config     │                                  │ ")
	log.Println("info: ├─────────────────────────────┴──────────────────────────────────┤")
	log.Println("info: │   visit us at www.talkkonnect.com and github.com/talkkonnect   │")
	log.Println("info: └────────────────────────────────────────────────────────────────┘")

	log.Println("info: IP Address & Session Information")
	b.pingconnectedserver()
	localAddresses()

	origMuted, _ := volume.GetMuted(OutputDevice)
	if origMuted {
		log.Println("info: Speaker Currently Muted")
	} else {
		log.Println("info: Speaker Currently Not Muted")
	}

}

func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("error: localAddresses %v\n", err.Error()))
		return
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			log.Print(fmt.Errorf("error: localAddresses %v\n", err.Error()))
			continue
		}

		for _, a := range addrs {
			if i.Name != "lo" {
				log.Printf("info: %v %v\n", i.Name, a)
			}
		}
	}
}

func (b *Talkkonnect) pingconnectedserver() {

	resp, err := gumble.Ping(b.Address, time.Second*1, time.Second*5)

	if err != nil {
		log.Println(fmt.Sprintf("warn: Ping Error ", err))
		return
	}

	major, minor, patch := resp.Version.SemanticVersion()

	log.Println("info: Server Address:         ", resp.Address)
	log.Println("info: Server Ping:            ", resp.Ping)
	log.Println("info: Server Version:         ", major, ".", minor, ".", patch)
	log.Println("info: Server Users:           ", resp.ConnectedUsers, "/", resp.MaximumUsers)
	log.Println("info: Server Maximum Bitrate: ", resp.MaximumBitrate)
}
