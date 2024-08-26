import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/chat.dart';
import '../repositories/chat_repository.dart';

class CreateChatUsecase extends UseCase<Chat, CreateChatParams> {
  final ChatRepository repository;
  CreateChatUsecase(this.repository);

  @override
  Future<Either<Failure, Chat>> call(CreateChatParams p) async {
    return repository.createChat(p.userId);
  }
}

class CreateChatParams extends Equatable {
  final String userId;

  const CreateChatParams(this.userId);
  
  @override
  List<Object> get props => [userId];
}