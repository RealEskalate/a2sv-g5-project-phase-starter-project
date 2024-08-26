import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class MyChatById {
  final ChatRepository _chatRepository;

  MyChatById({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<Either<Failure, ChatEntity>> execute(String chatId) {
    return _chatRepository.myChatbyId(chatId);
  }
}
