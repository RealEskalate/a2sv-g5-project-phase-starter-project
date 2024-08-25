import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../repository/chat_repository.dart';

class DeleteChatUseCase {
  late ChatRepository chatRepository;
  DeleteChatUseCase(this.chatRepository);

  Future<Either<Failure, void>> call(String chatId) async {
    final deleted = await chatRepository.deleteMessage(chatId);
    return Right(deleted);
  }
}
