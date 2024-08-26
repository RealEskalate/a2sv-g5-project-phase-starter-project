
import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/data/models/chat_model.dart';
import 'package:e_commerce_app/features/chat/data/models/message_model.dart';

abstract class LocalContrat{ 

  Future<void> cacheGetChatByIdLocal(ChatModel message);
  Future<List<MessageModel>> getChatByIdLocal(String userId);

  Future <void> cacheGetAllChatsLocal(List<ChatModel> chat);
  Future<List<ChatModel>> getAllChatLocal();


}