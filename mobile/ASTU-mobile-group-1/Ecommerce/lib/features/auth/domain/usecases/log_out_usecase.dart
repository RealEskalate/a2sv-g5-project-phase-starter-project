import 'dart:async';

import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../repositories/auth_repository.dart';

class LogOutUsecase {
  final AuthRepository repository;

  LogOutUsecase({required this.repository});

  Future<Either<Failure, bool>> call() async {
    return await repository.logOut();
  }
}
