import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../entities/message.dart';
import '../repositories/chat_repository.dart';

class GetChatUsecase {
  ChatRepository repository;
  GetChatUsecase(this.repository);

  Future<Stream<Either<Failure, Message>>> call(GetChatParams p) async {
    return repository.getChat(p.chatId);
  }
}

class GetChatParams extends Equatable {
  final String chatId;

  const GetChatParams(this.chatId);
  
  @override
  List<Object> get props => [chatId];
}