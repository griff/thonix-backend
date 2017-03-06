package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/griff/thonix/shell"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

type WsConn struct {
	conn   *websocket.Conn
	reader io.Reader
}

func (ws *WsConn) Read(buf []byte) (n int, err error) {
	if ws.reader == nil {
		_, ws.reader, err = ws.conn.NextReader()
		if err != nil {
			return 0, err
		}
	}
	n, err = ws.reader.Read(buf)
	if err == io.EOF {
		ws.reader = nil
		return n, nil
	}
	return n, err
}

func (ws *WsConn) Write(buf []byte) (n int, err error) {
	err = ws.conn.WriteMessage(websocket.BinaryMessage, buf)
	n = len(buf)
	return
}

func (ws *WsConn) Close() error {
	return ws.conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseGoingAway, "connection ended"))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var processes = shell.NewProcessManager()

type ShellHandler struct {
	shell string
	pm *shell.ProcessManager
	upgrader websocket.Upgrader
}

func NewShellHandler(shellPath string) *ShellHandler {
	return &ShellHandler {
		shell: shellPath,
		pm: shell.NewProcessManager(),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (sh *ShellHandler) MakeProcess(w http.ResponseWriter, r *http.Request) {
	c := exec.Command(sh.shell, "--login")
	_, err := sh.pm.Start(c)
	if err != nil {
		log.Printf("Error starting process %v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}
	log.Printf("Making process %d", c.Process.Pid)
	pid := fmt.Sprintf("%d", c.Process.Pid)
	w.Write([]byte(pid))
}

func (sh *ShellHandler) ShellWSHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["pid"], 10, 0)
	if err != nil {
		log.Printf("Invalid pid %v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("Making websocket %v", pid)
	process := sh.pm.Get(int(pid))
	if process != nil {
		log.Printf("Found no process with pid %d\n", pid)
		http.NotFound(w, r)
		return
	}

	log.Printf("Making websocket %v", process)
	conn, err := sh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}

	wsConn := &WsConn{conn, nil}
	process.Attach(wsConn, wsConn, true)
}
