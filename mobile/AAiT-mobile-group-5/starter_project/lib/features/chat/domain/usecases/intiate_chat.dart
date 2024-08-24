import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/core/usecases/usecases.dart';
import 'package:starter_project/features/chat/domain/entities/chat_entity.dart';
import 'package:starter_project/features/chat/domain/repositories/chat_repository.dart';

class InitiateChatUsecase implements UseCase<Chat, IntiateChatParams> {
  final ChatRepository repository;

  InitiateChatUsecase(this.repository);

  //  it return the created(intiated ) chat for the asked user

  @override
  Future<Either<Failure, Chat>> call(IntiateChatParams params) async {
    return await repository.intiateChat(params.userId);
  }
}

class IntiateChatParams extends Equatable {
  const IntiateChatParams({required this.userId});
  final String userId;
  @override
  List<Object?> get props => [userId];
}
