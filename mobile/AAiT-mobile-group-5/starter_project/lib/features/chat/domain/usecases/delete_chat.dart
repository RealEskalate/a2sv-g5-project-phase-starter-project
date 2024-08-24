import 'package:dartz/dartz.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/core/usecases/usecases.dart';
import 'package:starter_project/features/chat/domain/repositories/chat_repository.dart';

//  Deleting The Chat By Accept The chat Id
class DeleteChatUsecase implements UseCase<void, DeleteChatParams> {
  DeleteChatUsecase(this.repository);
  ChatRepository repository;

  @override
  Future<Either<Failure, void>> call(DeleteChatParams params) async {
    return await repository.deleteChat(params.chatId);
  }
}

class DeleteChatParams {
  const DeleteChatParams({required this.chatId});
  final String chatId;

  List<Object?> get props => [chatId];
}
