import 'package:dartz/dartz.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/core/usecases/usecases.dart';
import 'package:starter_project/features/chat/domain/entities/message_entity.dart';
import 'package:starter_project/features/chat/domain/repositories/chat_repository.dart';

//  Deleting The Chat By Accept The chat Id
class GetAllMessages implements UseCase<List<Message>, GetMessagesParams> {
  GetAllMessages(this.repository);
  ChatRepository repository;

  @override
  Future<Either<Failure, List<Message>>> call(GetMessagesParams params) async {
    return await repository.getAllChatMessages(params.chatId);
  }
}

class GetMessagesParams {
  const GetMessagesParams({required this.chatId});
  final String chatId;

  List<Object?> get props => [chatId];
}
