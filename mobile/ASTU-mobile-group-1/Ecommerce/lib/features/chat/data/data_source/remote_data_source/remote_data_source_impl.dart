import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;

import '../../../../../core/constants/constants.dart';
import '../../../../../core/error/exception.dart';
import '../../../../../core/error/failure.dart';
import '../../model/chat_model.dart';
import '../../model/message_model.dart';
import 'remote_data_source.dart';

class RemoteDataSourceImpl extends RemoteDataSource{

  final http.Client client;
  final String accessToken;
  RemoteDataSourceImpl(this.accessToken, {required this.client});

  @override
  Future<bool> deleteChat(String chatId) async{
    try{
   final response = await client.delete(Uri.parse(Urls.getChatById(chatId)),
       headers: {
        'Authorization': 'Bearer $accessToken',
        'Content-Type': 'application/json',
      },);
      if(response.statusCode==204){
        return true;
      }
      else{
        return false;  
      }
    }
    catch(e){
      throw Exception(e.toString());
    }

 
      
  }

  @override
  Future<List<ChatModel>> getAllChats() async{
    try{
      final response = await client.get(Uri.parse(Urls.baseChat),
      headers: {
        'Authorization': 'Bearer $accessToken',
        'Content-Type': 'application/json',
      },);

      if(response.statusCode==200){
        final result = json.decode(response.body)['data'];
        final List<ChatModel>  answer= [];
        answer.addAll(result.map((json) => ChatModel.fromJson(json)).toList());
        return answer;
      }
      else{
        throw Exception() ;
      }
    }
    catch(e){
    throw ServerException();
  }}

  @override
  Future<ChatModel> getChatById(String chatId)async  {
    try{
    final response = await client.post(Uri.parse(Urls.getChatById(chatId)),
    headers: {
        'Authorization': 'Bearer $accessToken',
        'Content-Type': 'application/json',
      },);
      if(response.statusCode==200){
        final result = json.decode(response.body)['data'];
        return ChatModel.fromJson(result);
      }
      else{
        throw Exception() ;
      }
    }
    catch(e){
      throw Exception(e.toString());
    }
  }
  @override
  Future<ChatModel> initiateChat(String recieverId) async {
    try{
    final response = await client.post(Uri.parse(Urls.baseChat),
    headers: {
        'Authorization': 'Bearer $accessToken',
        'Content-Type': 'application/json',
      },
     body: json.encode({'userId': recieverId})
    );
    if(response.statusCode==201){
      final result = json.decode(response.body)['data'];
      return ChatModel.fromJson(result);
    }
    else{
      throw Exception() ;
    }}
    catch(e){
      throw Exception(e.toString());
    }
  }

    @override
  Future<List<MessageModel>> getChatMessages(String chatId) async{
    try{
     final response = await client.delete(Uri.parse(Urls.getChatById(chatId)),
       headers: {
        'Authorization': 'Bearer $accessToken',
        'Content-Type': 'application/json',
      },);
      if(response.statusCode==200){
        final result = json.decode(response.body)['data'];
        final List<MessageModel>  answer= [];
        answer.addAll(result.map((json) => MessageModel.fromJson(json)).toList());
        return answer;
      }
      else{
        throw Exception() ;
      }
    }
    catch(e){
      throw Exception(e.toString());
    }
 
    
  }

  @override
  Future<void> sendMessage(String chatId, String message, String type) {
    
    throw UnimplementedError();
  }
}