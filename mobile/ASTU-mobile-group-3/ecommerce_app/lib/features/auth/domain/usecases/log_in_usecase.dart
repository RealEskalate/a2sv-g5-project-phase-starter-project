import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../entities/user_entity.dart';
import '../repositories/auth_repository.dart';

class LogInUsecase {
  final AuthRepository authRepository;

  LogInUsecase({required this.authRepository});

  Future<Either<Failure, bool>> execute(UserEntity user) async {
    return await authRepository.logIn(user);
  }
}
