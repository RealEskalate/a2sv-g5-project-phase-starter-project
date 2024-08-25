import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class MyChatUsecase {
  final ChatRepository _chatRepository;

  const MyChatUsecase({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<Either<Failure, List<ChatEntity>>> execute() {
    return _chatRepository.myChat();
  }
}
