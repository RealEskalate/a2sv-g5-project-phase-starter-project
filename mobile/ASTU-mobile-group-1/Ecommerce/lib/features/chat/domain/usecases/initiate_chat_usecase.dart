import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class InitiateChatUsecase {
  final ChatRepository _chatRepository;

  InitiateChatUsecase({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<Either<Failure, ChatEntity>> initiateChat(String userId) {
    return _chatRepository.initiateChat(userId);
  }
}
