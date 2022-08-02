# Threadbare

## Crawling Collections for Material Culture 

[Jason A. Heppler](https://jasonheppler.org), [Roy Rosenzweig Center for History and New Media](https://rrchnm.org), George Mason University

### About this repository

This repository contains code for a web crawler for the [Connecting Threads](https://connectingthreads.co.uk) project co-directed by Meha Priyadarshini and Deepthi Murali. *Threadbare* is a play on words: by revealing (making *bare*) the metadata for the *Threads* project. 

Connecting Threads is a born-digital project on the use of Indian and
Indian-imitation textiles by the African diaspora in the Americas in the
eighteenth and nineteenth centuries. Knitting together material from three collections — Victoria and Albert Museum, London, University of Glasgow Archives & Special Collections, and the Smithsonian Cooper Hewitt Design Museum, New York — the project re-centers the locus of global textile trade from White Euro-American markets to African diaspora markets.

### License

 All code is copyrighted &copy; 2022 Jason Heppler. Code is licensed [CC0 1.0 Universal](https://github.com/chnm/threadbare/blob/main/LICENSE).

## API Response

For checking the values for consuming the API, the following will help: 

```bash
curl -X GET "https://api.collection.cooperhewitt.org/rest/?method=cooperhewitt.exhibitions.getObjects?access_token=$THREADBARE_KEY&query=India%20textiles" \ 
  -H "Accept: application/json" | jq
```
