import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../entities/user.dart';
import '../repositories/user_repository.dart';

class RegisterParams {
  late String name;
  late String email;
  late String password;

  RegisterParams({required this.email, required this.password, required this.name});
}


class RegisterUser extends UseCase< User, RegisterParams> {
  final UserRepository userRepository;

  RegisterUser(this.userRepository);

  @override
  Future<Either<Failure, User>> call(RegisterParams params) async {
    return userRepository.registerUser(params.email, params.password, params.name);
  }
}