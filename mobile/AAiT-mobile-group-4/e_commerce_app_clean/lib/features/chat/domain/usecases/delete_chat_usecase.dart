import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../repositories/chat_repository.dart';

class DeleteChatUsecase extends UseCase<void, DeleteChatParams> {
  final ChatRepository repository;
  DeleteChatUsecase(this.repository);

  @override
  Future<Either<Failure, void>> call(DeleteChatParams p) async {
    return repository.deleteChat(p.chatId);
  }
}

class DeleteChatParams extends Equatable {
  final String chatId;

  const DeleteChatParams(this.chatId);
  
  @override
  List<Object> get props => [chatId];
}