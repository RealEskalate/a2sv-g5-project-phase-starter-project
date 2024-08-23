import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../repository/auth_repository.dart';

class GetUser {
  final AuthRepository authrepository;

  GetUser(this.authrepository);
  Future<Either<Failure, String>> execute() {
    return authrepository.getCurrentUser();
  }
}
