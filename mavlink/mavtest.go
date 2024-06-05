package mavlink

import (
	"fmt"
	"github.com/bluenviron/gomavlib/v3"
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/ardupilotmega"
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
	"log"
	"math/rand"
)

func MavTest() {
	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointTCPClient{"-"},
		},
		Dialect:     ardupilotmega.Dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to communicate with the target
		OutSystemID: 125,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	// print incoming messages
	for evt := range node.Events() {
		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			frm.Message()
			log.Printf("received: id=%d, %+v\n", frm.Message().GetID(), frm.Message())

			switch frm.Message().GetID() {
			case 27:
				if msg, ok := frm.Message().(*ardupilotmega.MessageRawImu); ok {
					fmt.Println("+", msg)
					node.WriteMessageTo(frm.Channel, &ardupilotmega.MessageHilGps{
						//TimeUsec: msg.TimeUsec + 1000,
						FixType: uint8(common.GPS_FIX_TYPE_RTK_FIXED),
						Lat:     100300300,
						Lon:     100300300,
						//SatellitesVisible: 30,
						Id: uint8(rand.Int()),
						//Yaw:               0,
					})
				}
			}
		}
	}
}
