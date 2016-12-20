package main

import (
    "fmt"
    "github.com/kr/pty"
    "github.com/gorilla/mux"
    "github.com/gorilla/websocket"
    "io"
    "log"
    "net/http"
    "os"
    "os/exec"
)

type Process struct {
    Logs []byte
    Command *exec.Cmd
    Pty *os.File
    Connections []*websocket.Conn
}

func (process *Process) AddConnection(conn *websocket.Conn) {
    log.Printf("Log: %s", string(process.Logs))
    conn.WriteMessage(websocket.TextMessage, process.Logs)
    process.Connections = append(process.Connections, conn)
}

func (process *Process) WriteMessage(msg []byte) {
    process.Logs = append(process.Logs, msg...)
    //log.Println("Logs is now %v", string(process.Logs))
    for i, conn := range process.Connections {
        //log.Printf("Writing to conn %v", string(msg))
        err := conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            process.Connections = append(process.Connections[:i], process.Connections[i+1:]...)
            conn.Close()
        }
    }
}

var processes map[string]*Process = make(map[string]*Process)
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func makeProcess(w http.ResponseWriter, r *http.Request) {
    c := exec.Command("bash", "--login")
    f, err := pty.Start(c)
    log.Printf("Making process %d", c.Process.Pid)
    process := &Process{
        Logs: make([]byte, 0),
        Command: c,
        Pty: f,
        Connections: make([]*websocket.Conn, 0),
    }
    pid := fmt.Sprintf("%d", c.Process.Pid)
    processes[pid] = process
    if err != nil {
        buf := []byte(err.Error())
        process.Logs = append(process.Logs, buf...)
    } else {
        go func() {
            buf := make([]byte, 1024)
            for {
                log.Println("Reading from pty")
                n, err := f.Read(buf)
                if err != nil {
                    buf = []byte(err.Error())
                    process.WriteMessage(buf)
                    for _, conn := range process.Connections {
                        conn.WriteMessage(websocket.CloseMessage,
                            websocket.FormatCloseMessage(websocket.CloseGoingAway, "program ended"))
                        conn.Close()
                    }
                    process.Connections = make([]*websocket.Conn, 0)
                    log.Println(err)
                    return
                } else {
                    msg := buf[:n]
                    //log.Println("Read from pty %d %v", n, string(msg))
                    process.WriteMessage(msg)
                }
            }
        }()
    }
    w.Write([]byte(pid))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pid := vars["pid"]
    log.Printf("Making websocket %v", pid)
    process := processes[pid]
    log.Printf("Making websocket %v", process)
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    process.AddConnection(conn)
    go func() {
        for {
            messageType, r, err := conn.NextReader()
            if err != nil {
                log.Printf("%v %v", messageType, err)
                return;
            }
            io.Copy(process.Pty, r)
        }
    }()
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/terminals/{pid}", wsHandler)
    r.HandleFunc("/terminals", makeProcess).Methods("POST")
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/target")))
    http.Handle("/", r)
    log.Println("Application listening on port 3001")
    log.Fatal(http.ListenAndServe(":3001", nil))
}


