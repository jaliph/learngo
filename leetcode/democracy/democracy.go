package democracy

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sort"
	"strings"
	"time"
)

type NodeData struct {
	Id       string
	Source   string
	Weight   int
	State    string
	Last     int
	Voters   []NodeData
	Channels []string
}

type Message struct {
	Content string `json:"content"`
	Data    string `json:"data"`
}

type Chunk struct {
	Chunk   string `json:"chunk"`
	Id      string `json:"id"`
	Total   int    `json:"total"`
	Counter int    `json:"counter"`
}

type NodeConfig struct {
	Interval      int
	Timeout       int
	Source        string
	Peers         []string
	Weight        int
	Id            string
	Channels      []string
	MaxPacketSize int
}

type Democracy struct {
	Nodes      map[string]NodeData
	ChunkData  map[string][]Chunk
	Config     *NodeConfig
	Id         string
	Weight     int
	State      string
	Connection *net.UDPConn
}

func NewDemocracy(nc *NodeConfig) (*Democracy, chan bool) {
	d := &Democracy{
		Config:    nc,
		Nodes:     map[string]NodeData{},
		ChunkData: map[string][]Chunk{},
		Id:        nc.Id,
		Weight:    nc.Weight,
		State:     "citizen",
	}

	connection, err := net.ListenUDP("udp", &net.UDPAddr{Port: 0})
	if err != nil {
		fmt.Println("Error creating Source connection:", err)
		os.Exit(1)
	}
	d.Connection = connection

	ch := d.Start()
	d.Hello()
	return d, ch
}

func (d *Democracy) Hello() {
	go func() {
		for {
			d.SendData("hello", NodeData{}, "")
			// d.Check()
			time.Sleep(time.Duration(d.Config.Interval) * time.Millisecond)
		}
	}()
}

func (d *Democracy) ProcessEvent(msg Message) {
	if msg.Content == "hello" {
		var receivedNode NodeData
		err := json.Unmarshal([]byte(msg.Data), &receivedNode)
		fmt.Println("Received Node Data", receivedNode)
		if err != nil {
			fmt.Println("Error decoding Node Data from Hello:", err)
			return
		}
		d.AddToNodeList(receivedNode)
	}
}

func (d *Democracy) AddToNodeList(nc NodeData) {
	nc.Last = time.Now().Second()
	d.Nodes[nc.Id] = nc

	flag := true
	for _, peer := range d.Config.Peers {
		if peer == nc.Source {
			flag = false
		}
	}

	if flag && !IsUniversal(nc.Source) {
		d.Config.Peers = append(d.Config.Peers, nc.Source)
	}
	// fmt.Println(d)
}

func (d *Democracy) SendData(event string, data NodeData, id string) {
	if event == "hello" {
		data = NodeData{
			Id:     d.Id,
			Source: d.Config.Source,
			Weight: d.Weight,
			State:  d.State,
		}
	}
	dataEncoded, _ := json.Marshal(data)
	message := Message{
		Content: event,
		Data:    string(dataEncoded),
	}

	encoded, _ := json.Marshal(message)
	chunks := d.GenerateChunks(encoded)

	// Send all the Chunks
	for _, chunk := range chunks {
		jsonData, _ := json.Marshal(chunk)
		for _, peer := range d.Config.Peers {
			sendAddr, err := net.ResolveUDPAddr("udp", peer)
			if err != nil {
				fmt.Println("Error resolving UDP address:", err)
				continue
			}
			_, err = d.Connection.WriteToUDP(jsonData, sendAddr)
			if err != nil {
				fmt.Println("Error sending data:", err)
			}
			fmt.Println("Sent message:", chunk, " --> ", sendAddr)
		}

	}
}

func (d *Democracy) Start() chan bool {
	ch := make(chan bool)
	go func(ch chan<- bool) {
		source := d.Config.Source

		s := strings.Split(source, ":")
		if s[0] == "0.0.0.0" {
			source = fmt.Sprintf(":%s", s[1])
		}

		fmt.Println("Source Address is ", source)
		listenAddr, err := net.ResolveUDPAddr("udp", source)
		if err != nil {
			fmt.Println("Error resolving UDP address:", err)
			os.Exit(1)
		}

		conn, err := net.ListenUDP("udp", listenAddr)
		if err != nil {
			fmt.Println("Error creating Source connection:", err)
			os.Exit(1)
		}

		defer conn.Close()
		fmt.Println("Starting to listen on the socket")
		buf := make([]byte, 1024)
		for {
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Error reading from UDP:", err)
				break
			}

			// Decode JSON data
			var receivedMsg Chunk
			err = json.Unmarshal(buf[:n], &receivedMsg)
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				continue
			}

			fmt.Printf("Received message from %s: %v\n", addr.String(), receivedMsg)
			if _, ok := d.ChunkData[receivedMsg.Id]; !ok {
				d.ChunkData[receivedMsg.Id] = []Chunk{}
			}
			d.ChunkData[receivedMsg.Id] = append(d.ChunkData[receivedMsg.Id], receivedMsg)

			if len(d.ChunkData[receivedMsg.Id]) == receivedMsg.Total {
				sort.Slice(d.ChunkData[receivedMsg.Id], func(i, j int) bool {
					return d.ChunkData[receivedMsg.Id][i].Id < d.ChunkData[receivedMsg.Id][j].Id
				})

				fmt.Println("Decoded chunks", d.ChunkData[receivedMsg.Id])

				content := []byte{}
				for _, ch := range d.ChunkData[receivedMsg.Id] {
					content = append(content, []byte(ch.Chunk)...)
				}

				var msg Message
				err = json.Unmarshal([]byte(content), &msg)
				if err != nil {
					fmt.Println("Error decoding Final JSON:", err)
				}

				fmt.Println("Final decoded msg ", msg)
				d.ProcessEvent(msg)
				delete(d.ChunkData, receivedMsg.Id)
			}
		}
		fmt.Println("Finishing election now!!")
		ch <- true
	}(ch)

	return ch
}

func (d *Democracy) GenerateChunks(data []byte) []Chunk {
	chunks := []Chunk{}
	shortId, _ := GenerateShortID()
	if len(data) > d.Config.MaxPacketSize {
		count := Ceil(len(data), d.Config.MaxPacketSize)
		for i := 0; i < count; i++ {
			last := (i + 1) * d.Config.MaxPacketSize
			if last > len(data) {
				last = len(data)
			}
			chunks = append(chunks, Chunk{
				Id:      shortId,
				Chunk:   string(data[i*d.Config.MaxPacketSize : last]),
				Total:   count,
				Counter: i,
			})
		}
	} else {
		chunks = append(chunks, Chunk{
			Id:      shortId,
			Chunk:   string(data),
			Total:   1,
			Counter: 0,
		})
	}
	return chunks
}
