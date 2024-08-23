import 'package:dartz/dartz.dart';

import '../repositories/message_repository.dart';

class DeleteMessage {
  final MessageRepository repository;

  DeleteMessage(this.repository);

  Future<Either<void, void>> execute(String chatId, String messageId) {
    return repository.deleteMessage(chatId, messageId);
  }
}
