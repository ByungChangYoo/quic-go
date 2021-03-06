// Package logging defines a logging interface for quic-go.
// This package should not be considered stable
package logging

import (
	"net"
	"time"

	"github.com/lucas-clemente/quic-go/internal/congestion"
	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/wire"
)

type (
	// A ByteCount is used to count bytes.
	ByteCount = protocol.ByteCount
	// A ConnectionID is a QUIC Connection ID.
	ConnectionID = protocol.ConnectionID
	// The EncryptionLevel is the encryption level of a packet.
	EncryptionLevel = protocol.EncryptionLevel
	// The KeyPhase is the key phase of the 1-RTT keys.
	KeyPhase = protocol.KeyPhase
	// The PacketNumber is the packet number of a packet.
	PacketNumber = protocol.PacketNumber
	// The Perspective is the role of a QUIC endpoint (client or server).
	Perspective = protocol.Perspective
	// The StreamID is the stream ID.
	StreamID = protocol.StreamID
	// The StreamNum is the number of the stream.
	StreamNum = protocol.StreamNum
	// The StreamType is the type of the stream (unidirectional or bidirectional).
	StreamType = protocol.StreamType
	// The VersionNumber is the QUIC version.
	VersionNumber = protocol.VersionNumber
	// The Header is the QUIC packet header, before removing header protection.
	Header = wire.Header
	// The ExtendedHeader is the QUIC packet header, after removing header protection.
	ExtendedHeader = wire.ExtendedHeader
	// The TransportParameters are QUIC transport parameters.
	TransportParameters = wire.TransportParameters

	// The RTTStats contain statistics used by the congestion controller.
	RTTStats = congestion.RTTStats
)

const (
	// PerspectiveServer is used for a QUIC server
	PerspectiveServer Perspective = protocol.PerspectiveServer
	// PerspectiveClient is used for a QUIC client
	PerspectiveClient Perspective = protocol.PerspectiveClient
)

const (
	// EncryptionInitial is the Initial encryption level
	EncryptionInitial EncryptionLevel = protocol.EncryptionInitial
	// EncryptionHandshake is the Handshake encryption level
	EncryptionHandshake EncryptionLevel = protocol.EncryptionHandshake
	// Encryption1RTT is the 1-RTT encryption level
	Encryption1RTT EncryptionLevel = protocol.Encryption1RTT
	// Encryption0RTT is the 0-RTT encryption level
	Encryption0RTT EncryptionLevel = protocol.Encryption0RTT
)

const (
	// StreamTypeUni is a unidirectional stream
	StreamTypeUni = protocol.StreamTypeUni
	// StreamTypeBidi is a bidirectional stream
	StreamTypeBidi = protocol.StreamTypeBidi
)

// A Tracer traces events.
type Tracer interface {
	TracerForServer(odcid ConnectionID) ConnectionTracer
	TracerForClient(odcid ConnectionID) ConnectionTracer
}

// A ConnectionTracer records events.
type ConnectionTracer interface {
	StartedConnection(local, remote net.Addr, version VersionNumber, srcConnID, destConnID ConnectionID)
	ClosedConnection(CloseReason)
	SentTransportParameters(*TransportParameters)
	ReceivedTransportParameters(*TransportParameters)
	SentPacket(hdr *ExtendedHeader, packetSize ByteCount, ack *AckFrame, frames []Frame)
	ReceivedVersionNegotiationPacket(*Header)
	ReceivedRetry(*Header)
	ReceivedPacket(hdr *ExtendedHeader, packetSize ByteCount, frames []Frame)
	ReceivedStatelessReset(token *[16]byte)
	BufferedPacket(PacketType)
	DroppedPacket(PacketType, ByteCount, PacketDropReason)
	UpdatedMetrics(rttStats *RTTStats, cwnd ByteCount, bytesInFLight ByteCount, packetsInFlight int)
	LostPacket(EncryptionLevel, PacketNumber, PacketLossReason)
	UpdatedPTOCount(value uint32)
	UpdatedKeyFromTLS(EncryptionLevel, Perspective)
	UpdatedKey(generation KeyPhase, remote bool)
	DroppedEncryptionLevel(EncryptionLevel)
	SetLossTimer(TimerType, EncryptionLevel, time.Time)
	LossTimerExpired(TimerType, EncryptionLevel)
	LossTimerCanceled()
	// Close is called when the connection is closed.
	Close()
}
