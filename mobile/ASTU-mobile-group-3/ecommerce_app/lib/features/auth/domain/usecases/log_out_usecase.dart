import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../repositories/auth_repository.dart';

class LogOutUsecase {
  final AuthRepository repository;
  LogOutUsecase({required this.repository});

  Future<Either<Failure, bool>> execute() {
    return repository.logOut();
  }
}
