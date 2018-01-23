package p2p

import (
	"sync/atomic"
	"time"

	"github.com/Frank-gh/simple_blockchain/blockchain"
	"github.com/Frank-gh/tcpnetwork"
	"github.com/golang/glog"
)

type FuncHandler func([]byte, *tcpnetwork.Connection)

var (
	StopSrv         = make(chan struct{})
	StopCli         = make(chan struct{})
	serverConnected int32
	Peer            *peer
	timer           = time.NewTicker(1 * time.Second)
)

type peer struct {
	svrPeers map[string]*tcpnetwork.Connection
	cliPeers map[string]*tcpnetwork.Connection
	peerName string
}

func init() {
	Peer = &peer{
		svrPeers: make(map[string]*tcpnetwork.Connection),
		cliPeers: make(map[string]*tcpnetwork.Connection),
		peerName: "",
	}
	AddHandler("index", onIndex)
	AddHandler("block", onBlock)
}

func (this *peer) AddServer(sName string, pServer *tcpnetwork.Connection) {
	this.svrPeers[sName] = pServer
}
func (this *peer) DelServer(sName string) {
	delete(this.svrPeers, sName)
}
func (this *peer) SetPeerName(pName string) {
	this.peerName = pName
}

func (this *peer) AddClient(sName string, pClient *tcpnetwork.Connection) {
	this.cliPeers[sName] = pClient
}
func (this *peer) DelClient(sName string) {
	delete(this.cliPeers, sName)
}

func OpenPort(host, port string) string {
	var err error
	addr := host + ":" + port
	server := tcpnetwork.NewTCPNetwork(1024, tcpnetwork.NewStreamProtocol4())
	err = server.Listen(addr)
	if nil != err {
		return err.Error()
	}
	Peer.SetPeerName(addr)
	go runServer(server)
	return "Listening on  " + host + ":" + port
}

func Connect(host, port string) string {
	var err error
	addr := host + ":" + port
	client := tcpnetwork.NewTCPNetwork(1024, tcpnetwork.NewStreamProtocol4())
	conn, err := client.Connect(addr)
	if nil != err {
		return err.Error()
	}
	go runClient(client, conn)
	return "Connected on  " + host + ":" + port
}

func runServer(server *tcpnetwork.TCPNetwork) {
	var activeConn *tcpnetwork.Connection = nil
	for {
		select {
		case evt, ok := <-server.GetEventQueue():
			{
				if !ok {
					return
				}

				switch evt.EventType {
				case tcpnetwork.KConnEvent_Connected:
					{
						Peer.AddClient(evt.Conn.GetRemoteAddress(), evt.Conn)
						Peer.SetPeerName(evt.Conn.GetLocalAddress())
						sendIndex(blockchain.BlockChain.Index(), evt.Conn)
						activeConn = evt.Conn
						glog.Info("Client ", evt.Conn.GetRemoteAddress(), " connected")
					}
				case tcpnetwork.KConnEvent_Close:
					{
						Peer.DelClient(evt.Conn.GetRemoteAddress())
						glog.Info("Client ", evt.Conn.GetRemoteAddress(), " disconnected")
					}
				case tcpnetwork.KConnEvent_Data:
					{
						text := string(evt.Data)
						glog.Info("server recv : ", text)
						handleInput(evt.Data, evt.Conn)
					}
				}
			}
		case <-timer.C:
			{
				if activeConn != nil {
					sendIndex(blockchain.BlockChain.Index(), activeConn)
				}
			}
		case <-StopSrv:
			{
				return
			}
		}
	}
}

func runClient(client *tcpnetwork.TCPNetwork, cliConn *tcpnetwork.Connection) {
	var activeConn *tcpnetwork.Connection = nil
EVENTLOOP:
	for {
		select {
		case evt, ok := <-client.GetEventQueue():
			{
				if !ok {
					return
				}
				switch evt.EventType {
				case tcpnetwork.KConnEvent_Connected:
					{
						// save server
						Peer.AddServer(evt.Conn.GetRemoteAddress(), cliConn)
						Peer.SetPeerName(evt.Conn.GetLocalAddress())
						sendIndex(blockchain.BlockChain.Index(), evt.Conn)
						glog.Info("Connected to server")
						activeConn = evt.Conn
						atomic.StoreInt32(&serverConnected, 1)
					}
				case tcpnetwork.KConnEvent_Close:
					{
						// delete server
						Peer.DelServer(evt.Conn.GetRemoteAddress())
						glog.Info("Disconnected from server")
						activeConn = nil
						atomic.StoreInt32(&serverConnected, 0)
						break EVENTLOOP
					}
				case tcpnetwork.KConnEvent_Data:
					{
						text := string(evt.Data)
						glog.Info("client recv : ", text)
						handleInput(evt.Data, evt.Conn)
					}
				}
			}
		case <-timer.C:
			{
				if activeConn != nil {
					sendIndex(blockchain.BlockChain.Index(), activeConn)
				}
			}
		case <-StopCli:
			{
				return
			}
		}
	}
}
