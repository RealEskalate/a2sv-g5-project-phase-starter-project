import 'package:dartz/dartz.dart';
import '../../../../core/errors/failure.dart';
import '../entity/auth_entity.dart';


abstract class AuthRepository {
  Future<Either<Failure, AuthResponseEntity>> login(AuthEntity authEntity);
  Future<Either<Failure, void>> register(SentUserEntity senduserentity);
  Future<Either<Failure, UserEntity>> getUserProfile();
  Future<Either<Failure, void>> logout();
}
