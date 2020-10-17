package main

import (
    "fmt"
    "log"
    "net/http"
    //"net/url"
    "encoding/json"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Meeting struct{

	Id      		string `json:"Id"`
	Title			string `json:"Title"`
	Participants	string `json:"Participants"`
	Start_Time	string `json:"Start_Time"`
	End_Time string `json:"End_Time"`
	Creation_Timestamp string `json:"Creation_Timestamp"`
}

type Participant struct{

	Name      string `json:"Name"`
	Email	string `json:"Email"`
	RSVP 	string `json:"RSVP"`
	
}


var Meetings []Meeting

var Participants []Participant

var a string

func Handler(w http.ResponseWriter, r *http.Request) {  
    fmt.Printf(w,"Req: %s %s", r.URL.Host, r.URL.Path)
    return r.URL.Host
}


func returnmeeting(w http.ResponseWriter, r *http.Request){
	//a =Handler(w http.ResponseWriter, r *http.Request)
    fmt.Println("Endpoint Hit: returnmeeting")
    json.NewEncoder(w).Encode(Meetings) //Meetings[1]
}

/*func getHost(r *http.Request) string {
    if r.URL.IsAbs() {
        host := r.Host
        // Slice off any port information.
        if i := strings.Index(host, "/"); i != -1 {
            host = host[:i]
        }
        return host
    }
    return r.URL.Host
} 
*/
func homePage(w http.ResponseWriter, r *http.Request){	

    fmt.Fprintf(w, "Welcome to the HomePage!")
    //fmt.Fprintf("%s",r.HOST)
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/meeting",Handler)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
   	Meetings = []Meeting{
        Meeting{Id: "1", Title: "Goal 1", Participants: "a,b,e", Start_Time: "10pm",End_Time:"11pm",Creation_Timestamp:"11/10/2020 11am"},
        Meeting{Id: "2", Title: "Goal 2", Participants: "a,c", Start_Time: "1pm",End_Time:"2pm",Creation_Timestamp:"12/10/2020 11am"},
        Meeting{Id: "3", Title: "Goal 3", Participants: "e", Start_Time: "2pm",End_Time:"3pm",Creation_Timestamp:"06/10/2020 11am"},

    }
    Participants = []Participant{
        Participant{Name: "a", Email: "a@sid.com", RSVP: "yes"},
        Participant{Name: "b", Email: "b@sid.com", RSVP: "no"},
        Participant{Name: "c", Email: "c@sid.com", RSVP: "yes"},
        Participant{Name: "e", Email: "e@sid.com", RSVP: "no resp"},
    }
    
    handleRequests()
}
