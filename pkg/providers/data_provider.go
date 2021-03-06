package providers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v8/pkg/app"
	api "github.com/pojntfx/liwasc/pkg/api/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Event struct {
	Message string
	Time    time.Time
}

type ScannerMetadata struct {
	Subnets []string
	Device  string
}

type Port struct {
	// Internal metadata
	createdAt time.Time
	priority  int64

	// Data
	PortNumber        int64
	TransportProtocol string

	// Public metadata
	ServiceName             string
	Description             string
	Assignee                string
	Contact                 string
	RegistrationDate        string
	ModificationDate        string
	Reference               string
	ServiceCode             string
	UnauthorizedUseReported string
	AssignmentNotes         string
}

type Node struct {
	// Internal metadata
	createdAt            time.Time
	priority             int64
	lastNodeWakePriority int64

	// Data
	MACAddress       string
	IPAddress        string
	PoweredOn        bool
	PortScanRunning  bool
	LastPortScanDate time.Time
	Ports            []Port
	NodeWakeRunning  bool
	LastNodeWakeDate time.Time

	// Public metadata
	Vendor       string
	Registry     string
	Organization string
	Address      string
	Visible      bool
}

type Network struct {
	ScannerMetadata  ScannerMetadata
	NodeScanRunning  bool
	LastNodeScanDate time.Time
	Nodes            []Node
	Events           []Event
}

type DataProviderChildrenProps struct {
	Network Network

	TriggerNetworkScan func(nodeScanTimeout int64, portScanTimeout int64, macAddress string)
	StartNodeWake      func(nodeWakeTimeout int64, macAddress string)

	Error   error
	Recover func()
	Ignore  func()
}

type DataProvider struct {
	app.Compo

	AuthenticatedContext   context.Context
	MetadataService        api.MetadataServiceClient
	NodeAndPortScanService api.NodeAndPortScanServiceClient
	NodeWakeService        api.NodeWakeServiceClient
	Children               func(DataProviderChildrenProps) app.UI

	network     Network
	networkLock sync.Mutex

	err error
}

func (c *DataProvider) Render() app.UI {
	return c.Children(DataProviderChildrenProps{
		Network: func() Network {
			sortedNetwork := c.network

			sort.Slice(sortedNetwork.Nodes, func(i, j int) bool {
				ipA := net.ParseIP(sortedNetwork.Nodes[i].IPAddress)
				ipB := net.ParseIP(sortedNetwork.Nodes[j].IPAddress)

				return bytes.Compare(ipA, ipB) < 0
			})

			for _, node := range sortedNetwork.Nodes {
				sort.Slice(node.Ports, func(i, j int) bool {
					protoA := node.Ports[i].TransportProtocol
					protoB := node.Ports[j].TransportProtocol

					return strings.Compare(protoA, protoB) < 0
				})

				sort.Slice(node.Ports, func(i, j int) bool {
					portA := node.Ports[i].PortNumber
					portB := node.Ports[j].PortNumber

					return portA-portB < 0
				})
			}

			return sortedNetwork
		}(),

		TriggerNetworkScan: c.triggerNetworkScan,
		StartNodeWake:      c.startNodeWake,

		Error:   c.err,
		Recover: c.recover,
		Ignore:  c.ignore,
	})
}

func (c *DataProvider) triggerNetworkScan(nodeScanTimeout int64, portScanTimeout int64, macAddress string) {
	// Optimistic UI
	c.dispatch(func() {
		// Set the node scan bool
		c.network.NodeScanRunning = true

		// Scan for one address
		if macAddress != "" {
			for i, node := range c.network.Nodes {
				if node.MACAddress == macAddress {
					// Set the port scan bool
					c.network.Nodes[i].PortScanRunning = true
					c.network.Nodes[i].Ports = []Port{}
				}
			}
		} else {
			c.network.Nodes = []Node{}
		}
	})

	// Start the node scan
	if _, err := c.NodeAndPortScanService.StartNodeScan(c.AuthenticatedContext, &api.NodeScanStartMessage{
		NodeScanTimeout: nodeScanTimeout,
		PortScanTimeout: portScanTimeout,
		MACAddress:      macAddress,
	}); err != nil {
		c.panic(err)

		return
	}
}

