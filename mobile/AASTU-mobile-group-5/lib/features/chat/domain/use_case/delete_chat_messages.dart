import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../repositories/chat_repository.dart';
import '../repositories/message_repository.dart';


class DeleteChatMessageUseCase implements UseCase<void, DeleteParams> {
  final MessageRepository repository;

  DeleteChatMessageUseCase(this.repository);

  @override
  Future<Either<Failure, void>> call(DeleteParams params) async {
    return await repository.deleteChatMessage(params.messageId);
  }
}

class DeleteParams {
  final String messageId;

  DeleteParams({required this.messageId});
}
