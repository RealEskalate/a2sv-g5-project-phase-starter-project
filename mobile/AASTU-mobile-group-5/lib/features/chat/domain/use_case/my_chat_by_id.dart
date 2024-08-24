import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/no_para_use_case.dart';
import '../entities/chat_entity.dart';
import '../repositories/chat_repository.dart';


class MyChatById extends NoParamsUseCase<Future<Either<Failure, ChatEntity>>> {
  final ChatRepository repository;

  MyChatById(this.repository);

  @override
  Future<Either<Failure, ChatEntity>> call() async {
    return await repository.myChatById();
  }
}