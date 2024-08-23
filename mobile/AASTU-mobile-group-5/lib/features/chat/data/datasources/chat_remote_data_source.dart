import 'package:http/http.dart' as http;

import '../../domain/entities/chat_entity.dart';
import '../models/chat_model.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatModel>> myChats();
  Future<ChatEntity> MyChatById(String chatId);
  Future<void> deleteChat(String chatId);
  Future<ChatEntity> initiateChat(String sellerId);
}

class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  late final http.Client client;
  ChatRemoteDataSourceImpl({
    required this.client,
  });
  @override
  Future<void> deleteChat(String chatId) {
    // TODO: implement deleteChat
    throw UnimplementedError();
  }

  @override
  Future<ChatEntity> initiateChat(String sellerId) {
    // TODO: implement initiateChat
    throw UnimplementedError();
  }
  
  @override
  Future<ChatEntity> MyChatById(String chatId) {
    // TODO: implement MyChatById
    throw UnimplementedError();
  }
  
  @override
  Future<List<ChatModel>> myChats() {
    // TODO: implement myChats
    throw UnimplementedError();
  }
}
