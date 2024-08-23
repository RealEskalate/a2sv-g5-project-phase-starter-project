import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/message_entity.dart';

abstract class MessageRepository {
  Future<Either<Failure, List<MessageEntity>>> getMessages(String chatId);
  Future<Either<Failure, MessageEntity>> sendMessage(String chatId, String message);
  Future<Either<Failure, void>> deleteMessage(String chatId, String messageId);
}