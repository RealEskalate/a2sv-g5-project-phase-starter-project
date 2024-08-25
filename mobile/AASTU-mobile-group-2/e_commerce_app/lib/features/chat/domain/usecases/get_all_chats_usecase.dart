import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';

import '../entities/chat_entity.dart';

class GetAllChatsUseCase {
  final ChatRepository _chatRepository;

  GetAllChatsUseCase(this._chatRepository);

  Future<Either<Failure,List<ChatEntity>>> execute() {
    return _chatRepository.getAllChats();
  }
}