import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../entities/user_entity.dart';
import '../repositories/auth_repository.dart';

class GetMeUsecase {
  final AuthRepository authRepository;

  GetMeUsecase({required this.authRepository});

  Future<Either<Failure, UserEntity>> execute() async {
    return await authRepository.getMe();
  }
}
