import 'package:dartz/dartz.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/features/chat/domain/entities/chat_entity.dart';
import 'package:starter_project/features/chat/domain/entities/message_entity.dart';

abstract class ChatRepository {
  Future<Either<Failure, void>> deleteChat(String chatId);
  Future<Either<Failure, List<Chat>>> getAllChat();
  Future<Either<Failure, Chat>> getSingleChat(String chatId);
  Future<Either<Failure, Chat>> intiateChat(String userId);
  Future<Either<Failure, List<Message>>> getAllChatMessages(String chatId);
}
