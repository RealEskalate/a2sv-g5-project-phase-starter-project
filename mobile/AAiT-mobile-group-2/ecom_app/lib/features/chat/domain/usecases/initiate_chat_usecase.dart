import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/core/usecase/usecase.dart';

import 'package:ecom_app/features/chat/domain/entities/user_chat_entity.dart';
import 'package:ecom_app/features/chat/domain/repositories/chat_repository.dart';

class InitiateChatUsecase implements UseCase<UserChatEntity, NoParams>{
  final ChatRepository chatRepository;

  InitiateChatUsecase(this.chatRepository);
  @override
  Future<Either<Failure, UserChatEntity>> call(NoParams params) {
    return chatRepository.initiateChat();
  }
}