import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/chat/domain/entities/message_entity.dart';
import 'package:ecom_app/features/chat/domain/entities/user_chat_entity.dart';

abstract class ChatRepository {
  Future<Either<Failure, List<UserChatEntity>>> getChats();
  Future<Either<Failure, UserChatEntity>> initiateChat();
  Future<Either<Failure, List<MessageEntity>>> getMessages(String ChatID);
}