import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../entities/message_entity.dart';
import '../repositories/message_repository.dart';

class GetMessages {
  final MessageRepository messageRepository;

  GetMessages(this.messageRepository);

  Future<Either<Failure, List<MessageEntity>>> execute(String chatId){
    return messageRepository.getMessages(chatId);
  }
}