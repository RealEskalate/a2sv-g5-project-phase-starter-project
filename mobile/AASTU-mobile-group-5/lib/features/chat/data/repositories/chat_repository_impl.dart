import 'package:dartz/dartz.dart';

import '../../../../core/error/exceptions.dart';
import '../../../../core/failure/failure.dart';
import '../../domain/repositories/chat_repository.dart';
import '../datasources/chat_remote_data_source.dart';
import '../datasources/local_data_source.dart';
import '../models/chat_model.dart';

class ChatRepositoryImpl implements ChatRepository {
  final ChatRemoteDataSource remoteDataSource;
  final ChatLocalDataSource localDataSource;

  ChatRepositoryImpl({
    required this.remoteDataSource,
    required this.localDataSource,
  });

  @override
  Future<Either<Failure, ChatModel>> myChatById(String id) async {
    try {
     
      final remoteChatResult = await remoteDataSource.getChatById(id);
      return remoteChatResult.fold(
        (failure) async {
         
          final localChatResult = await localDataSource.getChatById(id);
          return localChatResult;
        },
        (chat) async {
          
          await localDataSource.cacheChat(chat as ChatModel);
          return Right(chat);
        },
      );
    } on ServerException {
      return Left(ServerFailure(message: 'server Error'));
    }
  }
}
