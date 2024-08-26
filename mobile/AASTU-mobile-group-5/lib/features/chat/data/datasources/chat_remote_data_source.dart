import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;

import '../../../../core/failure/failure.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/chat_message_entity.dart';
import '../models/chat_model.dart';
import 'websocket_service.dart';

abstract class ChatRemoteDataSource {
  Future<Either<Failure,List<ChatModel>>> myChats();
  Future<Either<Failure,  ChatModel>>getChatById(String chatId);
  Future<void> deleteChat(String chatId);
  Future<Either<Failure,ChatModel>> initiateChat(String sellerId);
}
class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final WebSocketService webSocketService;
  final http.Client client;

  ChatRemoteDataSourceImpl({
    required this.webSocketService,
    required this.client,
  });

  @override
  Future<void> deleteChat(String chatId) async {
    
    webSocketService.sendMessage('delete_chat', {'chatId': chatId});
  }

  @override
  Future<Either<Failure,ChatModel>> initiateChat(String sellerId) async {
   
    final response = webSocketService.sendMessage('initiate_chat', {'sellerId': sellerId});
    
    throw UnimplementedError();
  }

   @override
  Future<Either<Failure, ChatModel>> getChatById(String id) async {
    try {
      final response = await client.get(
        Uri.parse('https://example.com/api/chats/$id'), 
        headers: {
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final chatModel = ChatModel.fromJson(json.decode(response.body));
        return Right(chatModel);
      } else {
        return Left(ServerFailure as Failure);
      }
    } catch (e) {
      return Left(ServerFailure as Failure);
    }
  }
}

  // @override
  // Future<List<ChatModel>> myChats() async {
  //   // Fetch all chats using WebSocket
  //   final response = webSocketService.sendMessage('get_chats', {});
  //   return (response as List)
  //       .map((chat) => ChatModel.fromJson(chat))
  //       .toList();
  // }
  
  @override
  Future<ChatEntity> MyChatById(String chatId) {
    // TODO: implement MyChatById
    throw UnimplementedError();
  }

// import 'package:http/http.dart' as http;

// import '../../domain/entities/chat_entity.dart';
// import '../models/chat_model.dart';

// abstract class ChatRemoteDataSource {
//   Future<List<ChatModel>> myChats();
//   Future<ChatEntity> MyChatById(String chatId);
//   Future<void> deleteChat(String chatId);
//   Future<ChatEntity> initiateChat(String sellerId);
// }

// class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
//   late final http.Client client;
//   ChatRemoteDataSourceImpl({
//     required this.client,
//   });
//   @override
//   Future<void> deleteChat(String chatId) {
//     // TODO: implement deleteChat
//     throw UnimplementedError();
//   }

//   @override
//   Future<ChatEntity> initiateChat(String sellerId) {
//     // TODO: implement initiateChat
//     throw UnimplementedError();
//   }
  
//   @override
//   Future<ChatEntity> MyChatById(String chatId) {
//     // TODO: implement MyChatById
//     throw UnimplementedError();
//   }
  
//   @override
//   Future<List<ChatModel>> myChats() {
//     // TODO: implement myChats
//     throw UnimplementedError();
//   }
// }
