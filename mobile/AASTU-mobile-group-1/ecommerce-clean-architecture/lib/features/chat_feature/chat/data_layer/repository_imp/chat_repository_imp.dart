import 'package:ecommerce/features/auth/data/data_sources/remote_data_source.dart';

import '../../../../../core/error/failure.dart';
import '../../domain/entity/message.dart';
import 'package:dartz/dartz.dart';

import '../../domain/repository/chat_repository.dart';
import '../data_source/remote_abstract.dart';
import '../model/chat_model.dart';

class ChatRepositoryImp extends ChatRepository {
  RemoteAbstract remoteAbstract;
  ChatRepositoryImp({required this.remoteAbstract});
  @override
  Future<Either<Failure, String>> chatRoom(String receiverId) async{
    var token =await getToken();
    return remoteAbstract.chatRoom(token!,receiverId);
  }

  @override
  Future<Either<Failure, void>> deleteMessage(String chatId) {
    // TODO: implement deleteMessage
    throw UnimplementedError();
  }

  @override
  Stream<Either<Failure, List<Message>>> getMessages(String chatId) {
    // TODO: implement getMessages
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, void>> sendMessage(Message message) {
    // TODO: implement sendMessage
    throw UnimplementedError();
  }

  @override
  Stream<Either<Failure, List<ChatModel>>> getChatHistory() async*{
    var token =await getToken();
   
    yield* remoteAbstract.getChatHistory(token!);
  }
}
