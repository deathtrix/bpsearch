var page = require('webpage').create(),
    system = require('system'), address;
page.settings.userAgent = 'Mozilla/5.0+(compatible;+Baiduspider/2.0;++http://www.baidu.com/search/spider.html)';
phantom.cookiesEnabled = true;

page.open('<<URL>>', function (status) {
  page.includeJs('https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js',
    function() {
      var data = page.evaluate(function() {
        // var el = document.getElementById('main');
        // var el = document.getElementsByTagName('h3')[0];
        // var cells = document.querySelectorAll('.tpo tr tr')[4].querySelectorAll('td');

        // parse all elements
        // $('*').each(function (i, el) {
        //   if ((el.innerText != '') && (!$(el).is('script')) && (!$(el).is('style'))) {
        //     console.log(el.innerText);
        //   }
        // });
        
        var el = $('#main')[0];
        var cssObj = window.getComputedStyle(el, null);
        var size = cssObj.getPropertyValue('font-size');
        var weight = cssObj.getPropertyValue('font-weight');
        var fontSize = parseFloat(size);
        
        // var txt;
        // for (i = 0; i < cssObj.length; i++) { 
        //   cssObjProp = cssObj.item(i)
        //   txt += cssObjProp + " = " + cssObj.getPropertyValue(cssObjProp) + "\n";
        // }

        return el.innerText;
      });
      
      console.log(data);
      phantom.exit();
    }
  );
});
