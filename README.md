# `ghanamps` A tiny web scraper that returns a JSON of all the members of parliament in Ghana.

### LIVE WEB VERSION
You can simply make a request to this URL to get the JSON:  
[https://ghanamps.herokuapp.com](https://ghanamps.herokuapp.com)  


### INSTALLATION (With Go already installed)
If you already have [Go](https://go.dev/) on your machine you can simply install this tool by running: 
```bash
go install github.com/yeboahnanaosei/ghanamps/cmd/ghanamps@latest
```
This will install a binary `ghanamps` to the `bin` folder in your `$GOPATH`. You can
then run the tool by simply issuing a command from your terminal
```bash
ghanamps
```
This will return a JSON array of all the current members of Ghana's parliament. Example below:
```json
[
    {
     "name": "ADAMA  SULEMANA",
     "party": "NDC",
     "constituency": "TAIN CONSTITUENCY",
     "region": "BONO REGION",
     "photo": "https://www.parliament.gh/epanel/mps/photos/198.jpg",
     "profile": "https://www.parliament.gh/mps?mp=198"
   },
   {
     "name": "ADELAIDE  NTIM",
     "party": "NPP",
     "constituency": "NSUTA/KWAMAN BEPOSO CONSTITUENCY",
     "region": "ASHANTI REGION",
     "photo": "https://www.parliament.gh/epanel/mps/photos/164.jpg",
     "profile": "https://www.parliament.gh/mps?mp=164"
   },
   {
     "name": "AGNES NAA MOMO LARTEY",
     "party": "NDC",
     "constituency": "KROWOR CONSTITUENCY",
     "region": "GREATER ACCRA REGION",
     "photo": "https://www.parliament.gh/epanel/mps/photos/64.jpg",
     "profile": "https://www.parliament.gh/mps?mp=64"
   },
   {
     "name": "AKWASI  DARKO BOATENG",
     "party": "NPP",
     "constituency": "BOSOME FREHO CONSTITUENCY",
     "region": "ASHANTI REGION",
     "photo": "https://www.parliament.gh/epanel/mps/photos/133.jpg",
     "profile": "https://www.parliament.gh/mps?mp=133"
   },
   {
     "name": "AKWASI  KONADU",
     "party": "NPP",
     "constituency": "MANHYIA NORTH CONSTITUENCY",
     "region": "ASHANTI REGION",
     "photo": "https://www.parliament.gh/epanel/mps/photos/145.jpg",
     "profile": "https://www.parliament.gh/mps?mp=145"
   }
   ...
]
```

### INSTALLATION (Without Go installed)
1. Go to the [releases page](https://github.com/yeboahnanaosei/ghanamps/releases) and download the appropriate binary for your system.
2. Extract the binary from the zip archive
3. Put the binary in your PATH environment variable.
    > If you don't know how to do this, here is a useful link to guide you on how to do it: [https://zwbetz.com/how-to-add-a-binary-to-your-path-on-macos-linux-windows/](https://zwbetz.com/how-to-add-a-binary-to-your-path-on-macos-linux-windows/)  

4. Launch your command prompt and run `ghanamps`
