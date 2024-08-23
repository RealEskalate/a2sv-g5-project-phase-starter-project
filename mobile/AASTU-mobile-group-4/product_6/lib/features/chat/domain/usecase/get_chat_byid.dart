import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../entity/chat_entity.dart';
import '../repository/chat_repository.dart';

class GetChatByIdUsecase {
  final ChatRepository repository;

  GetChatByIdUsecase(this.repository);

 Future<Either<Failure, ChatEntity>> getChatById(String chatId) async {
    return repository.getChatById(chatId);
  }
}
