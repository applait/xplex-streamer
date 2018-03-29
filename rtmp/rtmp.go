package rtmp

import (
	"fmt"
	"log"
	"sync"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format"
	"github.com/nareix/joy4/format/rtmp"
)

func init() {
	format.RegisterAll()
}

func handlePublish(conn *rtmp.Conn) {
	log.Printf("Publish from %s", conn.URL)
	defer conn.Close()

	streams, err := conn.Streams()
	if err != nil {
		log.Printf("Error getting streams from connection. Reason: %s\n", err)
		return
	}

	for _, s := range streams {
		fmt.Printf("Stream type %s. Audio: %t Video: %t\n", s.Type(), s.Type().IsAudio(), s.Type().IsVideo())
	}
	urls := []string{
		"rtmp://x.rtmp.youtube.com/live2/58d5-9hcw-023c-1214",
		"rtmp://live.twitch.tv/app/live_160285039_A9wlHmdH3yb3r6x9z45IX8uYzss50G",
	}

	// l.Lock()
	q := pubsub.NewQueue()
	defer q.Close()
	if err = q.WriteHeader(streams); err != nil {
		log.Printf("Error writing header to pubsub queue. Reason: %s\n", err)
	}
	log.Printf("Written header to queue")

	go func() {
		err := avutil.CopyPackets(q, conn)
		if err != nil {
			log.Printf("Error copying packets to queue. Reason: %s\n", err)
		}
	}()

	wait := &sync.WaitGroup{}

	for _, u := range urls {
		wait.Add(1)
		go func(u string) {
			defer wait.Done()
			origin := q.Latest()
			log.Printf("Dialing %s\n", u)
			dst, err := rtmp.Dial(u)
			defer dst.Close()
			if err != nil {
				log.Printf("Error connecting to destination %s. Reason: %s\n", u, err)
				return
			}
			if err = dst.WriteHeader(streams); err != nil {
				log.Printf("Error writing header to %s. Reason: %s\n", u, err)
				return
			}
			err = avutil.CopyPackets(dst, origin)
			if err != nil {
				log.Printf("Error pushing stream. Reason: %s\n", err)
			}
		}(u)
	}
	wait.Wait()

}

// Start bootstraps and launches RTMP server
func Start(port int) {
	server := &rtmp.Server{Addr: fmt.Sprintf(":%d", port)}
	server.HandlePublish = handlePublish
	log.Printf("Launching RTMP server on port %d", port)
	server.ListenAndServe()
}
