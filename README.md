TODO
-----
   - skeleton
      - indexer - get formula as parameter from go, save to storage
      - add settings page to frontend (split by categories), save settings to disk - indexing weights, etc.
      - save keywords in AVL with hash key (SHA1)
      - update frontend - show results from Storage (AVL)
      - add symspell to search (fuzzy)
   - P2P
      - function merge 2 AVL trees
      - DHT - sharing indexes (hash->ip sau hash->urls)
      - peer protocol (functions: get keyword(s), dump)
      - encrypt peer transfer
   - refactor
      - change AVL serialization (http://www.geeksforgeeks.org/serialize-deserialize-binary-tree/, https://www.cs.usfca.edu/~brooks/S04classes/cs245/lectures/lecture11.pdf)
      - use Indexer struct instead of map
      - check for duplicate links in Crawler
      - separate repos - crawler, indexer, storage (AVL)
      - add tests for crawler, indexer, storage
      - merge different forms of the same word ('work', 'works', 'worked', etc.)

Docs
-----
   - http://www.yacy-websearch.net/wiki/index.php/En:FAQ
   - http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-1
   - http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-2
   - https://github.com/armon/go-chord
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
