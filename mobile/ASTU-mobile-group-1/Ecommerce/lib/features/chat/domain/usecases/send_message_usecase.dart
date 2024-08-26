import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/message_entity.dart';
import '../repositories/chat_repository.dart';

class SendMessageUsecase {
  final ChatRepository _chatRepository;

  SendMessageUsecase({required ChatRepository chatRepository})
      : _chatRepository = chatRepository;

  Future<void> call(
      {required String chatId, required String message, required String type}) {
    return _chatRepository.sendMessage(
        chatId: chatId, message: message, type: type);
  }
}
