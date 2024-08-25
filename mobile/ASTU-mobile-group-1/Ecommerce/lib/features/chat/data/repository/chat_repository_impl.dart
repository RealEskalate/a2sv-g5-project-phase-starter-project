import 'dart:async';
import 'dart:developer';

import 'package:dartz/dartz.dart';


import '../../../../core/error/failure.dart';
import '../../../../core/network/network_info.dart';

import '../../domain/repositories/chat_repository.dart';
import '../data_source/local_data_source.dart';
import '../data_source/remote_data_source.dart';

class ChatRepositoryImpl implements ChatRepository {
  final ChatRemoteDataSource chatRemoteDataSource;
  final ChatLocalDataSource chatLocalDataSource;
  final NetworkInfo networkInfo;
  ChatRepositoryImpl(
      {required this.networkInfo,
      required this.chatRemoteDataSource,
      required this.chatLocalDataSource});


  
}
