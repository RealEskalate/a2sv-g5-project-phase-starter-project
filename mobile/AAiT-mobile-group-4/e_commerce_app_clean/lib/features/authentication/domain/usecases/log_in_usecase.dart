import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/log_in.dart';
import '../repositories/auth_repo.dart';

class LogInUsecase extends UseCase<void, LogInParams> {

  final AuthRepository authRepository;
  
  LogInUsecase({required this.authRepository});

  @override
  Future<Either<Failure, void>> call(LogInParams p) async {
    return await authRepository.logIn(p.logInEntity);
  }
}

class LogInParams extends Equatable {
  final LogInEntity logInEntity;

  const LogInParams({required this.logInEntity});
  
  @override
  List<Object?> get props => [logInEntity];
}