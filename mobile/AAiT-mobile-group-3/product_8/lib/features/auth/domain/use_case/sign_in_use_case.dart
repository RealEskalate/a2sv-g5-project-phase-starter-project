import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/sign_in_user_entitiy.dart';
import '../repositories/auth_repository.dart';

class SignInUseCase implements UseCase<Unit, SignInParams> {
  final AuthRepository authRepository;

  SignInUseCase({required this.authRepository});

  @override
  Future<Either<Failure, Unit>> call(SignInParams params) async {
    return  authRepository.signIn(params.signInUserEntitiy);
  }
}

class SignInParams extends Equatable {
  final SignInUserEntitiy signInUserEntitiy;

  SignInParams({required this.signInUserEntitiy});

  @override
  List<Object> get props => [signInUserEntitiy];
} 