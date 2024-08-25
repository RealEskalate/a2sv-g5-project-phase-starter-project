// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'package:dartz/dartz.dart';

import '../../../../core/Error/failure.dart';
import '../../domain/entity/chat_entity.dart';
import '../../domain/repository/chat_repo.dart';
import '../datasource/chat_local_data.dart';
import '../datasource/chat_remote_data.dart';

class ChatRepoImpl implements ChatRepositories {
  ChatLocalData chatLocalData;
  ChatRemoteData chatRemoteData;

  ChatRepoImpl({
    required this.chatLocalData,
    required this.chatRemoteData,
  });

  @override
  Future<Either<Failure, bool>> deleteMessages(String id) async {
    try {
      final result = await chatRemoteData.deleteChats(id);
      return Right(result);
    } on ConnectionFailur catch (e) {
      throw ConnectionFailur(message: e.toString());
    } catch (e) {
      throw ServerFailure(message: e.toString());
    }
  }

  @override
  Future<Either<Failure, ChatEntity>> getChatById(String chatId) async {
    try {
      final result = await chatRemoteData.getChatById(chatId);
      return Right(result);
    } on ConnectionFailur catch (e) {
      throw ConnectionFailur(message: e.toString());
    } catch (e) {
      throw ServerFailure(message: e.toString());
    }
  }

  @override
  Future<Either<Failure, List<ChatEntity>>> getMyChat() async {
    try {
      final result = await chatRemoteData.getMychats();
      return Right(result);
    } on ConnectionFailur catch (e) {
      throw ConnectionFailur(message: e.toString());
    } catch (e) {
      throw ServerFailure(message: e.toString());
    }
  }

  @override
  Future<Either<Failure, bool>> initiateChat(String userId) async {
    try {
      final result = await chatRemoteData.initiate(userId);
      return Right(result);
    } on ConnectionFailur catch (e) {
      throw ConnectionFailur(message: e.toString());
    } catch (e) {
      throw ServerFailure(message: e.toString());
    }
  }
}
