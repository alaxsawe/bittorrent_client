package main

import(
  "fmt"
  "github.com/codegangsta/cli"
  "os"
)


func getTorrent(filename string) { 
  metaInfo := parseTorrentFile(filename)
  fmt.Printf("announce: %v\n", metaInfo.Announce)
  /* fmt.Printf("announce-list: %v\n", metaInfo.AnnounceList) */
  fmt.Printf("creation date %v\n", metaInfo.CreationDate)
  fmt.Printf("comment: %v\n", metaInfo.Comment)
  fmt.Printf("createdBy: %v\n", metaInfo.CreatedBy)
  fmt.Printf("encoding: %v\n", metaInfo.Encoding)
  fmt.Printf("Info Hash: %v\n", metaInfo.InfoHash)
  /* fmt.Printf("Info: %v\n", metaInfo.Info) */

  queryTracker(metaInfo)
}

func main() {
  app := cli.NewApp()
  app.Name = "bittorrent_client"
  app.Usage = "Get cool stuff"
  app.Flags = []cli.Flag {
    cli.StringFlag{"file, f", "filename", "filename"},
  }
  app.Action = func(c *cli.Context) {
    fmt.Println("Lets torrent")
    getTorrent(c.String("file"))
  }
  app.Run(os.Args)
}

