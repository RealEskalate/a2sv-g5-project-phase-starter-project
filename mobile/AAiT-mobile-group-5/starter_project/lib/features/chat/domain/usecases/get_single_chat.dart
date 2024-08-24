import 'package:dartz/dartz.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/core/usecases/usecases.dart';
import 'package:starter_project/features/chat/domain/repositories/chat_repository.dart';

//  Deleting The Chat By Accept The chat Id
class GetSingleChat implements UseCase<void, GetChatParams> {
  GetSingleChat(this.repository);
  ChatRepository repository;

  @override
  Future<Either<Failure, void>> call(GetChatParams params) async {
    return await repository.getSingleChat(params.chatId);
  }
}

class GetChatParams {
  const GetChatParams({required this.chatId});
  final String chatId;

  List<Object?> get props => [chatId];
}
