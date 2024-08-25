import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';

import '../repositories/chat_repository.dart';

class DeleteChatUsecase {
  final ChatRepository _chatRepository;

  DeleteChatUsecase({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<Either<Failure, bool>> execute(String chatId) {
    return _chatRepository.deleteChat(chatId);
  }
}
