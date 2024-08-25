import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';

import '../entities/chat_entity.dart';

abstract class ChatRepository {
  Future<Either<Failure,List<ChatEntity>>> getAllChats();
  Future<Either<Failure,ChatEntity>> createChatById(String sellerId);
  Future<Either<Failure,ChatEntity>> getChatById(String chatId);
  Future<Either<Failure,bool>> deleteChatById(String chatId);
}