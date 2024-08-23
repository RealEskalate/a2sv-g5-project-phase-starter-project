import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../entity/chat_entity.dart';
import '../repository/chat_repository.dart';

class GetAllChatsUsecase {
  final ChatRepository repository;

  GetAllChatsUsecase(this.repository);

Future<Either<Failure, List<ChatEntity>>> getAllChats()
async {
    return repository.getAllChats();
  }
}
