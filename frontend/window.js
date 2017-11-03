$(() => {
  
    $('#text-input').bind('input keyup', (e) => {
        if (e.keyCode == 13) {
            const text = $('#text-input').val()
            $('#result').text("searched "+text)
        }
  
    })
  
    $('#text-input').focus() // focus input box
  })
