import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;

import '../../../../../core/constants/constants.dart';
import '../../../../../core/error/failure.dart';
import '../../model/chat_model.dart';
import '../../model/message_model.dart';
import 'remote_data_source.dart';

class RemoteDataSourceImpl extends RemoteDataSource{

  final http.Client client;
  final String accessToken;
  RemoteDataSourceImpl(this.accessToken, {required this.client});

  @override
  Future<void> deleteChat(String chatId) async{
      client.delete(Uri.parse(Urls.getChatById(chatId)),
       headers: {
        'Authorization': 'Bearer $accessToken',
        'Content-Type': 'application/json',
      },);
  }

  @override
  Future<List<ChatModel>> getAllChats() {
    // TODO: implement getAllChats
    throw UnimplementedError();
  }

  @override
  Future<ChatModel> getChatById(String receiverId) {
    // TODO: implement getChatById
    throw UnimplementedError();
  }

  @override
  Stream<MessageModel> getChatMessages(String chatId) {
    
    // TODO: implement getChatMessages
    
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, String>> initiateChat(String recieverId) async {
    final response = await client.post(Uri.parse(Urls.baseChat),
    headers: {
        'Authorization': 'Bearer $accessToken',
        'Content-Type': 'application/json',
      },
     body: json.encode({'userId': recieverId})
    );
    if(response.statusCode==201){
      return Future(()=>Right(json.decode(response.body)['data']['id']));
    }
    else{
      return Future(()=>Left(ServerFailure(response.body)));
    }
  }

  @override
  Future<void> sendMessage(String chatId, String message, String type) {
    // TODO: implement sendMessage
    throw UnimplementedError();
  }

}