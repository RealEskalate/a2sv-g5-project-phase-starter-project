import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../repository/chat_repository.dart';

class DeleteChatUsecase {
  final ChatRepository repository;

  DeleteChatUsecase(this.repository);

 Future<Either<Failure, void>> deleteChat(String chatId) async {
    return repository.deleteChat(chatId);
  }
}
