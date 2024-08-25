import 'dart:io';

import 'package:ecommerce/core/error/exception.dart';
import 'package:ecommerce/features/auth/data/data_sources/remote_data_source.dart';
import 'package:ecommerce/features/chat_feature/chat/data_layer/model/message_model.dart';

import '../../../../../core/error/failure.dart';
import '../../domain/entity/message.dart';
import 'package:dartz/dartz.dart';

import '../../domain/repository/chat_repository.dart';
import '../data_source/remote_abstract.dart';
import '../model/chat_model.dart';

class ChatRepositoryImp extends ChatRepository {
  final RemoteAbstract remoteAbstract;
  // final String _token;

  ChatRepositoryImp({required this.remoteAbstract}) {
    _initializeSocket();
  }
  var _token =
      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJlZGlldEBnbWFpbC5jb20iLCJzdWIiOiI2NmM1ZGM2NzAwNzgxNjk3ZGM3OThmNDQiLCJpYXQiOjE3MjQ1MDM2MTQsImV4cCI6MTcyNDkzNTYxNH0.M6M0gkmlLHw6FxnBynyu1P_TIZGUqZrQGf0mOTxDumU';

  Future<void> _initializeSocket() async {
    // _token = await getToken() ?? '';
    if (_token.isNotEmpty) {
      remoteAbstract.initializeSocket(_token);
      // remoteAbstract.initializeSocket(_token);
    } else {
      throw ServerException("Token retrieval failed");
    }
  }

  @override
  Future<Either<Failure, String>> chatRoom(String receiverId) async {
    if (_token.isEmpty) {
      await _initializeSocket();
    }
    return remoteAbstract.chatRoom(_token, receiverId);
  }

 @override
  Future<Either<Failure, void>> deleteMessage(String chatId) async {
      try {
        await remoteAbstract.deleteMessage(chatId,_token);
        return Right(null);
      } on ServerException {
        return Left(ServerFailure('An error has occurred'));
      } on SocketException {
        return Left(ConnectionFailure('Failed to connect to the network'));
      }
   
  }

  @override
  Stream<MessageModel> getMessages(String chatId) async* {
    if (_token.isEmpty) {
      await _initializeSocket();
    }
    
    yield* remoteAbstract.getMessages(chatId, _token);
   
  }

  @override
  void sendMessage(String chatId, String content, String type) {
    if (_token.isEmpty) {
      _initializeSocket().then((_) {
        remoteAbstract.sendMessage(chatId, content, type);
      }).catchError((e) {
        throw ServerException("Socket Initialization Error: $e");
      });
    } else {
      remoteAbstract.sendMessage(chatId, content, type);
    }
  }

  @override
  Stream<Either<Failure, List<ChatModel>>> getChatHistory() async* {
    if (_token.isEmpty) {
      await _initializeSocket();
    }
    yield* remoteAbstract.getChatHistory(_token);
  }
}
