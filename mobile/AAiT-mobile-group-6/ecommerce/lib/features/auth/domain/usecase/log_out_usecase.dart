
import 'package:dartz/dartz.dart';

import '../../../../core/base_usecase.dart';
import '../../../../core/error/failures.dart';
import '../../../product/domain/usecase/get_all_product.dart';
import '../repository/auth_repo.dart';


class LogoutUsecase extends UseCase<void, NoParams> {
  final AuthRepository repository;

  LogoutUsecase({required this.repository});

  @override
  Future<Either<Failure, void>> call(NoParams param) async {
    return await repository.logOut();
  }
  
  @override
  Future<Either<Failure, void>> execute(NoParams params) {
    // TODO: implement execute
    throw UnimplementedError();
  }
}

