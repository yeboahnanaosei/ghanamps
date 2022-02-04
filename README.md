# `ghanamps` A simple cli and server that returns a JSON of all the members of parliament in Ghana.

### LIVE WEB VERSION
You can simply make a request to this URL to get the JSON:  
[https://ghanamps.herokuapp.com](https://ghanamps.herokuapp.com)  


### INSTALLATION (With Go already installed)
If you already have [Go](https://go.dev/) on your machine you can simply install this tool by running: 
```bash
go install github.com/yeboahnanaosei/ghanamps/cmd/ghanamps@latest
```
This will install a binary `ghanamps` to the `bin` folder in your `$GOPATH`.  

The tool has two subcommands `members` for all members and `leaders` for all the leaders of parliament.  

### GET MEMBERS
To get members you can run the tool by simply issuing the following command from your terminal.
```bash
ghanamps members
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

To get members of a particular party, you can issue the following command:
```bash
ghanamps members -party npp
```

```bash
ghanamps members -party ndc
```

```bash
ghanamps members -party independent
```

### GET LEADERS
To get the leadership of parliament, you can run the tool by simply issuing the following command from your terminal:
```bash
ghanamps leaders
```
This will return a JSON array of all the current leaders of Ghana's parliament. Example below:
```json
[
   {
     "name": "Rt. Hon. Alban Sumana Kingsford Bagbin",
     "title": "Speaker of Parliament",
     "photo": "https://www.parliament.gh/epanel/mps/leaders/MrSpeaker.jpg"
   },
   {
     "name": "Hon. Joseph Osei-Owusu",
     "title": "First Deputy Speaker",
     "photo": "https://www.parliament.gh/epanel/mps/photos/132.jpg"
   },
   {
     "name": "Hon. Andrew Asiamah Amoako",
     "title": "Second Deputy Speaker",
     "photo": "https://www.parliament.gh/epanel/mps/photos/128.jpg"
   },
   ...
]
```

### INSTALLATION (Without Go installed)
1. Go to the [releases page](https://github.com/yeboahnanaosei/ghanamps/releases) and download the appropriate binary for your system.
1. Extract the binary from the zip archive
1. Put the binary in your PATH environment variable and make sure it is executable (Linux & macOS)
    > If you don't know how to do this, here is a useful link to guide you: [https://zwbetz.com/how-to-add-a-binary-to-your-path-on-macos-linux-windows/](https://zwbetz.com/how-to-add-a-binary-to-your-path-on-macos-linux-windows/)  

4. Launch your command prompt and run:
> 1. `ghanamps members` for all members
> 2. `ghanamps leaders` for all leaders
> 3. `ghanamps -h` for help


## Want to contribute?
Of course! [Read this](https://github.com/yeboahnanaosei/ghanamps/blob/master/CONTRIBUTING.md)
