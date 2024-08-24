import 'package:http/http.dart' as http;

import '../../domain/entities/chat_entity.dart';
import '../models/chat_model.dart';
import 'websocket_service.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatModel>> myChats();
  Future<ChatEntity> MyChatById(String chatId);
  Future<void> deleteChat(String chatId);
  Future<ChatEntity> initiateChat(String sellerId);
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
    // Use WebSocket or HTTP to delete a chat
    webSocketService.sendMessage('delete_chat', {'chatId': chatId});
  }

  @override
  Future<ChatEntity> initiateChat(String sellerId) async {
    // Use WebSocket or HTTP to initiate a chat
    final response = webSocketService.sendMessage('initiate_chat', {'sellerId': sellerId});
    // return ChatModel.fromJson(response);
    throw UnimplementedError();
  }

  @override
  Future<ChatEntity> myChatById(String chatId) async {
    // Fetch chat by ID using WebSocket
    final response = webSocketService.sendMessage('get_chat_by_id', {'chatId': chatId});
    // return ChatModel.fromJson(response);
    throw UnimplementedError();
  }

  @override
  Future<List<ChatModel>> myChats() async {
    // Fetch all chats using WebSocket
    final response = webSocketService.sendMessage('get_chats', {});
    return (response as List)
        .map((chat) => ChatModel.fromJson(chat))
        .toList();
  }
  
  @override
  Future<ChatEntity> MyChatById(String chatId) {
    // TODO: implement MyChatById
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
