import 'package:dartz/dartz.dart';

import '../../../../core/base_usecase.dart';
import '../../../../core/error/failures.dart';
import '../entities/log_in_entity.dart';
import '../repository/auth_repo.dart';


class LogInUseCase implements UseCase<String, LogInParams> {
  final AuthRepository authRepository;
  LogInUseCase(this.authRepository);

  @override
  Future<Either<Failure, String>> call(LogInParams params) async {
    return await authRepository.login(params.login);
  }
  
  @override
  Future<Either<Failure, String>> execute(LogInParams params) {
    // TODO: implement execute
    throw UnimplementedError();
  }
}

class LogInParams {
  final LogInEntity login;
  LogInParams(this.login);
}
