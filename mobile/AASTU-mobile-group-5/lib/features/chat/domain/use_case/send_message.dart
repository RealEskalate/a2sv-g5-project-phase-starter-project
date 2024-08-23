import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../repositories/message_repository.dart';

class SendMessage{
  final MessageRepository repository;

  SendMessage(this.repository);

  Future<Either<Failure, void>> execute(String chatId, String message) {
    return repository.sendMessage(chatId, message);
  }
}