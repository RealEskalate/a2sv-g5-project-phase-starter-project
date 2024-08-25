import 'dart:convert';

import 'package:ecommerce_app_ca_tdd/core/constants/constants.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

abstract class ChatRemoteDataSource {

  Future <ChatEntity> getMyChatById(String id);
  Future<String> initiateChat(String userId);
  Future<String> deleteChat(String chatId);
  Future<List<ChatModel>> getChatMessages(String chatId);
  Future<List<ChatEntity>> getAllChats();
 
}
class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final http.Client client;
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
  Future<List<ChatModel>> getChatMessages(String chatId) async{
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
      return (jsonDecode(response.body)['data'] as List)
          .map((e) => ChatModel.fromJson(e))
          .toList();
    } else {
      throw Exception('Failed to Fetch Chat');
    }
  }
  @override
  Future<String> initiateChat(String userId) async{
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
    print(response.statusCode);
    if (response.statusCode == 200) {
      var data = jsonDecode(response.body);
      return data['data']['_id'];
    } else {
      throw Exception('Failed to initiate chat');
    }
  }

}