import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../repositories/auth_repository.dart';

class LogOutUseCase implements UseCase<Unit, NoParams> {
  final AuthRepository authRepository;

  LogOutUseCase( this.authRepository);

  @override
  Future<Either<Failure, Unit>> call(NoParams params) async {
    return authRepository.logOut();
  }
}