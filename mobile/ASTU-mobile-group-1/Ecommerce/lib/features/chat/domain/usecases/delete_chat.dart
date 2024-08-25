import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';

import '../repositories/chat_repository.dart';

class DeleteChat {
  final ChatRepository _chatRepository;

  DeleteChat({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<Either<Failure, bool>> deleteChat(String chatId) {
    return _chatRepository.deleteChat(chatId);
  }
}
