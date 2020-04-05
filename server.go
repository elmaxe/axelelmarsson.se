package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "io"
    "fmt"
    "log"
    "os"
    "github.com/kjk/dailyrotate"
    "path/filepath"
    "time"
)

func main() {
    maintenance := false

    //---------------------------------
    // https://presstige.io/p/Logging-HTTP-requests-in-Go-233de7fe59a747078b35b82a1b035d36
    logDir := "logs"
    err := os.MkdirAll(logDir, 0755)
    if err != nil {
		log.Fatalf("os.MkdirAll()(")
	}
	pathFormat := filepath.Join(logDir, "2006-01-02.txt")
    err = openLogFile(pathFormat, onLogClose)
    
    if err != nil {
		log.Fatalf("openLogFile failed with '%s'\n", err)
    }
    
    defer closeLogFile()
    //---------------------------------



    if (maintenance) {
        http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
            io.WriteString(w, "Server maintenance, come back later.")        
        })

    } else {
        //Handle och handlefunc tar inte regex, orkar inte skriva något skit som fixar det
        //Därav kodduplicering

        //Moist
        moist, _ := url.Parse("http://192.168.10.171:8002")
        http.Handle("moist.axelelmarsson.se/", httputil.NewSingleHostReverseProxy(moist))
        http.Handle("www.moist.axelelmarsson.se/", httputil.NewSingleHostReverseProxy(moist))
        
        //Base page
        basepage, _ := url.Parse("http://192.168.10.171:8001")
        http.Handle("axelelmarsson.se/", httputil.NewSingleHostReverseProxy(basepage))
        http.Handle("www.axelelmarsson.se/", httputil.NewSingleHostReverseProxy(basepage))

	//Olprovning
        olprovning, _ := url.Parse("http://192.168.10.171:8003")
        http.Handle("olprovning.axelelmarsson.se/", httputil.NewSingleHostReverseProxy(olprovning))
        http.Handle("www.olprovning.axelelmarsson.se/", httputil.NewSingleHostReverseProxy(olprovning))

        //Elmarsson.se
        http.HandleFunc("www.elmarsson.se/", elmarsson)
        http.HandleFunc("elmarsson.se/", elmarsson)
    }
    
    http.ListenAndServe(":8000", logRequest(http.DefaultServeMux))
    // http.ListenAndServe(":8000", nil)
}

func elmarsson(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Här var det tomt.")        
}

//---------------------------------
// https://presstige.io/p/Logging-HTTP-requests-in-Go-233de7fe59a747078b35b82a1b035d36

var (
	logFile *dailyrotate.File
)

func openLogFile(pathFormat string, onClose func(string, bool)) error {
	w, err := dailyrotate.NewFile(pathFormat, onLogClose)
	if err != nil {
		return err
	}
	logFile = w
	return nil
}

func onLogClose(path string, didRotate bool) {
	fmt.Printf("we just closed a file '%s', didRotate: %v\n", path, didRotate)
	if !didRotate {
		return
	}
	// process just closed file e.g. upload to S3 for backup
	go func() {
		// if processing takes a long time, do it in background
	}()
}

func closeLogFile() error {
	return logFile.Close()
}


func writeToLog(msg string) error {
	_, err := logFile.Write([]byte(msg))
	return err
}


func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        err := writeToLog(fmt.Sprintf("%s %s %s %s %s %d\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, r.Host, r.URL, r.ContentLength))
        if err != nil {
            log.Fatalf("writeToLog() failed with '%s'\n", err)
        }
		handler.ServeHTTP(w, r)
	})
}
//---------------------------------
