import 'dart:io';

import 'package:socket_io_client/socket_io_client.dart' as IO;

class SocketManager {
  late IO.Socket socket;

  SocketManager() {
    initializeSocket();
  }
  
  void initializeSocket() {
    socket = IO.io('https://g5-flutter-learning-path-be.onrender.com/');

    socket.connect();

    socket.onConnect((_){
      print('Connect to network');

    });

    socket.onDisconnect((_) {
      print('disconnected');
    });

    socket.onError((error) {
      print('error: $error');
    });


  }


IO.Socket get s => socket;

  void receiveMessage(Function(Map <String, dynamic>) callback) {
    socket.on('receved_message', (data) {
      callback(data);
    });
    
  }

}