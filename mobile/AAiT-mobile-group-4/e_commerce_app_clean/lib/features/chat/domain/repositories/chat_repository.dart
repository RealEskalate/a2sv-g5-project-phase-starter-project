import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/chat.dart';
import '../entities/message.dart';

abstract class ChatRepository {
  Future<Either<Failure, List<Chat>>> getChats();
  Future<Either<Failure, List<Message>>> getChat(String chatId);
  Future<Either<Failure, Chat>> createChat(String userId);
  Future<Either<Failure, void>> deleteChat(String chatId);
  Future<Either<Failure, void>> sendMessage(Message message);
}