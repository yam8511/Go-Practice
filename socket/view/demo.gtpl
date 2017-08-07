<!doctype html>
<html>
  <head>
    <title>Socket.IO chat</title>
    <style>
      * { margin: 0; padding: 0; box-sizing: border-box; }
      body { font: 13px Helvetica, Arial; }
      form { background: #000; padding: 3px; position: fixed; bottom: 0; width: 100%; }
      form input { border: 0; padding: 10px; width: 90%; margin-right: .5%; }
      form button { width: 9%; background: rgb(130, 224, 255); border: none; padding: 10px; }
      #messages { list-style-type: none; margin: 0; padding: 0; }
      #messages li { padding: 5px 10px; }
      #messages li:nth-child(odd) { background: #eee; }
    </style>
  </head>
  <body>
    <h3>{{.}}</h3>
    <ul id="messages"></ul>
    <form>
      <input id="m" autofocus autocomplete="off" /><button>Send</button>
    </form>
    <script src="/asset/socket.io-1.3.7.js"></script>
    <script src="/asset/jquery-1.11.1.js"></script>
    <script>
    $(document).ready(function () {
      var socket = io();
      $('form').submit(function(){
        // socket.emit('chat message with ack', $('#m').val(), function(data){
        //   $('#messages').append($('<li>').text('ACK CALLBACK: ' + data));
        // });
        socket.emit('chat message', $('#m').val());
        $('#m').val('');
        return false;
      });
      socket.on('chat message', function(msg){
        $('#messages').prepend($('<li>').text(msg));
      });
    });
    </script>
  </body>
</html>
