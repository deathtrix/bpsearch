$(() => {
  
    $('#text-input').bind('input keyup', (e) => {
        if (e.keyCode == 13) {
            const text = $('#text-input').val()
            $.get("http://localhost:3333/search/?keywords="+text, function(data, status) {
                results = JSON.parse(data);
                $('#result').html("");
                for (i = 0; i < results.length; i++) {
                    var display = '<div class="row"><div class="col-xs-1 col-sm-1 col-md-1 col-lg-1">&nbsp;</div><div class="col-xs-11 col-sm-11 col-md-11 col-lg-11"><a href="'+results[i]+'">'+results[i]+'</a></div></div>';
                    $('#result').append(display);
                }
            });
        }
    })
  
    $('#text-input').focus() // focus input box
  })
