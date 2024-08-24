import 'package:dartz/dartz.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/core/usecases/usecases.dart';
import 'package:starter_project/features/chat/domain/entities/chat_entity.dart';
import 'package:starter_project/features/chat/domain/repositories/chat_repository.dart';

class GetAllChatUsecase implements UseCase<List<Chat>, NoParams> {
  ChatRepository repository;
  GetAllChatUsecase(this.repository);

  Future<Either<Failure, List<Chat>>> call(NoParams params) async {
    return await repository.getAllChat();
  }
}