func (c *DataProvider) startNodeWake(nodeWakeTimeout int64, macAddress string) {
	// Optimistic UI
	c.dispatch(func() {
		for i, node := range c.network.Nodes {
			if node.MACAddress == macAddress {
				// Set the node wake bool
				c.network.Nodes[i].NodeWakeRunning = true
			}
		}
	})

	// Start the node wake
	if _, err := c.NodeWakeService.StartNodeWake(c.AuthenticatedContext, &api.NodeWakeStartMessage{
		NodeWakeTimeout: nodeWakeTimeout,
		MACAddress:      macAddress,
	}); err != nil {
		c.panic(err)

		return
	}
}

func (c *DataProvider) recover() {
	// Clear the error
	c.err = nil

	// Resubscribe
	c.OnMount(app.Context{})

	c.Update()
}

func (c *DataProvider) ignore() {
	// Only clear the error
	c.err = nil

	c.Update()
}

func (c *DataProvider) panic(err error) {
	// Set the error
	c.err = err

	c.Update()
}

func (c *DataProvider) dispatch(action func()) {
	c.networkLock.Lock()

	action()

	c.Update()
	c.networkLock.Unlock()
}

func (c *DataProvider) OnMount(context app.Context) {
	// Initialize network struct
	c.dispatch(func() {
		c.network = Network{
			ScannerMetadata: ScannerMetadata{
				Subnets: []string{},
				Device:  "",
			},
			NodeScanRunning:  false,
			LastNodeScanDate: time.Unix(0, 0),
			Nodes:            []Node{},
			Events:           []Event{},
		}
	})

	// Get scanner metadata
	go func() {
		// Fetch from service
		scannerMetadata, err := c.MetadataService.GetMetadataForScanner(c.AuthenticatedContext, &emptypb.Empty{})
		if err != nil {
			c.panic(err)

			return
		}

		// Write to struct
		c.dispatch(func() {
			c.network.ScannerMetadata.Device = scannerMetadata.GetDevice()
			c.network.ScannerMetadata.Subnets = scannerMetadata.GetSubnets()
		})
	}()

	// Subscribe to node scans
	go func() {
		// Get stream from service
		nodeScanStream, err := c.NodeAndPortScanService.SubscribeToNodeScans(c.AuthenticatedContext, &emptypb.Empty{})
		if err != nil {
			c.panic(err)

			return
		}

		// Process stream
		for {
			// Receive scan from stream
			nodeScan, err := nodeScanStream.Recv()
			if err != nil {
				c.panic(err)

				return
			}

			// Parse the scan's date
			nodeScanCreatedAt, err := time.Parse(time.RFC3339, nodeScan.GetCreatedAt())
			if err != nil {
				c.panic(err)

				return
			}

			// Only continue evaluation if this scan is newer
			if nodeScanCreatedAt.After(c.network.LastNodeScanDate) || nodeScanCreatedAt.Equal(c.network.LastNodeScanDate) {
				c.dispatch(func() {
					// Log the node scan
					c.network.Events = append(c.network.Events, Event{
						Time:    nodeScanCreatedAt,
						Message: fmt.Sprintf("Received node scan for subnets %v", strings.Join(c.network.ScannerMetadata.Subnets, ", ")),
					})

					// Set the new latest scan date
					c.network.LastNodeScanDate = nodeScanCreatedAt

					// Update the scan indicator
					if nodeScan.Done {
						c.network.NodeScanRunning = false
					} else {
						c.network.NodeScanRunning = true
					}
				})

				// Subscribe to nodes
				go func(nsm *api.NodeScanMessage) {
					// Get stream from service
					nodeStream, err := c.NodeAndPortScanService.SubscribeToNodes(c.AuthenticatedContext, nsm)
					if err != nil {
						c.panic(err)

						return
					}

					// Process stream
					for {
						// Receive node from stream
						node, err := nodeStream.Recv()
						if err != nil {
							if err == io.EOF {
								break
							}

							c.panic(err)

							return
						}

						// Parse the node's date
						nodeCreatedAt, err := time.Parse(time.RFC3339, node.GetCreatedAt())
						if err != nil {
							c.panic(err)

							return
						}

						// Get the node's metadata
						nodeMetadata, err := c.MetadataService.GetMetadataForNode(c.AuthenticatedContext, &api.NodeMetadataReferenceMessage{
							MACAddress: node.GetMACAddress(),
						})
						if err != nil {
							if status, ok := status.FromError(err); ok && status.Code() == codes.NotFound {
								nodeMetadata = &api.NodeMetadataMessage{
									MACAddress:   node.GetMACAddress(),
									Vendor:       "",
									Registry:     "",
									Organization: "",
									Address:      "",
									Visible:      true, // The majority are visible, so set it as the default value
								}
							} else {
								c.panic(err)

								return
							}
						}

						c.dispatch(func() {
							// Only continue if this node is newer and has a higher priority
							lastKnownNodeIndex := -1
							for i, knownNode := range c.network.Nodes {
								if knownNode.MACAddress == node.GetMACAddress() {
									if (nodeCreatedAt.After(knownNode.createdAt) || nodeCreatedAt.Equal(knownNode.createdAt)) && knownNode.priority > node.GetPriority() {
										// Ignore the node
										return
									}

									lastKnownNodeIndex = i

									break
								}
							}

							// If an old node exists, remove it, but keep the current scan & wake state
							ports := []Port{} // Ignore old ports
							portScanRunning := false
							lastPortScanDate := time.Unix(0, 0)
							nodeWakeRunning := false
							lastNodeWakeDate := time.Unix(0, 0)
							lastNodeWakePriority := int64(0)
							if lastKnownNodeIndex != -1 {
								portScanRunning = c.network.Nodes[lastKnownNodeIndex].PortScanRunning
								lastPortScanDate = c.network.Nodes[lastKnownNodeIndex].LastPortScanDate
								nodeWakeRunning = c.network.Nodes[lastKnownNodeIndex].NodeWakeRunning
								lastNodeWakeDate = c.network.Nodes[lastKnownNodeIndex].LastNodeWakeDate
								lastNodeWakePriority = c.network.Nodes[lastKnownNodeIndex].lastNodeWakePriority

								c.network.Nodes = append(c.network.Nodes[:lastKnownNodeIndex], c.network.Nodes[lastKnownNodeIndex+1:]...)
							}

							// Add the new node
							c.network.Nodes = append(c.network.Nodes, Node{
								createdAt:            nodeCreatedAt,
								priority:             node.GetPriority(),
								lastNodeWakePriority: lastNodeWakePriority,

								MACAddress:       node.GetMACAddress(),
								IPAddress:        node.GetIPAddress(),
								PoweredOn:        node.GetPoweredOn(),
								PortScanRunning:  portScanRunning,
								LastPortScanDate: lastPortScanDate,
								Ports:            ports,
								NodeWakeRunning:  nodeWakeRunning,
								LastNodeWakeDate: lastNodeWakeDate,

								Vendor:       nodeMetadata.GetVendor(),
								Registry:     nodeMetadata.GetRegistry(),
								Organization: nodeMetadata.GetOrganization(),
								Address:      nodeMetadata.GetAddress(),
								Visible:      nodeMetadata.GetVisible(),
							})

							// Log the node
							c.network.Events = append(c.network.Events, Event{
								Time:    nodeCreatedAt,
								Message: fmt.Sprintf("Received node %v (%v)", node.GetMACAddress(), node.GetIPAddress()),
							})
						})

						// Subscribe to port scans
						go func(nm *api.NodeMessage) {
							// Get stream from service
							portScanStream, err := c.NodeAndPortScanService.SubscribeToPortScans(c.AuthenticatedContext, nm)
							if err != nil {
								c.panic(err)

								return
							}

							// Process stream
							for {
								// Receive scan from stream
								portScan, err := portScanStream.Recv()
								if err != nil {
									if err == io.EOF {
										break
									}

									c.panic(err)

									return
								}

								// Parse the scan's date
								portScanCreatedAt, err := time.Parse(time.RFC3339, portScan.GetCreatedAt())
								if err != nil {
									c.panic(err)

									return
								}

								// Check if this port scan is the newest one
								portScanIsNewest := false
								c.dispatch(func() {
									for i, currentNode := range c.network.Nodes {
										if currentNode.MACAddress == node.GetMACAddress() && (portScanCreatedAt.After(currentNode.LastPortScanDate) || portScanCreatedAt.Equal(currentNode.LastPortScanDate)) {
											portScanIsNewest = true

											// Set the new latest scan date
											c.network.Nodes[i].LastPortScanDate = portScanCreatedAt

											// Update the scan indicator
											if portScan.Done {
												c.network.Nodes[i].PortScanRunning = false
											} else {
												c.network.Nodes[i].PortScanRunning = true
											}

											break
										}
									}
								})

								// Only continue evaluation if this scan is newer
								if portScanIsNewest {
									// Log the port scan
									c.dispatch(func() {
										c.network.Events = append(c.network.Events, Event{
											Time:    portScanCreatedAt,
											Message: fmt.Sprintf("Received port scan for node %v (%v)", node.GetMACAddress(), node.GetIPAddress()),
										})
									})

									// Subscribe to ports
									go func(ps *api.PortScanMessage) {
										// Get stream from service
										portStream, err := c.NodeAndPortScanService.SubscribeToPorts(c.AuthenticatedContext, ps)
										if err != nil {
											c.panic(err)

											return
										}

										// Process stream
										for {
											// Receive port from stream
											port, err := portStream.Recv()
											if err != nil {
												if err == io.EOF {
													break
												}

												c.panic(err)

												return
											}

											// Parse the port's date
											portCreatedAt, err := time.Parse(time.RFC3339, port.GetCreatedAt())
											if err != nil {
												c.panic(err)

												return
											}

											// Get the port's metadata
											portMetadata, err := c.MetadataService.GetMetadataForPort(c.AuthenticatedContext, &api.PortMetadataReferenceMessage{
												PortNumber:        port.GetPortNumber(),
												TransportProtocol: port.GetTransportProtocol(),
											})
											if err != nil {
												if status, ok := status.FromError(err); ok && status.Code() == codes.NotFound {
													portMetadata = &api.PortMetadataMessage{
														ServiceName:             "",
														PortNumber:              port.GetPortNumber(),
														TransportProtocol:       port.GetTransportProtocol(),
														Description:             "",
														Assignee:                "",
														Contact:                 "",
														RegistrationDate:        "",
														ModificationDate:        "",
														Reference:               "",
														ServiceCode:             "",
														UnauthorizedUseReported: "",
														AssignmentNotes:         "",
													}
												} else {
													c.panic(err)

													return
												}
											}

											c.dispatch(func() {
												// Only continue if this port is newer and has a higher priority
												lastKnownNodeIndex := -1
												lastKnownPortIndex := -1
												for nodeIndex, knownNode := range c.network.Nodes {
													if knownNode.MACAddress == nm.GetMACAddress() {
														lastKnownNodeIndex = nodeIndex

														for portIndex, knownPort := range knownNode.Ports {
															if knownPort.PortNumber == port.PortNumber && knownPort.TransportProtocol == port.TransportProtocol {
																if (portCreatedAt.After(knownPort.createdAt) || portCreatedAt.Equal(knownPort.createdAt)) && knownPort.priority > port.GetPriority() {
																	// Ignore the port
																	return
																}

																lastKnownPortIndex = portIndex

																break
															}
														}

														break
													}
												}

												// Asynchronity; the node specified in the outer loops might not exist anymore
												if lastKnownNodeIndex != -1 {
													// If an old port exists, remove it.
													if lastKnownPortIndex != -1 {
														c.network.Nodes[lastKnownNodeIndex].Ports = append(c.network.Nodes[lastKnownNodeIndex].Ports[:lastKnownPortIndex], c.network.Nodes[lastKnownNodeIndex].Ports[lastKnownPortIndex+1:]...)
													}

													// Add the new port
													c.network.Nodes[lastKnownNodeIndex].Ports = append(c.network.Nodes[lastKnownNodeIndex].Ports, Port{
														createdAt: portCreatedAt,
														priority:  port.GetPriority(),

														PortNumber:        port.GetPortNumber(),
														TransportProtocol: port.GetTransportProtocol(),

														ServiceName:             portMetadata.ServiceName,
														Description:             portMetadata.Description,
														Assignee:                portMetadata.Assignee,
														Contact:                 portMetadata.Contact,
														RegistrationDate:        portMetadata.RegistrationDate,
														ModificationDate:        portMetadata.ModificationDate,
														Reference:               portMetadata.Reference,
														ServiceCode:             portMetadata.ServiceCode,
														UnauthorizedUseReported: portMetadata.UnauthorizedUseReported,
														AssignmentNotes:         portMetadata.AssignmentNotes,
													})

													// Log the port
													c.network.Events = append(c.network.Events, Event{
														Time:    portCreatedAt,
														Message: fmt.Sprintf("Received port %v/%v on node %v (%v)", port.GetPortNumber(), port.GetTransportProtocol(), node.GetMACAddress(), node.GetIPAddress()),
													})
												}
											})
										}
									}(portScan)
								}
							}
						}(node)
					}
				}(nodeScan)
			}
		}
	}()

	// Subscribe to node wakes
	go func() {
		retries := 0

		for {
			// Get stream from service
			nodeWakeStream, err := c.NodeWakeService.SubscribeToNodeWakes(c.AuthenticatedContext, &emptypb.Empty{})
			if err != nil {
				c.panic(err)

				return
			}

			// Process stream
			for {
				// Receive node wake from stream
				nodeWake, err := nodeWakeStream.Recv()
				if err != nil {
					c.panic(err)

					return
				}

				// Parse the node wake's date
				nodeWakeCreatedAt, err := time.Parse(time.RFC3339, nodeWake.GetCreatedAt())
				if err != nil {
					c.panic(err)

					return
				}

				// Update the node's wake status if it's newer and it's priority is higher
				nodeFound := false
				c.dispatch(func() {
					for i, currentNode := range c.network.Nodes {
						if currentNode.MACAddress == nodeWake.GetMACAddress() {
							nodeFound = true

							if (currentNode.LastNodeWakeDate.After(nodeWakeCreatedAt) || currentNode.LastNodeWakeDate.Equal(nodeWakeCreatedAt)) && currentNode.lastNodeWakePriority > nodeWake.GetPriority() {
								// Ignore the node wake

								return
							}

							// Set the new latest node wake date
							c.network.Nodes[i].LastNodeWakeDate = nodeWakeCreatedAt

							// Update the node wake indicator
							if nodeWake.Done {
								c.network.Nodes[i].NodeWakeRunning = false
								// If the scan is done, also set the powered on status
								c.network.Nodes[i].PoweredOn = nodeWake.GetPoweredOn()
							} else {
								c.network.Nodes[i].NodeWakeRunning = true
							}

							// Log the node wake
							c.network.Events = append(c.network.Events, Event{
								Time: nodeWakeCreatedAt,
								Message: fmt.Sprintf("Received node wake for node %v (%v)", nodeWake.GetMACAddress(), func() string {
									// In order for the node to be found, it has to have an IP address; this will never be `""`
									ip := ""

									for _, candidate := range c.network.Nodes {
										if candidate.MACAddress == nodeWake.GetMACAddress() {
											ip = candidate.IPAddress

											break
										}
									}

									return ip
								}()),
							})

							break
						}
					}
				})

				// If the node could not be found, the nodes have not been fetched yet
				if !nodeFound && retries <= 1000 {
					// Re-subscribe to the scan until the race has finished/the node exists
					// Unless there are manual interventions in the database diverging the node
					// and port scans from the node wakes this can't lead to an endless loop.
					// As a safety measure, it gives up after 1000 retries, in case they have.
					retries++

					time.Sleep(100 * time.Millisecond)

					if err := nodeWakeStream.CloseSend(); err != nil {
						c.panic(err)

						return
					}

					break
				}
			}
		}
	}()
}
