import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/message.dart';
import '../repositories/chat_repository.dart';

class GetChatUsecase extends UseCase<List<Message>, GetChatParams> {
  ChatRepository repository;
  GetChatUsecase(this.repository);

  @override
  Future<Either<Failure, List<Message>>> call(GetChatParams p) async {
    return repository.getChat(p.chatId);
  }
}

class GetChatParams extends Equatable {
  final String chatId;

  const GetChatParams(this.chatId);
  
  @override
  List<Object> get props => [chatId];
}