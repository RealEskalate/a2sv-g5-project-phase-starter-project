import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/login_entity.dart';
import '../repositories/auth_repository.dart';

class LoginUsecase implements UseCase<Unit, LoginParams>{
  final AuthRepository authRepository;
  LoginUsecase(this.authRepository);

  @override
  Future<Either<Failure, Unit>> call(LoginParams params) {
    return authRepository.login(params.loginEntity);
  }
}

class LoginParams extends Equatable {
  final LoginEntity loginEntity;

  LoginParams({required this.loginEntity});
  @override
  List<Object?> get props => [loginEntity];
}