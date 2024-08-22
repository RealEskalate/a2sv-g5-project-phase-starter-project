import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/register_entity.dart';
import '../repositories/auth_repository.dart';

class RegisterUsecase implements UseCase<Unit, RegisterParams> {
  final AuthRepository authRepository;
  RegisterUsecase(this.authRepository);

  @override
  Future<Either<Failure, Unit>> call(RegisterParams params) {
    return authRepository.register(params.registrationEntity);
  }
}

class RegisterParams extends Equatable{
  final RegistrationEntity registrationEntity;

  RegisterParams({required this.registrationEntity});

  @override
  List<Object?> get props => [registrationEntity];
}