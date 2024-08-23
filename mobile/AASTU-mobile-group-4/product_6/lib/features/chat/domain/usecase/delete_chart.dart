import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../repository/chat_repository.dart';

class DeleteChat {
  final ChatRepository repository;

  DeleteChat(this.repository);

 Future<Either<Failure, void>> deleteChat(String chatId) async {
    return repository.deleteChat(chatId);
  }
}
