import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/exception.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/core/network/network_info.dart';
import 'package:e_commerce_app/features/chat/data/data%20sources/Local%20data/local_contrat.dart';
import 'package:e_commerce_app/features/chat/data/data%20sources/remote%20data/remote_contrats.dart';
import 'package:e_commerce_app/features/chat/domain/entities/chat_entity.dart';
import 'package:e_commerce_app/features/chat/domain/repository/chat_repository.dart';
import 'package:flutter/services.dart';

class DataRepository extends ChatRepository{
  final  LocalContrat localContrat;
  final RemoteContrats remoteContrats;
  final NetworkInfo networkInfo;
  DataRepository({ required this.localContrat,required this.remoteContrats, required this.networkInfo});

  @override
  Future<Either<Failure, ChatEntity>> createChatById(String sellerId) async{
    if(await networkInfo.isConnected){
      try{
        final result = await remoteContrats.createChatByIdRemote(sellerId);
        return Right(result);
      }on ServerException{
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    }
    throw Exception("No network connection");
  }
   

  @override
  Future<Either<Failure, bool>> deleteChatById(String chatId) async{
    if (await networkInfo.isConnected) {
      try {
        final result = await remoteContrats.deleteChatByIdRemote(chatId);
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
        final result = remoteContrats.getAllChatsRemote();
        localContrat.CacheGetAllChatsLocal(result);

        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    }else{
      try {
        final result = localContrat.getAllChatsLocal();
        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
      
    }
    
  }

  @override
  Future<Either<Failure, ChatEntity>> getChatById(String chatId) async{
    if (await networkInfo.isConnected) {
      try {
        final result = remoteContrats.getChatByIdRemote(chatId);
        localContrat.catcheGetChatByIdLocal(result);
        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    }
    else{
      try {
        final result = localContrat.getChatByIdLocal(chatId);
        return Right(result);
      } on ServerException {
        return Left(ServerFailure("FAILED TO CONNECT TO SERVER"));
      }
    }
  }
}