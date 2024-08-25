import 'dart:convert';
import 'dart:developer';

import 'package:e_commerce_app/core/constants/constants.dart';

import '../../../../../core/failure/exception.dart';
import '../../models/chat_model.dart';
import '../../models/message_model.dart';
import 'remote_contrats.dart';
import 'package:http/http.dart' as http;
import 'stream_socket.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;
class ChatRemoteDataSourceImpl extends ChatRemoteDataSource {
  final http.Client client;
  final String _baseUrl;

  StreamSocket streamSocket = StreamSocket();
 // Replace 'YOUR_ACCESS_TOKEN' with the actual token you need
  final socket = IO.io(
    Urls.socketUrl,
    IO.OptionBuilder()
      .setTransports(['websocket']) 
      .disableAutoConnect() 
      .setExtraHeaders({
        'Authorization': 'Bearer YOUR_ACCESS_TOKEN', 
      })
      .build(),
  );
  ChatRemoteDataSourceImpl({
    required this.client,
  }) : _baseUrl = '${Urls.baseUrl}/chats';

  @override
  Future<bool> deleteChatById(String chatId) async {
    try {
      final response =
          await client.delete(Uri.parse('${Urls.baseUrl}/$chatId'));

      if (response.statusCode != 200) {
        print(response.body);
        // throw ServerException();
        return false;
      }
      else {
        return true;
      }
    } catch (e) {
      print(e);
      throw ServerException();
    }
  }

  @override
  Stream<MessageModel> getChatById(String chatId) {
    streamSocket.dispose();
    streamSocket = StreamSocket();

    client.get(Uri.parse('$_baseUrl/$chatId/messages')).then((response) {
      if (response.statusCode == 200) {
        final List<dynamic> messages = jsonDecode(response.body)['data'];

        for (var message in messages) {
          streamSocket.addResponse(MessageModel.fromJson(message));
        }
      } else {
        print(response.body);
        throw ServerException();
      }
    });

    socket.connect();

    socket.onConnect((_) {
      log('Connected to the socket server');
    });

    socket.onDisconnect((_) {
      log('Disconnected from the socket server');
    });

    socket.on('message:delivered', (data) {
      MessageModel message = MessageModel.fromJson(data);
      streamSocket.addResponse(message);
    });

    socket.on('message:received', (data) {
      MessageModel message = MessageModel.fromJson(data);
      streamSocket.addResponse(message);
    });

    return streamSocket.getResponse;
  }

  @override
  Future<ChatModel> createChatById(String sellerId) async {
    try {
      // final response = await client.post(_baseUrl, {
      //   'userId': receiver.id,
      // });
      final response = await client.post(Uri.parse(Urls.baseUrl), body: {
        'userId': sellerId,
      });

      if (response.statusCode == 200) {
        return ChatModel.fromJson(jsonDecode(response.body)['data']);
      } else {
        print(response.body); 
        throw ServerException();
      }
    } catch (e) {
      print(e);
      throw ServerException();
    }
  }

  @override
  Future<List<ChatModel>> getAllChat() async {
    try {
      final response = await client.get(Uri.parse(Urls.baseUrl));

      if (response.statusCode == 200) {
        final List<dynamic> chats = jsonDecode(response.body)['data'];
        return chats.map((e) => ChatModel.fromJson(e)).toList();
      } else {
        print(response.body);
        throw ServerException();
      }
    } catch (e) {
      print(e);
      throw ServerException();
    }
  }

  @override
  Future<bool> sendMessage(String chat, String message, String type) async{
    try{
    socket.emit('message:send', {
      'chatId': chat,
      'content': message,
      'type': type,
    }
    
    );
    return true;
    
    }catch(e){
      print(e);
      throw ServerException();
    }
  }
}
