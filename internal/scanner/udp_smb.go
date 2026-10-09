package scanner

import (
	"net"
	"time"

	"OrsoNetwork/internal/logger"

	"github.com/hirochachacha/go-smb2"
)

type SMBProbeResult struct {
	Found      bool
	EnumOK     bool
	Shares     []smb2.ShareInfo
	ErrorStage string
}

func ProbeSMB(
	ip string,
) SMBProbeResult {

	probeStart := time.Now()

	defer func() {

		logger.Debug(
			"SMB PROBE DURATION:",
			ip,
			time.Since(probeStart),
		)
	}()

	addr := ip + ":445"

	logger.Debug(
		"SMB CONNECT:",
		addr,
	)

	conn, err := net.DialTimeout(
		"tcp",
		addr,
		500*time.Millisecond,
	)

	if err != nil {

		logger.Debug(
			"SMB CONNECT ERROR:",
			ip,
			err,
		)

		return SMBProbeResult{
			Found:      false,
			ErrorStage: "tcp",
		}
	}

	defer conn.Close()

	logger.Debug(
		"SMB TCP CONNECTED:",
		ip,
	)

	dialer := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "guest",
			Password: "",
		},
	}

	logger.Debug(
		"SMB SESSION:",
		ip,
		"USER: guest",
	)

	session, err := dialer.Dial(conn)

	if err != nil {

		logger.Debug(
			"SMB SESSION ERROR:",
			ip,
			err,
		)

		return SMBProbeResult{
			Found:      false,
			ErrorStage: "session",
		}
	}

	defer session.Logoff()

	logger.Debug(
		"SMB SESSION ESTABLISHED:",
		ip,
	)

	shares, err := session.ListShares()

	if err != nil {

		logger.Debug(
			"SMB LIST SHARES ERROR:",
			ip,
			err,
		)

		return SMBProbeResult{
			Found:      true,
			EnumOK:     false,
			ErrorStage: "enum",
		}
	}

	logger.Info(
		"SMB FOUND:",
		ip,
		"SHARES:",
		shares,
	)

	return SMBProbeResult{
		Found:  true,
		EnumOK: true,
		Shares: shares,
	}
}
