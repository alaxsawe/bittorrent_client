package main

import(
  bencode "code.google.com/p/bencode-go"
  "bytes"
  "crypto/sha1"
  "fmt"
  "net/http"
  "net/url"
	"os"
  "log"
)

type MetaInfo struct {
  Info         map[string]interface {}
  InfoHash     string
  Announce     string
  /* AnnounceList []interface {} "announce-list" */
  CreationDate string     "creation date"
  Comment      string
  CreatedBy    string "created by"
  Encoding     string
}

func parseTorrentFile(filename string) (MetaInfo){
  file, err := os.Open(filename) // For read access.
  if err != nil {
    log.Fatal(err)
  }
  data, err := bencode.Decode(file)
  if err != nil {
    log.Fatal(err)
  }
  hash, ok := data.(map[string]interface{})
  if !ok {
    log.Fatal(err)
  }

  var b bytes.Buffer
  err = bencode.Marshal(&b, hash["info"])
  if err != nil {
    log.Fatal(err)
  }

  infoHash := sha1.New()
  infoHash.Write(b.Bytes())

  return MetaInfo { 
    hash["info"].(map[string]interface {}),
    string(infoHash.Sum(nil)),
    getString(hash, "announce"),
    /* hash["announce-list"].([]interface {}), */
    getString(hash, "creation date"),
    getString(hash, "comment"),
    getString(hash, "created by"),
    getString(hash, "encoding"),
  }
}

func queryTracker(info MetaInfo) {
  u, _ := url.Parse(info.Announce)
  uq := u.Query()
  uq.Add("info_hash", info.InfoHash)
  uq.Add("peer_id", peerId())
  uq.Add("port", "6881")
  uq.Add("uploaded", "0")
  uq.Add("downloaded", "0")
  uq.Add("left", "100")
  uq.Add("compact", "1")
  u.RawQuery = uq.Encode()
  resp, _ := http.Get(u.String())
	data, _ := bencode.Decode(resp.Body)
  hash, ok := data.(map[string]interface{})
  if !ok {
    log.Fatal("Well Ill be")
  }
	tr := TrackerResponse{
		getString(hash, "interval"),
		getString(hash, "min interval"),
		getString(hash, "peers"),
	}
  fmt.Printf("response: %v\n\n", tr)
  fmt.Printf("peers: %v\n\n", tr.GetPeerAddresses())
}
