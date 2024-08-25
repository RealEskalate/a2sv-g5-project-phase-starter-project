// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'dart:async';
import 'dart:convert';


import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/Error/failure.dart';
import '../../../../core/const/const.dart';
import '../../domain/entity/chat_entity.dart';
import '../../domain/entity/message_entity.dart';
import '../model/chat_model.dart';

abstract interface class ChatRemoteData {
  Future<List<ChatEntity>> getMychats();
  Future<ChatEntity> getChatById(String chatId);
  Future<bool> deleteChats(String id);
  Future<bool> initiate(String id);
}

class ChatRemoteDataImpl implements ChatRemoteData {
  final http.Client client;
  SharedPreferences sharedPreferences;

  ChatRemoteDataImpl({
    required this.client,
    required this.sharedPreferences,
  });

  @override
  Future<bool> deleteChats(String id) async{
    try {
      await client.delete(
        Uri.parse(ChatApi.deleteChatApi(id)),
        headers: {
          'Authorization': 'Bearer ${sharedPreferences.getString('key')}',
        },
      );
      
      return Future.value(true);
    } on ConnectionFailur catch (e) {
      throw ConnectionFailur(message: e.toString());
    } catch (e) {
      throw ServerFailure(message: e.toString());
    }
  }

  @override
  Future<ChatEntity> getChatById(String chatId) async {
    try {
      final respond = await client.get(
        Uri.parse(ChatApi.chatByIdApi(chatId)),
        headers: {
           'Authorization': 'Bearer ${sharedPreferences.getString('key')}',
        }
      );
      if (respond.statusCode == 200){
        final data = await client.get(
          Uri.parse(ChatApi.getMessagesApi(chatId)),
          headers: {
           'Authorization': 'Bearer ${sharedPreferences.getString('key')}',
        }
        );
        if (data.statusCode == 200){
          List<Map<String,String>> chatData = [];
          final jsonDecode = json.decode(data.body);
          dynamic messages = jsonDecode['data'];
          for (dynamic message in messages){
            Map<String,String> temp = {
              'senderId': message['sender']['_id'],
              'content':  message['content']
            };
            chatData.add(temp);
          }
          Map<String, dynamic> toEntity = {
            'messageId' : chatId,
            'messages' : chatData
          };
          final MessageModel result = MessageModel.fromJson(toEntity);
          final MessageEntity messageEntity = result.toEntity();
          final dynamic mainData = json.decode(respond.body);
          final Map<String,dynamic> jsonData  = {
            'senderId' : mainData['data']['user1']['_id'],
            'senderName' : mainData['data']['user1']['name'],
            'recieverId' : mainData['data']['user2']['id'],
            'recieverName' : mainData['data']['user2']['name'],
            'chatId' : chatId,
            'messages' : messageEntity
          };
          final ChatModel datas = ChatModel.fromJson(jsonData);
          final ChatEntity chatEntity = datas.toEntity();
          return Future.value(chatEntity);


        }
      }
      throw const ConnectionFailur(message: 'try again');
    } on ConnectionFailur catch (e) {
      throw ConnectionFailur(message: e.toString());
    } catch (e) {
      throw ServerFailure(message: e.toString());
    }
  }

  @override
  Future<List<ChatEntity>> getMychats() {
    throw UnimplementedError();
  }

  @override
  Future<bool> initiate(String id) async {
    try {
      final result =  await client.post(
        Uri.parse(ChatApi.startChatApi()),
        body:  {
          'userId': id
        },
        headers: {
           'Authorization': 'Bearer ${sharedPreferences.getString('key')}',
        }

      );
      if(result.statusCode == 201){
        return Future.value(true);
      }
      return Future.value(false);
    } on ConnectionFailur catch (e) {
      throw ConnectionFailur(message: e.toString());
    } catch (e) {
      throw ServerFailure(message: e.toString());
    }
  }
}
