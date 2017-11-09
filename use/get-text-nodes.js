function textNodesUnder(el){
  var n, a=[], walk=document.createTreeWalker(el,NodeFilter.SHOW_TEXT,null,false);
  while(n=walk.nextNode()) a.push(n);
  return a;
}

function textNodesUnder(node){
  var all = [];
  for (node=node.firstChild;node;node=node.nextSibling){
    if (node.nodeType==3) all.push(node);
    else all = all.concat(textNodesUnder(node));
  }
  return all;
}

function nativeTreeWalker() {
    var walker = document.createTreeWalker(
        document.body, 
        NodeFilter.SHOW_TEXT, 
        null, 
        false
    );

    var node;
    var textNodes = [];

    while(node = walker.nextNode()) {
        textNodes.push(node.nodeValue);
    }
}

function recurse(element)
{
    if (element.childNodes.length > 0) 
        for (var i = 0; i < element.childNodes.length; i++) 
            recurse(element.childNodes[i]);

    if (element.nodeType == Node.TEXT_NODE && /\S/.test(element.nodeValue)) 
        doSomething(element);
}
var html = document.getElementsByTagName('html')[0];
recurse(html);


function getTextNodesIn(elem, opt_fnFilter) {
  var textNodes = [];
  if (elem) {
    for (var nodes = elem.childNodes, i = nodes.length; i--;) {
      var node = nodes[i], nodeType = node.nodeType;
      if (nodeType == 3) {
        if (!opt_fnFilter || opt_fnFilter(node, elem)) {
          textNodes.push(node);
        }
      }
      else if (nodeType == 1 || nodeType == 9 || nodeType == 11) {
        textNodes = textNodes.concat(getTextNodesIn(node, opt_fnFilter));
      }
    }
  }
  return textNodes;
}
var menu = document.getElementById('divMenu');
getTextNodesIn(menu, function(textNode, parent) {
  if (/^\s+$/.test(textNode.nodeValue)) {
    parent.removeChild(textNode);
  }
});

