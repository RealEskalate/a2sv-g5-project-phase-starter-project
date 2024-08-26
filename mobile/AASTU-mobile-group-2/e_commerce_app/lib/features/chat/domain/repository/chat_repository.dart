import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';

import '../entities/chat_entity.dart';

abstract class ChatRepository {
  Future<Either<Failure,List<ChatEntity>>> getAllChats();
  Future<Either<Failure,ChatEntity>> createChatById(String sellerId);

  Future<Either<Failure,List<MessageEntity>>> getMessagesById(String chatId);
  Future<Either<Failure,bool>> deleteChatById(String chatId);
  Future<Either<Failure,bool>> sendMessage(String chatId, String message, String content);
}