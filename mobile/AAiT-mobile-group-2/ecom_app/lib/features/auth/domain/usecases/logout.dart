import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../repositories/auth_repository.dart';

class LogoutUsecase implements UseCase<Unit, NoParams> {
  final AuthRepository authRepository;
  LogoutUsecase(this.authRepository);

  @override
  Future<Either<Failure, Unit>> call(NoParams params) {
    return authRepository.logout();
  }
}
