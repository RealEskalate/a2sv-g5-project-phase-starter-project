import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';

import '../entities/chat_entity.dart';

class GetChatByUserId {
  final ChatRepository _chatRepository;

  GetChatByUserId(this._chatRepository);

  Future<Either<Failure,ChatEntity>> execute(String chatId) async {
    return _chatRepository.getChatById(chatId);
  }
}