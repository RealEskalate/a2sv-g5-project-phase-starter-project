import 'package:socket_io_client/socket_io_client.dart' as IO;

import '../../../../core/constants/constants.dart';
import '../../../auth/data/data_source/auth_local_data_source.dart';

abstract class SocketIOService {
  IO.Socket get socket;
  void connect();
  void disconnect();
}

class SocketIOServiceImpl implements SocketIOService {
  late IO.Socket socket;
  final AuthLocalDataSource authLocalDataSource;

  SocketIOServiceImpl({required this.authLocalDataSource}) {
    socket = IO.io(AppData.chatserver, <String, dynamic>{
      'transports': ['websocket'],
      'autoConnect': false,
      'extraHeaders': {
        'Authorization': 'Bearer ${authLocalDataSource.getToken()}'
      }
    });
  }

  @override
  void connect() {
    socket.connect();
  }

  @override
  void disconnect() {
    socket.disconnect();
  }
}
