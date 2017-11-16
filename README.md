TODO
-----
   - local
      - add symspell to search (fuzzy)
      - save keywords in AVL with hash key (SHA1), electron - sha1(keyword) before searching (on server /search/)
   - search tunnel
      - nginx reverse proxy for go webserver
      - serve static files from electron
      - modify static files address and port to access proxy
   - P2P
      - function merge 2 AVL trees
      - DHT
      - replace storage with AVL
   - refactor
      - change AVL serialization (http://www.geeksforgeeks.org/serialize-deserialize-binary-tree/, https://www.cs.usfca.edu/~brooks/S04classes/cs245/lectures/lecture11.pdf)
      - check for duplicate links in Crawler
      - merge different forms of the same word ('work', 'works', 'worked', etc.)
      - separate repos - crawler, indexer, storage (AVL)
      - encrypt peer transfer
      - add tests for crawler, indexer, storage

Docs
-----
   - http://www.yacy-websearch.net/wiki/index.php/En:FAQ
   - http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-1
   - http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-2
   - http://cs.brown.edu/courses/cs138/s17/syllabus.html
   - http://cs.brown.edu/courses/cs138/s17/content/projects/chord.pdf
   - https://medium.com/@sent0hil/consistent-hashing-a-guide-go-implementation-fe3421ac3e8f
   - https://www.slideshare.net/jsimnz/chord-dht
   - https://blog.savoirfairelinux.com/en-ca/2015/ring-opendht-a-distributed-hash-table/
   - http://infolab.stanford.edu/~backrub/google.html
   - https://moz.com/blog/search-engine-algorithm-basics
   - http://www.ardendertat.com/2011/05/30/how-to-implement-a-search-engine-part-1-create-index/
   - https://www.elastic.co/guide/en/elasticsearch/guide/current/inverted-index.html
   - https://en.m.wikipedia.org/wiki/Search_engine_indexing
   - https://www.google.ro/amp/s/www.maketecheasier.com/how-bittorrent-dht-peer-discovery-works/amp/
   - https://github.com/r-medina/gmaj
   - https://github.com/prettymuchbryce/kademlia
   - https://github.com/nictuku/dht

Docs headless browser
-----
   - https://github.com/k4s/phantomgo
   - https://github.com/PuerkitoBio/goquery - misses functions to get element attributes

Docs src
-----
   - https://github.com/automenta/kelondro
   - https://github.com/yacy/yacy_search_server/blob/8303e15419e789cad94b94a1d65e00f9627cd5f1/source/net/yacy/search/query/SearchEvent.java
   - https://github.com/yacy/yacy_search_server/blob/dd9cb06d250d8bbfc798c23ab8779a92018557f1/source/net/yacy/kelondro/data/word/WordReferenceVars.java


Use
-----
   - DHT (Chord ???)
   - symspell - handle search errors (https://github.com/heartszhang/symspell, https://github.com/sajari/fuzzy)
   - crawler
   - indexer (Bleve, https://github.com/nassor/studies-blevesearch ???)
   - storage: AVL trees

Sell points
-----
   - concurrent crawler
   - binary protocol (custom, protobuf ???)
   - distributed
   - modular
   - extensibile
   - separate algoritm to be easy updateable/replaceable
   - support dark web
   - portal to web and easy to make personal portals (tor also)
   - works even if much of the internet infrastructure is down (depends as little as possible on other protocols or infrastructure)
   - support ML - distributed - works like flink/storm ( stream processing ??? )
   - resistent to node-loss
   - resistent to hacking
   - algorithms for search, ranking - voted by consensus mechanism
   - platform - possible to build something else than search (see ethereum)
   - index other protocols, not just http - ftp, samba, torrent?, etc.
   - security - a node cant do much harm (inside threat) - it can only see what was searched but other data like IP and other identification must be masked
   - limit available resources for software
   - option for backend to run on raspberry and access it from there (with a browser)
   - every node selects what resources it uses, what protocols it indexes and what limits it has (CPU, RAM, internet/intranet/dark)
   - must work on many environments (linux, osx, win, mobile???)
   - easy to use

Build
-----
   - frontend: electron-packager .
   - backend: go build main.go
