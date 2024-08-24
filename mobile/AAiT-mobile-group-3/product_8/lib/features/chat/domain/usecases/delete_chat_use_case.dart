import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../repositories/chat_repository.dart';

class DeleteChatUseCase extends UseCase<Unit, DeleteChatParams> {
  final ChatRepository chatRepository;

  DeleteChatUseCase(this.chatRepository);

  @override
  Future<Either<Failure, Unit>> call(DeleteChatParams params) async {
    return await chatRepository.deleteChat(params.chatId);
  }
}

class DeleteChatParams extends Equatable {
  final String chatId;

  const DeleteChatParams({required this.chatId});

  @override
  List<Object> get props => [chatId];
}
