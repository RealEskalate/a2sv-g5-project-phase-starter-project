import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/core/usecase/usecase.dart';
import 'package:ecom_app/features/chat/domain/entities/message_entity.dart';
import 'package:ecom_app/features/chat/domain/repositories/chat_repository.dart';
import 'package:equatable/equatable.dart';

class GetMessagesUsecase
    implements UseCase<List<MessageEntity>, GetMessagesParams> {
  final ChatRepository chatRepository;

  GetMessagesUsecase(this.chatRepository);
  @override
  Future<Either<Failure, List<MessageEntity>>> call(GetMessagesParams params) {
    return chatRepository.getMessages(params.ChatID);
  }
}

class GetMessagesParams extends Equatable {
  final String ChatID;

  GetMessagesParams({required this.ChatID});

  @override
  List<Object?> get props => [ChatID];
}
