// +build linux

package journald

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/docker/docker/api/types/backend"
	"github.com/docker/docker/daemon/logger"
)

func (s *journald) Close() error {
	return nil
}

type JournalEntry struct {
	Pid                     int    `json:"_PID,string,omitempty"`
	Uid                     int    `json:"_UID,string,omitempty"`
	Gid                     int    `json:"_GID,string,omitempty"`
	Comm                    string `json:"_COMM,omitempty"`
	Exe                     string `json:"_EXE,omitempty"`
	Cmdline                 string `json:"_CMDLINE,omitempty"`
	CapEffective            string `json:"_CAP_EFFECTIVE,omitempty"`
	AuditSession            int    `json:"_AUDIT_SESSION,string,omitempty"`
	AuditLoginId            string `json:"_AUDIT_LOGINID,omitempty"`
	SystemdGroup            string `json:"_SYSTEMD_CGROUP,omitempty"`
	SystemdSession          string `json:"_SYSTEMD_SESSION,omitempty"`
	SystemdUnit             string `json:"_SYSTEMD_UNIT,omitempty"`
	SystemdUserInit         string `json:"_SYSTEMD_USER_INIT,omitempty"`
	SystemdOwnerUid         string `json:"_SYSTEMD_OWNER_UID,omitempty"`
	SystemdSlice            string `json:"_SYSTEMD_SLICE,omitempty"`
	SelinuxContext          string `json:"_SELINUX_CONTEXT,omitempty"`
	SourceRealtimeTimestamp int64  `json:"_SOURCE_REALTIME_TIMESTAMP,string,omitempty"`
	BootId                  string `json:"_BOOT_ID,omitempty"`
	MachineId               string `json:"_MACHINE_ID,omitempty"`
	Hostname                string `json:"_HOSTNAME,omitempty"`
	Transport               string `json:"_TRANSPORT,omitempty"`
	Cursor                  string `json:"__CURSOR,omitempty"`
	RealtimeTimestamp       int64  `json:"__REALTIME_TIMESTAMP,string,omitempty"`
	MonotonicTimestamp      int64  `json:"__MONOTONIC_TIMESTAMP,string,omitempty"`
	CoredumpUnit            string `json:"COREDUMP_UNIT,omitempty"`
	CoredumpUserInit        string `json:"COREDUMP_USER_INIT,omitempty"`
	ObjectPid               int    `json:"OBJECT_PID,string,omitempty"`
	ObjectUid               int    `json:"OBJECT_UID,string,omitempty"`
	ObjectGid               int    `json:"OBJECT_GID,string,omitempty"`
	ObjectComm              string `json:"OBJECT_COMM,omitempty"`
	ObjectExe               string `json:"OBJECT_EXE,omitempty"`
	ObjectCmdline           string `json:"OBJECT_CMDLINE,omitempty"`
	ObjectAuditSession      string `json:"OBJECT_AUDIT_SESSION,omitempty"`
	ObjectAuditLoginId      string `json:"OBJECT_AUDIT_LOGINID,omitempty"`
	ObjectSystemdCgroup     string `json:"OBJECT_SYSTEMD_CGROUP,omitempty"`
	ObjectSystemdSession    string `json:"OBJECT_SYSTEMD_SESSION,omitempty"`
	ObjectSystemdUnit       string `json:"OBJECT_SYSTEMD_UNIT,omitempty"`
	ObjectSystemdUserInit   string `json:"OBJECT_SYSTEMD_USER_INIT,omitempty"`
	ObjectSystemdOwnerUid   int    `json:"OBJECT_SYSTEMD_OWNER_UID,string,omitempty"`
	Message                 string `json:"MESSAGE,omitempty"`
	MessageId               int    `json:"MESSAGE_ID,string,omitempty"`
	Priority                int    `json:"PRIORITY,string,omitempty"`
	CodeFile                string `json:"CODE_FILE,omitempty"`
	CodeLine                string `json:"CODE_LINE,omitempty"`
	CodeFunc                string `json:"CODE_FUNC,omitempty"`
	ErrNo                   int    `json:"ERRNO,string,omitempty"`
	SyslogFacility          string `json:"SYSLOG_FACILITY,omitempty"`
	SyslogIdentifier        string `json:"SYSLOG_IDENTIFIER,omitempty"`
	ContainerId             string `json:"CONTAINER_ID,omitempty"`
	ContainerFullId         string `json:"CONTAINER_ID_FULL,omitempty"`
	ContainerName           string `json:"CONTAINER_NAME,omitempty"`
}

func (s *journald) readLogs(logWatcher *logger.LogWatcher, config logger.ReadConfig) {
	j_args := []string{
		"journalctl",
		"--output", "json",
		"--identifier", s.vars["CONTAINER_ID"],
		"--utc",
	}

	if config.Follow {
		j_args = append(j_args, "--follow")
	}
	if config.Tail >= 0 {
		j_args = append(j_args, "--lines", fmt.Sprint(config.Tail))
	}
	if !config.Since.IsZero() {
		j_args = append(j_args, "--since", config.Since.Format("2006-01-02 15:04:05"))
	}
	if !config.Until.IsZero() {
		j_args = append(j_args, "--until", config.Until.Format("2006-01-02 15:04:05"))
	}

	cmd := exec.Command(j_args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logWatcher.Err <- errors.Wrap(err, "failed to run journalctl")
		return
	}
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	var (
		done bool
	)
	for scanner.Scan() && !done {
		raw := scanner.Text()
		var entry JournalEntry
		err := json.Unmarshal(raw, &e)
		if err != nil {
			// Ignore blank lines
		}

		timestamp := time.Unix(e.RealtimeTimestamp/1000000, (int64(e.RealtimeTimestamp)%1000000)*1000)

		source := ""
		if e.Priority == 6 { // info
			source = "stdout"
		} else if e.Priority == 1 { // err
			source = "stderr"
		}

		select {
		case <-logWatcher.WatchConsumerGone():
			done = true
			return
		case logWatcher.Msg <- &logger.Message{
			Line:      e.Message,
			Source:    source,
			Timestamp: timestamp,
			Attrs:     []backend.LogAttr{},
		}:
			// noop
		}
	}
}

func (s *journald) ReadLogs(config logger.ReadConfig) *logger.LogWatcher {
	logWatcher := logger.NewLogWatcher()
	go s.readLogs(logWatcher, config)
	return logWatcher
}
