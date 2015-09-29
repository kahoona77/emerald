package bot

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/kahoona77/emerald/models"
)

func (ib *IrcBot) startDownload(fileEvent *models.DccFileEvent, startPos int64, settings *models.EmeraldSettings) {
	//check if the file is in the queue
	download := ib.pending[fileEvent.FileName]
	if download == nil {
		log.Printf("Could not find download-file '%v' in pending list ", fileEvent.FileName)
		return
	}

	//remove from pending list
	delete(ib.pending, fileEvent.FileName)

	file := getTempFile(fileEvent, settings)

	// set start position
	var totalBytes int64
	totalBytes = startPos
	file.Seek(startPos, 0)

	// make a write buffer
	w := bufio.NewWriter(file)

	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	conn, err := ib.createConnection(fileEvent)
	if err != nil {
		return
	}

	var complete bool
	var inBuf = make([]byte, 1024)
	counter := 0

	//read-loop
	for {
		//read a chunk
		n, err := conn.Read(inBuf)
		if err != nil {
			if err == io.EOF {
				complete = true
			} else {
				log.Printf("Read error: %s", err)
			}
			break
		}
		totalBytes += int64(n)

		// write to File
		if _, err := w.Write(inBuf[:n]); err != nil {
			log.Printf("Write to file error: %s", err)
			break
		}

		//Send back an acknowledgement of how many bytes we have got so far.
		//Convert bytesTransfered to an "unsigned, 4 byte integer in network byte order", per DCC specification
		outBuf := makeOutBuffer(totalBytes)
		if _, err = conn.Write(outBuf); err != nil {
			log.Printf("Write error: %s", err)
			break
		}

		if err = w.Flush(); err != nil {
			log.Printf("Flush error: %s", err)
			break
		}

		// only send updates every 500 KB
		counter++
		if counter == 500 {
			ib.updateChan <- models.DccUpdate{fileEvent.FileName, totalBytes, fileEvent.Size, models.UPDATE}
			counter = 0
		}
	}
	conn.Close()

	if complete {
		ib.updateChan <- models.DccUpdate{fileEvent.FileName, totalBytes, fileEvent.Size, models.COMPLETE}
	} else {
		ib.updateChan <- models.DccUpdate{fileEvent.FileName, totalBytes, fileEvent.Size, models.FAIL}
	}

}

func (ib *IrcBot) createConnection(fileEvent *models.DccFileEvent) (net.Conn, error) {
	//connect
	tcpConn, err := net.Dial("tcp", fileEvent.IP.String()+":"+fileEvent.Port)
	if err != nil {
		log.Printf("Connect error: %v", err)
		return nil, err
	}

	//add to throttled pool
	conn, err := ib.connPool.AddConn(tcpConn)
	if err != nil {
		log.Printf("Error while adding to connection pool: %s", err)
		tcpConn.Close()
		return nil, err
	}
	return conn, nil
}

func getTempFile(fileEvent *models.DccFileEvent, settings *models.EmeraldSettings) *os.File {
	filename := filepath.FromSlash(settings.TempDir + "/" + fileEvent.FileName)
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		fo, err := os.Create(filename)
		if err != nil {
			log.Printf("File create error: %s", err)
		}
		return fo
	}

	fo, err := os.OpenFile(filename, os.O_WRONLY, 0777)
	if err != nil {
		log.Printf("File open error: %s", err)
	}
	return fo
}

func fileExists(fileEvent *models.DccFileEvent, settings *models.EmeraldSettings) (bool, int64) {
	filename := filepath.FromSlash(settings.TempDir + "/" + fileEvent.FileName)
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, 0
	}
	return true, info.Size()
}

func makeOutBuffer(totalBytes int64) []byte {
	var bytes = make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(totalBytes))
	return bytes
}
