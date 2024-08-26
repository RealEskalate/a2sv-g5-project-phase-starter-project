import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/exception.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/core/network/network_info.dart';
import 'package:e_commerce_app/features/chat/data/data%20sources/Local%20data/local_contrat.dart';
import 'package:e_commerce_app/features/chat/data/data%20sources/remote%20data/remote_contrats.dart';
import 'package:e_commerce_app/features/chat/data/models/message_model.dart';
import 'package:e_commerce_app/features/chat/domain/entities/chat_entity.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';
import 'package:flutter/services.dart';

class ChatRepositoryImpl extends ChatRepository {
  final LocalContrat localContrat;
  final ChatRemoteDataSource remoteContrats;
  final NetworkInfo networkInfo;
  ChatRepositoryImpl(
      {required this.localContrat,
      required this.remoteContrats,
      required this.networkInfo});

  @override
  Future<Either<Failure, ChatEntity>> createChatById(String sellerId) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteContrats.createChatById(sellerId);
        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    }
    throw Exception("No network connection");
  }

  @override
  Future<Either<Failure, bool>> deleteChatById(String chatId) async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteContrats.deleteChatById(chatId);
        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    }
    throw Exception("No network connection");
  }

  @override
  Future<Either<Failure, List<ChatEntity>>> getAllChats() async {
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteContrats.getAllChat();
        localContrat.cacheGetAllChatsLocal(result);

        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    } else {
      try {
        final result = await localContrat.getAllChatLocal();

        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    }
  }

  // @override
  // Future<Stream<Either<Failure, MessageModel>>> getMessagesById(String chatId) async{
  //   if (await networkInfo.isConnected) {
  //     try {
  //       final result = await remoteContrats.getMessagesById(chatId);
  //       localContrat.cacheGetChatByIdLocal(result);
  //       return Right(result);
  //     } on ServerException {
  //       return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
  //     }
  //   }
  //   else{
  //     try {
  //       final result = localContrat.getChatByIdLocal(chatId);
  //       return Right(result);
  //     } on ServerException {
  //       return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
  //     }
  //   }
  // }

  // @override
  // Future<Either<Failure, List<MessageModel>>> getMessagesById (
  //     String chatId) async {
  //   if (await networkInfo.isConnected) {
  //     final result = await remoteContrats.getMessagesById(chatId);
  //     return Right(result);
  //   } else {
  //     return Left(Failure("network failure to get messages"));
  //   }
  // }

  @override
  Future<Either<Failure, bool>> sendMessage(
      String chatId, String message, String content) async {
    // TODO: implement sendMessage
    if (await networkInfo.isConnected) {
      try {
        final result =
            await remoteContrats.sendMessage(chatId, message, content);
            print("result $result");
        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    } else {
      return Right(false);
    }
  }
  
  @override
  Future<Either<Failure, List<MessageEntity>>> getMessagesById(String chatId) async{

 
    if (await networkInfo.isConnected) {
      final result = await remoteContrats.getMessagesById(chatId);
      return Right(result);
    } else {
      return Left(Failure("network failure to get messages"));
    }
  
    // TODO: implement getMessagesById
  
  }
}
