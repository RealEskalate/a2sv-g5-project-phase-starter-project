import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/no_para_use_case.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';

class MyChats extends NoParamsUseCase<Future<Either<Failure, List<ChatEntity>>>> {
  final ChatRepository repository;

  MyChats(this.repository);

  @override
  Future<Either<Failure, List<ChatEntity>>> call() async {
    return await repository.myChats();
  }
}