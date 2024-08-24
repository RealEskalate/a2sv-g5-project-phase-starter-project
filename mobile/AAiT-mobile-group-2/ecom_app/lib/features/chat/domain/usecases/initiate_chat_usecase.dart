import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/core/usecase/usecase.dart';

import 'package:ecom_app/features/chat/domain/entities/user_chat_entity.dart';
import 'package:ecom_app/features/chat/domain/repositories/chat_repository.dart';
import 'package:equatable/equatable.dart';

class InitiateChatUsecase implements UseCase<UserChatEntity, InitiateChatParams>{
  final ChatRepository chatRepository;

  InitiateChatUsecase(this.chatRepository);
  @override
  Future<Either<Failure, UserChatEntity>> call(InitiateChatParams params) {
    return chatRepository.initiateChat(params.UserID);
  }
}

class InitiateChatParams extends Equatable {
  final String UserID;

  InitiateChatParams({required this.UserID});

  @override
  List<Object?> get props => [UserID];
}
