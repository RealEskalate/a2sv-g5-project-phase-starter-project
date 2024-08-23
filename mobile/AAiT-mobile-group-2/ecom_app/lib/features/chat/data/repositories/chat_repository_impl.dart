import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/chat/domain/entities/message_entity.dart';
import 'package:ecom_app/features/chat/domain/entities/user_chat_entity.dart';
import 'package:ecom_app/features/chat/domain/repositories/chat_repository.dart';

class ChatRepositoryImpl extends ChatRepository{
  


  @override
  Future<Either<Failure, List<UserChatEntity>>> getChats() {
    // TODO: implement getChats
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, List<MessageEntity>>> getMessages(String ChatID) {
    // TODO: implement getMessages
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, UserChatEntity>> initiateChat() {
    // TODO: implement initiateChat
    throw UnimplementedError();
  }
}