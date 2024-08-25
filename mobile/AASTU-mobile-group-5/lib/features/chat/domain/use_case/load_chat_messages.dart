
import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../entities/chat_message_entity.dart';
import '../repositories/chat_repository.dart';
import '../repositories/message_repository.dart';


class LoadChatMessagesUseCase implements UseCase<List<ChatMessageEntity>, LoadParams> {
  final MessageRepository repository;

  LoadChatMessagesUseCase(this.repository);

  @override
  Future<Either<Failure, List<ChatMessageEntity>>> call(LoadParams params) async {
    return repository.getChatMessages(params.chatId) as Future<Either<Failure, List<ChatMessageEntity>>>;
  }
}

class LoadParams {
  final String chatId;

  LoadParams({required this.chatId});
}
