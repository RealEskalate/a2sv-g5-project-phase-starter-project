import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/repositories/chat_repository.dart';
import '../datasources/chat_remote_data_source.dart';
import '../models/chat_model.dart';

class ChatRepositoryImpl implements ChatRepository {
  final ChatRemoteDataSource remoteDataSource;

  ChatRepositoryImpl(this.remoteDataSource);

  @override
  Future<Either<Failure, List<ChatEntity>>> myChats() async {
    // Your implementation here...
  }

  @override
  Future<Either<Failure, ChatEntity>> myChatById() async {
    return remoteDataSource.myChatById();
    throw UnimplementedError();
  }

  @override
  Future<Either<Failure, void>> deleteChat(String id) {
    return remoteDataSource.deleteChat(id);
  }

  @override
  Future<Either<Failure, ChatModel>> initiateChat(String sellerId) async {
    return remoteDataSource.initiateChat(sellerId);
    
  }
}
