import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';

import '../entities/chat_entity.dart';

class CreateChatUsecase {
  final ChatRepository _chatRepository;

  CreateChatUsecase(this._chatRepository);

  Future<Either<Failure,ChatEntity>> execute( String sellerId) {
    return _chatRepository.createChatById(sellerId);
  }
}