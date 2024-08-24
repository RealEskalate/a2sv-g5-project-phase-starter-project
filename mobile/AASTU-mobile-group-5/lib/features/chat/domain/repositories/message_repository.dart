import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/message_entity.dart';

abstract class MessageRepository {
  Future<Either<Failure, List<MessageEntity>>> getChatMessages(String messageId);
  Future<Either<Failure, void>> deleteChatMessage(String messageId);
  Future<Either<Failure, void>> sendMessage(String messageId, MessageEntity message);
  
}