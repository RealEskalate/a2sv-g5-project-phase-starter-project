import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../entity/chat_entity.dart';
import '../repository/chat_repository.dart';

class GetAllChats {
  final ChatRepository repository;

  GetAllChats(this.repository);

Future<Either<Failure, List<ChatEntity>>> getAllChats()
async {
    return repository.getAllChats();
  }
}
