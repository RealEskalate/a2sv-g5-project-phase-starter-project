import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class MyChat {
  final ChatRepository _chatRepository;

  MyChat({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<Either<Failure, List<ChatEntity>>> myChat() {
    return _chatRepository.myChat();
  }
}