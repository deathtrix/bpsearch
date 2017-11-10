var page = require('webpage').create(),
    system = require('system'), address;
page.settings.userAgent = 'Mozilla/5.0+(compatible;+Baiduspider/2.0;++http://www.baidu.com/search/spider.html)';
phantom.cookiesEnabled = true;

page.open('<<URL>>', function (status) {
  page.includeJs('https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js',
    function() {
      var data = page.evaluate(function() {
        function textNodesUnder(el) {
          var n, a = [], walk = document.createTreeWalker(el, NodeFilter.SHOW_TEXT, { acceptNode: function(node) {
            if ( ! /^\s*$/.test(node.data) && 
               (node.parentNode.nodeName !== 'SCRIPT') &&
               (node.parentNode.nodeName !== 'STYLE') ) {
                  return NodeFilter.FILTER_ACCEPT;
            }
          }
        }, false);
          while (n = walk.nextNode()) a.push(n);
          return a;
        }

        function findParents(el) {
            var els = [];
            el = el.parentNode;
            while (el) {
              els.unshift(el);
              el = el.parentNode;
            }
            els.shift();
            return els;
        }

        function calculate(data, formula) {
          return eval(formula);
        }

        // TODO: remove if not used
        function onlyUnique(value, index, self) { 
          return self.indexOf(value) === index;
        }

        function each(obj, callback) {
          var length, i = 0;
        
          for ( i in obj ) {
            if ( callback.call( obj[ i ], i, obj[ i ] ) === false ) {
              break;
            }
          }
        
          return obj;
        }
      
        function indexElements() {
          var elems = [];
          var els = textNodesUnder(document.getElementsByTagName('body')[0]);
          els.forEach(function (e) {
            var elem = {};
            var parents = [];
            // $(e).parents().each(function(i2, e2) {  // using jQuery
            findParents(e).forEach(function(e2) {
              parents.push(e2.tagName);
            });

            var uniqueWords = {};
            var text = e.textContent.trim();
            // TODO: more complex regex for symbol removal
            text = text.replace(/[^0-9a-zA-Z \/\\\.@]/g, "");
            var words = text.split(' ');
            // TODO - remove empty string - words.splice(words.indexOf(''), 1);
            words.forEach(function (word) {
              word = word.toLowerCase();
              if (typeof uniqueWords[word] === "undefined") {
                uniqueWords[word] = 1;
              } else {
                uniqueWords[word]++;
              }
            });

            var cssObj = window.getComputedStyle(e.parentNode, null);

            elem.parents = parents;
            elem.words = uniqueWords;
            elem.sz = parseFloat(cssObj.getPropertyValue('font-size'));
            elem.wg = cssObj.getPropertyValue('font-weight');

            elems.push(elem);
          });

          return elems;
        }

        // TODO - add size, weight and parents to formula
        function calcScores(elems) {
          var uniqueWords = {};

          elems.forEach(function (el) {
            each(el.words, function (w, c) {
              var score = c;
              if (typeof uniqueWords[w] === "undefined") {
                uniqueWords[w] = c;
              } else {
                uniqueWords[w] += score;
              }
            });
          });

          return uniqueWords;
        }

        var elems = indexElements();
        var scores = calcScores(elems);
        
        return JSON.stringify(scores);
      });
      
      console.log(data);
      phantom.exit();
    }
  );
});
