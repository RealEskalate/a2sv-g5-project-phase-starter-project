import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';

class InitiateChatUsecase {
  final String _userId;

  InitiateChatUsecase({required String userId}) : _userId = userId;

  Future<Either<Failure, String>> initiateChat(String userId) {
    throw UnimplementedError();
  }
}
