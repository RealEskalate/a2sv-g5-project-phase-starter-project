import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';
class DeleteChatByIdUseCase {
  final ChatRepository _chatRepository;

  DeleteChatByIdUseCase(this._chatRepository);

  Future<Either<Failure,bool>> execute(String chatId) {
    return _chatRepository.deleteChatById(chatId);
  }
}