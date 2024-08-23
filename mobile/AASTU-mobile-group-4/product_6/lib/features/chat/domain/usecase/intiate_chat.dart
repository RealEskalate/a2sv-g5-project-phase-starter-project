import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../entity/chat_entity.dart';
import '../repository/chat_repository.dart';

class InitiateChat {
  final ChatRepository repository;

  InitiateChat(this.repository);

    Future<Either<Failure, ChatEntity>> initiateChat(String userId)
 async {
    return repository.initiateChat(userId);
  }
}
