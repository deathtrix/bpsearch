TODO
-----
   - basic skeleton
      - cum pasez url catre indexer ???
      - struct indexat tree[keyword] = {url1:weight1, url2:weight2} - map[string]int - struct idx {url string, weight int}, []idx ???
      
      - save keywords in AVL with hash key (SHA1)
      - update frontend - show results from Storage (AVL)
      - add symspell to search (fuzzy)
   - indexer
      - add settings page to frontend (split by categories), save settings to disk - indexing weights, etc.
      - refactor indexing algorithm (use trees of tokens and weight rules), CSS ???
   - P2P
      - function merge 2 AVL trees
      - DHT - sharing indexes (hash->ip sau hash->urls)
      - peer protocol (functions: get keyword(s), dump)
      - encrypt peer transfer
   - refactor
      - check for duplicate links in Crawler
      - change AVL serialization (http://www.geeksforgeeks.org/serialize-deserialize-binary-tree/, https://www.cs.usfca.edu/~brooks/S04classes/cs245/lectures/lecture11.pdf)
      - separate repos - crawler, indexer, storage (AVL)

Docs
-----
   - http://www.yacy-websearch.net/wiki/index.php/En:FAQ
   - http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-1
   - http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-2
   - https://github.com/armon/go-chord
   - http://cs.brown.edu/courses/cs138/s17/syllabus.html
   - https://medium.com/@sent0hil/consistent-hashing-a-guide-go-implementation-fe3421ac3e8f
   - http://blog.notdot.net/2009/11/Implementing-a-DHT-in-Go-part-1
   - http://cs.brown.edu/courses/cs138/s17/content/projects/chord.pdf
   - https://www.slideshare.net/jsimnz/chord-dht

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
