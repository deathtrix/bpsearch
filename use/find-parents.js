$( "li.item-a" ).parents()

function findUpTag(el, tag) {
    while (el.parentNode) {
        el = el.parentNode;
        if (el.tagName === tag)
            return el;
    }
    return null;
}

var a = document.getElementById("target");
var els = [];
while (a) {
    els.unshift(a);
    a = a.parentNode;
}
