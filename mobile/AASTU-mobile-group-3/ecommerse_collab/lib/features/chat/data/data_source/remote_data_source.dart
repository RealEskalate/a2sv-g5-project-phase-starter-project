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
  
  ChatRemoteDataSourceImpl({required this.client});
  var baseUrl = "https://g5-flutter-learning-path-be.onrender.com/api/v3/chats";

  @override
  Future<ChatModel> chatById(String chatId)  {


    throw UnimplementedError();


  }

  @override
  Future<void> deleteMessage(String chatId) async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    var headers = {
      'Authorization': "Bearer ${sharedPreferences.getString('token')}",
      'Content-Type': 'application/json',
    };

    final response = await client.delete(Uri.parse('$baseUrl/$chatId'), headers: headers);

    print(response.body);
    if (response.statusCode != 200){
      throw Exception();
    }
    
  }

  @override
  Future<List<MessageModel>> getChatMessage(String chatId) {
    // TODO: implement getChatMessage
    throw UnimplementedError();
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

    final response = await client.post(Uri.parse(baseUrl), headers: headers, body:json.encode({
      'userId' : userId
    }) );
    
    if (response.statusCode == 201) {
      final jsonData = json.decode(response.body) as Map<String, dynamic>;
      final jsonFinal = jsonData['data'];
      return ChatModel.fromJson(jsonFinal);
    } else {
      throw Exception();
    }


  }
}
