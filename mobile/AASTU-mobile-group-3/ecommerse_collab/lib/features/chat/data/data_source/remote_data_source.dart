import 'dart:async';
import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;

import '../../../../socket/socket.dart';
import '../models/chat_model.dart';
import '../models/message_model.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatModel>> getMyChats();
  Future<ChatModel> initiateChat(String userId);
  Future<ChatModel> chatById(String chatId);
  Future<void> deleteMessage(String chatId);
  Future<List<MessageModel>> getChatMessage(String chatId);
}

class ChatRemoteDataSourceImpl extends ChatRemoteDataSource {
  final http.Client client;
  final SocketIoManager _socketIoManager;

  ChatRemoteDataSourceImpl({required this.client, required SocketIoManager socketIoManager})
      : _socketIoManager = socketIoManager;

  void initSocket(String serverUrl, String token){
    _socketIoManager.init();
  }

  void sendMessage(String chatId, String content){
    final message = {
      'chatId': chatId,
      'content' : content,
      'type' : 'text'
    };
    _socketIoManager.sendMessage('message:send', message);
  }

  void handleIncomingMessages(){
    _socketIoManager.subscribe('message:received', (data){
      final message = MessageModel.fromJson(json.decode(data));
      print('New message received: ${message.content}');
    });
  }

  
  var baseUrl = "https://g5-flutter-learning-path-be.onrender.com/api/v3/chats";

  @override
  Future<ChatModel> chatById(String chatId) async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    var headers = {
      'Authorization': "Bearer ${sharedPreferences.getString('token')}",
      'Content-Type': 'application/json',
    };

    final response =
        await client.get(Uri.parse('$baseUrl/$chatId'), headers: headers);

    if (response.statusCode == 200) {
      final jsonData = json.decode(response.body) as Map<String, dynamic>;
      return ChatModel.fromJson(jsonData['data']);
    } else {
      throw Exception("Failed to load chat by ID");
    }
  }

  @override
  Future<void> deleteMessage(String chatId) async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    var headers = {
      'Authorization': "Bearer ${sharedPreferences.getString('token')}",
      'Content-Type': 'application/json',
    };

    final response =
        await client.delete(Uri.parse('$baseUrl/$chatId'), headers: headers);

    print(response.body);
    if (response.statusCode != 200) {
      throw Exception();
    }
  }

  @override
  Future<List<MessageModel>> getChatMessage(String chatId) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    var headers = {
      'Authorization': "Bearer ${prefs.getString('token')}",
      'Content-Type': 'application/json',
    };

    final response = await client.get(Uri.parse('$baseUrl/$chatId/messages'),
        headers: headers);

    if (response.statusCode == 200) {
      final jsonData = json.decode(response.body) as Map<String, dynamic>;
      List<MessageModel> messages = [];

      for (var messgeData in jsonData['data']) {
        messages.add(MessageModel.fromJson(messgeData));
      }
      return messages;
    } else {
      throw Exception("Failed to load chat messages");
    }
  }

  @override
  Future<List<ChatModel>> getMyChats() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    var headers = {
      'Authorization': "Bearer ${prefs.getString('token')}",
      'Content-Type': 'application/json',
    };
    final response = await client.get(Uri.parse(baseUrl), headers: headers);

    if (response.statusCode == 200) {
      final jsonData = json.decode(response.body) as Map<String, dynamic>;

      List<ChatModel> chatModels = [];
      for (var data in jsonData['data']) {
        chatModels.add(ChatModel.fromJson(data));
      }
      return chatModels;
    } else {
      throw Exception();
    }
  }

  @override
  Future<ChatModel> initiateChat(String userId) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    var headers = {
      'Authorization': "Bearer ${prefs.getString('token')}",
      'Content-Type': 'application/json',
    };

    final response = await client.post(Uri.parse(baseUrl),
        headers: headers, body: json.encode({'userId': userId}));

    if (response.statusCode == 201) {
      final jsonData = json.decode(response.body) as Map<String, dynamic>;
      final jsonFinal = jsonData['data'];
      return ChatModel.fromJson(jsonFinal);
    } else {
      throw Exception();
    }
  }
}
