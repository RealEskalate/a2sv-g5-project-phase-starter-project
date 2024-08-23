
import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/chat_feature/chat/domain/entity/message.dart';

import '../../../../../core/error/failure.dart';

abstract class ChatRepository {
  Future<Either<Failure,void>>sendMessage(Message message);
  Stream<Either<Failure,List<Message>>> getMessages(String chatId);
  Future<Either<Failure,void>> deleteMessage(String chatId);
  Future<Either<Failure,String>> chatRoom(String receiverId);
}