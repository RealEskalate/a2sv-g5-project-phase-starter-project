import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/sign_up_user_entitiy.dart';
import '../repositories/auth_repository.dart';

 class SignUpUseCase implements UseCase<Unit, SignUpParams> {
  final AuthRepository authRepository;
  SignUpUseCase({required this.authRepository});

  @override
  Future<Either<Failure, Unit>> call(SignUpParams params) async{
    return  authRepository.signUp(params.signUpUserEntitiy);
  }
}

class SignUpParams extends Equatable {
  final SignUpUserEntitiy signUpUserEntitiy;

  SignUpParams({required this.signUpUserEntitiy});

  @override
  List<Object> get props => [signUpUserEntitiy];
}