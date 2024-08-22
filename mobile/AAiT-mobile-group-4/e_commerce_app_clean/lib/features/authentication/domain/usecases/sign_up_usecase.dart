import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/sign_up.dart';
import '../repositories/auth_repo.dart';

class SignUpUsecase extends UseCase<void,GetParams>{
  final AuthRepository authRepository;
  
  SignUpUsecase({required this.authRepository});
  @override
  Future<Either<Failure, void>> call(GetParams p) async {
    return await authRepository.signUp(p.signUpEntity);
  }
}
class GetParams extends Equatable {
  final SignUpEntity signUpEntity;

  const GetParams({required this.signUpEntity});
  
  @override
  List<Object?> get props => [signUpEntity];
}