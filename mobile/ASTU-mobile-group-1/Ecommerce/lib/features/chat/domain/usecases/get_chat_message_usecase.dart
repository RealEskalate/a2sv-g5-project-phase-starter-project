import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/message_entity.dart';
import '../repositories/chat_repository.dart';

class GetChatMessageUsecase {
  final ChatRepository _chatRepository;

  GetChatMessageUsecase({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<Either<Failure, List<MessageEntity>>> getChatMessages(String chatId) {
    return _chatRepository.getChatMessages(chatId);
  }
}
