import '../../../../core/error/exception.dart';
import '../../../../core/error/failure.dart';
import '../../domain/entities/chat.dart';
import '../../domain/entities/message.dart';
import 'package:dartz/dartz.dart';

import '../../domain/repositories/chat_repository.dart';
import '../data_sources/chat_data_source.dart';
import '../models/message_model.dart';

class ChatRepositoryImpl implements ChatRepository {
  final ChatDataSource dataSource;

  ChatRepositoryImpl({required this.dataSource});

  @override
  Future<Either<Failure, Chat>> createChat(String userId) async {
    try {
      var chat = await dataSource.createChat(userId);
      return Right(chat);
    } on ServerException {
      return const Left(ServerFailure('Server error'));
    } 
    catch(e) {
      return const Left(ServerFailure('Unknow Failure'));
    }
  }

  @override
  Future<Either<Failure, void>> deleteChat(String chatId) async {
    try {
      await dataSource.deleteChat(chatId);
      return const Right(null);
    } on ServerException {
      return const Left(ServerFailure('Server Error'));
    }
  }

  @override
  Future<Stream<Either<Failure, Message>>> getChat(String chatId) async {
    try {
      final so = await dataSource.getChat(chatId);
      return so.map((msg) => Right(msg));
    } catch(e) {
      var err = e as ServerException;
      return Stream.value(Left(ServerFailure(err.message ?? 'Server Error')));
    }
  }

  @override
  Future<Either<Failure, List<Chat>>> getChats() async {
    try {
      return Right(await dataSource.getChats());
    } on ServerException {
      return const Left(ServerFailure('Server error'));
    }
  }

  @override
  Future<Either<Failure, void>> sendMessage(Message message) async {
    try {
      await dataSource.sendMessage(MessageModel.fromEntity(message));
      return const Right(null);
    } on ServerException {
      return const Left(ServerFailure('Server Error'));
    }
  }  
}