import 'dart:convert';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class SocketIoManager {
  late IO.Socket _socket;

  SocketIoManager({
    required String serverUrl,
    required String token,
    String nameSpace = '/',
    required Function(dynamic) onConnect,
    required Function(dynamic) onDisconnect,
    required Function(dynamic) onError,
    required Function(dynamic) onMessageDelivered,
    required Function(dynamic) onMessageReceived,
  }) {
   
    _socket = IO.io(serverUrl, IO.OptionBuilder()
        .setTransports(['websocket']) 
        .setPath(nameSpace) 
        .setQuery({'token': token}) 
        .enableAutoConnect()
        .setExtraHeaders({'Authorization': 'Bearer $token'}) 
        .build());

    // Listening for connection events
    _socket.onConnect(onConnect ?? (_) => print('Connected'));
    _socket.onDisconnect(onDisconnect ?? (_) => print('Disconnected'));
    _socket.onError(onError ?? (data) => print('Socket error: $data'));

    // Listening for specific server events
    _socket.on('message:delivered', onMessageDelivered ?? (data) {
      print('Message delivered: $data');
    });

    _socket.on('message:received', onMessageReceived ?? (data) {
      print('Message received: $data');
    });
  }

 
  void init() => _socket.connect();

  
  void subscribe(String event, Function(dynamic) callback) {
    _socket.on(event, callback);
  }

  void sendMessage(String event, Map<String, dynamic> message) {
    _socket.emit(event, json.encode(message));
  }

  void disconnect() {
    _socket.disconnect();
  }
}
