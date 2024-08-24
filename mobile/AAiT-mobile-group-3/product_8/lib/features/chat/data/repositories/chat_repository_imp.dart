
import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/message_entity.dart';

import '../../domain/repositories/chat_repository.dart';



class ChatRepositoryImp implements ChatRepository{
	final ChatLocalDataSource localdatasource;
	final ChatRemoteDataSource remotedatasource;

	const ChatRepositoryImp({required this.localdatasource, required this.remotedatasource});
  @override
  Future<Either<Failure, Unit>> deleteChat(String chatId) {
    // TODO: implement deleteChat
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, List<MessageEntity>>> getChatMessages(String chatId) {
    // TODO: implement getChatMessages
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, ChatEntity>> initiateChat(String sellerId) {
    // TODO: implement initiateChat
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, ChatEntity>> myChatById(String chatId) {
    // TODO: implement myChatById
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, List<ChatEntity>>> myChats() {
    // TODO: implement myChats
    throw UnimplementedError();
  }

}
