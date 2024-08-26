import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class MyChatById {
  final ChatRepository repository;

  MyChatById(this.repository);

  Future<Either<Failure, ChatEntity>> execute(String id) async {
    return await repository.myChatById(id);
  }
  
 
}
