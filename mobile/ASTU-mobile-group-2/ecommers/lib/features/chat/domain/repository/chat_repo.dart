
import 'package:dartz/dartz.dart';
import '../../../../core/Error/failure.dart';
import '../entity/chat_entity.dart';



abstract class ChatRepositories{


  Future<Either<Failure,ChatEntity>> getChatById(String chatId);

  Future<Either<Failure,List<ChatEntity>>> getMyChat();
 

  Future<Either<Failure,bool>> initiateChat(String userId);

  Future<Either<Failure,bool>> deleteMessages(String id);
  
}