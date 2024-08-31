import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/chat.dart';
import '../repositories/chat_repository.dart';

class GetChatsUsecase extends UseCase<List<Chat>, NoParams> {
  final ChatRepository repository;
  GetChatsUsecase(this.repository);

  @override
  Future<Either<Failure, List<Chat>>> call(p) async {
    return await repository.getChats();
  }
}