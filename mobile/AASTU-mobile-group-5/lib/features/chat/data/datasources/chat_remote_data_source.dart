import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;

import '../../../../core/constants/api_url.dart';
import '../../../../core/failure/failure.dart';
import '../../domain/entities/chat_entity.dart';
import '../models/chat_model.dart';
import 'websocket_service.dart';

abstract class ChatRemoteDataSource {
  Future<Either<Failure, ChatModel>> initiateChat(String sellerId);
  Future<Either<List, ChatModel>> myChats();
  Future<Either<Failure, ChatEntity>> myChatById();
  Future<Either<Failure, void>> deleteChat(String id);
  
}
class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final WebSocketService webSocketService;
  final http.Client client;

  ChatRemoteDataSourceImpl({
    required this.webSocketService,
    required this.client,
  });

  @override
  Future<Either<Failure, void>> deleteChat(String id) async {
    final url = Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v3/chats/$id');
    try {
      final response = await client.delete(url);

      if (response.statusCode == 204) {
        return Right(null);
      } else {
        return Left(ServerFailure(message: 'Failed to delete chat'));
      }
    } catch (e) {
      return Left(ServerFailure(message: 'Failed to delete chat'));
    }
  }

  @override
  Future<Either<Failure, ChatModel>> initiateChat(String sellerId) async {
    final url = Uri.parse(Urls.baseUrl);
    try {
      final response = await client.post(
        url,
        headers: {'Content-Type': 'application/json'},
        body: json.encode({'sellerId': sellerId}),
      );

      if (response.statusCode == 201) {
        final data = json.decode(response.body)['data'];
        final chatModel = ChatModel.fromJson(data);
        return Right(chatModel);
      } else {
        return Left(ServerFailure(message: 'Failed to initiate chat'));
      }
    } catch (e) {
      return Left(ServerFailure(message: 'Failed to initiate chat'));
    }
  }

  @override
  Future<Either<Failure, ChatEntity>> myChatById() async {
    // Fetch chat by ID using WebSocket
    final response = webSocketService.sendMessage('get_chat_by_id', {});
    // return ChatModel.fromJson(response);
    throw UnimplementedError();
  }

  @override
  Future<Either<List, ChatModel>> myChats() async {
    // Fetch all chats using WebSocket
    // final response = webSocketService.sendMessage('get_chats', {});
    // return (response as List)
    //     .map((chat) => ChatModel.fromJson(chat))
    //     .toList();
    throw UnimplementedError();
  }

  
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
