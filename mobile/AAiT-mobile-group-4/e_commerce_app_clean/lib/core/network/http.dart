import 'package:http/http.dart';
import 'package:socket_io_client/socket_io_client.dart' as io;

class CustomHttp extends BaseClient {
  final Client client;
  late io.Socket socket;
  final Map<String, String> header = {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  };

  set setAuthToken(String token) {
    header['Authorization'] = 'Bearer $token';
    socket = io.io('https://g5-flutter-learning-path-be.onrender.com', <String, dynamic>{
      'transports': ['websocket'],
      'autoConnect': false,
      'extraHeaders': {'authorization': 'Bearer $token'},
    });
  }
  CustomHttp({required this.client,});

  @override
  Future<StreamedResponse> send(BaseRequest request) async {
    request.headers.addAll(header);
    return client.send(request);
  }
}