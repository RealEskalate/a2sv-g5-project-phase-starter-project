import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';

import '../entities/chat_entity.dart';


class GetMessagesById {
  final ChatRepository _chatRepository;

  GetMessagesById(this._chatRepository);

  Future<Either<Failure, List<MessageEntity>>> execute(
      String chatId) async {
    return await _chatRepository.getMessagesById(chatId);
  }
}

