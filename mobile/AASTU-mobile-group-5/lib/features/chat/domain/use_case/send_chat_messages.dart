import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../entities/chat_message_entity.dart';
import '../entities/message_entity.dart';
import '../repositories/chat_repository.dart';
import '../repositories/message_repository.dart';


class SendChatMessageUseCase implements UseCase<void, SendParams> {
  final MessageRepository repository;

  SendChatMessageUseCase(this.repository);

  @override
  Future<Either<Failure, void>> call(SendParams params) async {
    return await repository.sendMessage(params.chatId, params.message as MessageEntity);
  }
}

class SendParams {
  final String chatId;
  final ChatMessageEntity message;

  SendParams({required this.chatId, required this.message});
}
