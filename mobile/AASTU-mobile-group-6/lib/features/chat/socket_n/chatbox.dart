import 'package:socket_io_client/socket_io_client.dart' as IO;

class SocketService {
  late IO.Socket socket;

  void connectToServer() {
    socket = IO.io('https://g5-flutter-learning-path-be.onrender.com/', <String, dynamic>{
      'transports': ['websocket'],
      'autoConnect': false,
    });

    socket.connect();

    // Event listeners
    socket.onConnect((_) {
      print('Connected to server');
    });

    socket.onDisconnect((_) {
      print('Disconnected from server');
    });

  }

  void sendMessage(String message) {
    
    socket.emit('message', message);
  }

  void dispose() {
    socket.dispose();
  }
}
