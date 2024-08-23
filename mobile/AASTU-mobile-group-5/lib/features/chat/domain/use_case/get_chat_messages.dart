import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/message_entity.dart';
import '../repositories/message_repository.dart';

class GetChatMessages{
  final MessageRepository repository;

  GetChatMessages(this.repository);

  Future<Either<Failure, List<MessageEntity>>> call(String messageId) async {
    return await repository.getChatMessages(messageId);
  }
}