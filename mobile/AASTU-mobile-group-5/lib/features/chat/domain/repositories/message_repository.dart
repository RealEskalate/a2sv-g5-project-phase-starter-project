import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/message_entity.dart';

abstract class MessageRepository {
  Future<Either<Failure, List<MessageEntity>>> getChatMessages(String messageId);
}