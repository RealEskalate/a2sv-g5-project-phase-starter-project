
import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/core/network/network_info.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/data_sources/remote_data_source/remote_data_source.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/models/message_model.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/repository/repository.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/pages/HomeChat.dart';

class ChatRepositoryImpl implements ChatRepository {
  final ChatRemoteDataSource remoteDataSource;
  final NetworkInfo networkInfo;

  ChatRepositoryImpl(this.remoteDataSource, this.networkInfo);

  @override
  Future<Either<Failure,List<ChatEntity>>> getAllChats() async {
     if(await networkInfo.isConnected){
       try{
        final products = await remoteDataSource.getAllChats();
        return Right(products);
      } on ServerFailure {
        return Left(ServerFailure("Server Failure"));
      }
     }else{
        return Left(ServerFailure('No Internet Connection'));
     }
  }
  
  @override
  Future<Either<Failure,ChatEntity>> getMyChatById(String userid) async {
     if(await networkInfo.isConnected){
       try{
        final user = await remoteDataSource.getMyChatById(userid);
        return Right(user);
      } on ServerFailure {
        return Left(ServerFailure("Server Failure"));
      }
     }else {
      return Left(ServerFailure('No User Found'));
     }
  }// Get User info




  @override
  Future<Either<Failure,Stream<MessageModel>>> getChatMessages(String chatId) async {
    if(await networkInfo.isConnected){
       try{
        final product = await remoteDataSource.getChatMessages(chatId);
        return Right(product);
      } on ServerFailure {
        return Left(ServerFailure('Server: Failed to get product'));
      }
     }else {
      return Left(ServerFailure('No Internet Connection'));
     }
  }// Get Product impl

  @override
  Future<Either<Failure,ChatEntity>> initiateChat(String userId) async {
    try{
        final results = await remoteDataSource.initiateChat(userId);
        return Right(results);
      } catch(e) {
        return Left(ServerFailure(e.toString()));
      } 
  } // Add Product impl


  @override
  Future<Either<Failure,String>> deleteChat(String chatId) async {
    try{
        final products = await remoteDataSource.deleteChat(chatId);
        return Right(products);
      } on Exception {
        return Left(ServerFailure('Server: Failed to delete product'));
      } 
      } // Delete Product impl

  @override
  Future<Either<Failure, Unit>> sendChat(
      String chatId, String message, String type) async {
    try{
      remoteDataSource.sendMessage(chatId, message, type);
      return const Right(unit);

    }on Exception {
        return Left(ServerFailure('Server: Failed to delete product'));
      } 
    }
  
}
