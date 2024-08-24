import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/message_entity.dart';
import '../repositories/chat_repository.dart';

class GetChatMessagesUseCase
    extends UseCase<List<MessageEntity>, GetChatMessagesParams> {
  final ChatRepository chatRepository;

  GetChatMessagesUseCase(this.chatRepository);

  @override
  Future<Either<Failure, List<MessageEntity>>> call(
      GetChatMessagesParams params) async {
    return await chatRepository.getChatMessages(params.chatId);
  }
}

class GetChatMessagesParams extends Equatable {
  final String chatId;

  const GetChatMessagesParams({required this.chatId});

  @override
  List<Object> get props => [chatId];
}
