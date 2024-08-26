import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/message_entity.dart';
import '../repositories/chat_repository.dart';

class GetMessageUsecase {
  final ChatRepository _chatRepository;

  GetMessageUsecase({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<Either<Failure, Stream<MessageEntity>>> call() {
    return _chatRepository.getMessages();
  }
}
