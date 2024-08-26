import 'dart:convert';

import 'package:ecommerce_app_ca_tdd/core/constants/constants.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/data_sources/socket/stream.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/models/message_model.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:http/http.dart' as http;
import 'package:ecommerce_app_ca_tdd/features/chat/socket_n/chatbox.dart';
import 'package:shared_preferences/shared_preferences.dart';

abstract class ChatRemoteDataSource {

  Future <ChatEntity> getMyChatById(String id);
  Future<ChatEntity> initiateChat(String userId);
  Future<String> deleteChat(String chatId);
  Future<List<Message>> getChatMessages(String chatId);
  Future<List<ChatEntity>> getAllChats();
  void sendMessage(String chat, String message, String type);
 
}
class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final http.Client client;

  StreamSocket streamSocket = StreamSocket();
  SocketService socket = SocketService();

  ChatRemoteDataSourceImpl({required this.client});

  @override
  Future<String> deleteChat(String chatId) async{
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');
    var head = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',      
    };
    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
    final response = await client.delete(Uri.parse(url+'/$chatId'),headers: head);
    print(response.statusCode);
    if (response.statusCode == 200) {
      return 'Product deleted successfully';
    } else {
      throw Exception('Failed to delete product');
    }
  }
  @override
  Future<List<ChatEntity>> getAllChats() async{
    
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');
    var head = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',      
    };

    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
    final response = await client.get(Uri.parse(url),headers: head);
    print(response.statusCode);
    if (response.statusCode == 200) {
     return (jsonDecode(response.body)['data'] as List)
          .map((e) => ChatEntity.fromJson(e))
          .toList();
    } else {
      throw Exception('Failed to Fetch Chat');
    }
    
  }
  @override
  Future<ChatEntity> getMyChatById(String id) async{
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');
    var head = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',      
    };
    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
    final response = await client.get(Uri.parse(url+'/$id'),headers: head);
    print(response.statusCode);
    if (response.statusCode == 200) {
      return ChatEntity.fromJson(jsonDecode(response.body)['data']);
    } else {
      throw Exception('Failed to load data');
    }
  }
  @override
  Future<List<Message>> getChatMessages(String chatId) async{
    streamSocket.dispose();
    streamSocket = StreamSocket();
    
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');
    var head = {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',      
    };
    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
    final response = await client.get(Uri.parse(url+'/$chatId/messages'),headers: head);

    print(response.statusCode);
    if (response.statusCode == 200) {
      final List<dynamic> messages = jsonDecode(response.body)['data'];
      for (var message in messages) {
          streamSocket.addResponse(MessageModel.fromJson(message));
        }
      return (jsonDecode(response.body)['data'] as List)
          .map((e) => Message.fromJson(e))
          .toList();
    } else {
      throw Exception('Failed to Fetch Chat');
    }
  }
  @override
  Future<ChatEntity> initiateChat(String userId) async{
    var temp = await SharedPreferences.getInstance();
    var temp2 = temp.getString('access_token');
    var head =  {
      'Authorization': 'Bearer $temp2',
      'Content-Type': 'application/json',      
    };
    var body = jsonEncode({
        "userId": userId,
        });
    var url = 'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
    final response = await client.post(Uri.parse(url),headers: head,body: body);
    if (response.statusCode == 201) {
      var data = jsonDecode(response.body);
      return ChatEntity.fromJson(data['data']);
    } else {
      throw Exception('Failed to initiate chat');
    }
  }
  @override
  void serverConnect() {
    SocketService().connectToServer();
  }
  @override
  void sendMessage(String chat, String message, String type) {
    serverConnect();
    SocketService().sendMessage(chat, message, type);


  }

}